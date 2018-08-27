package config

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"qaPlatform/models/Util"
	"qaPlatform/models/mapper"
)

/**

 *@描述 数据库配置

 *@创建人  吉喆

 *@创建时间  2018/08/06 11:07

 *@备注

 */
type dbConfig struct {
	//数据库驱动
	dbDriver string
	//数据库IP
	dbIP string
	//数据库端口
	dbPort string
	//数据库用户名
	dbUserName string
	//数据库密码
	dbPassword string
	//连接的数据库
	dbDatabase string
}

/**

 *@描述 数据库配置初始化

 *@创建人  吉喆

 *@创建时间  2018/08/06 11:07

 *@备注

 */
func init() {
	//读取配置文件信息
	logs.Info("数据库配置文件初始化init start")
	dbConf := &dbConfig{}
	dbConf.Init()
	dbConf.RegisterDb()
	mapper.DbDriver = dbConf.dbDriver
	logs.Info("数据库配置文件初始化init end")
}

/**

 *@描述 数据库配置参数初始化

 *@创建人  吉喆

 *@创建时间  2018/08/07 1:33

 *@备注

 */
func (this *dbConfig) Init() {
	this.dbDriver = beego.AppConfig.String(baseConf.runMod + "::" + DB_DRIVER)
	this.dbIP = beego.AppConfig.String(baseConf.runMod + "::" + DB_IPADDR)
	this.dbPort = beego.AppConfig.String(baseConf.runMod + "::" + DB_PORT)
	this.dbUserName = beego.AppConfig.String(baseConf.runMod + "::" + DB_USER_NAME)
	this.dbPassword = beego.AppConfig.String(baseConf.runMod + "::" + DB_PASSWORD)
	this.dbDatabase = beego.AppConfig.String(baseConf.runMod + "::" + DB_DATABASE)
}

/**

 *@描述 数据库注册

 *@创建人  吉喆

 *@创建时间  2018/08/07 1:34

 *@备注

 */
func (this *dbConfig) RegisterDb() {
	err := orm.RegisterDriver(this.dbDriver, orm.DRMySQL)
	if nil != err {
		logs.Error("注册数据库失败,err=" + err.Error())
		panic("注册数据库失败,err=" + err.Error())
	}

	strUtil := Util.CreateStrUtil()
	nameArray := strUtil.Split(this.dbDatabase, ",")

	for k, v := range nameArray {
		logs.Info(v + "注册数据库--start")
		if 0 == k {
			dataSource := this.dbUserName + ":" + this.dbPassword + "@tcp(" + this.dbIP + ":" + this.dbPort + ")/" + v + "?charset=utf8&allowNativePasswords=true"
			//注册连接数据库
			err = orm.RegisterDataBase("default", this.dbDriver, dataSource)
			if nil != err {
				logs.Error("连接数据库失败,err=" + err.Error())
				panic("连接数据库失败,err=" + err.Error())
			}
		}
		dataSource := this.dbUserName + ":" + this.dbPassword + "@tcp(" + this.dbIP + ":" + this.dbPort + ")/" + v + "?charset=utf8&allowNativePasswords=true"
		//注册连接数据库
		err = orm.RegisterDataBase(v, this.dbDriver, dataSource)
		if nil != err {
			logs.Error("连接数据库失败,err=" + err.Error())
			panic("连接数据库失败,err=" + err.Error())
		}
		logs.Info(v + "注册数据库--end")
	}
}
