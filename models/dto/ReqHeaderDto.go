package dto

/**

 *@描述 银行请求报文头文件(API网关)

 *@创建人  吉喆

 *@创建时间  2018/08/23 14:41

 *@备注

 */
type ReqSvcHeaderDto struct {
	//调用方系统ID
	ConsumerId string `xml:"consumerId"`
	//交易码
	TranCode string `xml:"tranCode"`
	//交易流水号
	TranSeqNo string `xml:"tranSeqNo"`
	//全局流水号
	GlobalSeqNo string `xml:"globalSeqNo"`
	//交易日期
	TranDate string `xml:"tranDate"`
	//交易时间
	TranTime string `xml:"tranTime"`
	//交易柜员
	TranTellerNo string `xml:"tranTellerNo"`
	//机构代码
	BranchId string `xml:"branchId"`
	//会计日期
	AcctDate string `xml:"acctDate"`
	//源发起系统ID
	SourceSysId string `xml:"sourceSysId"`
	//渠道号
	Channel string `xml:"channel"`
}

/**

 *@描述 银行请求报文附件数据头文件(外部服务总线（OSB）)

 *@创建人  吉喆

 *@创建时间  2018/08/23 14:41

 *@备注

 */
type ReqSvcExtHeaderMapDto struct {
	//信息格式版本号
	Version string `xml:"version"`
	//报文类型
	RType string `xml:"type"`
	//安全值校验
	MacValue string `xml:"macValue"`
	//登陆sessionID
	SessionID string `xml:"sessionID"`
	//客户渠道终端地址
	ClientIP string `xml:"clientIP"`
	//客户终端MAC地址
	ClientMAC string `xml:"clientMAC"`
	//客户终端操作系统
	ClientOS string `xml:"clientOS"`
	//客户终端浏览器
	ClientBrowser string `xml:"clientBrowser"`
	//客户终端设备型号
	ClientNunitType string `xml:"clientNunitType"`
	//客户终端设备ID
	ClientTerminateNO string `xml:"clientTerminateNO"`
}
