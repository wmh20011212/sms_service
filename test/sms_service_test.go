package test

import (
	"sms/mock"
	"testing"
)

func TestMockSMSService_SendBulkSMS(t *testing.T) {
	// 创建模拟短信服务实例
	mockSMSService := mock.NewMockSMSService()

	// 模拟发送短信
	phoneNumbers := []string{"1234567890", "0987654321"}
	param := "1234"
	mockSMSService.SendBulkSMS(phoneNumbers, param)
}
