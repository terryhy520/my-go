package service

import (
	"fmt"
)

// GreetingService 提供问候相关功能
type GreetingService struct{}

// NewGreetingService 创建新的 GreetingService 实例
func NewGreetingService() *GreetingService {
	return &GreetingService{}
}

// SayHello 返回问候消息
func (s *GreetingService) SayHello(name string) string {
	return fmt.Sprintf("Hello, %s! Welcome to Go world.", name)
}
