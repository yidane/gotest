package main

import "testing"

func Test_connectionAndCreateIndex(t *testing.T) {
	if err := connectionAndCreateIndex(true); err != nil {
		t.Errorf("connectionAndCreateIndex() error = %v", err)
	}
}
