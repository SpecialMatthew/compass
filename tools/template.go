/*
     Company: 达梦数据库有限公司
  Department: 达梦技术公司/产品研发中心
      Author: 陈磊
      E-mail: chenlei@dameng.com
      Create: 2021/2/1 14:09
     Project: compass
     Package: tools
    Describe: Todo
*/
package tools

import (
	"bytes"
	"encoding/hex"
	"github.com/Masterminds/sprig/v3"
	pet "github.com/dustinkirkland/golang-petname"
	"k8s.io/klog/v2"
	"sigs.k8s.io/yaml"
	"strings"
	"text/template"
)

func ParseTemplate(name string, parameters interface{}) (string, error) {
	templates, err := template.New("default").Funcs(sprig.TxtFuncMap()).Funcs(buildFunctionMap()).ParseFiles(Files(GetEnv("TEMPLATES_PATH", "/etc/dmcca/templates"), "\\.gotmpl$")...)
	if err != nil {
		klog.Errorf("parse template error: %v", err)
		return "", err
	}
	buffer := new(bytes.Buffer)
	if err := templates.ExecuteTemplate(buffer, name, parameters); err != nil {
		klog.Errorf("template execute error: %v", err)
		return "", err
	}
	return buffer.String(), nil
}

func buildFunctionMap() template.FuncMap {
	return template.FuncMap{
		"toYaml": func(object interface{}) string {
			if b, err := yaml.Marshal(object); err != nil {
				return ""
			} else {
				return strings.TrimSpace(string(b))
			}
		},
		"toJson": ToJson,
		"snipe":  Snipe,
		"hitch":  Hitch,
		"filter": ArrayFilter,
		"include": func(name string, parameters interface{}) (result string) {
			result, err := ParseTemplate(name, parameters)
			if err != nil {
				klog.Errorf("template include error: %v", err)
				return ""
			}
			return result
		},
		"generateName": func() string {
			return pet.Generate(3, "-")
		},
		"hexenc": func(src string) string {
			return hex.EncodeToString([]byte(src))
		},
		"hexdec": func(src string) string {
			if dst, err := hex.DecodeString(src); err != nil {
				klog.Errorf("hex decode error: %v", err)
				return src
			} else {
				return string(dst)
			}
		},
	}
}
