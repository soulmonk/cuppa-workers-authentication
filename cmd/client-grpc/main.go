package main

import (
	"context"
	"flag"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/api/v1"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

func main() {
	// get configuration
	address := flag.String("server", ":9090", "gRPC server in format host:port")
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := v1.NewAuthenticationServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//
	//// Call Create
	//req1 := v1.SignUpRequest{
	//	Api: apiVersion,
	//	Username: "username",
	//	Email: "username@example.com",
	//	Password: "plainpassword",
	//}
	//res1, err := c.SignUp(ctx, &req1)
	//if err != nil {
	//	log.Fatalf("Create failed: %v", err)
	//}
	//log.Printf("Create result: <%+v>\n\n", res1)

	// Login
	req2 := v1.LoginRequest{
		Api:      apiVersion,
		Username: "username",
		Password: "plainpassword",
	}
	res2, err := c.Login(ctx, &req2)
	if err != nil {
		log.Fatalf("Login failed: %v", err)
	}
	log.Printf("Login result: <%+v>\n\n", res2)

	// Validate
	req3 := v1.ValidateRequest{
		Api:   apiVersion,
		Token: res2.Token,
	}
	res3, err := c.Validate(ctx, &req3)
	if err != nil {
		log.Fatalf("Validate failed: %v", err)
	}
	log.Printf("Validate result: <%+v>\n\n", res3)
}
