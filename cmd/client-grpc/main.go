package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/api/v1"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

var rootCmd = &cobra.Command{
	Use:   "client-grpc",
	Short: "Hugo is a very fast static site generator",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

var loginCmd = &cobra.Command{
	Use: "login",
}

func Execute() {

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// get configuration
	address := flag.String("server", ":9090", "gRPC server in format host:port")
	username := flag.String("username", "", "username")
	email := flag.String("email", "", "used in signup")
	action := flag.String("action", "", "action")
	flag.Parse()

	fmt.Printf("Action: \"%s\" \n", *action)

	// Set up a connection to the server.
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := v1.NewAuthenticationServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	switch *action {
	case "login":
		doLogin(ctx, c, username)
	case "signup":
		doSignUp(ctx, c, username, email)
	case "token":
		doTokenValidate(ctx, c)
	default:
		log.Printf("Possible options [login, signup, token]\n")
		return
	}
}

func doTokenValidate(ctx context.Context, c v1.AuthenticationServiceClient) {
	req := v1.ValidateRequest{
		Api:   apiVersion,
		Token: getPassword(),
	}
	res, err := c.Validate(ctx, &req)
	if err != nil {
		log.Fatalf("Validate failed: %v", err)
	}
	log.Printf("Validate result: <%+v>\n\n", res)
}

func doSignUp(ctx context.Context, c v1.AuthenticationServiceClient, username *string, email *string) {
	// Call Create
	req1 := v1.SignUpRequest{
		Api:      apiVersion,
		Username: *username,
		Email:    *email,
		Password: getPassword(),
	}
	res1, err := c.SignUp(ctx, &req1)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create result: <%+v>\n\n", res1)
}

func getPassword() string {
	password := ""
	prompt := &survey.Password{
		Message: "Please type your password",
	}
	survey.AskOne(prompt, &password)
	return password
}

func doLogin(ctx context.Context, c v1.AuthenticationServiceClient, username *string) {
	req2 := v1.LoginRequest{
		Api:      apiVersion,
		Username: *username,
		Password: getPassword(),
	}
	res2, err := c.Login(ctx, &req2)
	if err != nil {
		log.Fatalf("Login failed: %v", err)
	}
	log.Printf("Login result: <%+v>\n\n", res2)
}
