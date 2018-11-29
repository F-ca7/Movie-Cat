package main

import (
	_ "Movie-Cat/routers"
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/mysql"
	"Movie-Cat/models"
)

func init() {
	models.Init()
}


func main() {
	beego.Run()
}

