package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// cmds := []*exec.Cmd{
	// 	exec.Command("ps", "-ef"),
	// 	exec.Command("grep", "redis"),
	// 	exec.Command("grep", "-v", "grep"),
	// 	exec.Command("/bin/sh", "-c", "sar -u|grep -v -E 'Linux|平均时间|cpu'|tail -n 3"),
	// }
	cpuStat()

}

//返回三个字段，时间，%iowait，%idle
func cpuStat() []byte {
	cmd := exec.Command("/bin/sh", "-c", "sar -u|grep -v -E 'Linux|平均时间|cpu'|tail -n 3")
	cpustat, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(cpustat))
	return cpustat
}

//

func memStat() {

}

func ioStat() {

}
