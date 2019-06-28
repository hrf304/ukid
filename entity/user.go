package entity

import "xormt"

func init(){
	xormt.AddModel(new (User))
}

type User struct {
	Id string `xorm:"'id' varchar(36) pk"json:"id"`
	Name string `xorm:"'name' varchar(20) default('')"json:"name"`
	LoginId string `xorm:"'login_id' varchar(20) notnull"json:"login_id"`
	Major string `xorm:"'major' int default(110)"json:"major"`
}

func (u *User)TableName() string{
	return "user"
}
