package migration

import (
	"context"
	"github.com/jackc/pgx/v4"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

type migrationFile struct {
	name string
	data []byte
}

type migrationFiles []migrationFile

func (p migrationFiles) Len() int           { return len(p) }
func (p migrationFiles) Less(i, j int) bool { return p[i].name < p[j].name }
func (p migrationFiles) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort is a convenience method.
func (p migrationFiles) Sort() { sort.Sort(p) }

func Proceed(conn *pgx.Conn) error {
	var err error
	var count uint64
	if count, err = initTable(conn); err != nil {
		return err
	}

	log.Println("Existing migration count: ", count)

	var storedMigrations []string
	if storedMigrations, err = loadMigration(conn, count); err != nil {
		return err
	}

	log.Println("Loaded migrations from conn", storedMigrations)

	migrationData, err := getMigrationList(sort.StringSlice(storedMigrations))
	if err != nil {
		return err
	}

	if migrationData.Len() == 0 {
		log.Println("No new migrations")
		return nil
	}

	log.Println("Exists new migrations", len(migrationData))

	return applyMigration(conn, migrationData)
}

func applyMigration(conn *pgx.Conn, files migrationFiles) (err error) {
	ctx := context.Background()
	for _, file := range files {
		log.Println("Start process data: ", file.name)
		var query = string(file.data)
		log.Println("Run query:\n", query)
		if _, err = conn.Exec(ctx, query); err != nil {
			log.Println("Error applying migration", err.Error())
			return
		}

		if err = create(conn, file.name); err != nil {
			return
		}
	}

	return
}

func create(conn *pgx.Conn, name string) error {
	query := `INSERT INTO migrations (name, created_at) VALUES ($1, now()) RETURNING id`
	ctx := context.Background()
	var id uint64
	err := conn.
		QueryRow(ctx, query, name).
		Scan(&id)

	if err != nil {
		log.Println("Error on create note")
		return err
	}
	log.Println("Created new record:", id)

	return nil
}

func loadMigration(conn *pgx.Conn, count uint64) (migrations []string, err error) {
	migrations = make([]string, count)

	if count == 0 {
		return
	}

	ctx := context.Background()

	rows, err := conn.Query(ctx, `SELECT name FROM migrations`)

	if err != nil {
		log.Println("Error on executing query")
		return nil, err
	}

	defer rows.Close()

	var name string
	for rows.Next() {
		if err := rows.Scan(&name); err != nil {
			log.Println("Error corrupted while scanning migration:", err.Error())
			return nil, err
		}

		migrations = append(migrations, name)
	}
	if err := rows.Err(); err != nil {
		log.Println("Error on migrations rows:", err.Error())
		return nil, err
	}
	return migrations, err
}

func initTable(conn *pgx.Conn) (uint64, error) {

	ctx := context.Background()
	query := `SELECT COUNT(1) FROM migrations;`
	var count uint64
	err := conn.QueryRow(ctx, query).Scan(&count)

	if err != nil {
		// todo check error table dose not exists
		log.Println("Some error on get count from migration", err.Error())

		//return count, err

		query = `CREATE TABLE migrations(
  id serial not null
        constraint migrations_pk
            primary key,
  name varchar(255),
  created_at timestamp default now()
);
`
		if _, err := conn.Exec(ctx, query); err != nil {
			log.Println("Error initTable", err.Error())
			return count, err
		}
	}

	return count, nil
}

func getMigrationList(storedMigrations sort.StringSlice) (result migrationFiles, err error) {
	var files []os.FileInfo

	// from cwd => pkg/db/pg/migration
	const mathToMigrations = "./pkg/db/pg/migration/"
	if files, err = ioutil.ReadDir(mathToMigrations); err != nil {
		return nil, err
	}

	storedMigrations.Sort()
	var lastIdx = storedMigrations.Len()
	for _, file := range files {
		if file.Name() == "migration.go" || storedMigrations.Search(file.Name()) != lastIdx {
			continue
		}
		_migrationFile := migrationFile{
			name: file.Name(),
		}
		_migrationFile.data, err = ioutil.ReadFile(mathToMigrations + file.Name())
		if err != nil {
			return
		}
		result = append(result, _migrationFile)
		log.Println("filename in migration folder", file.Name())
	}
	result.Sort()
	return
}
