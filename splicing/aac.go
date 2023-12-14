package splicing

import (
	"os/exec"
	"processAll/util"
)

/*
ffmpeg -f concat -safe 0 -i work.txt -c copy x.aac
仅处理当前文件夹下的aac文件
*/
func SpicingAAC(out string) {
	cmd := exec.Command("ffmpeg", "-f", "concat", "-safe", "0", "-i", "work.txt", "-c", "copy", out)
	err := util.ExecCommand(cmd)
	if err != nil {
		panic(err)
	}
}
