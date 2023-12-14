package sim

import "testing"

func TestSimilar(t *testing.T) {
	p1 := "/Users/zen/github/processAVIWithXorm/sim/10003.jpg"
	p2 := "/Users/zen/github/processAVIWithXorm/sim/10003.jpg"
	Similar(p1, p2)
	//-16544.503349298768 相似
	//-678.9719753749642 不相似
	//1 相等
}

func TestLoad(t *testing.T) {
	p := "/Users/zen/github/processAVIWithXorm/sim/10003.jpg"
	image, err := loadImage(p)
	if err != nil {
		t.Log("panic:", err)
		return
	}
	t.Log(image)
}
