package entity

import "ukid/sqlgen"

/**************查询相关*******************/
type QueryResult struct {
	Page   int         `json:"page"`   //页码，从1开始
	TotalRows   int         `json:"totalRows"`   // 页行数
	TotalPages  int         `json:"totalPages"`  // 总页数
	Result interface{} `json:"result"` // 结果
}

type QueryParam struct {
	Rows       int           `form:"rows"json:"rows"`     // 页行数
	Page       int           `form:"page"json:"page"`     // 页码，从1开始
	Sort       string        `form:"sort"json:"sort"`     // 排序，类似与name,age desc
	Filter     sqlgen.Filter `form:"filter"json:"filter"` // 过滤条件
	Table      string        `form:"-"json:"-"`           // 表名
	Selects    []string      `form:"-"json:"-"`           // select 字段列表
	JoinParams []JoinParam   `form:"-"json:"-"`           // 连接表列表
}

type JoinParam struct {
	JoinOperator string //连接方式
	TabelName    string //连接表名
	Condition    string //连接条件
}

/*****************返回结果相关***************/
type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

