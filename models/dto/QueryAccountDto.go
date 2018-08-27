package dto

import "encoding/xml"

/**

 *@描述 账户状态查询、资产查询请求dto

 *@创建人  吉喆

 *@创建时间  2018/08/23 17:53

 *@备注

 */
type QueryAccountReqDto struct {
	XMLName      xml.Name               `xml:"ROOT"`
	ReqSvcHeader ReqSvcHeaderDto        `xml:"ReqSvcHeader"`
	SvcBody      queryAccountReqBodyDto `xml:"SvcBody"`
	//服务名称
	EsbServiceName string `xml:"esbServiceName"`
}

/**

 *@描述 账户状态查询、资产查询请求body Dto

 *@创建人  吉喆

 *@创建时间  2018/08/23 17:53

 *@备注

 */
type queryAccountReqBodyDto struct {
	//E账号
	EAccountNo string `xml:"eAccountNo"`
	//客户号
	CustNo string `xml:"custNo"`
}

/**

 *@描述 账户状态查询、资产查询返回dto

 *@创建人  吉喆

 *@创建时间  2018/08/23 17:53

 *@备注

 */
type QueryAccountRspDto struct {
	XMLName      xml.Name               `xml:"ROOT"`
	ReqSvcHeader RspSvcHeaderDto        `xml:"RspSvcHeader"`
	SvcBody      queryAccountRspBodyDto `xml:"SvcBody"`
	//服务名称
	SysChannelName string `xml:"sys_channel_name"`
}

/**

 *@描述 账户状态查询、资产查询返回body dto

 *@创建人  吉喆

 *@创建时间  2018/08/23 17:53

 *@备注

 */
type queryAccountRspBodyDto struct {
	//客户号
	CustNo string `xml:"custNo"`
	//冻结金额
	FreezeAmount string `xml:"freezeAmount"`
	//E账户类型
	EAccountType string `xml:"eAccountType"`
	//预留手机号
	SignificantPhone string `xml:"significantPhone"`
	//账户等级
	EAccountLevel string `xml:"eAccountLevel"`
	//E账户名称
	EAccountName string `xml:"eAccountName"`
	//开户机构
	AccountBranchId string `xml:"accountBranchId"`
	//E账户状态
	StatusId string `xml:"statusId"`
	//证件号
	InfoString string `xml:"infoString"`
	//客户名称
	CustomerName string `xml:"customerName"`
	//币种
	CurrencyUomId string `xml:"currencyUomId"`
	//产品账号
	FinAccountId string `xml:"finAccountId"`
	//可用余额
	AvailableBalance string `xml:"availableBalance"`
	//E账号
	EAccountNo string `xml:"eAccountNo"`
	//涉案账户标识
	AccountFlag string `xml:"accountFlag"`
	//联网核查结果
	IsInterconnectCheck string `xml:"isInterconnectCheck"`
	//开户日期
	FromDate string `xml:"fromDate"`
	//交易渠道状态
	ChannelStatus string `xml:"channelStatus"`
	//实际余额
	ActualBalance string `xml:"actualBalance"`
	//证件类型 默认1010-居民身份证
	ContactCertificateTypeId string `xml:"contactCertificateTypeId"`
	//核实状态
	VerifyStatus string `xml:"verifyStatus"`
}
