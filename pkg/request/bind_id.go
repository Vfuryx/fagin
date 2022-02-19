package request

import "github.com/gin-gonic/gin"

// 获取 uint 类型的 id
type uriUintIDRequest struct {
	ID uint `uri:"id" binding:"required,min=1"`
}

// ShouldBindURIUintID 绑定 uri uintID
func ShouldBindURIUintID(ctx *gin.Context) (uint, error) {
	var uri uriUintIDRequest
	err := ctx.ShouldBindUri(&uri)

	return uri.ID, err
}

// 获取 uint64 类型的 id
type uriUint64IdRequest struct {
	ID uint64 `uri:"id" binding:"required,min=1"`
}

// ShouldBindURIUint64ID 绑定 uri uint64ID
func ShouldBindURIUint64ID(ctx *gin.Context) (uint64, error) {
	var uri uriUint64IdRequest
	err := ctx.ShouldBindUri(&uri)

	return uri.ID, err
}
