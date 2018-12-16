package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"log"
	"strings"
	"Movie-Cat/utils"
)

type Movie struct {
	Id       int		`orm:"pk;auto"`
	Mid		 string
	Title    string
	Subtitle string
	Other    string
	Desc     string
	Year     string
	Area     string
	Tag      string
	Star     string
	Comment  string
	Quote    string
	Poster   string
}

type MovieDetail struct {
	Id		int
	Mid      string
	Actors  string
	Story	string
	ImdbLink	string
	Comment1 	string
	Comment2 	string
	Comment3 	string
	Trailer		string
}

type MovieSimilar struct {
	Id		int
	SId1	int
	SId2	int
	SId3	int
}

//获取所有的movie
func GetAllMovies()  ([]*Movie){
	o := orm.NewOrm()
	var movie Movie
	var movieList []*Movie
	_,err:=o.QueryTable(movie).All(&movieList)
	if err!=nil {
		fmt.Println(err)
	}
	if movieList==nil {
		log.Fatal("list empty!!")
	}else {
		fmt.Println(movieList)
	}
	return movieList
}

//返回是否存在以及movie
func GetMovieById(id int) (Movie) {
	o := orm.NewOrm()
	var movie Movie
	err := o.QueryTable(movie).Filter("id", id).One(&movie)
	if err!=nil {
		fmt.Println(err)
	}
	return  movie
}

//返回是否存在以及movie
func GetMovieByMId(mid string) (Movie) {
	o := orm.NewOrm()
	var movie Movie
	err := o.QueryTable(movie).Filter("mid", mid).One(&movie)
	if err!=nil {
		fmt.Println(err)
	}
	return  movie
}

//根据分类返回电影
func GetMoviesByCate(cate string)   (bool,[]*Movie) {
	o := orm.NewOrm()
	var movieList []*Movie
	query := fmt.Sprintf("select * from movie where tag like '%%%s%%'",cate)
	fmt.Println(query)
	n,err := o.Raw(query).QueryRows(&movieList)
	if err!=nil{
		fmt.Println(err)
	}
	return n!=0,movieList
}

func GetMoviesByKeyword(keyword string)   (bool,[]*Movie) {
	keyword = strings.TrimSpace(keyword)
	o := orm.NewOrm()
	var movieList []*Movie
	keys := strings.Split(keyword," ")
	var query string
	var q  = make([]string,len(keys))
	for i, key := range keys {
		q[i] = fmt.Sprintf("(tag like '%%%s%%' or title like '%%%s%%' or subtitle like '%%%s%%' or 'desc' like '%%%s%%')",
			key, key, key, key)

	}

	query = "select * from movie where " + strings.Join(q[:]," and ")
	fmt.Println(query)
	n,err := o.Raw(query).QueryRows(&movieList)
	if err!=nil{
		fmt.Println(err)
	}
	return n!=0,movieList
}

//详细信息By MId
func GetDetailByMid(mid string) *MovieDetail {
	o := orm.NewOrm()
	var detail *MovieDetail
	query := fmt.Sprintf("select * from `movie-detail` where mid=%s",mid)
	err := o.Raw(query).QueryRow(&detail)
	if err!=nil{
		fmt.Println(err)
	}
	return detail
}


//详细信息By Id
func GetDetailById(id int) *MovieDetail {
	o := orm.NewOrm()
	var detail *MovieDetail
	query := fmt.Sprintf("select * from `movie-detail` where id=%d",id)
	err := o.Raw(query).QueryRow(&detail)
	if err!=nil{
		fmt.Println(err)
	}
	return detail
}

//获取相似电影
func GetSimilarById(id int)  MovieSimilar{
	var sMovie MovieSimilar
	o := orm.NewOrm()
	query := fmt.Sprintf("select * from `movie-similarity` where id=%d",id)
	err := o.Raw(query).QueryRow(&sMovie)
	if err!=nil{
		fmt.Println(err)
	}
	return sMovie
}

//给电影分页
func PageRankingMovie(p int,size int)  utils.Page{
	o := orm.NewOrm()
	var movie Movie
	var movieList []*Movie
	_,err:=o.QueryTable(movie).Limit(size).Offset((p - 1) * size).All(&movieList)
	if err!=nil {
		fmt.Println(err)
	}
	if movieList==nil {
		log.Fatal("list empty!!")
	}else {
		fmt.Println(movieList)
	}
	count := 250
	return utils.PageUtil(count,p,size,movieList)
}