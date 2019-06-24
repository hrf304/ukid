package sqlgen

import (
	"fmt"
	"strings"
)

/**
 * @brief: rule 转换成 where 语句
 * @param1 tableName: 表名
 * @param2 rule: 规则
 * @return1 where语句
 * @return2 where语句对应的参数值列表
 */
func RuleToWhereStmt(tableName string, rule *Rule) (string, []interface{}) {
	datas := make([]interface{}, 0)
	if rule == nil {
		return "", datas
	}
	fData := FormatData(rule.Op, rule.Data)
	switch fData.(type) {
	case string:
		{
			if len(fData.(string)) >= 0 {
				datas = append(datas, fData)
			}
		}
	case []string:
		{
			fDataArr := fData.([]string)
			for _, arr := range fDataArr {
				if len(arr) > 0 {
					datas = append(datas, arr)
				}
			}
		}
	}

	whereStr := FormatOp(tableName, rule.Op, rule.Field, rule.Data)

	if len(whereStr) > 0 {
		return whereStr, datas
	} else {
		return "", datas
	}
}

func FormatOp(table, op, field, data string) string {
	if strings.Contains(field, ".") {
		// the field is "talbe.field"
		switch strings.ToLower(op) {
		case "eq":
			{
				return fmt.Sprintf(" %s = ? ", field)
			}
		case "ne":
			{
				return fmt.Sprintf(" %s <> ? ", field)
			}
		case "lt":
			{
				return fmt.Sprintf(" %s < ? ", field)
			}
		case "le":
			{
				return fmt.Sprintf(" %s <= ? ", field)
			}
		case "gt":
			{
				return fmt.Sprintf(" %s > ? ", field)
			}
		case "ge":
			{
				return fmt.Sprintf(" %s >= ? ", field)
			}
		case "bw":
			{
				return fmt.Sprintf(" %s LIKE ? ", field)
			}
		case "bn":
			{
				return fmt.Sprintf(" %s NOT LIKE ? ", field)
			}
		case "ew":
			{
				return fmt.Sprintf(" %s LIKE ? ", field)
			}
		case "en":
			{
				return fmt.Sprintf(" %s NOT LIKE ? ", field)
			}
		case "cn":
			{
				return fmt.Sprintf(" %s LIKE ? ", field)
			}
		case "nc":
			{
				return fmt.Sprintf(" %s NOT LIKE ? ", field)
			}
		case "nu":
			{
				return fmt.Sprintf(" %s IS NULL ", field)
			}
		case "nn":
			{
				return fmt.Sprintf(" %s IS NOT NULL ", field)
			}
		case "date-cn":
			{
				return fmt.Sprintf(" %s >= ? and %s <= ? ", field, field)
			}
		case "in":
			{
				dataArr := strings.Split(data, ",")
				params := ""
				for _, arr := range dataArr {
					if len(arr) > 0 {
						params += "?,"
					}
				}
				params = strings.TrimRight(params, ",")
				if len(params) > 0 {
					return fmt.Sprintf(" %s IN (%s)", field, params)
				}
			}
		case "ni":
			{
				dataArr := strings.Split(data, ",")
				params := ""
				for _, arr := range dataArr {
					if len(arr) > 0 {
						params += "?,"
					}
				}
				params = strings.TrimRight(params, ",")
				if len(params) > 0 {
					return fmt.Sprintf(" %s NOT IN (%s)", field, params)
				}
			}
		}
	} else {
		if table == "" {
			switch strings.ToLower(op) {
			case "eq":
				{
					return fmt.Sprintf(" `%s` = ? ", field)
				}
			case "ne":
				{
					return fmt.Sprintf(" `%s` <> ? ", field)
				}
			case "lt":
				{
					return fmt.Sprintf(" `%s` < ? ", field)
				}
			case "le":
				{
					return fmt.Sprintf(" `%s` <= ? ", field)
				}
			case "gt":
				{
					return fmt.Sprintf(" `%s` > ? ", field)
				}
			case "ge":
				{
					return fmt.Sprintf(" `%s` >= ? ", field)
				}
			case "bw":
				{
					return fmt.Sprintf(" `%s` LIKE ? ", field)
				}
			case "bn":
				{
					return fmt.Sprintf(" `%s` NOT LIKE ? ", field)
				}
			case "ew":
				{
					return fmt.Sprintf(" `%s` LIKE ? ", field)
				}
			case "en":
				{
					return fmt.Sprintf(" `%s` NOT LIKE ? ", field)
				}
			case "cn":
				{
					return fmt.Sprintf(" `%s` LIKE ? ", field)
				}
			case "nc":
				{
					return fmt.Sprintf(" `%s` NOT LIKE ? ", field)
				}
			case "nu":
				{
					return fmt.Sprintf(" `%s` IS NULL ", field)
				}
			case "nn":
				{
					return fmt.Sprintf(" `%s` IS NOT NULL ", field)
				}
			case "date-cn":
				{
					return fmt.Sprintf(" %s >= ? and %s <= ? ", field, field)
				}
			case "in":
				{
					dataArr := strings.Split(data, ",")
					params := ""
					for _, arr := range dataArr {
						if len(arr) > 0 {
							params += "?,"
						}
					}
					params = strings.TrimRight(params, ",")
					if len(params) > 0 {
						return fmt.Sprintf(" `%s` IN (%s)", field, params)
					}
				}
			case "ni":
				{
					dataArr := strings.Split(data, ",")
					params := ""
					for _, arr := range dataArr {
						if len(arr) > 0 {
							params += "?,"
						}
					}
					params = strings.TrimRight(params, ",")
					if len(params) > 0 {
						return fmt.Sprintf(" `%s` NOT IN (%s)", field, params)
					}
				}
			}
		} else {
			switch strings.ToLower(op) {
			case "eq":
				{
					return fmt.Sprintf(" %s.%s = ? ", table, field)
				}
			case "ne":
				{
					return fmt.Sprintf(" %s.%s <> ? ", table, field)
				}
			case "lt":
				{
					return fmt.Sprintf(" %s.%s < ? ", table, field)
				}
			case "le":
				{
					return fmt.Sprintf(" %s.%s <= ? ", table, field)
				}
			case "gt":
				{
					return fmt.Sprintf(" %s.%s > ? ", table, field)
				}
			case "ge":
				{
					return fmt.Sprintf(" %s.%s >= ? ", table, field)
				}
			case "bw":
				{
					return fmt.Sprintf(" %s.%s LIKE ? ", table, field)
				}
			case "bn":
				{
					return fmt.Sprintf(" %s.%s NOT LIKE ? ", table, field)
				}
			case "ew":
				{
					return fmt.Sprintf(" %s.%s LIKE ? ", table, field)
				}
			case "en":
				{
					return fmt.Sprintf(" %s.%s NOT LIKE ? ", table, field)
				}
			case "cn":
				{
					return fmt.Sprintf(" %s.%s LIKE ? ", table, field)
				}
			case "nc":
				{
					return fmt.Sprintf(" %s.%s NOT LIKE ? ", table, field)
				}
			case "nu":
				{
					return fmt.Sprintf(" %s.%s IS NULL ", table, field)
				}
			case "nn":
				{
					return fmt.Sprintf(" %s.%s IS NOT NULL ", table, field)
				}
			case "date-cn":
				{
					return fmt.Sprintf(" %s.%s >= ? and %s.%s <= ? ", table, field, table, field)
				}
			case "in":
				{
					dataArr := strings.Split(data, ",")
					params := ""
					for _, arr := range dataArr {
						if len(arr) > 0 {
							params += "?,"
						}
					}
					params = strings.TrimRight(params, ",")
					if len(params) > 0 {
						return fmt.Sprintf(" %s.%s IN (%s)", table, field, params)
					}
				}
			case "ni":
				{
					dataArr := strings.Split(data, ",")
					params := ""
					for _, arr := range dataArr {
						if len(arr) > 0 {
							params += "?,"
						}
					}
					params = strings.TrimRight(params, ",")
					if len(params) > 0 {
						return fmt.Sprintf(" %s.%s NOT IN (%s)", table, field, params)
					}
				}
			}
		}
	}
	return ""
}

func FormatData(op, data string) interface{} {
	switch strings.ToLower(op) {
	case "eq", "ne", "lt", "le", "gt", "ge":
		{
			return data
		}
	case "bw":
		{
			return fmt.Sprintf("%s%%", data)
		}
	case "bn":
		{
			return fmt.Sprintf("%s%%", data)
		}
	case "ew":
		{
			return fmt.Sprintf("%%%s", data)
		}
	case "en":
		{
			return fmt.Sprintf("%%%s", data)
		}
	case "cn":
		{
			return fmt.Sprintf("%%%s%%", data)
		}
	case "nc":
		{
			return fmt.Sprintf("%%%s%%", data)
		}
	case "in":
		{
			return strings.Split(data, ",")
		}
	case "ni":
		{
			return strings.Split(data, ",")
		}
	case "date-cn":
		{
			data = strings.TrimSpace(data)
			if data == "" {
				return []string{"", ""}
			} else {
				data = data[0:10]
				return []string{data + " 00:00:00", data + " 23:59:59"}
			}
		}
	}
	return ""
}

/**
 * @brief: 获取Rule值的类型，见常量_RULE_DATA_TYPE_**
 * @param1 data: 值
 * @return data的值类型
 * 			如果值以$开始, 表示当前用户信息相关值，返回user_value
 * 			如果值以#开始, 表示返回与对应字段值相关，返回field_value
 * 			如果值以普通字符开始, 表示当前为值，返回field_name
 */
func GetRuleDataType(data string) string {
	data = strings.TrimSpace(data)
	if strings.HasPrefix(data, "$") {
		return _RULE_DATA_TYPE_USER_VALUE
	} else if strings.HasPrefix(data, "#") {
		return _RULE_DATA_TYPE_FIELD_VALUE
	} else {
		return _RULE_DATA_TYPE_FIELD_NAME
	}
}