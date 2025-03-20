# Codereview Ai
![License](https://img.shields.io/badge/license-MIT-yellow)
## 项目介绍
:tada:这是一个练手项目，简单地实现了通过AI进行code review的能力
## 主要框架
- GO 1.22.5
- Gin v1.10.0
## 预期功能
|功能|状态|
|:----|:--:|
|提供接口进行代码审查|:white_check_mark:|
|自定义MD文件提示词功能|:white_check_mark:|
|RAG|:black_square_button:|
|Git Hook|:black_square_button:|
|Git Pull Request comment|:black_square_button:|
## 使用方法
- 在配置文件中完善你的AI配置
- 启动本项目
- 通过接口Post请求，将代码以参数传入，最后打印出结果
  ```json
  {
    "code_content":"your code"
  }
  ```
## 开源许可
MIT License