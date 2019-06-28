package xormtext

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"xormt"
)

var i = 0

func GetTenants()[] *xormt.TenantDBInfo{
	tdbs := make([]*xormt.TenantDBInfo, 0)

	item := xormt.TenantDBInfo{}
	item.Tid = "default"
	item.Name = "默认数据库"
	item.ConnStr = "root:aadmin123@tcp(localhost:3306)/test?charset=utf8&loc=Local"
	item.DriverName = "mysql"
	tdbs = append(tdbs, &item)

	item1 := xormt.TenantDBInfo{}
	item1.Tid = "tenant1"
	item1.Name = "租户1"
	item1.ConnStr = "root:aadmin123@tcp(localhost:3306)/test?charset=utf8&loc=Local"
	item1.DriverName = "mysql"
	tdbs = append(tdbs, &item1)

	fmt.Println(tdbs)

	return tdbs
}

func GetTenantId(ctx *gin.Context)string{
	i++
	if i % 2 == 0{
		return "default"
	}else{
		return "tenant1"
	}
}
