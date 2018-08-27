package enum

import "fmt"

type ManngerStatusEnum int

/**

 *@描述 返回数据管理表状态控制枚举

 *@创建人  吉喆

 *@创建时间  2018/08/24 14:34

 *@备注

 */
const (
	FAILED ManngerStatusEnum = iota
	SUCCESS
	TIME_OUT
)

/**

 *@描述 返回数据管理表状态控制枚Dto

 *@创建人  吉喆

 *@创建时间  2018/08/24 14:36

 *@备注

 */
type manngerStatusEnumDto struct {
	Name string
}

/**

 *@描述 返回数据管理表状态控制列表

 *@创建人  吉喆

 *@创建时间  2018/08/24 14:36

 *@备注

 */
var manngerStatusEnumList = map[ManngerStatusEnum]*manngerStatusEnumDto{
	FAILED:   &manngerStatusEnumDto{Name: "FAILE"},
	SUCCESS:  &manngerStatusEnumDto{Name: "SUCCESS"},
	TIME_OUT: &manngerStatusEnumDto{Name: "TIME_OUT"},
}

/**

 *@描述 根据code获取返回数据管理表状态控制枚举值

 *@创建人  吉喆

 *@创建时间  2018/08/24 14:36

 *@备注

 */
func GetManngerStatusEnumByCode(code int) (enum ManngerStatusEnum) {
	enum = ManngerStatusEnum(code)
	if nil == manngerStatusEnumList[enum] {
		str := fmt.Sprintf("枚举输入错误。输入code=%d", +enum)
		panic(str)
	}
	return enum
}

/**

 *@描述 返回值转字符串

 *@创建人  吉喆

 *@创建时间  2018/08/24 14:37

 *@备注

 */
func (this ManngerStatusEnum) ToString() string {
	return manngerStatusEnumList[this].Name
}
