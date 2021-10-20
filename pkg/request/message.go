package request

import "strings"

// 占位符 :attribute 替换
func replaceAttribute(message, value string) string {
	return strings.Replace(message, ":attribute", value, 1)
}

// 占位符 :input 替换
func replaceInput(message, value string) string {
	return strings.Replace(message, ":input", value, 1)
}

// ReplacerFunc 替换表单验证错误信息方法
type ReplacerFunc func(message, attribute, rule, parameters string) string

// 调起替换方法
func callReplacer(message, attribute, rule, parameters string) string {
	if callback, ok := replacers[rule]; ok {
		message = callback(message, attribute, rule, parameters)
	}
	return message
}

// 定义替换方法集 不可以更改
var replacers = map[string]ReplacerFunc{
	"max": replaceMax,
	"min": replaceMin,
}

// ReplacerMap 返回替换方法集
func ReplacerMap() map[string]ReplacerFunc {
	return replacers
}

// 替换 最大值
func replaceMax(message, attribute, rule, parameters string) string {
	return strings.Replace(message, ":max", parameters, 1)
}

// 替换 最小值
func replaceMin(message, attribute, rule, parameters string) string {
	return strings.Replace(message, ":min", parameters, 1)
}
