package logic

import (
    "context"
    "gotests/services/srv2"
    "gotests/services/srv2/pb"
)

type Server struct {
    pb.UnimplementedService2Server
    impl srv2.ISrv2
}

func NewServer(impl srv2.ISrv2) *Server {
    return &Server{
        impl: impl,
    }
}

func (srv *Server) Login(ctx context.Context, msg *pb.LoginMessage) (*pb.LoginResultMessage, error) {
    userInfo, result, err := srv.impl.Login(ctx, msg.UserName, msg.PassHash)
    if err != nil {
        return nil, err
    }

    ui := &pb.UserInfoMessage{
       Id:       7,
       FistName: userInfo.FistName,
       LastName: userInfo.LastName,
       Email:    userInfo.Email,
    }
    lrm := &pb.LoginResultMessage{
       UserInfo: ui,
       Result:   result,
    }
    return lrm, nil
}
