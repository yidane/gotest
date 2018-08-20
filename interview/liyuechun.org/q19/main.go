package main

type Student struct {
	name string
}

func main() {
	m := map[string]Student{"people": {"liyuechun"}}
	m["people"].name = "wuyanzu"
}