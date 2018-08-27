package controllers

import "github.com/astaxie/beego"

/**

 *@描述 账户检查相关控制

 *@创建人  吉喆

 *@创建时间  2018/08/18 18:12

 *@备注

 */
type AccountCheckingController struct {
	beego.Controller
}

/**

 *@描述 对账组状态查询 UPP4212T

 *@创建人  吉喆

 *@创建时间  2018/08/21 21:07

 *@备注

 */
func (this AccountCheckingController) QueryCheckingGroupStatus() {
}

/**

 *@描述 异步请求对账服务 UPP4213T

 *@创建人  吉喆

 *@创建时间  2018/08/21 21:10

 *@备注

 */
func (this AccountCheckingController) RequestCheckingAsync() {
}

/**

 *@描述 账户检查相关控制初始化函数

 *@创建人  吉喆

 *@创建时间  2018/08/21 21:08

 *@备注

 */
func init() {
	//4+1转账认证接口
	RegisterController("/api/v1.0/UPP4212T/TRU001", &AccountCheckingController{}, "*:QueryCheckingGroupStatus")
	//异步请求对账服务
	RegisterController("/api/v1.0/UPP4213T/TRU001", &AccountCheckingController{}, "*:RequestCheckingAsync")
}
