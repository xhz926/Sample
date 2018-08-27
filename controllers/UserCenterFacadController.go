package controllers

import "github.com/astaxie/beego"

/**

 *@描述 用户中心控制

 *@创建人  吉喆

 *@创建时间  2018/08/22 15:43

 *@备注

 */
type UserCenterFacadController struct {
	beego.Controller
}

/**

 *@描述 修改登录密码 ATS11052

 *@创建人  吉喆

 *@创建时间  2018/08/22 15:43

 *@备注

 */
func (this UserCenterFacadController) ChangeLoginPwd() {
}

/**

 *@描述 令牌验证 001009

 *@创建人  吉喆

 *@创建时间  2018/08/22 15:46

 *@备注

 */
func (this UserCenterFacadController) VerifyToken() {
}

/**

 *@描述 查询银行名称统一行号信息服务 UPP6103T

 *@创建人  吉喆

 *@创建时间  2018/08/22 15:50

 *@备注

 */
func (this UserCenterFacadController) QueryBankPartyid() {
}

/**

 *@描述 用户中心控制初始化

 *@创建人  吉喆

 *@创建时间  2018/08/22 15:44

 *@备注

 */
func init() {
	//修改交易密码
	RegisterController("/api/v1.0/BIP0003T/TRU001", &UserCenterFacadController{}, "*:ChangeLoginPwd")
	//令牌验证
	RegisterController("/api/v1.0/001009/TRU001", &UserCenterFacadController{}, "*:VerifyToken")
	//查询银行名称统一行号信息服务
	RegisterController("/api/v1.0/UPP6103T/TRU001", &UserCenterFacadController{}, "*:QueryBankPartyid")
}
