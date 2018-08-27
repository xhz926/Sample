package config

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

/**

 *@描述 日志控制类

 *@创建人  吉喆

 *@创建时间  2018/08/05 21:18

 *@备注

 */
type logConfig struct {
	//日志路径
	logPath string
}

/**

 *@描述 go的初始化函数

 *@创建人  吉喆

 *@创建时间  2018/08/05 21:21

 *@备注

 */
func init() {
	logs.Info("日志配置文件初始化init start")
	logConf := &logConfig{}
	logConf.init()
	logs.Info("日志配置文件初始化init end")
}

/**

 *@描述 日志配置文件初始化函数

 *@创建人  吉喆

 *@创建时间  2018/08/07 7:26

 *@备注

 */
func (this *logConfig) init() {
	this.logPath = beego.AppConfig.String(baseConf.runMod + "::" + LOG_PATH)
	logs.Info("输出日志路径为:", this.logPath)
	//设置日志初始化路径
	beego.SetLogger(logs.AdapterFile, `{"filename":"`+this.logPath+`"}`)
	logasync, err := beego.AppConfig.Bool(baseConf.runMod + "::" + LOG_ASYNC)
	if nil != err {
		logs.Error("日志异步输出初始化错误,err=" + err.Error())
		panic("日志异步输出初始化错误,err=" + err.Error())
	}
	if logasync {
		logs.Async(1e3)
	}
}
