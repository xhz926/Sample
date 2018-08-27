package mapper

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
)

//数据库驱动类型
var DbDriver string

const(
	QA="qa"
)

/**

 *@描述 基础数据库连接类

 *@创建人  吉喆

 *@创建时间  2018/08/06 22:02

 *@备注

 */
type baseMapper struct {
	db orm.Ormer
}

/**

 *@描述 基础数据库连接类初始化

 *@创建人  吉喆

 *@创建时间  2018/08/07 11:32

 *@备注

 */
func (this *baseMapper)init(){
	this.db=orm.NewOrm()
}

func (this *baseMapper)GetQueryBuilder()orm.QueryBuilder{
	qb,err:=orm.NewQueryBuilder(DbDriver)
	if nil !=err{
		logs.Error("数据库连接初始化失败,err="+err.Error())
		panic("数据库连接初始化失败,err="+err.Error())
	}
	return qb
}

