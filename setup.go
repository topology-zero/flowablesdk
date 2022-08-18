package flowablesdk

type Config struct {
	Url      string // 请求地址
	Username string // 用户名
	Password string // 密码
	TenantId int    // 多租户
}

var Configs Config

func Setup(url, username, password string, tenantId ...int) {
	if url[len(url)-1] == '/' {
		url = url[:len(url)-1]
	}
	Configs = Config{
		Url:      url,
		Username: username,
		Password: password,
	}

	if len(tenantId) > 0 {
		Configs.TenantId = tenantId[0]
	}
}
