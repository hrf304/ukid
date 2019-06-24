package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"ukid/entity"
	"ukid/util"
	"ukid/db"
)

type UserController struct {
}

func (c *UserController) Get(ctx *gin.Context) {
	var Resp struct {
		Code int         `json:"code"`
		Data interface{} `json:"data"`
		Msg  string      `json:"msg"`
	}
	Resp.Code = 200
	Resp.Msg = "hello world"
	Resp.Data = ctx.Param("id")

	ctx.JSON(http.StatusOK, Resp)
}

func (c *UserController) GetUsers(ctx *gin.Context) {
	queryParam := entity.QueryParam{}
	err := ctx.ShouldBindQuery(&queryParam)
	if err != nil {
		fmt.Println(err)
	} else {
		queryParam.Table = "sys_user"
		rowsSlice := make([]map[string]interface{}, 0)
		result, err := util.Query(db.DB, queryParam, &rowsSlice)
		if err != nil {
			ctx.JSON(500, &entity.Resp{500, err.Error(), nil})
		} else {
			ctx.JSON(200, &entity.Resp{200, "", result})
		}
	}
}
