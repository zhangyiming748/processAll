package mediaInfo

import "testing"

func TestAudio(t *testing.T) {
	ret := GetAudioMedia("/Users/zen/Pictures/譚詠麟 - 水中花 [6LJ2mJU4BpI].m4a")
	t.Logf("%+v\n", ret)
}
