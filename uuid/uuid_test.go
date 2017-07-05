package uuid

import (
	"crypto/rand"
	"fmt"
	"testing"
)

func Test_GetGUID(t *testing.T) {
	uuid := RandomUUID()
	fmt.Print(uuid)
}

func Test_GetUUID(t *testing.T) {
	list := make(map[string]int)
	total := 1000000
	i := 0
	for i < total {
		u := make([]byte, 16)
		l, err := rand.Read(u)
		if err != nil {
			continue
		}
		if l != 16 {
			continue
		}
		uuid := fmt.Sprintf("%x", u[:])
		if _, ok := list[uuid]; !ok {
			list[uuid] = 1
		} else {
			list[uuid] = list[uuid] + 1
		}
		i++
	}

	for k, v := range list {
		if v > 1 {
			fmt.Println(k, v)
		}
	}

	fmt.Println(len(list))
}

func Test_GetRandUUID(t *testing.T) {
	u := make([]byte, 16)
	rand.Read(u)
	r := fmt.Sprintf("%08x-%04x-%04x-%04x-%12x", u[:4], u[4:6], u[6:8], u[8:10], u[10:])
	fmt.Println(r)
}
