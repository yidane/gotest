package uuid

import "testing"
import "fmt"

func Test_GetUUID(t *testing.T) {
	uuid := RandomUUID()
	fmt.Println(uuid)
}
