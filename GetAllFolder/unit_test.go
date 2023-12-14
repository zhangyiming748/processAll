package GetAllFolder

import "testing"

/*
go test -run TestListFolders -v
*/
func TestListFolders(t *testing.T) {
	ret := List("/home/zen/Downloads")
	for _, d := range ret {
		t.Log(d)
	}
}
