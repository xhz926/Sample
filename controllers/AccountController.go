package controllers

import (
	"github.com/astaxie/beego"
	"hxyhBankTest/models/server"
)

/**

 *@描述 账户相关控制

 *@创建人  吉喆

 *@创建时间  2018/08/18 18:13

 *@备注

 */
type AccountController struct {
	beego.Controller
}

/**

 *@描述 账户状态查询、资产查询 BIP0054T

 *@创建人  吉喆

 *@创建时间  2018/08/21 21:16

 *@备注

 */

func (this AccountController) QueryAccount() {
	//创建管理服务实例
	rspManngerService:=server.CreateRspManngerService()
	//执行服务
	respStr:=rspManngerService.Execute(server.BIP0054T,this.Ctx.Request.Body)
	this.Ctx.WriteString(respStr)
}

/**

 *@描述 内部户查询 GHC1049T

 *@创建人  吉喆

 *@创建时间  2018/08/21 21:20

 *@备注

 */
func (this AccountController) QueryInnerAccount() {
}

/**

 *@描述 电子账户销户 BIP0139T

 *@创建人  吉喆

 *@创建时间  2018/08/21 21:23

 *@备注

 */
func (this AccountController) RemoveAccount() {
}

/**

 *@描述 销户前检查 BIP0177T

 *@创建人  吉喆

 *@创建时间  2018/08/21 21:24

 *@备注

 */
func (this AccountController) CheckedPreRemoveAccount() {
}

/**

 *@描述 账户状态查询、资产查询初始化

 *@创建人  吉喆

 *@创建时间  2018/08/21 22:01

 *@备注

 */
func init() {
	//4+1转账认证接口
	RegisterController("/api/v1.0/BIP0054T/TRU001", &AccountController{}, "*:QueryAccount")
	//内部户查询
	RegisterController("/api/v1.0/GHC1049T/TRU001", &AccountController{}, "*:QueryInnerAccount")
	//电子账户销户
	RegisterController("/api/v1.0/BIP0139T/TRU001", &AccountController{}, "*:RemoveAccount")
	//销户前检查
	RegisterController("/api/v1.0/BIP0177T/TRU001", &AccountController{}, "*:CheckedPreRemoveAccount")
}
