package logic

import (
    "context"
    "gotests/services/srv2"
    "log"
)

type ServerImpl struct {

}

func (srv *ServerImpl) Login(ctx context.Context, userName string, passHash string) (*srv2.UserInfo, bool, error) {
    log.Printf("Login username = %s, passHash = %s\n", userName, passHash)

    ui := &srv2.UserInfo{
       Id:       7,
       FistName: "SomeFirstName1",
       LastName: "SomeLastName",
       Email:    "some_email@gmail.com",
    }

    return ui, true, nil
}