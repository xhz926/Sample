package config

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"sync"
)

var once sync.Once
var baseConf *baseConfig

/**

 *@描述 基础配置文件类

 *@创建人  吉喆

 *@创建时间  2018/08/06 22:01

 *@备注

 */
type baseConfig struct {
	//设置运行环境mod
	runMod string
}

/**

 *@描述 基础配置文件初始化

 *@创建人  吉喆

 *@创建时间  2018/08/06 22:01

 *@备注

 */
func init() {
	once.Do(func() {
		baseConf = &baseConfig{}
		baseConf.runMod = beego.AppConfig.String(RUN_MODE)
		logs.Info("运行模式为:", baseConf.runMod)
		if 0 == len(baseConf.runMod) {
			logs.Info("运行参数未指定,默认dev环境")
			baseConf.runMod = "dev"
		}
	})
}
