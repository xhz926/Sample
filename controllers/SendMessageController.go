package controllers

import "github.com/astaxie/beego"

/**

 *@描述 发送短信控制

 *@创建人  吉喆

 *@创建时间  2018/08/22 8:50

 *@备注

 */
type SendMessageController struct {
	beego.Controller
}

/**

 *@描述 发送短信 ZTSA24M01

 *@创建人  吉喆

 *@创建时间  2018/08/22 8:51

 *@备注

 */
func (this SendMessageController) SendMessage() {
}

/**

 *@描述 发送短信控制初始化

 *@创建人  吉喆

 *@创建时间  2018/08/22 9:02

 *@备注

 */
func init() {
	//产品目录查询
	RegisterController("/api/v1.0/ZTSA24M01/TRU001", &SendMessageController{}, "*:SendMessage")
}
