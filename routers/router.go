package routers

import (
	"Movie-Cat/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{},"get:Index")
	beego.Router("/index", &controllers.IndexController{},"get:Index")
	beego.Router("/register", &controllers.IndexController{},"get:ShowRegister")
	beego.Router("/register", &controllers.IndexController{},"post:Register")
	beego.Router("/logout", &controllers.IndexController{}, "get:Logout")

	beego.Router("/login", &controllers.IndexController{},"post:Login")


	beego.Router("/list", &controllers.MovieListController{},"get:ListMovies")
	beego.Router("/top250", &controllers.MovieListController{},"get:RankMovies")
	beego.Router("/imdb_top250", &controllers.MovieListController{},"get:RankMovies")
	beego.Router("/search", &controllers.MovieListController{},"get:SearchMovies")

	beego.Router("/movie", &controllers.DetailController{},"get:ShowDetails")
	beego.Router("/postcomment", &controllers.DetailController{},"post:PostComment")
}
