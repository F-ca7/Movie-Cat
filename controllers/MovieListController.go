package controllers

import (
	"github.com/astaxie/beego"
	"Movie-Cat/models"
	"strconv"
	"Movie-Cat/filters"
)

type MovieListController struct {
	beego.Controller
}

func (c *MovieListController) ListMovies()  {
	category := c.Ctx.Input.Query("category")

	flag,list := models.GetMoviesByCate(category)
	c.Data["Exist"] = flag
	c.Data["List"] = list
	isLogin, user := filters.IsLogin(c.Ctx)
	c.Data["IsLogin"] = isLogin
	c.Data["User"] = user
	c.TplName = "movie-list.html"
}

func (c *MovieListController) RankMovies()  {
	p, _ := strconv.Atoi(c.Ctx.Input.Query("p"))
	if p == 0 {
		p = 1
	}
	size, _ := beego.AppConfig.Int("page.size")
	c.Data["Exist"] = true
	c.Data["Page"] = models.PageRankingMovie(p,size)

	c.TplName = "movie-rank.html"
}

func (c *MovieListController) SearchMovies()  {
	key := c.Ctx.Input.Query("keyword")
	flag,list := models.GetMoviesByKeyword(key)
	c.Data["Exist"] = flag
	c.Data["List"] = list
	c.TplName = "movie-list.html"
}
