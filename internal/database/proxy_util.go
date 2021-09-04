package database

import (
	"github.com/sakura-rip/sakurabot-cli/internal/utils"
	"net/url"
)

// GetRandomFreeProxy returns array of unused proxies
func GetRandomFreeProxy(count int) []*Proxy {
	var proxies []*Proxy
	result := Client.Limit(count).Where(map[string]interface{}{"is_used": false}).Find(&proxies)
	if result.RowsAffected != int64(count) {
		utils.Logger.Fatal().Msg("cant get free proxy")
	}
	return proxies
}

// ParseProxyUrl parse proxy url to database table Proxy
func ParseProxyUrl(proxy string) *Proxy {
	u, _ := url.Parse(proxy)
	password, _ := u.User.Password()
	return &Proxy{
		URL:      proxy,
		UserId:   u.User.Username(),
		Password: password,
		Host:     u.Host,
		Port:     u.Port(),
	}
}
