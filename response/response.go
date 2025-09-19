package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 封装公共response给api层调用
func Failed(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, err)
}

func Success(ctx *gin.Context, ins any) {
	ctx.JSON(http.StatusOK, ins)
}
