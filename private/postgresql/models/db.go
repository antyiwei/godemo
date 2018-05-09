package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Student struct {
	Id   int64
	Name string
	Age  int
}

type UsersMessage struct {
	Id                int
	Title             string
	Content           string
	Uid               int
	CreateTime        time.Time
	MsgType           int
	IsRead            int
	ContentActivityId int
	IsshowApp         int
	IsshowPc          int
	IsshowWap         int
	ShareCounts       int
	Guid              string
	OpeTime           time.Time
	Muuid             string
	MsgClassify       string
	MsgId             int
}

func init() {
	orm.RegisterModel(new(Student), new(UsersMessage))
}
