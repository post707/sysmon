package main

import (
	"fmt"
	"os/exec"
)

//返回三个字段，时间，%iowait，%idle
func cpuStat() []byte {
	cmd := exec.Command("/bin/sh", "-c", `sar -u |tail -n 3|awk   'BEGIN {print "time %iowait %idle"} {print $1" "$6" "$8}'`)
	cpustat, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(cpustat))
	return cpustat
}

//

func memStat() []byte {

	cmd := exec.Command("/bin/sh", "-c", `free -m|awk   'free -m|grep -v total|awk   'BEGIN {print "Flag TOTAL USED FREE"} {print $1" "$2" "$3" "$4}''`)
	mem, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(mem))
	return mem
}

func ioStat() []byte {
	cmd := exec.Command("/bin/sh", "-c", `iostat -x|tail -n +6|awk '{print $1" "$14}'`)
	io, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(io))
	return io
}
