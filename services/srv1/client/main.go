package main

import (
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "gotests/services/srv1/pb"
    "log"
)

const (
    address     = "localhost:33084"
)

func doCall1(c pb.Service1Client, ctx context.Context) {
    for i := 0; i < 100; i++ {
        msg0In := new(pb.Msg0)
        msg0In.Id1 = int64(i)
        msg0Out, err := c.DoCall1(ctx, msg0In)
        if err != nil {
            log.Printf("err = %v", err)
            continue
        }

        log.Printf("msg0Out id = %d", msg0Out.Id1)
    }
}

func doCallMsg1(c pb.Service1Client, ctx context.Context) {
    for i := 0; i < 100; i++ {
        msg0In := new(pb.Msg0)
        msg0In.Id1 = int64(i)
        msg0Out, err := c.DoCallMsg1(ctx, msg0In)
        if err != nil {
            log.Printf("err = %v", err)
            continue
        }

        log.Printf("msg0Out id1 = %d, id2 = %d, msg1 = %s, msg2 = %s",
            msg0Out.Id1, msg0Out.Id2, msg0Out.Msg1, msg0Out.Msg2)
    }
}

func doCallMsg4(c pb.Service1Client, ctx context.Context) {
    for i := 0; i < 100; i++ {
        msg0In := new(pb.Msg0)
        msg0In.Id1 = int64(i)
        msg0Out, err := c.DoCallMsg4(ctx, msg0In)
        if err != nil {
            log.Printf("err = %v", err)
            continue
        }

        msg3 := msg0Out.Msg3
        msg1 := msg3.Msg1

        log.Printf("msg4 id = %d, msg3 id = %d, msg1 id1 = %d, id2 = %d, msg1 = %s, msg2 = %s",
            msg0Out.Id1, msg3.Id1, msg1.Id1, msg1.Id2, msg1.Msg1, msg1.Msg2)
    }
}

func main() {
    cliConn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer cliConn.Close()
    c := pb.NewService1Client(cliConn)

    ctx := context.Background()


    //doCall1(c, ctx)
    //doCallMsg1(c, ctx)
    doCallMsg4(c, ctx)

    log.Println("done")
}
