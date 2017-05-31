package models

import (
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/yidane/gotest/beegoTest/models/grid/data"
	"github.com/yidane/gotest/beegoTest/models/grid/style"
)

//SaleReport 销售报表
type SaleReport struct {
	NumericalOrder float64
	DataDate       time.Time
	PaymentStatus  bool
	PaymentAmount  float64
	Consignee      string
}

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
		var NumericalOrder string
		var DataDate string
		var PaymentStatus interface{}
		var PaymentAmount interface{}
		var Consignee interface{}
		err = rs.Scan(&NumericalOrder, &DataDate, &PaymentStatus, &PaymentAmount, &Consignee)
		if err != nil {
			panic(err)
		}
		rows = append(rows, data.DefineRow{CellCollection: []interface{}{NumericalOrder, DataDate, PaymentStatus, PaymentAmount, Consignee}})
	}
	define.RowCollection = rows

	//生成Grid
	g := style.Grid{RowCollection: []style.GridRow{style.CreateHeader("NumericalOrder", "DataDate", "PaymentStatus", "PaymentAmount", "Consignee"),
		style.CreateBoundRow("NumericalOrder", "DataDate", "PaymentStatus", "PaymentAmount", "Consignee")}}
	g.FrozenRows = 1

	return define, g
}
