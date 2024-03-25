package main

import (
	"fmt"
	"time"

	"github.com/bdgca/hardinfo"
)

func main() {
	hardinfo.PrintHardInfo()
	fmt.Println("==============================硬件信息查询完成==================================")
	for {
		time.Sleep(time.Second * 10)
	}
}
