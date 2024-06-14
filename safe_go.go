package tool

import (
	"bytes"
	"fmt"
	"reflect"
	"runtime"
)

// SafeRun 安全的运行go
func SafeRun(f func()) {
	defer Recover()
	f()
}

// SafeRunParams 带参数安全的运行go
func SafeRunParams(fun interface{}, args ...interface{}) (_ error) {
	v := reflect.ValueOf(fun)
	go func() {
		defer Recover()
		switch v.Kind().String() {
		case "func":
			pps := make([]reflect.Value, 0, len(args))
			for _, arg := range args {
				pps = append(pps, reflect.ValueOf(arg))
			}
			v.Call(pps)
		default:
			err := fmt.Errorf("func is not func,type=%v", v.Kind().String())
			logn.Errorf("error=%v", err)
		}
	}()
	return
}

func Recover() {
	if r := recover(); r != nil {
		stacks := stack(2)
		fmt.Errorf("[SafeRun] %s", string(stacks))
	}
}

// stack returns a nicely formatted stack frame, skipping skip frames.
func stack(skip int) []byte {
	buf := new(bytes.Buffer)
	for i := skip; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fmt.Fprintf(buf, "%s:%d\n", file, line)
	}
	return buf.Bytes()
}
