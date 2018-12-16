package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"time"
)

type Comment struct {
	Id 			int		`orm:"pk;auto"`
	MovieId		int
	Content		string
	Username	string
	InTime		time.Time	`orm:"auto_now_add;type(datetime)"`
}

func SaveComment(comment Comment) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(&comment)
	return id
}

func FindCommentsByMoiveId(id int) []*Comment {
	o := orm.NewOrm()
	var comment Comment
	var comments []*Comment
	_,err:=o.QueryTable(comment).RelatedSel().Filter("MovieId", id).OrderBy( "-InTime").Limit(3).All(&comments)
	if err!=nil {
		fmt.Println(err)
	}
	return comments
}