package style

//GridCell define the report minimum unit
type GridCell struct {
	Value         interface{}
	TextAlign     string
	VerticalAlign string
	Wrap          bool
	Italic        bool
	Format        string
	FontSize      int
	FontFamily    string
	Enable        bool
	Color         string
	BorderTop     string
	BorderRight   string
	BorderBottom  string
	Bold          bool
	Background    string
	ColumnName    string
	CellProperty  GridCellProperty
}
