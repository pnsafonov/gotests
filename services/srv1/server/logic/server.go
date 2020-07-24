package logic

import (
    "fmt"
    "golang.org/x/net/context"
    "gotests/services/srv1/pb"
    "log"
)

type Server struct {
    pb.UnimplementedService1Server
}


func (srv *Server) DoCall1(ctx context.Context, msg0 *pb.Msg0) (*pb.Msg0, error) {
    log.Printf("DoCall1 msg0.Id1 = %d", msg0.Id1)

    msgOut := new(pb.Msg0)
    msgOut.Id1 = msg0.Id1 * 1
    return msgOut, nil
}
func (srv *Server) DoCallMsg1(ctx context.Context, msg0 *pb.Msg0) (*pb.Msg1, error) {
    log.Printf("DoCallMsg1 msg0.Id1 = %d", msg0.Id1)

    msgOut := new(pb.Msg1)
    msgOut.Id1 = msg0.Id1 * 1
    msgOut.Id2 = msg0.Id1 * 1
    msgOut.Msg1 = "msg1"
    msgOut.Msg2 = "msg2"
    return msgOut, nil
}
func (srv *Server) DoCallMsg2(ctx context.Context, msg0 *pb.Msg0) (*pb.Msg2, error) {
    log.Printf("DoCallMsg2 msg0.Id1 = %d", msg0.Id1)

    msgOut := new(pb.Msg2)
    msgOut.Id = msg0.Id1 * 2
    msgOut.Msg = "msg"
    return msgOut, nil
}
func (srv *Server) DoCallMsg3(ctx context.Context, msg0 *pb.Msg0) (*pb.Msg3, error) {
    log.Printf("DoCallMsg3 msg0.Id1 = %d", msg0.Id1)

    msgOut := new(pb.Msg3)
    msgOut.Id1 = msg0.Id1 * 3

    msg1Out := new(pb.Msg1)
    msg1Out.Id1 = msg0.Id1 * 1
    msg1Out.Id2 = msg0.Id1 * 1
    msg1Out.Msg1 = "msg1"
    msg1Out.Msg2 = "msg2"

    msgOut.Msg1 = msg1Out

    return msgOut, nil
}
func (srv *Server) DoCallMsg4(ctx context.Context, msg0 *pb.Msg0) (*pb.Msg4, error) {
    log.Printf("DoCallMsg4 msg0.Id1 = %d", msg0.Id1)

    msgOut := new(pb.Msg4)
    msgOut.Id1 = msg0.Id1 * 4

    msg1Out := new(pb.Msg1)
    msg1Out.Id1 = msg0.Id1 * 1
    msg1Out.Id2 = msg0.Id1 * 1
    msg1Out.Msg1 = fmt.Sprintf("msg1_%d", msg0.Id1)
    msg1Out.Msg2 = fmt.Sprintf("msg2_%d", msg0.Id1)

    msg3Out := new(pb.Msg3)
    msg3Out.Id1 = msg0.Id1 * 3
    msg3Out.Msg1 = msg1Out

    msgOut.Msg3 = msg3Out

    return msgOut, nil
}