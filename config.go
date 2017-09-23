package gopher

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"os"
	"runtime"

	"github.com/bradrydzewski/go.auth"
)

type ConfigStruct struct {
	Host                       string `json:"host"`
	Port                       int    `json:"port"`
	DB                         string `json:"db"`
	CookieSecret               string `json:"cookie_secret"`
	SendMailPath               string `json:"sendmail_path"`
	SmtpUsername               string `json:"smtp_username"`
	SmtpPassword               string `json:"smtp_password"`
	SmtpHost                   string `json:"smtp_host"`
	SmtpAddr                   string `json:"smtp_addr"`
	FromEmail                  string `json:"from_email"`
	Superusers                 string `json:"superusers"`
	TimeZoneOffset             int64  `json:"time_zone_offset"`
	AnalyticsFile              string `json:"analytics_file"`
	StaticFileVersion          int    `json:"static_file_version"`
	GoGetPath                  string `json:"go_get_path"`
	PackagesDownloadPath       string `json:"packages_download_path"`
	PublicSalt                 string `json:"public_salt"`
	CookieSecure               bool   `json:"cookie_secure"`
	GithubClientId             string `json:"github_auth_client_id"`
	GithubClientSecret         string `json:"github_auth_client_secret"`
	GithubLoginRedirect        string `json:"github_login_redirect"`
	GithubLoginSuccessRedirect string `json:"github_login_success_redirect"`
	DeferPanicApiKey           string `json:"deferpanic_api_key"`
	GtCaptchaId                string `json:"gt_captcha_id"`
	GtPrivateKey               string `json:"gt_private_key"`
	GoDownloadPath             string `json:"go_download_path"`
	LiteIDEDownloadPath        string `json:"liteide_download_path"`
	ImagePath                  string `json:"image_path"`
}

var (
	Config        ConfigStruct
	analyticsCode template.HTML // 网站统计分析代码
	shareCode     template.HTML // 分享代码
	goVersion     = runtime.Version()
)

func parseJsonFile(path string, v interface{}) {
	file, err := os.Open(path)
	if err != nil {
		logger.Fatal("配置文件读取失败:", err)
	}
	defer file.Close()
	dec := json.NewDecoder(file)
	err = dec.Decode(v)
	if err != nil {
		logger.Fatal("配置文件解析失败:", err)
	}
}

func getDefaultCode(path string) (code template.HTML) {
	if path != "" {
		content, err := ioutil.ReadFile(path)
		if err != nil {
			logger.Fatal("文件 " + path + " 没有找到")
		}
		code = template.HTML(string(content))
	}
	return
}

func configGithubAuth() {
	if Config.GithubClientId == "" || Config.GithubClientSecret == "" {
		logger.Fatal("没有配置github应用的参数")
	}
	auth.Config.CookieSecret = []byte(Config.CookieSecret)
	auth.Config.LoginRedirect = Config.GithubLoginRedirect
	auth.Config.LoginSuccessRedirect = Config.GithubLoginSuccessRedirect
	auth.Config.CookieSecure = Config.CookieSecure
	if !auth.Config.CookieSecure {
		logger.Println("注意,cookie_secure设置为false,只能在本地环境下测试")
	}
	githubHandler = auth.Github(Config.GithubClientId, Config.GithubClientSecret, "user")
}
