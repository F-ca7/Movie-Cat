package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id 		 int		`orm:"pk;auto"`
	Username string		`orm:"unique"`
	Password string
	Email 	string
	Token 	string 		`orm:"unique"`
	Avatar 	string
	InTime    time.Time `orm:"auto_now_add;type(datetime)"`
}

func Login(username string, password string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Username", username).Filter("Password", password).One(&user)
	return err != orm.ErrNoRows, user
}

func SaveUser(user *User) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(user)
	return id
}

func UpdateUser(user *User) {
	o := orm.NewOrm()
	o.Update(user)
}

func FindUserByUserName(username string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Username", username).One(&user)
	return err != orm.ErrNoRows, user
}

func FindUserByToken(token string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Token", token).One(&user)
	return err != orm.ErrNoRows, user
}