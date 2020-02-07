package model

import "time"

type User struct {
	Id       int64     `json:"id" form:"id" xorm:"pk autoincr bigint(20)"`
	Mobile   string    `json:"mobile" form:"mobile" xorm:"varchar(40)"`
	Passwd   string    `json:"passwd" form:"passwd" xorm:"varchar(40)"`
	Avatar   string    `json:"avatar" form:"avatar" xorm:"varchar(150)"`
	NickName string    `json:"nick_name" form:"nick_name" xorm:"varchar(40)"`
	Salt     string    `json:"salt" form:"salt" xorm:"varchar(40)"`
	Online   int       `json:"online" form:"online" xorm:"int(10)"`
	Token    string    `json:"token" form:"token" xorm:"varchar(40)"`
	Memo     string    `json:"memo" form:"memo" xorm:"varchar(150)"`
	Creatat  time.Time `json:"creatat" form:"creatat" xorm:"datetime"`
}
