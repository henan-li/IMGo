package controller

import (
	"net/http"
	"workplace/IMGo/service"
	"workplace/IMGo/structManager"
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

	resp := structManager.H{}
	if res == 1 {
		resp.Data = map[string]string{"mobile":info["mobile"],"avatar":info["avatar"],"nick_name":info["nick_name"],"memo":info["memo"]}
		resp.Msg = "register is done"
		resp.Code = 1
	} else {
		resp.Data = map[string]string{"mobile":info["mobile"],}
		resp.Msg = "register is failed, user already registered"
		resp.Code = -1
	}
	utils.Resp(writer, resp)
}
