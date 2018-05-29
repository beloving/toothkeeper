package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"strings"
	"github.com/astaxie/beego"
	"time"
)
var NotExit = errors.New("not exit")

type User struct {
	Id       int64 `orm:"auto"`
	Phone    string
	Password string
	Profile  Profile  `orm:"-"`
	Create   time.Time
	Update   time.Time
}

type Profile struct {
	Gender  string
	Age     int
	Address string
	Email   string
	Head    string
	Username string
	Create   time.Time
	Update   time.Time
}

func AddUser(u User)(va int64,err error){
	_,error :=GetUser(&User{Username: u.Username})
	if error == NotExit{
		o:=orm.NewOrm()
		index,err:=o.Insert(&u)
		beego.Warn(" index=%d",index)
		if err != nil{
			return -1,err
		}else{
			return index,nil
		}
	}else{
		return -1,error
	}

}

func GetUser(uu *User) (u *User, err error) {
	var user User
	err = orm.NewOrm().QueryTable("user").Filter("username",uu.Username).One(&user)
	beego.Warn(err)
	if err ==nil {
		return uu, nil
	}else{
		return nil, NotExit
	}
}
func GetAllUsers() *User {
	users := new(User)
	orm.NewOrm().QueryTable(new(User)).All(users)
	return users
}

func UpdateUser(username string, uu *User) (a *User, err error) {
	user:=new(User)
	user.Username = username
	user,error :=GetUser(user)
	if error == nil{
		if len(uu.Password) != 0 && (strings.Compare(uu.Password,user.Password) !=0){
			user.Password = uu.Password
		}
		if len(uu.Address) != 0 && (strings.Compare(uu.Address,user.Address) !=0){
			user.Address = uu.Address
		}
		if uu.Age != 0 && uu.Age != user.Age{
			user.Age = uu.Age
		}
		if len(uu.Email) != 0 && (strings.Compare(uu.Email,user.Email) !=0){
			user.Email = uu.Email
		}
		if len(uu.Gender) != 0 && (strings.Compare(uu.Gender,user.Gender) !=0){
			user.Gender = uu.Gender
		}
		_,err :=orm.NewOrm().Update(user)
		if err == nil {
			return user ,nil
		}else{
			return uu ,err
		}
	}else{
		return uu , error
	}
}

func Login(user User) bool {
	u,error :=GetUser(&user)
	if error == nil{
		return u.Password == user.Password
	}
	return false
}
