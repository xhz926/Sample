package controllers

import "github.com/astaxie/beego"

/**

 *@描述 开户相关控制(旧)

 *@创建人  吉喆

 *@创建时间  2018/08/21 22:33

 *@备注

 */
type HxOpenAccountConntroller struct {
	beego.Controller
}

/**

 *@描述 开通网银 PBS01010

 *@创建人  吉喆

 *@创建时间  2018/08/21 22:32

 *@备注

 */
func (this HxOpenAccountConntroller) OpenCyberBank() {
}

/**

 *@描述 查询银行卡 IAT0403T

 *@创建人  吉喆

 *@创建时间  2018/08/21 22:46

 *@备注

 */
func (this HxOpenAccountConntroller) QueryBankCard() {
}

/**

 *@描述 查询客户身份标识 PBS01006

 *@创建人  吉喆

 *@创建时间  2018/08/21 22:48

 *@备注

 */
func (this HxOpenAccountConntroller) QueryCyberBank() {
}

/**

 *@描述 登录密码校验 PBS01001

 *@创建人  吉喆

 *@创建时间  2018/08/21 22:49

 *@备注

 */
func (this HxOpenAccountConntroller) CheckLoginPwd() {
}

/**

 *@描述 转加密 EEY1

 *@创建人  吉喆

 *@创建时间  2018/08/22 14:29

 *@备注

 */
func (this HxOpenAccountConntroller) PasswordTrans001T() {
}

/**

 *@描述 开户相关控制(旧)初始化

 *@创建人  吉喆

 *@创建时间  2018/08/21 22:34

 *@备注

 */
func init() {
	//开通网银
	RegisterController("/api/v1.0/PBS01010/TRU001", &HxOpenAccountConntroller{}, "*:OpenCyberBank")
	//查询银行卡
	RegisterController("/api/v1.0/IAT0403T/TRU001", &HxOpenAccountConntroller{}, "*:QueryBankCard")
	//查询客户身份标识
	RegisterController("/api/v1.0/PBS01006/TRU001", &HxOpenAccountConntroller{}, "*:QueryCyberBank")
	//登录密码校验
	RegisterController("/api/v1.0/PBS01001/TRU001", &HxOpenAccountConntroller{}, "*:CheckLoginPwd")
	//转加密
	RegisterController("/api/v1.0/EEY1/TRU001", &HxOpenAccountConntroller{}, "*:PasswordTrans001T")
}
