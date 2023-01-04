package ihelp

import (
	"bytes"
	"fmt"
	"log"
	"runtime/debug"
	"time"
)

// ErrCatch 错误捕获
func ErrCatch() {
	if err := recover(); err != nil {
		errs := debug.Stack()
		log.Printf("错误：%v", err)
		log.Printf("追踪：%s", string(errs))
	}
}

// ShowVersion 输出格式化版本信息
func ShowVersion(text string) {
	text += " " + time.Now().Format("2006/01/02 15:04:05")
	alen := len([]uint8(text))
	line := bytes.Buffer{}
	for i := 0; i < alen; i++ {
		line.Write([]byte("="))
	}
	output := line.String() + "\n" + text + "\n" + line.String() + "\n"
	fmt.Println(output)
}

// Debug kv打印
func Debug(data interface{}) {
	fmt.Printf("%+v\n", data)
}
