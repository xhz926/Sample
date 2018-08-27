package server

import (
	"github.com/astaxie/beego/logs"
	"io"
	"io/ioutil"
	"reflect"
	"encoding/xml"
)

const (
	//结构体默认声明的请求数据字段
	REQ_DATA="ReqData"
)

/**

 *@描述 返回值接口

 *@创建人  吉喆

 *@创建时间  2018/08/25 9:08

 *@备注

 */
type RspIntereface interface {
	DoExecute() string
}

/**

 *@描述 返回值统一控制类

 *@创建人  吉喆

 *@创建时间  2018/08/25 9:02

 *@备注

 */
type rspManngerService struct {
}

func CreateRspManngerService() *rspManngerService {
	return &rspManngerService{}
}

func (this rspManngerService) Execute(code TransCodeEnum,requestData io.Reader)(result string) {
	logs.Info(code.ToString()+"接口处理请求开始")
	defer func() {
		if err:=recover();nil!=err{
			result="请求异常,"+err.(string)
		}
		logs.Info(code.ToString()+"接口处理请求结束")
	}()
	//入参判断
	obj:=code.GetObj()
	body,err:=ioutil.ReadAll(requestData)

	if nil !=err{
		logs.Error(code.ToString()+"接口,body数据解析失败。err="+err.Error())
		panic(code.ToString()+"接口,body数据解析失败。err="+err.Error())
	}

	if 0 ==len(body){
		logs.Error(code.ToString()+"接口,body为空,请输入正确请求")
		panic(code.ToString()+"接口,body为空,请输入正确请求")
	}

	t:=reflect.TypeOf(obj)
	if _,ok:=t.Elem().FieldByName(REQ_DATA);!ok{
		logs.Error("不存在指定的成员变量,请确认是否已经声明了,ReqData变量")
		panic("系统错误.")
	}

	if _,ok:=obj.(RspIntereface);!ok{
		logs.Error("不存在指定接口,请确认对象是否已经实现RspIntereface接口内容")
		panic("系统错误.")
	}

	var rspIntereface RspIntereface
	rspIntereface=obj.(RspIntereface)

	sf,_:=t.Elem().FieldByName(REQ_DATA)
	v:=reflect.New(sf.Type.Elem())

	err=xml.Unmarshal(body,v.Interface())
	if nil !=err{
		logs.Error("解析xml错误,err="+err.Error())
		panic("解析xml错误,err="+err.Error())
	}
	//logs.Info("test=",test)
	logs.Info("v=",v)

	_=rspIntereface
	return ""
}

func (this rspManngerService)buildRequestData([]byte)[]byte{
	return nil
}

