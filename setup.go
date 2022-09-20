package flowablesdk

type Config struct {
	Url      string // 请求地址
	Username string // 用户名
	Password string // 密码
	TenantId int    // 多租户
}

var Configs Config

func Setup(c Config) {
	if c.Url[len(c.Url)-1] == '/' {
		c.Url = c.Url[:len(c.Url)-1]
	}
	Configs.Url = c.Url
	Configs.Username = c.Username
	Configs.Password = c.Password
	Configs.TenantId = c.TenantId
}
