// package main

// import (
// 	"fmt"
// )

// func main() {
// 	//testArgument()

// 	//sort()

// 	// s := GetString()
// 	// fmt.Println(CheckString(s))

// 	for i := 0; i < 5; i++ {
// 		defer fmt.Println(i)
// 	}

// 	arr := map[string]bool{
// 		"yidane": true,
// 	}

// 	fmt.Println(arr["yidane"])
// 	fmt.Println(arr["yinsw"])

// 	var x int64 = 1<<63 - 1
// 	fmt.Printf("%d	%x;	%v	%x\n", x, x, int64(x), int64(x))

// 	var y uint64 = 1<<64 - 1
// 	fmt.Printf("%d	%x;	%v	%x\n", y, y, int64(y), int64(y))

// 	fmt.Println(uint(0))
// 	fmt.Println(^uint(0))
// 	fmt.Println(1<<63 - 1)
// 	fmt.Println(int(^uint(0) >> 1))

// 	/*
// 		4
// 		3
// 		2
// 		1
// 		0
// 	*/
// }

package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":1718", "http	service	address") //	Q=17,	R=18
var templ = template.Must(template.New("qr").Parse(templateStr))

func main() {
	flag.Parse()
	http.Handle("/", http.HandlerFunc(QR))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
func QR(w http.ResponseWriter, req *http.Request) {
	templ.Execute(w, req.FormValue("s"))
}

const templateStr = `
<html>
<head>
<title>QR	Link	Generator</title>
</head>
<body>
{{if	.}}
<img	src="http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl={{.}}"	
/>
<br>
{{.}}
<br>
<br>
{{end}}
<form	action="/"	name=f	method="GET"><input	maxLength=1024	size=70
name=s	value=""	title="Text	to	QR	Encode"><input	type=submit
value="Show	QR"	name=qr>
</form>
</body>
</html>`
