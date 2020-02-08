package controller

import (
	"net/http"
	"workplace/IMGo/args"
	"workplace/IMGo/model"
	"workplace/IMGo/service"
	"workplace/IMGo/utils"
)

var (
	contactService = service.ContactServiceConstruct()
)

type contactController struct {
}

func (c *contactController) AddFriends(writer http.ResponseWriter, request *http.Request) {
	var arg args.ContactArg
	utils.Bind(request, &arg)
	err := contactService.AddFriend(arg.Userid, arg.Dstid)
	if err != nil {
		utils.RespFail(writer, err.Error())
	} else {
		utils.RespOk(writer, nil, "好友添加成功")
	}
}

func (c *contactController) JoinGroup(writer http.ResponseWriter, request *http.Request) {
	var arg args.ContactArg
	//如果这个用的上,那么可以直接
	utils.Bind(request, &arg)
	err := contactService.JoinCommunity(arg.Userid, arg.Dstid)
	//todo 刷新用户的群组信息
	//AddGroupId(arg.Userid,arg.Dstid)
	if err != nil {
		utils.RespFail(writer, err.Error())
	} else {
		utils.RespOk(writer, nil, "")
	}
}

func (c *contactController) LoadGroup(writer http.ResponseWriter, request *http.Request) {
	var arg args.ContactArg
	//如果这个用的上,那么可以直接
	utils.Bind(request, &arg)
	comunitys := contactService.SearchComunity(arg.Userid)
	utils.RespOkList(writer, comunitys, len(comunitys))
}

func (c *contactController) LoadFriend(writer http.ResponseWriter, request *http.Request) {

	var arg args.ContactArg
	//如果这个用的上,那么可以直接
	utils.Bind(request, &arg)

	users := contactService.SearchFriend(arg.Userid)
	utils.RespOkList(writer, users, len(users))
}

func (c *contactController) CreateGroup(writer http.ResponseWriter, request *http.Request) {
	var arg model.Community
	//如果这个用的上,那么可以直接
	utils.Bind(request, &arg)
	com, err := contactService.CreateCommunity(arg)
	if err != nil {
		utils.RespFail(writer, err.Error())
	} else {
		utils.RespOk(writer, com, "")
	}
}

func ContactControllerConstruct() *contactController {

	return &contactController{}
}
