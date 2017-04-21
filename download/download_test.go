package download

import "testing"

func Test_Download(t *testing.T) {
	result, err := Download()
	if !result {
		if err != nil {
			t.Error(err)
		}
		t.Error("unkonw error")
	}
}

func Benchmark_Download(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Download()
	}
}
