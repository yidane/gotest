package models

import "testing"
import "fmt"
import "encoding/json"

func Test_GetDefine(t *testing.T) {
	_, _ = GetDefine()
}

func Test_GetTestDefine(t *testing.T) {
	r := GetTestDefine()
	if !r.IsSuccess {
		t.Error(r.Message)
	}

	fmt.Println(json.Marshal(r))
}
