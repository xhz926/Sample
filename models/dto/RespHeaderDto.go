package dto

/**

 *@描述 返回报文头文件

 *@创建人  吉喆

 *@创建时间  2018/08/23 15:21

 *@备注

 */
type RspSvcHeaderDto struct {
	//交易日期
	TranDate string `xml:"tranDate"`
	//会计日期
	AcctDate string `xml:"acctDate"`
	//交易时间
	TranTime string `xml:"tranTime"`
	//后台流水号
	BackendSeqNo string `xml:"backendSeqNo"`
	//后台系统ID
	BackendSysId string `xml:"backendSysId"`
	//全局流水号
	GlobalSeqNo string `xml:"globalSeqNo"`
	//服务返回代码
	ReturnCode string `xml:"returnCode"`
	//服务返回信息
	ReturnMsg string `xml:"returnMsg"`
	//用户语言
	LangCode string `xml:"langCode"`
	//扩展内容
	RsrvContent string `xml:"rsrvContent"`
}
