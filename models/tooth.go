package models

import "time"

type tooth struct {
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
	Create   time.Time
	Update   time.Time
}

