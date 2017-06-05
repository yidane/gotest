package style

//MergeInfo 描叙报表合并单元格
type MergeInfo struct {
	LeftCellRow     int
	LeftCellColumn  int
	RightCellRow    int
	RightCellColumn int
}

//NewMergeInfo 创建新合并对象
func NewMergeInfo(leftCellRow, leftCellColumn, rightCellRow, rightCellColumn int) (mergeInfo MergeInfo) {
	mergeInfo = MergeInfo{
		LeftCellRow:     leftCellRow,
		LeftCellColumn:  leftCellColumn,
		RightCellRow:    rightCellRow,
		RightCellColumn: rightCellColumn,
	}
	return
}
