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

    //if len(userName) != 0 {
    //    return nil, false, errors.New("fail in Login")
    //}

    ui := &srv2.UserInfo{
       Id:       7,
       FistName: "SomeFirstName1",
       LastName: "SomeLastName",
       Email:    "some_email@gmail.com",
    }

    //if len(userName) != 0 {
    //    return ui, true, errors.New("fail in Login")
    //}

    return ui, true, nil
}