package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// GetLocation 根基IP获取地址信息
func GetLocation(ip string) string {
	if ip == "127.0.0.1" || ip == "localhost" {
		return "内部IP"
	}
	resp, err := http.Get("https://restapi.amap.com/v3/ip?ip=" + ip + "&key=3fabc36c20379fbb9300c79b19d5d05e")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	s, err := io.ReadAll(resp.Body)
	m := make(map[string]string)
	err = json.Unmarshal(s, &m)
	if err != nil {
		return "未知位置"
	}
	if m["province"] == "" {
		return "未知位置"
	}
	return m["province"] + "-" + m["city"]
}
