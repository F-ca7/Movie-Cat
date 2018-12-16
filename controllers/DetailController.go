package controllers

import (
	"Movie-Cat/models"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
	"fmt"
	"math/rand"
	"time"
	"Movie-Cat/filters"
)

type DetailController struct {
	beego.Controller
}


func (c *DetailController) ShowDetails()  {
	id, _ := strconv.Atoi(c.Ctx.Input.Query("id"))

	movie := models.GetMovieById(id)
	sMovie := models.GetSimilarById(id)
	s1 := models.GetMovieById(sMovie.SId1)
	s2 := models.GetMovieById(sMovie.SId2)
	s3 := models.GetMovieById(sMovie.SId3)
	rand.Seed(time.Now().UnixNano())
	//获取评论
	comments := models.FindCommentsByMoiveId(id)
	c.Data["Comments"]	= comments
	//随机头像
	rand_ava := rand.Intn(7)+4
	//fmt.Printf("随机头像:%d\n",rand_ava)
	c.Data["Rand"] = rand_ava

	c.Data["Color"] = GetColorByTag(movie.Tag)
	c.Data["Movie"] = movie
	c.Data["Detail"] = models.GetDetailByMid(movie.Mid)

	c.Data["SMovie1"] = s1
	c.Data["SMovie2"] = s2
	c.Data["SMovie3"] = s3
	isLogin, user := filters.IsLogin(c.Ctx)
	c.Data["IsLogin"] = isLogin
	c.Data["User"] = user
	c.TplName = "movie-details.html"
}

func GetColorByTag(p string)  string{
	var tags = []string{"爱情","犯罪","恐怖","惊悚","动画","动作","冒险","喜剧"}
	if strings.Contains(p,tags[0]){			//"爱情"
		return "pink"
	}else if strings.Contains(p,tags[1]){	//"犯罪"
		return "red"
	}else if strings.Contains(p,tags[2]) || strings.Contains(p,tags[3]){		//"恐怖"
		return "cyan"
	} else if strings.Contains(p,tags[4]){		//"动画"
		return "blue"
	}else if strings.Contains(p,tags[5]) {		//"动作"
		return "graphite"
	}else if strings.Contains(p,tags[6]){		//"冒险"
		return "twany"
	}else if strings.Contains(p,tags[7]){		//"喜剧"
		return "lightgreen"
	} else {					//其他
		return "red"
	}
}

func (c *DetailController)PostComment	(){
	content,username := c.Input().Get("comment"),c.Input().Get("username")
	movie_id,_ := strconv.Atoi(c.Input().Get("movie-id"))
	comment := models.Comment{MovieId:movie_id, Content:content, Username:username}
	models.SaveComment(comment)
	fmt.Printf("%s 发布评论成功",username)
	url := fmt.Sprintf("/movie?id=%d",movie_id)

	c.Redirect(url,302)
}