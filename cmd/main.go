package main

import (
	"fmt"

	"github.com/terryhy520/my-go/internal/service"
)

func main() {
	// 创建服务实例
	greetingService := service.NewGreetingService()

	// 调用服务方法
	message := greetingService.SayHello("World")

	// 输出结果
	fmt.Println(message)
}
