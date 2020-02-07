package service

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
	"workplace/IMGo/model"
	"workplace/IMGo/utils"
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

	salt := fmt.Sprintf("%06d", rand.Int31n(10000))

	user.Mobile = info["mobile"]
	user.Passwd = utils.MakePasswd(info["passwd"], salt)
	user.Avatar = info["avatar"]
	user.Nickname = info["nick_name"]
	user.Salt = salt
	user.Online = 0
	user.Token = fmt.Sprintf("%08d", rand.Int31())
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

	user := model.User{}
	DbEngine.Where("mobile = ?",s).Get(&user)
	counts := 0
	if user.Id != 0{
		counts = 1
	}
	return counts
}

func (this *userService) Login(mobile string, passwd string) (user model.User,err error) {
	//首先通过手机号查询用户
	tmp :=model.User{

	}
	DbEngine.Where("mobile = ?",mobile).Get(&tmp)
	//如果没有找到
	if tmp.Id==0{
		return tmp,errors.New("该用户不存在")
	}
	fmt.Println(tmp)
	//查询到了比对密码
	if !utils.ValidatePasswd(passwd,tmp.Salt,tmp.Passwd){
		return tmp,errors.New("密码不正确")
	}
	//刷新token,安全
	str := fmt.Sprintf("%d",time.Now().Unix())
	token := utils.MD5Encode(str)
	tmp.Token = token
	//返回数据
	DbEngine.ID(tmp.Id).Cols("token").Update(&tmp)
	return tmp,nil
}
