package main

import (
	"fmt"
	"os/exec"
)


//返回三个字段，时间，%iowait，%idle
func cpuStat() []byte {
	cmd := exec.Command("/bin/sh", "-c", "sar -u|grep -v -E 'Linux|平均时间|cpu|Average'|tail -n 3|awk   '{print $1" "$6" "$8}'")
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
