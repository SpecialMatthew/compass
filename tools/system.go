/*
     Company: 达梦数据库有限公司
  Department: 达梦技术公司/产品研发中心
      Author: 陈磊
      E-mail: chenlei@dameng.com
      Create: 2021/2/5 08:41
     Project: compass
     Package: tools
    Describe: Todo
*/
package tools

import (
	"k8s.io/klog/v2"
	"os"
	"path/filepath"
	"regexp"
)

func GetEnv(name, def string) string {
	env, found := os.LookupEnv(name)
	if found {
		return env
	}
	return def
}

func Files(path, pattern string) (files []string) {
	regex, err := regexp.Compile(pattern)
	if err != nil {
		klog.Errorf("compile regex error: %v", err)
		return nil
	}
	if err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err == nil && regex.MatchString(info.Name()) {
			files = append(files, path)
		}
		return nil
	}); err != nil {
		klog.Errorf("recursive file error: %v", err)
		return nil
	}
	return files
}
