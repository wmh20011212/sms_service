package mock

import (
	"fmt"
	"log"
)

// MockSMSService 模拟短信服务
type MockSMSService struct{}

// NewMockSMSService 创建新的 MockSMSService 实例
func NewMockSMSService() *MockSMSService {
	return &MockSMSService{}
}

// SendSMS 模拟发送短信
func (m *MockSMSService) SendSMS(phoneNumber, param string, resultChan chan<- string) {
	// 模拟成功发送短信
	log.Printf("Mock sending SMS to %s with code %s\n", phoneNumber, param)
	resultChan <- fmt.Sprintf("Mock SMS sent successfully to %s", phoneNumber)
}

// SendBulkSMS 批量模拟发送短信
func (m *MockSMSService) SendBulkSMS(phoneNumbers []string, param string) {
	// 创建一个通道来接收每个协程的执行结果
	resultChan := make(chan string, len(phoneNumbers))

	// 遍历手机号列表，为每个手机号启动一个协程来模拟发送短信
	for _, phoneNumber := range phoneNumbers {
		go m.SendSMS(phoneNumber, param, resultChan)
	}

	// 等待所有协程完成，接收发送结果并打印
	for i := 0; i < len(phoneNumbers); i++ {
		result := <-resultChan // 从通道中接收每个协程的执行结果
		log.Println(result)    // 打印结果
	}
}
