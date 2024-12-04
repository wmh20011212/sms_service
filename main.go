package main

import (
	"fmt"
	"github.com/spf13/viper" 
	"log"
	"sms/service"
)

func main() {
	// 加载配置文件
	viper.SetConfigFile("config/config.yml")     // 配置文件路径
	if err := viper.ReadInConfig(); err != nil { // 读取配置文件
		log.Fatalf("Error reading config file, %s", err) // 配置文件读取失败则退出
	}

	// 从配置文件中读取配置
	accessKeyID := viper.GetString("aliyun.access_key_id")
	accessKeySecret := viper.GetString("aliyun.access_key_secret")
	signName := viper.GetString("aliyun.sign_name")
	templateCode := viper.GetString("aliyun.template_code")

	// 创建短信服务实例
	smsService := service.NewSMSService(accessKeyID, accessKeySecret, signName, templateCode)

	// 发送短信
	phoneNumbers := []string{"1234567890", "0987654321"}
	param := "1234"
	smsService.SendBulkSMS(phoneNumbers, param) // 批量发送短信

	fmt.Println("Finished sending all SMS.")
}
