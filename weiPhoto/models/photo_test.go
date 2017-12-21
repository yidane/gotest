package models

import (
	"testing"
)

func TestGetPhotoThumbList(t *testing.T) {
	files, f, err := GetPhotoThumbList(`E:\testimages`, 1)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(f)
	for _, f := range files {
		t.Log((*f).Name)
	}
}
