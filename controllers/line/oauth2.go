package controllersLine

import (
	"github.com/supertaihei02/example-golang-oauth2/lib/line"

	"github.com/astaxie/beego"
)

// Oauth2Controller Oauth2コントローラー
type Oauth2Controller struct {
	beego.Controller
}

// Get 認証する
func (c *Oauth2Controller) Get() {
	config := line.GetConnect()

	url := config.AuthCodeURL("")

	c.Redirect(url, 302)
}
