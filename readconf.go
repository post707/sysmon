package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const middle = "->"

type Config struct {
	Mymap  map[string]string
	strcet string
}

func (c *Config) InitConfig(path string) {

	c.Mymap = make(map[string]string)

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	readfile := bufio.NewReader(file)
	for {
		b, _, err := readfile.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			panic(err)
		}
		s := strings.TrimSpace(string(b))
		if strings.Index(s, "#") == 0 {
			continue
		}
		// if strings.Index(s, "//") == 0 {
		// 	continue
		// }
		n1 := strings.Index(s, "[")
		n2 := strings.LastIndex(s, "]")
		if n1 > -1 && n2 > -1 && n2 > n1+1 {
			c.strcet = strings.TrimSpace(s[n1+1 : n2])
			continue
		}

		// if len(c.strcet) == 0 {
		// 	continue
		// }
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}

		frist := strings.TrimSpace(s[:index])
		if len(frist) == 0 {
			continue
		}
		second := strings.TrimSpace(s[index+1:])

		pos := strings.Index(second, "\t#")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, " #")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, "\t//")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, " //")
		if pos > -1 {
			second = second[0:pos]
		}

		if len(second) == 0 {
			continue
		}

		key := c.strcet + middle + frist
		c.Mymap[key] = strings.TrimSpace(second)
	}
}

func (c Config) Read(node, key string) string {
	key = node + middle + key
	v, found := c.Mymap[key]
	if !found {
		return ""
	}
	return v
}
