package rename

import "testing"

// go test -v -run  TestRename ./

func TestRenameForDup(t *testing.T) {
	RenameForDup("/sdcard/Movies/bili", "蔡依林")
}
func TestClearName(t *testing.T) {
	cleanName("/home/zen/storage/AllBackup/20231121_093359/Music", "aac;WAV;ogg")
}
