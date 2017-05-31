package style

//Grid apply report style
type Grid struct {
	Name                  string
	FrozenRows            int
	FrozenColumns         int
	ShowGridLines         bool
	SubTotalColor         string
	TotalColor            string
	ColumnWidthCollection []float64
	RowCollection         []GridRow
}
