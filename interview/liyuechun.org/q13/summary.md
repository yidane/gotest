####13. 写出打印的结果。
~~~
type People struct {
	name string `json:"name"`
}

func main() {
	js := `{
		"name":"11"
	}`
	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("people: ", p)
}
~~~
输出内容如下：
~~~
people:  {}
~~~
p中的属性值为空的原因是因为，name的首字母小写，修改成大写，重新运行即可。

~~~
package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	Name string `json:"name"`
}

func main() {
	js := `{
        "name":"11"
    }`
	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("people: ", p)
}
~~~
运行结果如下：
~~~
people:  {11}

Process finished with exit code 0
~~~