package enum

import (
	"fmt"
	"hxyhBankTest/models/server"
)

/**

 *@描述 银行交易码Enum

 *@创建人  吉喆

 *@创建时间  2018/08/24 23:47

 *@备注

 */
type TransCodeEnum int

/**

 *@描述 银行交易码Enum dto

 *@创建人  吉喆

 *@创建时间  2018/08/24 23:51

 *@备注

 */
type transCodeEnumDto struct {
	Name    string
	Message string
	Obj interface{}
}

/**

 *@描述 银行交易码列表

 *@创建人  吉喆

 *@创建时间  2018/08/24 23:51

 *@备注

 */
var transCodeEnumList = map[TransCodeEnum]*transCodeEnumDto{
	UPP4212T:  &transCodeEnumDto{Name: "UPP4212T", Message: "对账组状态查询"},
	UPP4213T:  &transCodeEnumDto{Name: "UPP4213T", Message: "异步请求对账服务"},
	BIP0054T:  &transCodeEnumDto{Name: "BIP0054T", Message: "账户状态查询、资产查询",Obj:server.CreateQueryAccountService()},
	GHC1049T:  &transCodeEnumDto{Name: "GHC1049T", Message: "内部户查询"},
	BIP0139T:  &transCodeEnumDto{Name: "BIP0139T", Message: "电子账户销户"},
	BIP0177T:  &transCodeEnumDto{Name: "BIP0177T", Message: "销户前检查"},
	CRC0007T:  &transCodeEnumDto{Name: "CRC0007T", Message: "银联水晶球个人风险评估"},
	CRC0004T:  &transCodeEnumDto{Name: "CRC0004T", Message: "同盾信息收集"},
	CRC0005T:  &transCodeEnumDto{Name: "CRC0005T", Message: "汇法信息收集"},
	DWB0314T:  &transCodeEnumDto{Name: "DWB0314T", Message: "同盾风险核查概要信息查询"},
	DWB0311T:  &transCodeEnumDto{Name: "DWB0311T", Message: "银联身份核查信息查询"},
	DWB0316T:  &transCodeEnumDto{Name: "DWB0316T", Message: "汇法网个人核查详细信息查询"},
	DWB0315T:  &transCodeEnumDto{Name: "DWB0315T", Message: "汇法网个人核查基本信息查询"},
	PBS01010:  &transCodeEnumDto{Name: "PBS01010", Message: "开通网银"},
	IAT0403T:  &transCodeEnumDto{Name: "IAT0403T", Message: "查询银行卡"},
	PBS01006:  &transCodeEnumDto{Name: "PBS01006", Message: "查询客户身份标识"},
	EEY1:      &transCodeEnumDto{Name: "EEY1", Message: "转加密"},
	PBS01001:  &transCodeEnumDto{Name: "PBS01001", Message: "登录密码校验"},
	FLS02089:  &transCodeEnumDto{Name: "FLS02089", Message: "产品目录查询"},
	FLS12009:  &transCodeEnumDto{Name: "FLS12009", Message: "签约信息查询"},
	FLS12007:  &transCodeEnumDto{Name: "FLS12007", Message: "产品申购/追加"},
	FLS12005:  &transCodeEnumDto{Name: "FLS12005", Message: "委托存单查询"},
	FLS12008:  &transCodeEnumDto{Name: "FLS12008", Message: "申购全部撤销"},
	TPS01012:  &transCodeEnumDto{Name: "TPS01012", Message: "根据卡BIN查银行名称"},
	PBS02074:  &transCodeEnumDto{Name: "PBS02074", Message: "申请令牌"},
	PBS02076:  &transCodeEnumDto{Name: "PBS02076", Message: "绑定账号并五要素验证"},
	PBS02077:  &transCodeEnumDto{Name: "PBS02077", Message: "设置密码并开户"},
	PBS02078:  &transCodeEnumDto{Name: "PBS02078", Message: "发送短信验证"},
	PBS02079:  &transCodeEnumDto{Name: "PBS02079", Message: "身份证OCR比对"},
	PBS02080:  &transCodeEnumDto{Name: "PBS02080", Message: "上送客户端错误信息"},
	PBS02081:  &transCodeEnumDto{Name: "PBS02081", Message: "活体检测信息录入"},
	PBS02082:  &transCodeEnumDto{Name: "PBS02082", Message: "联网核查"},
	PBS02075:  &transCodeEnumDto{Name: "PBS02075", Message: "二类户开户校验"},
	PBS02083:  &transCodeEnumDto{Name: "PBS02083", Message: "三照对比"},
	ZTSA24M01: &transCodeEnumDto{Name: "ZTSA24M01", Message: "发送短信"},
	UPP3104T:  &transCodeEnumDto{Name: "UPP3104T", Message: "查询签约协议信息"},
	UPP2211T:  &transCodeEnumDto{Name: "UPP2211T", Message: "账户金额冻结"},
	UPP2212T:  &transCodeEnumDto{Name: "UPP2212T", Message: "账户金额解冻"},
	UPP1101T:  &transCodeEnumDto{Name: "UPP1101T", Message: "内部转账流程"},
	UPP6102T:  &transCodeEnumDto{Name: "UPP6102T", Message: "放款流程"},
	UPP6101T:  &transCodeEnumDto{Name: "UPP6101T", Message: "路由试算"},
	UPP1601T:  &transCodeEnumDto{Name: "UPP1601T", Message: "支付订单概要查询"},
	EAS1011T:  &transCodeEnumDto{Name: "EAS1011T", Message: "查询活期户交易明细列表"},
	BIP0003T:  &transCodeEnumDto{Name: "BIP0003T", Message: "修改交易密码"},
	BIP0008T:  &transCodeEnumDto{Name: "BIP0008T", Message: "校验交易密码"},
	ATS11052:  &transCodeEnumDto{Name: "ATS11052", Message: "修改登录密码"},
	UPP6103T:  &transCodeEnumDto{Name: "UPP6103T", Message: "查询银行名称统一行号信息服务"},
	_001009:   &transCodeEnumDto{Name: "001009", Message: "修改登录密码"},
}

func GetTrancCodeEunmByCode(code int) (enum TransCodeEnum) {
	enum = TransCodeEnum(code)
	if nil == transCodeEnumList[enum] {
		str := fmt.Sprintf("枚举输入错误。输入code=%d", +enum)
		panic(str)
	}
	return enum
}

/**

 *@描述 银行交易码枚举

 *@创建人  吉喆

 *@创建时间  2018/08/24 23:51

 *@备注

 */
const (
	UPP4212T  TransCodeEnum = iota //对账组状态查询
	UPP4213T                       //异步请求对账服务
	BIP0054T                       //账户状态查询、资产查询
	GHC1049T                       //内部户查询
	BIP0139T                       //电子账户销户
	BIP0177T                       //销户前检查
	CRC0007T                       //银联水晶球个人风险评估
	CRC0004T                       //同盾信息收集
	CRC0005T                       //汇法信息收集
	DWB0314T                       //同盾风险核查概要信息查询
	DWB0311T                       //银联身份核查信息查询
	DWB0316T                       //汇法网个人核查详细信息查询
	DWB0315T                       //汇法网个人核查基本信息查询
	PBS01010                       //开通网银
	IAT0403T                       //查询银行卡
	PBS01006                       //查询客户身份标识
	EEY1                           //转加密
	PBS01001                       //登录密码校验
	FLS02089                       //产品目录查询
	FLS12009                       //签约信息查询
	FLS12007                       //产品申购/追加
	FLS12005                       //委托存单查询
	FLS12008                       //申购全部撤销
	TPS01012                       //根据卡BIN查银行名称
	PBS02074                       //申请令牌
	PBS02076                       //绑定账号并五要素验证
	PBS02077                       //设置密码并开户
	PBS02078                       //发送短信验证
	PBS02079                       //身份证OCR比对
	PBS02080                       //上送客户端错误信息
	PBS02081                       //活体检测信息录入
	PBS02082                       //联网核查
	PBS02075                       //二类户开户校验
	PBS02083                       //三照对比
	ZTSA24M01                      //发送短信
	UPP3104T                       //查询签约协议信息
	UPP2211T                       //账户金额冻结
	UPP2212T                       //账户金额解冻
	UPP1101T                       //内部转账流程
	UPP6102T                       //放款流程
	UPP6101T                       //路由试算
	UPP1601T                       //支付订单概要查询
	EAS1011T                       //查询活期户交易明细列表
	BIP0003T                       //修改交易密码
	BIP0008T                       //校验交易密码
	ATS11052                       //修改登录密码
	UPP6103T                       //查询银行名称统一行号信息服务
	_001009                        //修改登录密码
)

/**

 *@描述 获取银行交易码名称

 *@创建人  吉喆

 *@创建时间  2018/08/24 23:52

 *@备注

 */
func (this TransCodeEnum) ToString() string {
	return transCodeEnumList[this].Name
}

/**

 *@描述 获取银行交易码描述

 *@创建人  吉喆

 *@创建时间  2018/08/24 23:52

 *@备注

 */
func (this TransCodeEnum) GetMessage() string {
	return transCodeEnumList[this].Message
}

/**

 *@描述 获取注册对象

 *@创建人  吉喆

 *@创建时间  2018/08/25 16:06

 *@备注

 */
func (this TransCodeEnum) GetObj() interface{} {
	return transCodeEnumList[this].Obj
}


