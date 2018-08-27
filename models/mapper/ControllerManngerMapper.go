package mapper

import "hxyhBankTest/models/dao"

/**

 *@描述 返回数据管理表mapper

 *@创建人  吉喆

 *@创建时间  2018/08/24 17:42

 *@备注

 */
type controllerManngerMapper struct {
	baseMapper
}

/**

 *@描述 创建返回数据管理表实例

 *@创建人  吉喆

 *@创建时间  2018/08/24 17:42

 *@备注

 */
func CreateControllerManngerMapper() *controllerManngerMapper {
	cmm := &controllerManngerMapper{}
	cmm.init()
	cmm.db.Using(QA)
	return cmm
}

func (this controllerManngerMapper) SelectcontrollerManngerList() []dao.ControllerManngerDao {
	selectDataStr := ""
	qb := this.GetQueryBuilder()
}
