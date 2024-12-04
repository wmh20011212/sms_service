## 项目简介

本项目是一个 Go 语言编写的短信发送服务示例，支持使用协程并发发送短信。项目中包含了真实的短信发送服务和模拟服务（用于测试）。模拟服务可以在没有实际短信服务（如阿里云、Twilio）账号的情况下完成本地功能测试。

------

## 功能特点

- **支持并发发送短信**：利用 Go 的协程，实现高效的批量短信发送。
- **模块化设计**：代码分为服务层、模拟层和主程序，便于扩展和维护。
- **支持模拟测试**：通过 Mock 服务模拟短信发送，便于开发和测试。
- **配置化**：使用 `config.yml` 文件管理短信服务的配置信息。

------

## 项目结构

```
bash复制代码
/sms
  /config
    config.yml         # 配置文件，存储短信服务的配置信息
  /mock
    mock_sms_service.go # 模拟短信发送服务
  /service
    sms_service.go     # 短信发送服务的业务逻辑
  main.go              # 主程序，初始化并启动短信发送功能
  /test
    sms_service_test.go # 单元测试文件
```

------

## 环境依赖

在运行本项目之前，请确保已安装以下环境：

1. [Go](https://go.dev/) (1.19 或更高版本)
2. 配置管理工具 `viper`（通过 Go 模块自动安装）
3. 阿里云短信服务 SDK（可选，仅用于真实短信发送）

------

## 配置说明

配置文件 `config/config.yml` 包含短信服务所需的相关信息：

```
yml
aliyun:
  access_key_id: "your-access-key-id"       # 阿里云 Access Key ID
  access_key_secret: "your-access-key-secret" # 阿里云 Access Key Secret
  sign_name: "your-sign-name"                # 短信签名
  template_code: "your-template-code"        # 短信模板编号
```

- 如果使用模拟服务，此文件可以任意填写，不影响测试功能。

------

## 使用说明

### 1. 运行主程序

主程序会默认使用模拟服务来发送短信：

```
bash

go run main.go
```

输出示例：

```
css
Mock sending SMS to 1234567890 with code 1234
Mock sending SMS to 0987654321 with code 1234
Finished sending all SMS.
```

### 2. 切换到真实短信服务

如需切换到真实短信服务：

1. 确保 `config.yml` 文件填写了正确的阿里云短信服务信息。
2. 在 `main.go` 中，将 `mockSMSService` 替换为 `smsService`：

```
go
// smsService := service.NewSMSService(accessKeyID, accessKeySecret, signName, templateCode)
// smsService.SendBulkSMS(phoneNumbers, param)
```

1. 运行程序即可发送真实短信。

------

### 3. 运行测试

项目提供了单元测试，用于验证模拟服务的功能：

```
bash

go test ./test
```

测试输出示例：

```
diff
=== RUN   TestMockSMSService_SendBulkSMS
--- PASS: TestMockSMSService_SendBulkSMS (0.00s)
PASS
ok  	sms/test	0.001s
```
