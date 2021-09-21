package utils

// BToI 布尔转数字
func BToI(b bool) int {
	if b {
		return 1
	}
	return 0
}

// IToB 数字转布尔
func IToB(b int) bool {
	if b != 0 {
		return true
	}
	return false
}
