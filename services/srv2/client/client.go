package client

import (
    "context"
    "github.com/go-kit/kit/endpoint"
    "google.golang.org/grpc"
    "gotests/services/srv2"
    "gotests/services/srv2/pb"

    grpctransport "github.com/go-kit/kit/transport/grpc"
)

type Client struct {
    call endpoint.Endpoint
}

func (cli *Client) Login(ctx context.Context, userName string, passHash string) (*srv2.UserInfo, bool, error) {
    resp, err := cli.call(ctx, loginRequest{ userName: userName, passHash: passHash})
    if err != nil {
        return nil, false, err
    }
    response := resp.(loginResponse)
    return response.UserInfo, response.Result, nil
}

type loginRequest struct {
    userName string
    passHash string
}

type loginResponse struct {
    UserInfo *srv2.UserInfo
    Result   bool
}

func encodeGRPCLoginRequest(_ context.Context, request interface{}) (interface{}, error) {
    req := request.(loginRequest)
    return &pb.LoginMessage{UserName: req.passHash, PassHash: req.passHash}, nil
}

func decodeGRPCLoginResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
    reply := grpcReply.(*pb.LoginResultMessage)
    ui := &srv2.UserInfo{
        Id:         reply.UserInfo.Id,
        FistName:   reply.UserInfo.FistName,
        LastName:   reply.UserInfo.LastName,
        Email:      reply.UserInfo.Email,
    }
    return loginResponse{UserInfo: ui, Result: reply.Result}, nil
}

func NewClient(conn *grpc.ClientConn) *Client {
    cli := grpctransport.NewClient(conn,
        "pb.Service2", "Login",
        encodeGRPCLoginRequest,
        decodeGRPCLoginResponse,
        pb.LoginResultMessage{},
    )
    ep := cli.Endpoint()
    return &Client{call: ep}
}
