package xhamster

import "testing"

func TestFile2Shell(t *testing.T) {
	File2Shell()
}

/*
go test -v -run  TestSplitLinks ./
*/
func TestSplitLinks(t *testing.T) {
	fp := "/mnt/e/git/ProcessAVI/xhamster/list.txt"
	SplitLinks(fp)
}
