/*
     Company: 达梦数据库有限公司
  Department: 达梦技术公司/产品研发中心
      Author: 陈磊
      E-mail: chenlei@dameng.com
      Create: 2021/2/1 11:22
     Project: compass
     Package: tools
    Describe: Todo
*/
package tools

import (
	"k8s.io/klog/v2"
	"reflect"
)

func ArrayContains(array []string, element string) bool {
	for _, s := range array {
		if s == element {
			return true
		}
	}
	return false
}

func ArrayRemove(array []string, element string) (result []string) {
	for _, s := range array {
		if s != element {
			result = append(result, s)
		}
	}
	return result
}

func ArrayFilter(array interface{}, path string, value interface{}) (result []interface{}) {
	if obj := reflect.ValueOf(array); obj.Kind() == reflect.Slice {
		for i := 0; i < obj.Len(); i++ {
			data := Snipe(obj.Index(i).Interface(), path)
			dv := reflect.ValueOf(data)
			if dv.Kind() == reflect.ValueOf(value).Kind() {
				switch dv.Kind() {
				case reflect.String:
					if dv.String() == value.(string) {
						result = append(result, obj.Index(i).Interface())
					}
				case reflect.Bool:
					if dv.Bool() == value.(bool) {
						result = append(result, obj.Index(i).Interface())
					}
				case reflect.Int:
					if dv.Int() == value.(int64) {
						result = append(result, obj.Index(i).Interface())
					}
				default:
					klog.V(4).Infof("filter not supported type: %v, %v, %v", array, path, dv.Kind())
				}
			} else {
				klog.V(4).Infof("filter different type: %v, %v, %v, %v", array, path, data, value)
			}
		}
	} else {
		klog.Errorf("filter error type: %v", obj.Kind())
	}
	klog.V(4).Infof("filter result: %v, %v, %v, %v", array, path, value, result)
	return result
}
