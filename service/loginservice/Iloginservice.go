package loginservice

import (
	"QuizBattle/handler"
)

//ILoginServices to do
type ILoginServices interface {
	Login() bool
}

//Login to do
func Login(loginservice ILoginServices) bool {
	return loginservice.Login()
}

//LoginFactory to do
func LoginFactory(_id string, _pass string) ILoginServices {
	var loginModel ILoginServices

	switch {
	case handler.IsEmailValid(_id):
		loginModel = NewEmailLogin(_id, _pass)
		break
	case handler.IsMobileNumberValid(_id):
		loginModel = NewMobileLogin(_id, _pass)
		break
	default:
		loginModel = NewUsernameLogin(_id, _pass)
		break
	}
	return loginModel
}