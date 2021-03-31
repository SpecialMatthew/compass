/*
     Company: 达梦数据库有限公司
  Department: 达梦技术公司/产品研发中心
      Author: 陈磊
      E-mail: chenlei@dameng.com
      Create: 2021/2/6 15:52
     Project: compass
     Package: tools
    Describe: Todo
*/
package tools

func IfNil(object, reserve interface{}) interface{} {
	if object == nil {
		return reserve
	}
	return object
}
