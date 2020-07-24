package srv2

import "context"

type UserInfo struct {
    Id       int64
    FistName string
    LastName string
    Email    string
}

type ISrv2 interface {
    Login(ctx context.Context, userName string, passHash string) (*UserInfo, bool, error)
}
