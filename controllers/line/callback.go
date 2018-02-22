package controllersLine

import (
	"log"
	"context"
	"github.com/supertaihei02/example-golang-oauth2/lib/line"

	"github.com/astaxie/beego"
)

// CallbackController コールバックコントローラ
type CallbackController struct {
	beego.Controller
}

// CallbackRequest コールバックリクエスト
type CallbackRequest struct {
	Code  string `form:"code"`
	State int    `form:"state"`
}

// Get コールバックする
func (c *CallbackController) Get() {
	c.StartSession()

	request := CallbackRequest{}
	if err := c.ParseForm(&request); err != nil {
		panic(err)
	}
	log.Printf("request %#v", request)

	config := line.GetConnect()

	context := context.Background()

	log.Printf("config %#v", config)
	token, err := line.GetAccessToken(context, request.Code)
	//tok, err := config.Exchange(context, request.Code)
	if err != nil {
		panic(err)
	}
	log.Printf("token %#v", token)

	// TODO トークン使っていろいろする
	profile := line.Profile{}
	if err = line.GetProfileToken(context, token.AccessToken, &profile); err != nil {
		panic(err)
	}
	log.Printf("profile %#v", profile)

	c.Data["AccessToken"] = token.AccessToken
	c.Data["DisplayName"] = profile.DisplayName
	c.Data["ID"] = profile.ID
	c.Data["PictureUrl"] = profile.PictureUrl
	c.Data["StatusMessage"] = profile.StatusMessage
	c.Data["Code"] = request.Code
	c.TplName = "line/callback.tpl"
}
