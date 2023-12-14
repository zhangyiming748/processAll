package chang

import "testing"

func init() {
	go RunNumGoroutineMonitor()
}

/*
go test -v -run TestMaster ./
*/
func TestMaster(t *testing.T) {
	master()
}
