# 短信发送服务项目

## 项目简介

本项目是一个基于 Go 语言的短信发送服务，使用 **阿里云短信服务** 进行短信的发送。项目支持批量短信发送，并采用协程的方式提高并发处理效率，适用于需要快速、高效发送大量短信的场景

------

## 目录结构

```
/sms
  /config
    config.yml         # 配置文件，存储阿里云短信服务的配置信息
  /service
    sms_service.go     # 短信发送服务的核心逻辑
  main.go              # 主程序入口
  go.mod
```

------

## 环境要求

- **Go 版本**：1.18 或更高版本
- **阿里云账户**：需要开通阿里云短信服务，并获取 Access Key 和 Secret。
- 依赖库：
  - 阿里云短信 SDK (`github.com/aliyun/alibaba-cloud-sdk-go`)
  - 配置管理库 (`github.com/spf13/viper`)

------

## 配置说明

请在 `config/config.yml` 文件中填写阿里云短信服务的配置信息：

```
yml
aliyun:
  access_key_id: "your-access-key-id"         # 阿里云 Access Key ID
  access_key_secret: "your-access-key-secret" # 阿里云 Access Key Secret
  sign_name: "your-sign-name"                 # 短信签名
  template_code: "your-template-code"         # 短信模板编号
```

### 参数说明

- **access_key_id**：从阿里云控制台获取的 Access Key ID。
- **access_key_secret**：从阿里云控制台获取的 Access Key Secret。
- **sign_name**：短信签名，例如 "阿里云"。
- **template_code**：短信模板编号，例如 "SMS_123456789"。

------

## 安装与运行

### 1. 下载项目

```
bash
git clone https://github.com/your-repo/sms-service.git
cd sms-service
```

### 2. 安装依赖

确保已启用 Go Modules，然后运行以下命令安装依赖：

```
bash

go mod tidy
```

### 3. 配置项目

根据实际情况修改 `config/config.yml` 文件，填写正确的阿里云配置。

### 4. 运行项目

运行以下命令启动项目：

```
bash

go run main.go
```

------

## 使用说明

1. 在 `main.go` 中可以自定义需要发送短信的手机号列表：

   ```
   go复制代码phoneNumbers := []string{"1234567890", "0987654321"} // 示例手机号
   param := "1234" // 示例验证码
   smsService.SendBulkSMS(phoneNumbers, param)
   ```

2. 运行程序后，控制台会输出短信发送结果，包括成功和失败的详细信息。

------

## 日志与调试

- 项目会在控制台输出每条短信的发送结果，包括：
  - 发送成功的信息。
  - 发送失败的错误原因。

