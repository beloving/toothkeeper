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
	_,error :=GetUser(&User{Phone: u.Phone})
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
	err = orm.NewOrm().QueryTable("user").Filter("username",uu.Phone).One(&user)
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

func UpdateUser(phone string, uu *User) (a *User, err error) {
	user:=new(User)
	user.Phone = phone
	user,error :=GetUser(user)
	if error == nil{
		if len(uu.Password) != 0 && (strings.Compare(uu.Password,user.Password) !=0){
			user.Password = uu.Password
		}
		if len(uu.Profile.Address) != 0 && (strings.Compare(uu.Profile.Address,user.Profile.Address) !=0){
			user.Profile.Address = uu.Profile.Address
		}
		if uu.Profile.Age != 0 && uu.Profile.Age != user.Profile.Age{
			user.Profile.Age = uu.Profile.Age
		}
		if len(uu.Profile.Email) != 0 && (strings.Compare(uu.Profile.Email,user.Profile.Email) !=0){
			user.Profile.Email = uu.Profile.Email
		}
		if len(uu.Profile.Gender) != 0 && (strings.Compare(uu.Profile.Gender,user.Profile.Gender) !=0){
			user.Profile.Gender = uu.Profile.Gender
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
