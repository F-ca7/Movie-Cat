package controllers

import (
	"github.com/astaxie/beego"
	"Movie-Cat/models"
	"math/rand"
	"time"
	"fmt"
	"github.com/sluu99/uuid"
	"Movie-Cat/filters"
	"crypto/md5"
)
var list = []string{
	"1291552","1291560","1291561","1292001","1292063","1292220","1292370","1292720",
	"1292722","1297359","1298624","1307914","1652587","1889243","1929463","3319755",
	"21318488","1907966","5989818","3395373"}

type IndexController struct {
	beego.Controller
}

type Cover struct {
	ImgPath		string
	Sentence	string
}


func (c *IndexController) Index()  {
	movies := GetRandomBanner()
	movie1 := models.GetMovieById(7)
	movie2 := models.GetMovieById(27)
	movie3 := models.GetMovieById(33)
	movie4 := models.GetMovieById(18)
	selected := GetRandomSelected()
	c.Data["Cover1"] = Cover{"p2189370932.webp",movie1.Quote}
	c.Data["Cover2"] = Cover{"p2151543123.webp",movie2.Quote}
	c.Data["Cover3"] = Cover{"p1325714821.webp",movie3.Quote}
	c.Data["Cover4"] = Cover{"p2233084908.webp",movie4.Quote}
	c.Data["Banner0"] = movies[0]
	c.Data["Banner1"] = movies[1]
	c.Data["Banner2"] = movies[2]
	c.Data["Banner3"] = movies[3]
	c.Data["Banner4"] = movies[4]
	c.Data["Selected0"] = selected[0]
	c.Data["Selected1"] = selected[1]
	c.Data["Selected2"] = selected[2]
	c.Data["Selected3"] = selected[3]
	c.Data["Selected4"] = selected[4]
	isLogin, user := filters.IsLogin(c.Ctx)
	c.Data["IsLogin"] = isLogin
	c.Data["User"] = user
	c.TplName = "index.html"
}

func GetRandomBanner()  []models.Movie{
	index := generateRandomNumber(0, len(list),5)
	var movies = make([]models.Movie,5)
	for i:=range movies{
		movies[i] = models.GetMovieByMId(list[index[i]])
	}
	return movies
}

func GetRandomSelected()  []*models.MovieDetail{
	index := generateRandomNumber(1, 250,5)
	var movies = make([]*models.MovieDetail,5)
	for i:=range movies{
		movies[i] = models.GetDetailById(index[i])
	}
	return movies
}

//生成count个[start,end)结束的不重复的随机数
func generateRandomNumber(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}
	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn(end - start) + start
		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
		}
	}

	return nums
}

func (c *IndexController)ShowRegister  (){
	c.TplName = "register.html"
}

func (c *IndexController)Login () {
	username, password:= c.Input().Get("username"), c.Input().Get("password")
	data := []byte(password)
	has := md5.Sum(data)
	md5pass := fmt.Sprintf("%x", has) //将[]byte转成16进制
	if flag, user := models.Login(username, md5pass); flag {
		c.SetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"), user.Token, 30*24*60*60, "/", beego.AppConfig.String("cookie.domain"), false, true)
		fmt.Printf("用户%s登陆成功", username)
	} else {
		fmt.Println("用户名或密码错误")
	}
	c.Redirect("/index", 302)
}

func (c *IndexController)Register  (){
	username, password,email := c.Input().Get("username"), c.Input().Get("password"),c.Input().Get("email")
	if len(username) == 0 || len(password) == 0 {
		fmt.Println("注册时用户名密码不能为空")
		c.Redirect("/register", 302)
	} else if flag, _ := models.FindUserByUserName(username); flag {
		fmt.Println("用户名已被注册")
		c.Redirect("/register", 302)
	} else {
		token := uuid.Rand().Hex()
		count := rand.Intn(10)+1
		avatar := fmt.Sprintf("/static/img/avatar/%d.jpg",count)
		data := []byte(password)
		has := md5.Sum(data)
		md5pass := fmt.Sprintf("%x", has) //将[]byte转成16进制
		user := models.User{Username: username,Email:email ,Password: md5pass, Avatar: avatar, Token: token}
		models.SaveUser(&user)
		//7*24*60*60是7天；路径是"/"；domain是localhost；https关；http开
		c.SetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"), token, 7 * 24 * 60 * 60, "/", beego.AppConfig.String("cookie.domain"), false, true)
		fmt.Printf("用户%s注册成功",username)
		c.Redirect("/", 302)
	}
}

//登出，即把有效时间改为-1
func (c *IndexController) Logout() {
	c.SetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"), "", -1, "/", beego.AppConfig.String("cookie.domain"), false, true)
	c.Redirect("/", 302)
}