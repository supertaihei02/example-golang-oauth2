package routers

import (
	"github.com/supertaihei02/example-golang-oauth2/controllers"
	"github.com/supertaihei02/example-golang-oauth2/controllers/facebook"
	"github.com/supertaihei02/example-golang-oauth2/controllers/github"
	"github.com/supertaihei02/example-golang-oauth2/controllers/google"
	"github.com/supertaihei02/example-golang-oauth2/controllers/twitter"

	"github.com/astaxie/beego"
	"github.com/supertaihei02/example-golang-oauth2/controllers/line"
)

func init() {
	beego.Router("/google/oauth2", &controllersGoogle.Oauth2Controller{})
	beego.Router("/google/callback", &controllersGoogle.CallbackController{})
	beego.Router("/twitter/oauth", &controllersTwitter.Oauth2Controller{})
	beego.Router("/twitter/callback", &controllersTwitter.CallbackController{})
	beego.Router("/twitter/post", &controllersTwitter.TweetController{})
	beego.Router("/facebook/oauth2", &controllersFacebook.Oauth2Controller{})
	beego.Router("/facebook/callback", &controllersFacebook.CallbackController{})
	beego.Router("/line/oauth2", &controllersLine.Oauth2Controller{})
	beego.Router("/line/callback", &controllersLine.CallbackController{})
	beego.Router("/github/oauth2", &controllersGithub.Oauth2Controller{})
	beego.Router("/github/callback", &controllersGithub.CallbackController{})

	beego.Router("/", &controllers.MainController{})
}
