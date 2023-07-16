package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

const UnknownLocation = "未知位置"

// GetLocation 根基IP获取地址信息
func GetLocation(ip string, key string) string {
	if ip == "127.0.0.1" || ip == "localhost" {
		return "内部IP"
	}

	resp, err := http.Get("https://restapi.amap.com/v3/ip?ip=" + ip + "&key=" + key)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	s, err := io.ReadAll(resp.Body)
	if err != nil {
		return UnknownLocation
	}

	m := make(map[string]string)

	err = json.Unmarshal(s, &m)
	if err != nil {
		return UnknownLocation
	}

	if m["province"] == "" {
		return UnknownLocation
	}

	return m["province"] + "-" + m["city"]
}
