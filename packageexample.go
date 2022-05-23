package main

import (
	"Go/packageexample/mypackage" // 커스텀 패키지
	"fmt"                         // 기본 패키지

	"github.com/guptarohit/asciigraph" //외부 모듈
)

func main() {
	mypackage.SayHello()
	data := []float64{3, 4, 5, 6}
	graph := asciigraph.Plot(data)

	fmt.Println(graph)
}
