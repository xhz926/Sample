package controllers

import "github.com/astaxie/beego"

/**

 *@描述 交易流程控制

 *@创建人  吉喆

 *@创建时间  2018/08/22 9:02

 *@备注

 */
type TransactionController struct {
	beego.Controller
}

/**

 *@描述 查询签约协议信息 UPP3104T

 *@创建人  吉喆

 *@创建时间  2018/08/22 9:08

 *@备注

 */
func (this TransactionController) GdhxRecharge() {
}

/**

 *@描述 账户金额冻结 UPP2211T

 *@创建人  吉喆

 *@创建时间  2018/08/22 9:10

 *@备注

 */
func (this TransactionController) FreezeAmount() {
}

/**

 *@描述 账户金额解冻 UPP2212T

 *@创建人  吉喆

 *@创建时间  2018/08/22 9:59

 *@备注

 */
func (this TransactionController) UnFreezeAmount() {
}

/**

 *@描述 内部转账流程 UPP1101T

 *@创建人  吉喆

 *@创建时间  2018/08/22 10:01

 *@备注

 */
func (this TransactionController) GdhxInternalTransfer() {
}

/**

 *@描述 放款流程 UPP6102T

 *@创建人  吉喆

 *@创建时间  2018/08/22 15:33

 *@备注

 */
func (this TransactionController) GdhxLoan() {
}

/**

 *@描述 路由试算 UPP6101T

 *@创建人  吉喆

 *@创建时间  2018/08/22 15:33

 *@备注

 */
func (this TransactionController) GdhxWithdraw() {
}

/**

 *@描述 支付订单概要查询 UPP1601T

 *@创建人  吉喆

 *@创建时间  2018/08/22 15:35

 *@备注

 */
func (this TransactionController) GdhxPayOrderSummary() {
}

/**

 *@描述 查询活期户交易明细列表 EAS1011T

 *@创建人  吉喆

 *@创建时间  2018/08/22 15:41

 *@备注

 */
func (this TransactionController) QueryTransactionFlow() {
}

/**

 *@描述 交易控制初始化

 *@创建人  吉喆

 *@创建时间  2018/08/22 9:03

 *@备注

 */
func init() {
	//查询签约协议信息
	RegisterController("/api/v1.0/UPP3104T/TRU001", &TransactionController{}, "*:GdhxRecharge")
	//账户金额冻结
	RegisterController("/api/v1.0/UPP2211T/TRU001", &TransactionController{}, "*:FreezeAmount")
	//账户金额解冻
	RegisterController("/api/v1.0/UPP2212T/TRU001", &TransactionController{}, "*:UnFreezeAmount")
	//转账流程
	RegisterController("/api/v1.0/UPP1101T/TRU001", &TransactionController{}, "*:GdhxInternalTransfer")
	//放款流程
	RegisterController("/api/v1.0/UPP6102T/TRU001", &TransactionController{}, "*:GdhxLoan")
	//路由试算
	RegisterController("/api/v1.0/UPP6101T/TRU001", &TransactionController{}, "*:GdhxWithdraw")
	//查询活期户交易明细列表
	RegisterController("/api/v1.0/EAS1011T/TRU001", &TransactionController{}, "*:QueryTransactionFlow")
}
