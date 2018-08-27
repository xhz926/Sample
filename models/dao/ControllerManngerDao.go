package dao

import "github.com/astaxie/beego/orm"

/**

 *@描述 返回数据管理表Dao

 *@创建人  吉喆

 *@创建时间  2018/08/24 17:44

 *@备注

 */
type ControllerManngerDao struct {
	//id编号
	Id int64
	//五要素标识:0失败 1:成功 2:超时
	FiveElementsFlag int
	//联网核查标识:0失败 1:成功 2:超时
	NetworkVerificationFlag int
	//开户标识:0失败 1:成功 2:超时
	OpenAccountFlag int
	//开通网银标识:0失败 1:成功 2:超时
	OnlineBankingFlag int
}

func (this *ControllerManngerDao) TableName() string {
	return CONTROLLER_MANNGER
}

func init() {
	orm.RegisterModel(new(ControllerManngerDao))
}
