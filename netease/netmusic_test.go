package main

import "testing"

func Test_Decrypt(t *testing.T) {
	text := "yidane"
	key := createKey()
	p, err := AESDecrypt([]byte(text), []byte(key))
	if err != nil {
		t.Fatal(err)
	}

	t.Log(p)
}

func Test_GetKey(t *testing.T) {
	key := createKey()
	if len(key) != 16 {
		t.Error(key)
	}
}
