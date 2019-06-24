package util

import (
	"github.com/go-xorm/xorm"
	"strings"
	"ukid/entity"
	"fmt"
)

func Query(db *xorm.Engine, param entity.QueryParam, rowsSlicePtr interface{})(*entity.QueryResult, error){
	pageResult := entity.QueryResult{}

	// 结果session
	dbSession := db.Table(param.Table)
	if len(param.Selects) > 0{
		dbSession = dbSession.Select(param.Table + ".*")
	}else{
		dbSession = dbSession.Select(strings.Join(param.Selects, ","))
	}

	// 总行数session
	totalRows := 0
	totalRowsSession := db.Table(param.Table)
	totalRowsSession.Select("count(*) as totalPages")

	// 默认join之影响结果session
	for _, join := range param.JoinParams {
		dbSession = dbSession.Join(join.JoinOperator, join.TabelName, join.Condition)
	}

	// 获取条件语句
	whereStr, datas, err := param.Filter.GenWhereStmt("")
	if err != nil {
		return nil, err
	}
	fmt.Println("---------------------->", whereStr)

	dbSession = dbSession.Where(whereStr, datas...)
	totalRowsSession = totalRowsSession.Where(whereStr, datas...)

	// 排序，只影响结果session
	sorts := strings.Split(param.Sort, ",")
	for _, order := range sorts {
		order = strings.TrimSpace(order)
		if len(order) == 0 {
			continue
		}
		dbSession = dbSession.OrderBy(order)
	}

	// 获取结果
	err = dbSession.Limit(param.Rows, (param.Page-1)*param.Rows).Find(rowsSlicePtr)
	if err != nil{
		return nil, err
	}
	// 获取行数
	_, err = totalRowsSession.Get(&totalRows)
	if err != nil{
		return nil, err
	}

	fmt.Println(totalRows)

	pageResult.Result = rowsSlicePtr
	pageResult.TotalRows = totalRows
	pageResult.Page = param.Page
	pageResult.TotalPages = totalRows / param.Rows
	if totalRows % param.Rows != 0{
		pageResult.TotalRows += 1
	}

	return &pageResult, nil
}
