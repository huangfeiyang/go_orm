package log

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

/*
定义error和info日志实例，info为蓝色，error为红色
使用log.Lshortfile 支持显示文件名和代码行号
暴露Error、Errorf、Info、Infof四个方法
*/
var (
	errorLog = log.New(os.Stdout, "\033[31m,[error]\033[0m ", log.LstdFlags | log.Lshortfile)
	infoLog = log.New(os.Stdout, "\033[34m,[info]\033[0m ", log.LstdFlags | log.Lshortfile)
	loggers = []*log.Logger{errorLog, infoLog}
	mu sync.Mutex
)

//log methods
var (
	Error = errorLog.Println
	Errorf = errorLog.Printf
	Info = infoLog.Println
	Infof = infoLog.Printf
)


/*
设置日志层级(InfoLevel,ErrorLevel,Disabled)
*/

//log levels
const (
	InfoLevel = iota
	ErrorLevel
	Disabled
)

//SetLevel controls log level
func SetLevel(level int) {
	mu.Lock()
	defer mu.Unlock()

	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}

	if ErrorLevel < level {
		errorLog.SetOutput(ioutil.Discard)
	}

	if InfoLevel < level {
		infoLog.SetOutput(ioutil.Discard)
	}
}