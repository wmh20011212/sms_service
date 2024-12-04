package service

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"log"
)

type SMSService struct {
	AccessKeyID     string
	AccessKeySecret string
	SignName        string
	TemplateCode    string
}

func NewSMSService(accessKeyID, accessKeySecret, signName, templateCode string) *SMSService {
	return &SMSService{
		AccessKeyID:     accessKeyID,
		AccessKeySecret: accessKeySecret,
		SignName:        signName,
		TemplateCode:    templateCode,
	}
}

// SendSMS 发送短信的函数，使用协程异步发送
// resultChan 是用来接收发送结果的通道
func (s *SMSService) SendSMS(phoneNumber, param string, resultChan chan<- string) {
	// 使用阿里云 API 创建短信客户端
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", s.AccessKeyID, s.AccessKeySecret)
	if err != nil {
		// 如果创建客户端失败，通过通道返回错误信息
		resultChan <- fmt.Sprintf("Error creating client: %v", err)
		return
	}

	// 创建发送短信的请求
	request := dysmsapi.CreateSendSmsRequest()
	request.PhoneNumbers = phoneNumber                          // 设置接收短信的手机号
	request.SignName = s.SignName                               // 设置短信签名
	request.TemplateCode = s.TemplateCode                       // 设置短信模板编号
	request.TemplateParam = fmt.Sprintf(`{"code":"%s"}`, param) // 设置短信模板中的参数（如验证码）

	// 发送短信请求
	response, err := client.SendSms(request)
	if err != nil {
		// 如果发送失败，通过通道返回错误信息
		resultChan <- fmt.Sprintf("Error sending SMS: %v", err)
		return
	}

	// 如果短信发送失败，返回错误信息
	if response.Code != "OK" {
		resultChan <- fmt.Sprintf("Failed to send SMS: %s", response.Message)
		return
	}

	// 如果短信发送成功，返回成功信息
	resultChan <- fmt.Sprintf("SMS sent successfully to %s: %s", phoneNumber, response.Message)
}

// SendBulkSMS 批量发送短信，使用协程异步处理多个手机号
func (s *SMSService) SendBulkSMS(phoneNumbers []string, param string) {
	// 创建一个通道来接收每个协程的执行结果
	resultChan := make(chan string, len(phoneNumbers))

	// 遍历手机号列表，为每个手机号启动一个协程来发送短信
	for _, phoneNumber := range phoneNumbers {
		go s.SendSMS(phoneNumber, param, resultChan)
	}

	// 等待所有协程完成，接收发送结果并打印
	for i := 0; i < len(phoneNumbers); i++ {
		result := <-resultChan // 从通道中接收每个协程的执行结果
		log.Println(result)    // 打印结果
	}
}
