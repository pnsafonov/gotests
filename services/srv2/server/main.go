package main

import (
    "google.golang.org/grpc"
    "gotests/services/srv2/pb"
    "gotests/services/srv2/server/logic"
    "log"
    "net"
)

const (
    port = ":31141"
)

func main() {
    srv := new(logic.Server)

    lis, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterService2Server(s, srv)
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}