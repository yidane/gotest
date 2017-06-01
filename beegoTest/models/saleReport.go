package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/yidane/gotest/beegoTest/models/grid/data"
	"github.com/yidane/gotest/beegoTest/models/grid/style"
)

//GetDefine 获取报表数据
func GetDefine() (data.Define, style.Grid) {
	db, err := sql.Open("mysql", "dbn_admin:dbn002385@tcp(172.16.200.65:4501)/NXin_Qlw_Business?charset=utf8")
	defer db.Close()
	if err != nil {
		panic(err)
	}
	rs, err := db.Query(`Select  Distinct 
						NXin_Qlw_Business.SA_Order.NumericalOrder,
						NXin_Qlw_Business.SA_Order.DataDate,
						NXin_Qlw_Business.SA_Order.PaymentStatus,
						NXin_Qlw_Business.SA_Order.PaymentAmount,
						NXin_Qlw_Business.SA_Order.Consignee From NXin_Qlw_Business.SA_Order
						Inner Join
						NXin_Qlw_Business.SA_OrderDetail ON NXin_Qlw_Business.SA_Order.NumericalOrder=NXin_Qlw_Business.SA_OrderDetail.NumericalOrder`)
	if err != nil {
		panic(err)
	}

	define := data.Define{}
	define.ColumnCollection = []string{"NumericalOrder", "DataDate", "PaymentStatus", "PaymentAmount", "Consignee"}
	rows := []data.DefineRow{}
	for rs.Next() {
		var NumericalOrder *int64
		var DataDate *string
		var PaymentStatus *string
		var PaymentAmount *string
		var Consignee *string
		err = rs.Scan(&NumericalOrder, &DataDate, &PaymentStatus, &PaymentAmount, &Consignee)
		if err != nil {
			panic(err)
		}
		rows = append(rows, data.DefineRow{CellCollection: []interface{}{NumericalOrder, DataDate, PaymentStatus, PaymentAmount, Consignee}})
	}
	define.RowCollection = rows

	//生成Grid
	g := style.Grid{
		RowCollection: []style.GridRow{style.CreateHeader("DataDate", "NumericalOrder", "PaymentStatus", "PaymentAmount", "Consignee")},
	}
	dataDateCell := style.GridCell{
		Value:         "[" + "DataDate" + "]",
		TextAlign:     "center",
		VerticalAlign: "center",
		Link:          "http://www.baidu.com",
		CellProperty: style.GridCellProperty{
			IsGroup: true,
		},
	}
	numericalOrderCell := style.GridCell{
		Value:        "[" + "NumericalOrder" + "]",
		CellProperty: style.GridCellProperty{},
	}
	paymentStatusCell := style.GridCell{
		Value:        "[" + "PaymentStatus" + "]",
		CellProperty: style.GridCellProperty{},
	}
	paymentAmountCell := style.GridCell{
		Value:        "[" + "PaymentAmount" + "]",
		CellProperty: style.GridCellProperty{},
	}
	consigneeCell := style.GridCell{
		Value:        "[" + "Consignee" + "]",
		CellProperty: style.GridCellProperty{},
	}
	g.RowCollection = append(g.RowCollection, style.GridRow{
		IsBindingRow: true,
		Height:       30,
		CellCollection: []style.GridCell{
			dataDateCell,
			numericalOrderCell,
			paymentStatusCell,
			paymentAmountCell,
			consigneeCell,
		},
	})
	g.FrozenRows = 1
	g.ColumnWidthCollection = []float64{0.2, 0.15, 0.15, 0.15, 0.15}

	return define, g
}
