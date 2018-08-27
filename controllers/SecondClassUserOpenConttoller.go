package controllers

import "github.com/astaxie/beego"

/**

 *@描述 新开户接口控制

 *@创建人  吉喆

 *@创建时间  2018/08/21 23:26

 *@备注

 */
type SecondClassUserOpenConttoller struct {
	beego.Controller
}

/**

 *@描述 根据卡BIN查银行名称 TPS01012

 *@创建人  吉喆

 *@创建时间  2018/08/21 23:24

 *@备注

 */
func (this SecondClassUserOpenConttoller) QueryBankInfoByBIN() {
}

/**

 *@描述 申请令牌 PBS02074

 *@创建人  吉喆

 *@创建时间  2018/08/21 23:27

 *@备注

 */
func (this SecondClassUserOpenConttoller) ApplyForAccessToken() {
}

/**

 *@描述 绑定账号并五要素验证 PBS02076

 *@创建人  吉喆

 *@创建时间  2018/08/21 23:28

 *@备注

 */
func (this SecondClassUserOpenConttoller) BindAccountNoAndCheckFiveElements() {
}

/**

 *@描述 设置密码并开户 PBS02077

 *@创建人  吉喆

 *@创建时间  2018/08/21 23:29

 *@备注

 */
func (this SecondClassUserOpenConttoller) SetPasswordAndOpenAccount() {
}

/**

 *@描述 发送短信验证(仅开户) PBS02078

 *@创建人  吉喆

 *@创建时间  2018/08/21 23:32

 *@备注

 */
func (this SecondClassUserOpenConttoller) SendMessageIdCode() {
}

/**

 *@描述 身份证OCR比对 PBS02079

 *@创建人  吉喆

 *@创建时间  2018/08/21 23:33

 *@备注

 */
func (this SecondClassUserOpenConttoller) IdCardOcrContrast() {
}

/**

 *@描述 上送客户端错误信息 PBS02080

 *@创建人  吉喆

 *@创建时间  2018/08/21 23:34

 *@备注

 */
func (this SecondClassUserOpenConttoller) ClientMsg() {
}

/**

 *@描述 活体检测信息录入 PBS02081

 *@创建人  吉喆

 *@创建时间  2018/08/21 23:35

 *@备注

 */
func (this SecondClassUserOpenConttoller) LivingCheck() {
}

/**

 *@描述 联网核查 PBS02082

 *@创建人  吉喆

 *@创建时间  2018/08/21 23:36

 *@备注

 */
func (this SecondClassUserOpenConttoller) NetworkCheck() {
}

/**

 *@描述 二类户开户校验 PBS02075

 *@创建人  吉喆

 *@创建时间  2018/08/21 23:37

 *@备注

 */
func (this SecondClassUserOpenConttoller) OpenVerifyTwo() {
}

/**

 *@描述 三照对比 PBS02083

 *@创建人  吉喆

 *@创建时间  2018/08/21 23:40

 *@备注

 */
func (this SecondClassUserOpenConttoller) ThreePhotosContrast() {
}

/**

 *@描述 新开户接口控制初始化

 *@创建人  吉喆

 *@创建时间  2018/08/21 23:26

 *@备注

 */
func init() {
	//产品目录查询
	RegisterController("/api/v1.0/TPS01012/TRU001", &SecondClassUserOpenConttoller{}, "*:QueryBankInfoByBIN")
	//申请令牌
	RegisterController("/api/v1.0/PBS02074/TRU001", &SecondClassUserOpenConttoller{}, "*:ApplyForAccessToken")
	//绑定账号并五要素验证
	RegisterController("/api/v1.0/PBS02076/TRU001", &SecondClassUserOpenConttoller{}, "*:BindAccountNoAndCheckFiveElements")
	//设置密码并开户
	RegisterController("/api/v1.0/PBS02077/TRU001", &SecondClassUserOpenConttoller{}, "*:SetPasswordAndOpenAccount")
	//发送短信验证
	RegisterController("/api/v1.0/PBS02078/TRU001", &SecondClassUserOpenConttoller{}, "*:SendMessageIdCode")
	//身份证OCR比对
	RegisterController("/api/v1.0/PBS02079/TRU001", &SecondClassUserOpenConttoller{}, "*:IdCardOcrContrast")
	//上送客户端错误信息
	RegisterController("/api/v1.0/PBS02080/TRU001", &SecondClassUserOpenConttoller{}, "*:ClientMsg")
	//活体检测信息录入
	RegisterController("/api/v1.0/PBS02081/TRU001", &SecondClassUserOpenConttoller{}, "*:LivingCheck")
	//联网核查
	RegisterController("/api/v1.0/PBS02082/TRU001", &SecondClassUserOpenConttoller{}, "*:NetworkCheck")
	//二类户开户校验
	RegisterController("/api/v1.0/PBS02075/TRU001", &SecondClassUserOpenConttoller{}, "*:OpenVerifyTwo")
	//三照对比
	RegisterController("/api/v1.0/PBS02083/TRU001", &SecondClassUserOpenConttoller{}, "*:ThreePhotosContrast")
}
