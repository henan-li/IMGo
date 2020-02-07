package service

import (
	"strconv"
	"time"
	"workplace/IMGo/model"
)

var (
	user = &model.User{}
)

type userService struct {
	Table string
}

func UserServiceConstruct() *userService {
	return &userService{
		Table: "user",
	}
}

//func (this *userService) validateCheck(mobile string) {
//	// check mobile number
//	sql := "select count(*) as phone_count from user where mobile = "+mobile
//	res, err := DbEngine.QueryString(sql)
//
//	fmt.Println(res[0]["phone_count"],err)
//}
func (this *userService) CreateUser(info map[string]string) (affected int64) {

	user.Mobile = info["mobile"]
	user.Passwd = info["passwd"]
	user.Avatar = info["avatar"]
	user.Nickname = info["nick_name"]
	user.Salt = "zxcasd"
	user.Online = 0
	user.Token = "zxcasd"
	user.Memo = info["memo"]
	user.Createat = time.Now()

	counts := this.checkUser(info["mobile"])

	if counts >= 1 {
		affected = 0
	} else {
		affected, _ = DbEngine.Insert(user)
	}
	return affected
}

func (this *userService) DeleteUser(id int64) {

}

func (this *userService) FindUser(conditions interface{}) {

}

func (this *userService) UpdateUser(id int64, kv map[string]interface{}) {

}

func (this *userService) checkUser(s string) int {

	sql := "select count(*) as num_count from user where mobile = " + s
	res, _ := DbEngine.QueryString(sql)
	counts, _ := strconv.Atoi(res[0]["num_count"])
	return counts
}
