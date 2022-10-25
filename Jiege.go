package main

import (
	"fmt"
	"time"
)

func 欢迎来我家玩() chan string {
	time.Sleep(1 * time.Second)
	a := make(chan string, 1)
	return a
}
func 打电动() {
	fmt.Println("输了啦,都是你害的!")
}
func 诸神黄昏(杰哥说 chan string) {
	杰哥说 <- "登dua郎"
	x := <-杰哥说
	fmt.Println(x)
}
func main() {
	打电动()
	go 诸神黄昏(欢迎来我家玩())
	time.Sleep(2 * time.Second) //再给我两秒钟,让我把main函数结成冰
}
