package service

import (
	"testing"
)

func TestGreetingService_SayHello(t *testing.T) {
	// 创建测试用例
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "NormalCase", input: "Alice", expected: "Hello, Alice! Welcome to Go world."},
		{name: "EmptyName", input: "", expected: "Hello, ! Welcome to Go world."},
	}

	// 创建服务实例
	service := NewGreetingService()

	// 执行测试
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := service.SayHello(tt.input)
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}
