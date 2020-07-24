package logic

import (
    "context"
    "gotests/services/srv2/pb"
    "log"
)

type Server struct {
    pb.UnimplementedService2Server
}

func (srv *Server) Login(_ context.Context, msg *pb.LoginMessage) (*pb.LoginResultMessage, error) {
    log.Printf("Login username = %s, passHash = %s\n", msg.UserName, msg.PassHash)

    ui := &pb.UserInfoMessage{
        Id:       7,
        FistName: "SomeFirstName1",
        LastName: "SomeLastName",
        Email:    "some_email@gmail.com",
    }

    result := &pb.LoginResultMessage{
        UserInfo: ui,
        Result:   true,
    }

    return result, nil
}
