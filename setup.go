package flowablesdk

import "strings"

type Config struct {
	Url           string // 请求地址
	Username      string // 用户名
	Password      string // 密码
	ProcessPrefix string // 流程API前缀
	FormPrefix    string // 表单API前缀
}

var Configs Config

func Setup(c Config) {
	if len(c.Url) == 0 {
		c.Url = "http://127.0.0.1:8080"
	} else {
		c.Url = strings.Trim(c.Url, "/")
	}

	if len(c.ProcessPrefix) == 0 {
		c.ProcessPrefix = "/service"
	} else {
		c.ProcessPrefix = strings.Trim(c.ProcessPrefix, "/")
		c.ProcessPrefix = "/" + c.ProcessPrefix
	}

	if len(c.FormPrefix) == 0 {
		c.FormPrefix = "/form-api"
	} else {
		c.FormPrefix = strings.Trim(c.FormPrefix, "/")
		c.FormPrefix = "/" + c.FormPrefix
	}

	Configs.Url = c.Url
	Configs.Username = c.Username
	Configs.Password = c.Password
	Configs.ProcessPrefix = c.ProcessPrefix
	Configs.FormPrefix = c.FormPrefix
}
