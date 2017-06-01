package style

//GridCellProperty 为报表单元格定义附加属性
type GridCellProperty struct {
	OrderByDirection OrderByDirection
	CanCompare       bool
	Encording        string
	IsGroup          bool
	SubTotalType     AggregateType
	TotalType        AggregateType
	HyperLink        string
}

//OrderByDirection 定义排序规则
type OrderByDirection int

const (
	//None 定义默认排序规则
	None = iota
	//Asc 顺序
	Asc
	//Desc 逆序
	Desc
)

//AggregateType 定义统计规则
type AggregateType int

const (
	//Sum 求和
	Sum = 1
	//Avg 求平均值
	Avg = 2
	//Max 求最大值
	Max = 3
	//Min 求最小值
	Min = 4
)
