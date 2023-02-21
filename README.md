# ChatGPT AI 在线智能工具

## 一、概述

### 1.1 项目简介

使用`golang`封装了`ChatGPT`的`AI`接口，后续有时间会继续完善。

项目示例地址：<https://ai.jkapp.net>

### 1.2 如何运行

首先确保你的电脑上已经安装了`golang`环境，然后执行以下命令：

```bash
# 下载项目
git clone https://github.com/jikerdev/ai.git
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

- 接口示例地址: <https://ai.jkapp.net/completions?token=此处填写问题>

#### 2.1.1 请求参数

| 参数名 |   备注   |     示例值     |
| ------ | -------- | -------------- |
| token  | 关键问题 | 帮我写一首情诗 |

#### 2.1.2 示例

请求示例

```text
GET https://ai.jkapp.net/completions?token=帮我写一首情诗
```

响应示例

```json
{
  "code": 200,
  "data": {
    "result": "\n\n一场相遇至今无知，醉眼缘起挥舞定情思。\n情深似海把一切尽息，她内心惊喜不再深沉。\n多少春风吹凉秋雨，她笑容照亮心间梦想。\n轻声唤慰彩虹般安抚，爱的芬芳于心里发酵。\n岁月静好两情相伴，彼此从此甜蜜于心田。"
  },
  "msg": "success"
}
```
