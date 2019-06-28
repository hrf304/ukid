package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"ukid/router"
	"ukid/sqlgen"
	"ukid/xormtext"
	"xormt"
)

func main() {
	test()
	ginEngine := gin.Default()

	//db.InitDB()
	xormt.Init(xormtext.GetTenants, xormtext.GetTenantId)

	router.Register(ginEngine)

	ginEngine.Run(":8080")
}
func test(){
	jsonStr := `
		{
		    "groupOp":"AND",
		    "rules":[{"field":"major","op":"eq","data":"110"}],
		    "groups":[
				{
				    "groupOp":"OR",
				    "rules":[{"field":"name","op":"cn","data":"ad"},{"field":"login_id","op":"cn","data":"ad"}],
				    "groups":[]
				}
			]
		}
	`
	filter, err := sqlgen.ParseJson(jsonStr)
	if err != nil{
		fmt.Println(err)
		return
	}
	stmt, paramDatas, err := filter.GenWhereStmt("")
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(stmt)
	fmt.Println(paramDatas)
}
