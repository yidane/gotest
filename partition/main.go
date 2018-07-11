package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	const beginYear = 2016
	const endYear = 2021
	const sp = 100

	buf := bytes.Buffer{}

	for i := beginYear; i <= endYear; i++ {
		for q := 1; q <= 4; q++ {
			pname := fmt.Sprintf("%v%v", i, q)
			var qvalue string

			switch q {
			case 1:
				qvalue = fmt.Sprintf("%v0%v", i, 4)
			case 2:
				qvalue = fmt.Sprintf("%v0%v", i, 7)
			case 3:
				qvalue = fmt.Sprintf("%v%v", i, 10)
			case 4:
				qvalue = fmt.Sprintf("%v0%v", i+1, 1)
			}

			buf.WriteString(fmt.Sprintf("partition p%v values less than (%v01) (", pname, qvalue))
			for j := 0; j < sp; j++ {
				buf.WriteString(fmt.Sprintf("subpartition p%v_%v", pname, j))
				if j < sp-1 {
					buf.WriteString(" ,")
				}
			}
			buf.WriteString(")")
			if q < 4 {
				buf.WriteString(",")
			}
		}
		buf.WriteString(",")
	}

	buf.WriteString("partition p20180420 values less than MAXVALUE (")
	for j := 0; j < sp; j++ {
		buf.WriteString(fmt.Sprintf("subpartition p20180420_%v", j))
		if j < sp-1 {
			buf.WriteString(" ,")
		}
	}
	buf.WriteString(")")

	ioutil.WriteFile("1.txt", buf.Bytes(), os.ModeAppend)

	fmt.Println("succeed")
}
