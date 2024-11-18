package ominode

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

var binpath string

func init() {
	var rootpath string
	_, rootpath, _, _ = runtime.Caller(0)     //获取本代码文件地址
	binpath = filepath.Dir(rootpath) + "/bin" //获取redis所在文件夹
}

func StartRedis(port int, path string) {
	filePath := path + "/redis.conf"
	if !fileExists(filePath) {
		os.MkdirAll(path, 0755)
		createFileNX(filePath)
		//创建配置文件
		appendToFile(filePath, "bind 0.0.0.0\n")
		appendToFile(filePath, "protected-mode no\n")
		appendToFile(filePath, "databases 1\n")
		appendToFile(filePath, "daemonize yes\n")
		appendToFile(filePath, "port "+strconv.Itoa(port)+"\n")
		appendToFile(filePath, "dir "+path+"\n")
		appendToFile(filePath, "always-show-logo yes\n")
		appendToFile(filePath, "loglevel debug\n")
		appendToFile(filePath, "save 900 1\n")
		appendToFile(filePath, "save 300 10\n")
		appendToFile(filePath, "save 60 10000\n")
	} else {
		log.Println(filePath, "已存在，将其当前配置文件启动redis")
	}

	//启动redis
	go func() {
		exec.Command("cmd", "/C", fmt.Sprintf("start %s/redis-server "+filePath, binpath)).CombinedOutput()
	}()
	time.Sleep(100 * time.Millisecond)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

func createFileNX(filename string) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		file, err := os.Create(filename)
		if err != nil {
			return
		}
		defer file.Close()
	}
}

func appendToFile(filename string, s string) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	defer file.Close()
	s = s + "\n"
	if s == "" {
		return
	}
	_, err = io.WriteString(file, s)
	if err != nil {
		return
	}
}
