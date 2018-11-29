package controllers

import (
	"Movie-Cat/models"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
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
	c.Data["Color"] = GetColorByTag(movie.Tag)
	c.Data["Movie"] = movie
	c.Data["Detail"] = models.GetDetailByMid(movie.Mid)
	c.Data["SMovie1"] = s1
	c.Data["SMovie2"] = s2
	c.Data["SMovie3"] = s3
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