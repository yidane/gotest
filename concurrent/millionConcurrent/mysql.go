package millionConcurrent

import (
	"bytes"
	"database/sql"
	"runtime"

	"fmt"

	"time"

	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

//生成10000000条数据，插入mysql数据库

func genegrateData() {
	cpus := runtime.NumCPU()
	chn := make(chan time.Duration, cpus)
	for i := 0; i < cpus; i++ {
		go func(id int) {
			tNow := time.Now()
			db, err := sql.Open("mysql", "root:sasasasasa@tcp(localhost:3306)/test?charset=utf8")
			defer db.Close()
			if err != nil {
				panic(err)
			}
			batchSize := 1000000
			for j := 0; j < 25; j++ {
				now := time.Now()
				b := bytes.Buffer{}
				b.WriteString("insert into user(name,datetime) values")
				for k := 0; k < batchSize; k++ {
					if k == batchSize-1 {
						b.WriteString("('")
						b.WriteString(strconv.Itoa(id))
						b.WriteRune('_')
						b.WriteString(strconv.Itoa(k))
						b.WriteString("name")
						b.WriteString(strconv.Itoa(j))
						b.WriteString("',sysdate());")
					} else {
						b.WriteString("('")
						b.WriteString(strconv.Itoa(id))
						b.WriteRune('_')
						b.WriteString(strconv.Itoa(k))
						b.WriteString("name")
						b.WriteString(strconv.Itoa(j))
						b.WriteString("',sysdate()),")
					}
				}
				gStr := time.Now().Sub(now).String()
				_, err = db.Exec(b.String())
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(j, "	生成耗时：", gStr, "	插入耗时：", time.Now().Sub(now).String())
			}
			chn <- time.Now().Sub(tNow)
		}(i)
	}

	var fd time.Duration
	for i := 0; i < cpus; i++ {
		id := <-chn
		fmt.Println(id.String())
		if id > fd {
			fd = id
		}
	}

	fmt.Println("总耗时：", fd.String())
}
