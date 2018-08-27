package controllers

import "github.com/astaxie/beego"

/**

 *@描述 结构性存款控制

 *@创建人  吉喆

 *@创建时间  2018/08/21 23:12

 *@备注

 */
type PayDepositProductController struct {
	beego.Controller
}

/**

 *@描述 产品目录查询 FLS02089

 *@创建人  吉喆

 *@创建时间  2018/08/21 23:10

 *@备注

 */
func (this PayDepositProductController) DepositQueryDepositProduct() {
}

/**

 *@描述 签约信息查询 FLS12009

 *@创建人  吉喆

 *@创建时间  2018/08/21 23:16

 *@备注

 */
func (this PayDepositProductController) DepositQuerySign() {
}

/**

 *@描述 产品申购/追加 FLS12007

 *@创建人  吉喆

 *@创建时间  2018/08/21 23:18

 *@备注

 */
func (this PayDepositProductController) DepositContractAppend() {
}

/**

 *@描述 委托存单查询 FLS12005

 *@创建人  吉喆

 *@创建时间  2018/08/21 23:19

 *@备注

 */
func (this PayDepositProductController) DepositQueryReceipt() {
}

/**

 *@描述 申购全部撤销 FLS12008

 *@创建人  吉喆

 *@创建时间  2018/08/21 23:20

 *@备注

 */
func (this PayDepositProductController) DepositWholeCancel() {
}

/**

 *@描述 结构性存款控制初始化

 *@创建人  吉喆

 *@创建时间  2018/08/21 23:13

 *@备注

 */
func init() {
	//产品目录查询
	RegisterController("/api/v1.0/FLS02089/TRU001", &PayDepositProductController{}, "*:DepositQueryDepositProduct")
	//签约信息查询
	RegisterController("/api/v1.0/FLS12009/TRU001", &PayDepositProductController{}, "*:DepositQuerySign")
	//产品申购/追加
	RegisterController("/api/v1.0/FLS12007/TRU001", &PayDepositProductController{}, "*:DepositContractAppend")
	//委托存单查询
	RegisterController("/api/v1.0/FLS12005/TRU001", &PayDepositProductController{}, "*:DepositQueryReceipt")
	//申购全部撤销
	RegisterController("/api/v1.0/FLS12008/TRU001", &PayDepositProductController{}, "*:DepositWholeCancel")
}
