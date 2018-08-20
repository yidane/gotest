package main


type student struct {
	Name string
}

func f(v interface{}) {
	switch msg := v.(type) {
	case *student, student:
		msg.Name
	}
}

func main(){

}