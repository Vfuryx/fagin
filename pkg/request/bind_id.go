package request

import "github.com/gin-gonic/gin"

// 获取 uint 类型的 id
type uriUintIdRequest struct {
	Id uint `uri:"id" binding:"required,min=1"`
}

func ShouldBindUriUintID(ctx *gin.Context) (uint, error) {
	var uri uriUintIdRequest
	err := ctx.ShouldBindUri(&uri)
	return uri.Id, err
}

// 获取 uint64 类型的 id
type uriUint64IdRequest struct {
	Id uint64 `uri:"id" binding:"required,min=1"`
}

func ShouldBindUriUint64ID(ctx *gin.Context) (uint64, error) {
	var uri uriUint64IdRequest
	err := ctx.ShouldBindUri(&uri)
	return uri.Id, err
}
