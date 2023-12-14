package information

import (
	"runtime"
	"strings"
)

func GetMachineInfo() string {
	return strings.Join([]string{runtime.GOOS, runtime.GOARCH}, "/")
}
