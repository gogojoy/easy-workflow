package engine

import (
	"easy-workflow/pkg/dao"
	"encoding/json"
	"strings"
)

//判断传入字符是否是变量(以$开头)
func IsVariable(Key string) bool {
	if strings.HasPrefix(Key, "$") {
		return true
	}
	return false
}

//去掉变量前缀"$"
func RemovePrefix(variable string) string {
	return strings.Replace(variable, "$", "", -1)
}

//从proc_inst_variable表中查找变量，若有则返回,若无则返回false
func SetVariable(ProcessInstanceID int, variable string) (string, bool, error) {
	Key := RemovePrefix(variable)
	type result struct {
		Value string
	}
	var r result
	if _, err := dao.ExecSQL("SELECT `value` FROM `proc_inst_variable` "+
		"WHERE `proc_inst_id`=? AND `key`=? LIMIT 1", &r, ProcessInstanceID, Key); err == nil {

		exists := false
		if r.Value != "" {
			exists = true
		}

		return r.Value, exists, nil

	} else {
		return "", false, err
	}
}

//func FindAllVariables() []string {
//
//}

//将变量map生成kv对形式的json字符串，以便存入数据库
func VariablesMap2Json(Variables map[string]string) (string, error) {

	type kv struct {
		Key   string
		Value string
	}
	var kvs []kv
	for k, v := range Variables {
		kvs = append(kvs, kv{Key: k, Value: v})
	}
	j, err := json.Marshal(kvs)
	if err != nil {
		return "", err
	}
	return string(j), nil
}