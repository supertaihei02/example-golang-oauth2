package line

import (
	"context"
	"net/http"
	"net/url"

	"github.com/astaxie/beego"
	"golang.org/x/oauth2"
	"time"
	"strings"
	"io/ioutil"
	"io"
	"fmt"
	"log"
	"encoding/json"
)

/**
 developer guide
 https://developers.line.me/ja/docs/line-login/web/integrate-line-login-v2/
 */

const (
	authorizeEndpoint = "https://access.line.me/dialog/oauth/weblogin"
	tokenEndpoint     = "https://api.line.me/v2/oauth/accessToken"
	profileEndpoint		= "https://api.line.me/v2/profile"
)

// Line プロフィール
type Profile struct {
	ID              string `json:"userId"`
	DisplayName      string `json:"displayName"`
	PictureUrl string `json:"pictureUrl"`
	StatusMessage           string `json:"statusMessage"`
}

type Token struct {
	AccessToken string
	TokenType string
	RefreshToken string
	Expiry time.Time
	Raw interface{}
}

type tokenJSON struct {
	AccessToken  string         `json:"access_token"`
	TokenType    string         `json:"token_type"`
	RefreshToken string         `json:"refresh_token"`
	ExpiresIn    expirationTime `json:"expires_in"` // at least PayPal returns string, while most return number
	Expires      expirationTime `json:"expires"`    // broken Facebook spelling of expires_in
}

func (e *tokenJSON) expiry() (t time.Time) {
	if v := e.ExpiresIn; v != 0 {
		return time.Now().Add(time.Duration(v) * time.Second)
	}
	if v := e.Expires; v != 0 {
		return time.Now().Add(time.Duration(v) * time.Second)
	}
	return
}

type expirationTime int32

func (e *expirationTime) UnmarshalJSON(b []byte) error {
	var n json.Number
	err := json.Unmarshal(b, &n)
	if err != nil {
		return err
	}
	i, err := n.Int64()
	if err != nil {
		return err
	}
	*e = expirationTime(i)
	return nil
}

// GetConnect 接続を取得する
func GetConnect() *oauth2.Config {
	config := &oauth2.Config{
		ClientID:     beego.AppConfig.String("lineClientID"),
		ClientSecret: beego.AppConfig.String("lineClientSecret"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  authorizeEndpoint,
			TokenURL: tokenEndpoint,
		},
		RedirectURL: beego.AppConfig.String("lineRedirectURL"),
	}

	return config
}

/**
 * golang/x/oauth2ではclient_secretがうまく設定されず動作しないため手実装
 */
func GetAccessToken(ctx context.Context, code string) (*Token, error) {

	values := url.Values{}
	values.Set("grant_type", "authorization_code")
	values.Set("client_id", beego.AppConfig.String("lineClientID"))
	values.Set("client_secret", beego.AppConfig.String("lineClientSecret"))
	values.Set("code", code)
	values.Set("redirect_uri", beego.AppConfig.String("lineRedirectURL"))

	req, err := http.NewRequest(
		"POST",
		tokenEndpoint,
		strings.NewReader(values.Encode()),
	)
	if err != nil {
		return nil, err
	}

	// Content-Type 設定
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	r, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1<<20))
	log.Printf("body %#v", body)
	if err != nil {
		return nil, fmt.Errorf("oauth2: cannot fetch token: %v", err)
	}
	if code := r.StatusCode; code < 200 || code > 299 {
		return nil, fmt.Errorf("oauth2: cannot fetch token: %v\nResponse: %s", r.Status, body)
	}

	var token *Token
	// Json形式だとわかっている
	var tj tokenJSON
	if err = json.Unmarshal(body, &tj); err != nil {
		return nil, err
	}
	token = &Token{
		AccessToken:  tj.AccessToken,
		TokenType:    tj.TokenType,
		RefreshToken: tj.RefreshToken,
		Expiry:       tj.expiry(),
		Raw:          make(map[string]interface{}),
	}
	json.Unmarshal(body, &token.Raw)
	return token, nil
}

func GetProfileToken(ctx context.Context, at string, profile *Profile) error {
	req, err := http.NewRequest(
		"GET",
		profileEndpoint,
		nil,
	)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer {"+at+"}")
	client := &http.Client{}
	log.Printf("request %#v", req)
	r, err := client.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	log.Printf("non decoded response %#v", r)
	err = json.NewDecoder(r.Body).Decode(profile)
	if err != nil {
		return err
	}

	return nil
}