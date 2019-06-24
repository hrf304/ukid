package sqlgen

import (
	"encoding/json"
	"strings"
	"errors"
	"fmt"
)

/**
 * @brief: token
 */
type Token struct {
	Result       int    `json:"-"`
	AccessToken  string `json:"access_token"`
	LoginID      string `json:"login_id"`
	UserID       string `json:"user_id"`
	Domain       string `json:"domain"`
}

/**
 * @brief: 基础条件过
 */
type Filter struct {
	GroupOp string    `json:"groupOp"`	// 组之间操作（or,and)
	Rules   []*Rule   `json:"rules"`	// 过滤规则
	Groups  []*Filter `json:"groups"`	// 子级过滤
}

/**
 * @brief: 过滤规则
 */
type Rule struct {
	Field string `json:"field"`			// 字段名
	Op    string `json:"op"`			// 操作，eq,le......
	Data  string `json:"data"`			// 值
}

/**
 * @brief 根据json转换成Filter对象
 * @param1 jsonStr json格式字符串
 * @return1 Filter对象
 * @return2 错误信息， nil表示没有错误
 */
func ParseJson(jsonStr string) (*Filter, error) {
	filters := &Filter{}
	err := json.Unmarshal([]byte(jsonStr), filters)
	if err != nil {
		return nil, err
	}
	return filters, nil
}

/**
 * @brief: 生成where条件语句
 * @ @param1 tableName: 表名
 * @return1: where条件语句
 * @return2: 条件值列表
 * @return3: u错误信息
 */
func (f *Filter) GenWhereStmt(tableName string) (string, []interface{}, error) {

	whereStatement := ""
	groupWhereStatement := ""
	datas := make([]interface{}, 0)

	isAnd := true
	if strings.ToLower(strings.TrimSpace(f.GroupOp)) == "or" {
		isAnd = false
	}

	if len(f.Rules) > 0 {
		whereStr, ruleDatas := f.GenRuleWhereStmt(tableName, f.Rules[0])
		if whereStr != "" {
			whereStatement = whereStr
			for _, ruleData := range ruleDatas {
				datas = append(datas, ruleData)
			}
		} else {
			return whereStatement, datas, errors.New(fmt.Sprintf("the 0 rule is invalid"))
		}
		for i := 1; i < len(f.Rules); i++ {
			whereStr, ruleDatas = f.GenRuleWhereStmt(tableName, f.Rules[i])
			if whereStr != "" {
				if isAnd {
					whereStatement += " AND " + whereStr
				} else {
					whereStatement += " OR " + whereStr
				}
				for _, ruleData := range ruleDatas {
					datas = append(datas, ruleData)
				}
			} else {
				return whereStatement, datas, errors.New(fmt.Sprintf("the %d rule is invalid", i))
			}
		}
	}

	if len(f.Groups) > 0 {

		whereStr, groupDatas, err := f.Groups[0].GenWhereStmt(tableName)
		if err != nil {
			return whereStatement, datas, err
		}
		if whereStr != "" {
			groupWhereStatement = whereStr
			for _, groupData := range groupDatas {
				datas = append(datas, groupData)
			}
		} else {
			return whereStatement, datas, errors.New("the 0 group is invalid")
		}
		for i := 1; i < len(f.Groups); i++ {
			whereStr, groupDatas, err := f.Groups[i].GenWhereStmt(tableName)
			if err != nil {
				return whereStatement, datas, err
			}
			if whereStr != "" {
				if isAnd {
					groupWhereStatement += " AND " + whereStr
				} else {
					groupWhereStatement += " OR " + whereStr
				}
				for _, groupData := range groupDatas {
					datas = append(datas, groupData)
				}
			} else {
				return whereStatement, datas, errors.New(fmt.Sprintf("the %d group is invalid", i))
			}
		}
		if groupWhereStatement != "" {
			groupWhereStatement = "(" + groupWhereStatement + ")"
		}
	}
	if whereStatement == "" {
		whereStatement = groupWhereStatement
	} else {
		whereStatement = "(" + whereStatement + ")"
		if groupWhereStatement != "" {
			if isAnd {
				whereStatement += " AND " + groupWhereStatement
			} else {
				whereStatement += " OR " + groupWhereStatement
			}
		}
	}

	return whereStatement, datas, nil
}

/**
 * @brief: 根据rule生成where语句
 * @param1 tableName: 表名
 * @param2 field: 过滤规则
 * @return1: where条件语句
 * @return2: 条件值列表
 * @return3: u错误信息
 */
func (f *Filter) GenRuleWhereStmt(tableName string, field *Rule) (string, []interface{}) {
	switch GetRuleDataType(field.Field) {
	case _RULE_DATA_TYPE_FIELD_NAME:
		return RuleToWhereStmt(tableName, field)
	default:
		return "", make([]interface{}, 0)
	}
}
