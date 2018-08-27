package controllers

import "github.com/astaxie/beego"

/**

 *@描述 资产黑名单查询控制

 *@创建人  吉喆

 *@创建时间  2018/08/21 22:33

 *@备注

 */
type CreditController struct {
	beego.Controller
}

/**

 *@描述 银联水晶球个人风险评估 CRC0007T

 *@创建人  吉喆

 *@创建时间  2018/08/21 22:01

 *@备注

 */
func (this CreditController) UnionpayRiskAssessment() {
}

/**

 *@描述 同盾信息收集 CRC0004T

 *@创建人  吉喆

 *@创建时间  2018/08/21 22:05

 *@备注

 */
func (this CreditController) TongDunCollectInformation() {
}

/**

 *@描述 汇法信息收集 CRC0005T

 *@创建人  吉喆

 *@创建时间  2018/08/21 22:15

 *@备注

 */
func (this CreditController) HuiFaCollectInformation() {
}

/**

 *@描述 同盾风险核查概要信息查询 DWB0314T

 *@创建人  吉喆

 *@创建时间  2018/08/21 22:23

 *@备注

 */
func (this CreditController) TongDunQuery() {
}

/**

 *@描述 银联身份核查信息查询 DWB0311T

 *@创建人  吉喆

 *@创建时间  2018/08/21 22:24

 *@备注

 */
func (this CreditController) UnionpayIdentityVerification() {
}

/**

 *@描述 汇法网个人核查详细信息查询 DWB0316T

 *@创建人  吉喆

 *@创建时间  2018/08/21 22:28

 *@备注

 */
func (this CreditController) HuiFaQuery() {
}

/**

 *@描述 汇法网个人核查基本信息查询 DWB0315T

 *@创建人  吉喆

 *@创建时间  2018/08/21 22:29

 *@备注

 */
func (this CreditController) HuiFaSummaryQuery() {
}

/**

 *@描述 资产黑名单查询控制初始化

 *@创建人  吉喆

 *@创建时间  2018/08/21 22:32

 *@备注

 */
func init() {
	//银联水晶球个人风险评估
	RegisterController("/api/v1.0/CRC0007T/TRU001", &CreditController{}, "*:UnionpayRiskAssessment")
	//同盾信息收集
	RegisterController("/api/v1.0/CRC0004T/TRU001", &CreditController{}, "*:TongDunCollectInformation")
	//汇法信息收集
	RegisterController("/api/v1.0/CRC0005T/TRU001", &CreditController{}, "*:HuiFaCollectInformation")
	//同盾风险核查概要信息查询
	RegisterController("/api/v1.0/DWB0314T/TRU001", &CreditController{}, "*:TongDunQuery")
	//同盾风险核查概要信息查询
	RegisterController("/api/v1.0/DWB0314T/TRU001", &CreditController{}, "*:TongDunQuery")
	//银联身份核查信息查询
	RegisterController("/api/v1.0/DWB0311T/TRU001", &CreditController{}, "*:UnionpayIdentityVerification")
	//汇法网个人核查详细信息查询
	RegisterController("/api/v1.0/DWB0316T/TRU001", &CreditController{}, "*:HuiFaQuery")
	//汇法网个人核查基本信息查询
	RegisterController("/api/v1.0/DWB0315T/TRU001", &CreditController{}, "*:HuiFaSummaryQuery")
}
