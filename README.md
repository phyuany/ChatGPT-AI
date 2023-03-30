# ChatGPT AI 在线智能工具

## 一、概述

### 1.1 项目简介

使用`golang`封装了`ChatGPT`的`AI`接口，后续有时间会继续完善。

项目示例地址：**不再提供示例地址**

### 1.2 如何运行

首先确保你的电脑上已经安装了`golang`环境，然后执行以下命令：

```bash
# 下载项目
git clone https://github.com/jikerdev/ChatGPT-AI.git
# 安装依赖
cd ai
go get
# 添加环境变量
export GPT3_API_KEY="此处填写真实的OPENAI API KEY"
# 运行
go run main.go routes.go
```

## 二、接口列表

### 2.1 `/completions`

> 传入一个问题，ChatGPT返回回答内容

- 请求方式：`GET`

- 本地接口示例: <http://localhost/completions?token=此处填写问题>

#### 2.1.1 请求参数

| 参数名 |   备注   |     示例值     |
| ------ | -------- | -------------- |
| token  | 关键问题 | 帮我写一首情诗 |

#### 2.1.2 示例

本地请求示例

```text
GET http://localhost/completions?token=帮我写一首情诗
```

响应数据头为：`content-type: text/event-stream; charset=utf-8`
