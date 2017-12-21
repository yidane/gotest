package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

var buffer = bytes.Buffer{}

func main() {
	genegrateSQLFromCSV("201612.csv")
	for i := 1; i < 11; i++ {
		if i < 10 {
			genegrateSQLFromCSV(fmt.Sprintf("20170%d.csv", i))
		} else {
			genegrateSQLFromCSV(fmt.Sprintf("2017%d.csv", i))
		}
	}

	fmt.Println(buffer.String())
}

func genegrateSQLFromCSV(path string) {
	const sql string = "INSERT INTO card(`对账标志`,`交易日期`,`记账日期`,`交易摘要`,`卡号后四位`,`人民币金额`,`消费类别`,`备注`)VALUES"
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	buf := bytes.Buffer{}
	buf.WriteString("-- ")
	buf.WriteString(path)
	buf.WriteString("\n")
	buf.WriteString(sql)
	r := csv.NewReader(f)

	r.Read() //跳过表头
	contain := false
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			return
		}
		if contain {
			buf.WriteString(`,`)
		}
		buf.WriteString(`(`)
		contain = true
		for i := 0; i < len(line); i++ {
			if i == 5 {
				buf.WriteString(`'` + strings.Replace(line[i], `,`, ``, 1) + `'`)
			} else {
				buf.WriteString(`'` + strings.TrimSpace(line[i]) + `'`)
			}

			if i < len(line)-1 {
				buf.WriteString(`,`)
			}
		}
		buf.WriteString(`)`)
	}
	buf.WriteString(`;`)
	buf.WriteString("\n")

	buffer.WriteString(buf.String())
}
