package controller

import (
	"net/http"
	"workplace/IMGo/service"
	"workplace/IMGo/utils"
)

// 通常一个控制器对应一个service, service负责和BD交互, controller负责:数据接收,数据校验,数据转发,返回结果
var (
	userService = service.UserServiceConstruct()
)

type userController struct {
	Mobile string
	Pwd    string
}

func UserControllerConstruct() *userController {
	return &userController{}
}

func (this *userController) UserRegister(writer http.ResponseWriter, request *http.Request) {

	info := make(map[string]string)

	request.ParseForm()
	info["mobile"] = request.PostFormValue("mobile")
	info["passwd"] = request.PostFormValue("passwd")
	info["avatar"] = request.PostFormValue("avatar")
	info["nick_name"] = request.PostFormValue("nick_name")
	info["memo"] = request.PostFormValue("memo")
	res := userService.CreateUser(info)

	if res == 1 {
		delete(info, "passwd")
		utils.RespOk(writer, info, "register ok")
	} else {
		utils.RespFail(writer, "register fail")
	}
}

func (this *userController) UserLogin(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()

	mobile := request.PostForm.Get("mobile")
	passwd := request.PostForm.Get("passwd")

	user, err := userService.Login(mobile, passwd)

	if err!=nil{
		utils.RespFail(writer,err.Error())
	}else{
		utils.RespOk(writer,user,"")
	}
}
