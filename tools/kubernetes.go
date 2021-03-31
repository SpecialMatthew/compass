/*
     Company: 达梦数据库有限公司
  Department: 达梦技术公司/产品研发中心
      Author: 陈磊
      E-mail: chenlei@dameng.com
      Create: 2021/2/19 08:44
     Project: compass
     Package: tools
    Describe: Todo
*/
package tools

import (
	"bytes"
	"io"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/yaml"
	"strings"
)

func Decode(content string) (objects []*runtime.Object, err error) {
	decoder := yaml.NewYAMLOrJSONDecoder(strings.NewReader(content), 4096)
	for {
		extension := runtime.RawExtension{}
		if err := decoder.Decode(&extension); err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		} else {
			if object, err := runtime.Decode(unstructured.UnstructuredJSONScheme, bytes.TrimSpace(extension.Raw)); err != nil {
				return nil, err
			} else {
				objects = append(objects, &object)
			}
		}
	}
	return objects, nil
}
