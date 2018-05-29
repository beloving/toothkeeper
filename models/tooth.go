package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Tooth struct {
	Id int64 `orm:"auto"`
	Paramter1 float32
	Paramter2 int64
	Paramter3 string
	Paramter4 string
	Paramter5 string
	Paramter6 string
	Paramter7 string
	Paramter8 string
	Paramter9 string
	Paramter10 string
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
	Delete 	bool `orm:"default(false)"`
}

func CreateTooth(t Tooth)  (index int64,err error){
	t.Created = time.Now()
	t.Updated = time.Now()
	o:=orm.NewOrm()
	index,err = o.Insert(&t)
	return index,err
}

func GetAllTooths() (tooths *[]Tooth,err error){
	var ts *[]Tooth
	_,err=orm.NewOrm().QueryTable("Tooth").All(&ts)
	return ts,err
}
func FindOneTooth(t *Tooth) (tooth *Tooth,err error) {
	err = orm.NewOrm().Read(t)
	return t,err
}
func DeleteTooth(id int64) (err error) {
	o := orm.NewOrm()
	tooth := Tooth{Id: id,Delete:true,Updated:time.Now()}
	_,err=o.Update(&tooth,"Delete","Updated")
	return err
}