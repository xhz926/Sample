package server

import (
	"github.com/astaxie/beego/logs"
	"hxyhBankTest/models/dto"
)

/**

 *@描述 账户状态查询、资产查询服务

 *@创建人  吉喆

 *@创建时间  2018/08/23 17:52

 *@备注

 */
type queryAccountService struct {
	ReqData *dto.QueryAccountReqDto
}

/**

 *@描述 创建账户状态查询、资产查询服务

 *@创建人  吉喆

 *@创建时间  2018/08/23 18:07

 *@备注

 */
func CreateQueryAccountService() (obj *queryAccountService) {
	return &queryAccountService{ReqData:&dto.QueryAccountReqDto{}}
}


func (this queryAccountService)DoExecute()string{
	logs.Info("this=",this," DoExecute")
	return ""
}


