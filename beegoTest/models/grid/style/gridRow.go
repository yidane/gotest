package style

//GridRow 定义报表行对象
type GridRow struct {
	Height         int
	IsBindingRow   bool
	CellCollection []GridCell
}

//CreateHeader 构造表头
func CreateHeader(n ...string) (r GridRow) {
	r = GridRow{
		Height:         30,
		IsBindingRow:   false,
		CellCollection: []GridCell{},
	}

	for _, v := range n {
		r.CellCollection = append(r.CellCollection, GridCell{Value: v})
	}

	return
}

//CreateBoundRow 构造绑定行
func CreateBoundRow(n ...string) (r GridRow) {
	r = GridRow{
		Height:         30,
		IsBindingRow:   true,
		CellCollection: []GridCell{},
	}

	for _, v := range n {
		r.CellCollection = append(r.CellCollection, GridCell{Value: "[" + v + "]"})
	}

	return
}
