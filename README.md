# Codereview-Ai
![License](https://img.shields.io/badge/license-MIT-yellow)
## 项目介绍

***
:tada:这是一个练手项目，简单地实现了通过AI进行code review的能力。如果你喜欢，欢迎Star:star:
## 主要技术框架
- GO 1.22.5
- Gin v1.10.0
## 项目功能
- 提供了审查接口，传入代码即可，参考```route/code_review.go```
- 提供了```提示词```功能，你的提示词可以通过markdown文件的形式使用，参考```doc/system_prompt.md```
## 预期功能
:white_check_mark:提供接口进行代码审查
:white_check_mark:提供自定义提示词功能，便于自定义相关规范
:black_square_button:RAG
:black_square_button:Git Hook
:black_square_button:向Git Pull Request提交comment

## 开源许可
MIT License