package util

import (
	"fmt"
	"github.com/klauspost/cpuid/v2"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

var (
	concurrent      int32
	concurrentLimit = make(chan struct{}, 10)
)

func ReadDB() {
	atomic.AddInt32(&concurrent, 1)
	fmt.Printf("readDB并发度%v\n", atomic.LoadInt32(&concurrent))
	time.Sleep(200 * time.Microsecond)
	atomic.AddInt32(&concurrent, -1)
}
func handler(f func()) {
	concurrentLimit <- struct{}{}
	//readDB()
	f()
	<-concurrentLimit
	return
}
func TestLimitThreads(t *testing.T) {
	for i := 0; i < 100; i++ {
		go handler(ReadDB)
	}
	time.Sleep(3 * time.Second)
}

// go test -v -run TestKindOfCPU ./
func TestKindOfCPU(t *testing.T) {
	fmt.Println("Name:", cpuid.CPU.BrandName)
	fmt.Println("物理核心PhysicalCores:", cpuid.CPU.PhysicalCores)
	fmt.Println("每个核心的线程ThreadsPerCore:", cpuid.CPU.ThreadsPerCore)
	fmt.Println("逻辑核心LogicalCores:", cpuid.CPU.LogicalCores)
	fmt.Println("Family", cpuid.CPU.Family, "Model:", cpuid.CPU.Model, "Vendor ID:", cpuid.CPU.VendorID)
	fmt.Println("Features:", strings.Join(cpuid.CPU.FeatureSet(), ","))
	fmt.Println("Cacheline bytes:", cpuid.CPU.CacheLine)
	fmt.Println("L1 Data Cache:", cpuid.CPU.Cache.L1D, "bytes")
	fmt.Println("L1 Instruction Cache:", cpuid.CPU.Cache.L1I, "bytes")
	fmt.Println("L2 Cache:", cpuid.CPU.Cache.L2, "bytes")
	fmt.Println("L3 Cache:", cpuid.CPU.Cache.L3, "bytes")
	fmt.Println("Frequency", cpuid.CPU.Hz, "hz")

	// Test if we have these specific features:
	if cpuid.CPU.Supports(cpuid.SSE, cpuid.SSE2) {
		fmt.Println("We have Streaming SIMD 2 Extensions")
	}
	t.Logf("Vendor ID is : %v\ttype is %s\n", cpuid.CPU.VendorID, cpuid.CPU.VendorID)

	if fmt.Sprintf("%v", cpuid.CPU.VendorID) == "Intel" {
		t.Log(fmt.Sprintf("%v", cpuid.CPU.VendorID))
	} else {
		fmt.Println("not")
	}
}

func BenchmarkRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomWithSeed()
	}
}
func TestDuplicateBySlice(t *testing.T) {
	s := ReadByLine("E:\\git\\ProcessAVI\\util\\list.txt")
	r := DuplicateBySlice(s)
	result := []string{}
	for _, v := range r {
		prefix := "ytdlp --proxy 127.0.0.1:8889"
		suffix := strings.Join([]string{prefix, v}, " ")
		result = append(result, suffix)
	}
	WriteByLine("E:\\git\\ProcessAVI\\util\\plist.ps1", result)
}

func TestSave(t *testing.T) {
	size, err := GetSize("/Users/zen/Downloads/curl.go")
	if err != nil {
		return
	}
	t.Log(size)
}

func TestName(t *testing.T) {
	a := "张韶涵2023722寓言天津站不害怕afire内场前排高清饭拍我在漫漫长夜之中飞翔寻找属于我的那道星光" //31个汉字
	var s string
	t.Log(len(a))
	for i, char := range a {
		t.Logf("第%d个字符:%v\n", i+1, string(char))
		if i >= 124 {
			t.Log("截取124之前的完整字符")
			break
		} else {
			s = strings.Join([]string{s, string(char)}, "")
		}
	}
	t.Logf("截取的完整字符:%s\n", s)
}
