package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/superbool/dgoogauth"
)

// 定义命令行参数对应的变量
var Secret = flag.String("s", "", "请输入你的登录secret")

func main() {
	flag.Parse()
	fileName := os.Getenv("HOME") + "/.auth_secret"
	//fmt.Println(fileName)
	var secretBytes []byte
	if *Secret != "" {
		autoCreateFileIfNotExist(fileName)
		secretBytes = []byte(*Secret)
		err := ioutil.WriteFile(fileName, secretBytes, 0644)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		s, err := ioutil.ReadFile(fileName)
		secretBytes = s
		if err != nil {
			fmt.Println("打开密钥文件失败，请使用 dadaauth -s <secret> 重新设置密钥", err)
			return
		}
	}
	t0 := time.Now().Unix() / 30
	i := dgoogauth.ComputeCode(string(secretBytes), t0)
	fmt.Printf("%06d\n", i)

}

func autoCreateFileIfNotExist(fileName string) {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		_, err := os.Create(fileName)
		if err != nil {
			fmt.Println("创建文件失败", fileName, err)
		}
	}
}
