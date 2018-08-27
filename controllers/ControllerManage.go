package controllers

var ControllerList = make([]*ControllerDto, 0)

type ControllerDto struct {
	RootPath            string
	ControllerInterface interface{}
	Methods             string
}

/**

*@描述 注册Controller方法
*示例为
  RegisterController("/api/list",&RestController{},"*:ListFood")
  RegisterController("/api/create",&RestController{},"post:CreateFood")
  RegisterController("/api/update",&RestController{},"put:UpdateFood")
  RegisterController("/api/delete",&RestController{},"delete:DeleteFood")
*第一个参数为请求路径,第二个参数为注册Controller的指针，指针，指针重要的事情说3遍 第三个参数是【请求方法:调用函数名】

*此方法请在init函数中调用

*@创建人  吉喆

*@创建时间  2018/08/21 18:02

*@备注

*/
func RegisterController(rootPath string, controllerInterface interface{}, methods string) {
	c := &ControllerDto{}
	c.RootPath = rootPath
	c.ControllerInterface = controllerInterface
	c.Methods = methods
	ControllerList = append(ControllerList, c)
}
