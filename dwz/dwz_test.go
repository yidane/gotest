package main

import "testing"

func Test_GetLongURL(t *testing.T) {
	s, err := GetLongUrl(`tinyurl=http://dwz.cn/6pe6Pa&access_type=web`)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(*s)
}
