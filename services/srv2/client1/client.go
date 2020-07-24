package client

import (
    "context"
    "gotests/services/srv2"
    "gotests/services/srv2/pb"
)

type Client struct {
    cli pb.Service2Client
}

func NewClient(cli pb.Service2Client) *Client {
    return &Client{
        cli: cli,
    }
}

func (cli *Client) Login(ctx context.Context, userName string, passHash string) (*srv2.UserInfo, bool, error) {
    loginMsg := &pb.LoginMessage{
        UserName: userName,
        PassHash: passHash,
    }
    loginResult, err := cli.cli.Login(ctx, loginMsg)
    if err != nil {
        return nil, false, err
    }
    userInfo := loginResult.UserInfo
    ui := &srv2.UserInfo{
        Id:       userInfo.Id,
        FistName: userInfo.FistName,
        LastName: userInfo.LastName,
        Email:    userInfo.Email,
    }
    return ui, true, nil
}
