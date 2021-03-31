/*
     Company: 达梦数据库有限公司
  Department: 达梦技术公司/产品研发中心
      Author: 陈磊
      E-mail: chenlei@dameng.com
      Create: 2021/2/6 13:52
     Project: compass
     Package: tools
    Describe: Todo
*/
package tools

import (
	"crypto/sha1"
	"github.com/fatih/structs"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/json"
	"k8s.io/klog/v2"
	"strings"
)

func Snipe(object interface{}, path string) interface{} {
	klog.V(4).Infof("snipe: %v, %v", object, path)
	data, found, err := unstructured.NestedFieldNoCopy(structs.Map(object), strings.Split(path, ".")...)
	klog.V(4).Infof("snipe: %v, %v, %v", data, found, err)
	if err != nil {
		klog.Errorf("snipe error: %v, %v", path, err)
		return nil
	}
	if found {
		return data
	}
	return nil
}

func Hitch(object interface{}, path string, value interface{}) (result map[string]interface{}) {
	if object == nil {
		result = make(map[string]interface{})
	} else {
		result = structs.Map(object)
	}
	if i := strings.Index(path, "."); i == -1 {
		result[path] = value
	} else {
		result[path[:i]] = Hitch(nil, path[i+1:], value)
	}
	return result
}

func ToJson(object interface{}) string {
	str, err := json.Marshal(object)
	if err != nil {
		klog.Errorf("to json error: %v, %v", err)
		return ""
	}
	return string(str)
}

func Sha1sum(object interface{}) string {
	json, err := json.Marshal(object)
	if err != nil {
		klog.Errorf("compute SHA1 hashes error: %v", err)
		return ""
	}
	return string(sha1.New().Sum(json))
}
