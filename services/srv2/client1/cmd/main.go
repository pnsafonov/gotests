package main

import (
    "context"
    "google.golang.org/grpc"
    "gotests/services/srv2"
    client "gotests/services/srv2/client1"
    "gotests/services/srv2/pb"
    "log"
)

const (
    address     = "localhost:31141"
)

func login1(ctx context.Context, cli pb.Service2Client) {
    loginMsg := &pb.LoginMessage{
        UserName: "login1",
        PassHash: "pwd1",
    }

    loginResult, err := cli.Login(ctx, loginMsg)
    if err != nil {
        log.Printf("login1 err = %v\n", err)
        return
    }

    log.Printf("result = %v\n", loginResult.Result)

    ui := loginResult.UserInfo
    log.Printf("FirstName = %v, LastName = %v\n", ui.FistName, ui.LastName)
}

func login2(ctx context.Context, cli srv2.ISrv2) {
    ui, result, err := cli.Login(ctx, "login1", "pwd1")
    if err != nil {
        log.Printf("login1 err = %v\n", err)
        return
    }

    log.Printf("result = %v\n", result)
    log.Printf("FirstName = %v, LastName = %v\n", ui.FistName, ui.LastName)
}

func main() {
    cliConn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer cliConn.Close()
    c := pb.NewService2Client(cliConn)

    ctx := context.Background()
    cli := client.NewClient(c)

    //login1(ctx, c)
    login2(ctx, cli)

    log.Println("done")
}
