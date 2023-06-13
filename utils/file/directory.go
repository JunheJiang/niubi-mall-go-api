package file

import (
	"go.uber.org/zap"
	"niubi-mall/global"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, err
	}
	return false, err
}

// CreateDir --- err限定修饰返回名
func CreateDir(dirs ...string) (err error) {
	//Java快速转go --- for range 频繁使用
	for _, path := range dirs {
		exist, err := PathExists(path)
		if err != nil {
			return err
		}
		if !exist {
			global.GVA_LOG.Debug("create directory:" + path)
			//Java快速转go--- if 赋值操作后判断err 频繁使用
			if err := os.MkdirAll(path, os.ModePerm); err != nil {
				global.GVA_LOG.Error("create directory:"+path, zap.Any("error", err))
				return err
			}
		}
	}
	return err
}
