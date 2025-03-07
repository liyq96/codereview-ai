# codereview-ai
## 项目介绍
这是一个与AI相关的项目，简单实现了通过AI进行code review的能力，使用```go```开发
## 项目功能
1. 提供了```WebHook```接口，你可以将git仓库的webhook设置为本接口，参考:```route/code_review.go```
2. 提供了```提示词```功能，你的提示词可以通过markdown文件的形式使用，参考:```doc/system_prompt.md```
