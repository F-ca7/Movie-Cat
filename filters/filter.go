package filters

import (
	"github.com/astaxie/beego"

	"github.com/astaxie/beego/context"
	"Movie-Cat/models"
)

//判断是否登录
func IsLogin(ctx *context.Context) (bool, models.User) {
	token, flag := ctx.GetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"))
	var user models.User
	if flag {
		flag, user = models.FindUserByToken(token)
	}
	return flag, user
}

