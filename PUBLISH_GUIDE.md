# 发布到 GitHub 指南

本指南将帮助您将 HOTCOIN Go SDK 发布到 GitHub，让用户可以通过 `go get` 命令轻松安装和使用。

## 📋 准备工作

### 1. 检查项目结构

确保您的项目具有以下结构：

```
hotcoin_api_trade/
├── .gitignore          # Git忽略文件
├── LICENSE             # 开源许可证
├── README.md           # 项目说明文档
├── CHANGELOG.md        # 版本更新日志
├── go.mod              # Go模块定义
├── go.sum              # 依赖版本锁定
├── *.go                # Go源代码文件
└── examples/           # 示例代码目录
    └── basic/
        └── main.go
```

### 2. 验证代码质量

在发布前，确保代码质量：

```bash
# 运行所有测试
go test ./...

# 格式化代码
go fmt ./...

# 静态分析
go vet ./...

# 检查依赖
go mod tidy
go mod verify
```

## 🚀 发布流程

### 步骤1: 创建 GitHub 仓库

1. 登录 [GitHub](https://github.com)
2. 点击右上角的 "+" 按钮，选择 "New repository"
3. 设置仓库信息：
   - **Repository name**: `hotcoin-go-sdk` (推荐)
   - **Description**: `HOTCOIN永续合约API的官方Go SDK`
   - **Visibility**: Public (公开仓库，让用户可以访问)
   - ✅ 不要初始化 README、.gitignore 或 LICENSE（我们已经有了）

### 步骤2: 更新 go.mod 文件

确保 `go.mod` 中的模块路径与您的 GitHub 仓库路径一致：

```bash
# 如果您的 GitHub 用户名是 "yourusername"，仓库名是 "hotcoin-go-sdk"
# 修改 go.mod 第一行为：
module github.com/yourusername/hotcoin-go-sdk
```

**重要**: 将 `yourusername` 替换为您的实际 GitHub 用户名！

### 步骤3: 初始化 Git 仓库并推送代码

```bash
# 在项目根目录执行以下命令

# 1. 初始化 Git 仓库
git init

# 2. 添加所有文件
git add .

# 3. 创建初始提交
git commit -m "feat: 初始版本发布 - HOTCOIN Go SDK v1.0.0

- 完整的REST API支持
- WebSocket实时数据推送  
- HmacSHA256签名认证
- 详细的示例代码和文档"

# 4. 添加远程仓库（替换为您的实际仓库地址）
git remote add origin https://github.com/yourusername/hotcoin-go-sdk.git

# 5. 设置默认分支
git branch -M main

# 6. 推送代码到 GitHub
git push -u origin main
```

### 步骤4: 创建版本标签

为了让用户可以使用特定版本，需要创建 Git 标签：

```bash
# 创建 v1.0.0 标签
git tag -a v1.0.0 -m "Release v1.0.0

- 初始版本发布
- 完整的API功能支持
- WebSocket实时数据推送
- 详细的文档和示例"

# 推送标签到 GitHub
git push origin v1.0.0
```

### 步骤5: 在 GitHub 创建正式 Release

1. 访问您的 GitHub 仓库页面
2. 点击右侧的 "Releases" 
3. 点击 "Create a new release"
4. 填写以下信息：
   - **Tag version**: `v1.0.0`
   - **Release title**: `HOTCOIN Go SDK v1.0.0`
   - **Description**: 
     ```markdown
     ## 🎉 HOTCOIN Go SDK 首次发布！
     
     这是 HOTCOIN 永续合约 API 的官方 Go 语言 SDK 的首个正式版本。
     
     ### ✨ 主要功能
     - ✅ 完整的 REST API 支持
     - ✅ WebSocket 实时数据推送
     - ✅ HmacSHA256 签名认证
     - ✅ 自动重连和心跳机制
     - ✅ 详细的错误处理
     - ✅ 类型安全的数据结构
     - ✅ 丰富的示例代码
     - ✅ 完整的单元测试
     
     ### 📦 安装方法
     ```bash
     go get github.com/yourusername/hotcoin-go-sdk@v1.0.0
     ```
     
     ### 🚀 快速开始
     ```go
     import hotcoin "github.com/yourusername/hotcoin-go-sdk"
     
     client := hotcoin.NewClient("your_api_key", "your_secret_key")
     contracts, err := client.Market.GetContracts("")
     ```
     
     更多详细信息请查看 [README.md](./README.md)
     ```
5. 点击 "Publish release"

## 📝 更新 README.md 安装说明

发布后，更新 README.md 中的安装命令：

```bash
# 修改 README.md 中的安装命令为：
go get github.com/yourusername/hotcoin-go-sdk
```

然后提交更新：

```bash
git add README.md
git commit -m "docs: 更新安装说明中的仓库地址"
git push origin main
```

## 🎯 用户使用方式

发布后，用户可以通过以下方式使用您的 SDK：

### 安装

```bash
# 安装最新版本
go get github.com/yourusername/hotcoin-go-sdk

# 安装特定版本
go get github.com/yourusername/hotcoin-go-sdk@v1.0.0
```

### 使用

```go
package main

import (
    "fmt"
    "log"
    
    hotcoin "github.com/yourusername/hotcoin-go-sdk"
)

func main() {
    client := hotcoin.NewClient("your_api_key", "your_secret_key")
    
    contracts, err := client.Market.GetContracts("")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("获取到 %d 个合约\n", len(contracts))
}
```

## 🔄 后续版本发布

当您需要发布新版本时：

### 1. 更新代码和文档

```bash
# 修改代码
# 更新 CHANGELOG.md
# 更新版本号等
```

### 2. 提交更改

```bash
git add .
git commit -m "feat: 添加新功能 - xxx"
git push origin main
```

### 3. 创建新版本标签

```bash
git tag -a v1.1.0 -m "Release v1.1.0 - 添加新功能"
git push origin v1.1.0
```

### 4. 在 GitHub 创建新 Release

重复步骤5，但使用新的版本号。

## 📋 最佳实践

### 版本号规范
遵循 [语义化版本](https://semver.org/lang/zh-CN/) 规范：
- `v1.0.0` - 主版本.次版本.修订版本
- `v1.0.1` - 修复Bug
- `v1.1.0` - 新增功能
- `v2.0.0` - 重大变更

### 提交信息规范
使用清晰的提交信息：
- `feat: 添加新功能`
- `fix: 修复Bug`
- `docs: 更新文档`
- `test: 添加测试`
- `refactor: 重构代码`

### 文档维护
- 保持 README.md 最新
- 更新 CHANGELOG.md
- 提供丰富的示例代码
- 及时回复用户的 Issues 和 PR

## 🆘 常见问题

### Q: 模块路径错误怎么办？
A: 确保 `go.mod` 中的模块路径与 GitHub 仓库路径完全一致。

### Q: 用户报告 `go get` 失败？
A: 检查：
1. 仓库是否公开
2. 是否创建了版本标签
3. go.mod 路径是否正确

### Q: 如何撤回错误的版本？
A: 可以创建新版本修复问题，不建议删除已发布的标签。

---

🎉 **恭喜！** 按照这个指南，您的 HOTCOIN Go SDK 就可以成功发布到 GitHub，供全世界的 Go 开发者使用了！

如有问题，请查看 [Go Modules 官方文档](https://golang.org/ref/mod) 或 [GitHub 帮助文档](https://docs.github.com/)。