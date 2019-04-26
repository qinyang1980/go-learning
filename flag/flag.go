package flag

import (
	"flag"
)

// 定义命令行参数
var mode = flag.String("mode", "", "process mode")

func TestFlag() {

	// 解析命令行参数
	flag.Parse()

	flag.Usage()
}
