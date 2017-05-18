package main
import (
	"fmt"
	"time"
	"github.com/tong"
)

func main() {
	tw, err := tong.NewIdWorker(1)
	if err!= nil {
		fmt.Println(err)
	}
	t2 := time.Now()
	j:=0
	for i := 0; i < 1000000; i++ {
		if _, err := tw.NextId(); err != nil {
			fmt.Println(err)
			j++
		} else{
			//fmt.Println(id)
		}
	}
	//记录结束时间
	used2 := time.Since(t2)
	//输出执行时间，单位为毫秒。

	fmt.Println(used2)
	fmt.Println(j)
}