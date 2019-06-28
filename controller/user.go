package controller

import (
	"fmt"
	"net/http"
	"ukid/entity"
	"ukid/util"
	"xormt"
)

type UserController struct {
}

var i = 0

func (c *UserController) Get(ctx *xormt.MultiTenantContext) {
	i++

	user := &entity.User{}
	user.Id = util.UUID()
	user.Name = fmt.Sprintf("huangrf%d", i)
	user.LoginId = fmt.Sprintf("huangrf%d", i)
	user.Major = fmt.Sprintf("major%d", i)

	_, err := ctx.DB.InsertOne(user)
	if err != nil{
		ctx.JSON(500, &entity.Resp{500, err.Error(), nil})
	}else{
		ctx.JSON(http.StatusOK, &entity.Resp{http.StatusOK, "", user.Id})
	}
}

func (c *UserController) GetUsers(ctx *xormt.MultiTenantContext) {
	queryParam := entity.QueryParam{}
	err := ctx.ShouldBindQuery(&queryParam)
	if err != nil {
		fmt.Println(err)
	} else {
		queryParam.Table = "sys_user"
		rowsSlice := make([]map[string]interface{}, 0)
		result, err := util.Query(ctx.DB, queryParam, &rowsSlice)
		if err != nil {
			ctx.JSON(500, &entity.Resp{500, err.Error(), nil})
		} else {
			ctx.JSON(200, &entity.Resp{200, "", result})
		}
	}
}
