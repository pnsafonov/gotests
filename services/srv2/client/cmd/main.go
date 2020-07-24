package main

import (
    "context"
    "google.golang.org/grpc"
    "gotests/services/srv2"
    "gotests/services/srv2/client"
    "log"
)

const (
    address     = "localhost:31141"
)

func login1(ctx context.Context, cli srv2.ISrv2) {
    ui, result, err := cli.Login(ctx, "login1", "pwd1")
    if err != nil {
        log.Printf("login1 err = %v\n", err)
        return
    }

    log.Printf("result = %v\n", result)
    log.Printf("FirstName = %v, LastName = %v\n", ui.FistName, ui.LastName)
}

func main() {
    conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

    cli := client.NewClient(conn)
    ctx := context.Background()

    login1(ctx, cli)

    log.Println("done")
}
