package main

import (
	"context"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/api/authentication"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

const (
	usernameFlag = "username"
	emailFlag    = "email"
	serverFlag   = "server"
	userIdFlag   = "id"
)

var rootCmd = &cobra.Command{
	Use:   "client-grpc",
	Short: "simple cli to login, signup, validate token",
	Long:  ``,
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login user",
	Run: func(cmd *cobra.Command, args []string) {
		username := cmd.Flag(usernameFlag).Value.String()
		server := cmd.Flag(serverFlag).Value.String()

		conn, ctx, cancel := createConnection(server)

		defer conn.Close()
		defer cancel()

		doLogin(ctx, conn, username)
	},
}
var signupCmd = &cobra.Command{
	Use:   "signup",
	Short: "signup user",
	Run: func(cmd *cobra.Command, args []string) {
		username := cmd.Flag(usernameFlag).Value.String()
		email := cmd.Flag(emailFlag).Value.String()
		server := cmd.Flag(serverFlag).Value.String()

		conn, ctx, cancel := createConnection(server)

		defer conn.Close()
		defer cancel()

		doSignUp(ctx, conn, username, email)
	},
}
var activateCmd = &cobra.Command{
	Use:   "activate",
	Short: "activate user",
	Run: func(cmd *cobra.Command, args []string) {
		id := cmd.Flag(userIdFlag).Value.String()
		server := cmd.Flag(serverFlag).Value.String()

		conn, ctx, cancel := createConnection(server)

		defer func(conn *grpc.ClientConn) {
			err := conn.Close()
			if err != nil {
				log.Fatalln("Could not close conn")
			}
		}(conn)
		defer cancel()

		uid, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			// todo log
			fmt.Printf("Cannot parse id: %s", id)
			return
		}
		doActivate(ctx, conn, uid)
	},
}
var validateTokenCmd = &cobra.Command{
	Use:   "validate",
	Short: "validate user token",
	Run: func(cmd *cobra.Command, args []string) {
		server := cmd.Flag(serverFlag).Value.String()

		conn, ctx, cancel := createConnection(server)

		defer conn.Close()
		defer cancel()

		doTokenValidate(ctx, conn)
	},
}

func main() {
	rootCmd.PersistentFlags().StringP(serverFlag, "s", ":9090", "gRPC server in format host:port")

	loginCmd.Flags().StringP(usernameFlag, "u", "", "username")
	_ = loginCmd.MarkFlagRequired(usernameFlag)

	activateCmd.Flags().Int64(userIdFlag, 0, "user id")
	_ = activateCmd.MarkFlagRequired(userIdFlag)

	signupCmd.Flags().StringP(usernameFlag, "u", "", "username")
	signupCmd.Flags().StringP(emailFlag, "e", "", "email")
	_ = signupCmd.MarkFlagRequired(usernameFlag)
	_ = signupCmd.MarkFlagRequired(emailFlag)

	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(signupCmd)
	rootCmd.AddCommand(validateTokenCmd)
	rootCmd.AddCommand(activateCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func createConnection(address string) (conn *grpc.ClientConn, ctx context.Context, cancel context.CancelFunc) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	return conn, ctx, cancel
}

func doTokenValidate(ctx context.Context, conn *grpc.ClientConn) {
	c := authentication_v1.NewAuthenticationServiceClient(conn)

	req := authentication_v1.ValidateRequest{
		Api:   apiVersion,
		Token: getToken(),
	}
	res, err := c.Validate(ctx, &req)
	if err != nil {
		log.Fatalf("Validate failed: %v", err)
	}
	log.Printf("Validate result: <%+v>\n\n", res)
}

func doSignUp(ctx context.Context, conn *grpc.ClientConn, username string, email string) {
	c := authentication_v1.NewAuthenticationServiceClient(conn)
	// Call Create
	req := authentication_v1.SignUpRequest{
		Api:      apiVersion,
		Username: username,
		Email:    email,
		Password: getPassword(),
	}
	res, err := c.SignUp(ctx, &req)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create result: <%+v>\n\n", res)
}

func getPassword() string {
	password := ""
	prompt := &survey.Password{
		Message: "Please type your password",
	}
	survey.AskOne(prompt, &password, survey.WithValidator(survey.Required))
	return password
}

func getToken() string {
	token := ""
	prompt := &survey.Password{
		Message: "Please type user token",
	}
	survey.AskOne(prompt, &token, survey.WithValidator(survey.Required))
	return token
}

func doActivate(ctx context.Context, conn *grpc.ClientConn, id int64) {
	c := authentication_v1.NewAuthenticationServiceClient(conn)
	req := authentication_v1.ActivateRequest{
		Api:    apiVersion,
		Id:     id,
		Secret: getPassword(),
	}
	res, err := c.Activate(ctx, &req)
	if err != nil {
		log.Fatalf("Active failed: %v", err)
	}
	log.Printf("Active result: <%+v>\n\n", res)
}

func doLogin(ctx context.Context, conn *grpc.ClientConn, username string) {
	c := authentication_v1.NewAuthenticationServiceClient(conn)
	req := authentication_v1.LoginRequest{
		Api:      apiVersion,
		Username: username,
		Password: getPassword(),
	}
	res, err := c.Login(ctx, &req)
	if err != nil {
		log.Fatalf("Login failed: %v", err)
	}
	log.Printf("Login result: <%+v>\n\n", res)
}
