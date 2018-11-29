package logger_test

import (
	"base/logger"
	"fmt"
	"testing"
)

func TestNewLogSys(t *testing.T) {
	ret := logger.NewLogSys("log")
	if ret == nil {
		fmt.Println("NewLogSys Log Fail")
		return
	}

	ret.Close()
}

func TestSysError(t *testing.T) {
	//初始化LOG文件
	ret := logger.Instance().Init("syserr", 0, 0)
	if ret != true {
		fmt.Println("SysError Instance Log Fail")
		return
	}
	logger.SysError("%s:%s", "testError1", "testError2")

	logger.Instance().Close()
}

func TestSysLog(t *testing.T) {
	//初始化LOG文件
	ret := logger.Instance().Init("syslog", 0, 0)
	if ret != true {
		fmt.Println("SysLog Instance Log Fail")
		return
	}
	logger.SysLog("%v:%v", "testLog1", "testLog2")

	logger.Instance().Close()
}
