# K8s Copilot

一个基于 Cobra 的智能 K8s 运维助手，集成了 OpenAI 来辅助 Kubernetes 运维工作。

## 功能特性

1. **基础命令**
   - `hello`: 基础的 hello world 命令
   - `world`: 示例子命令（已弃用，使用 hello 替代）

2. **分析功能**
   - `analyze event`: 分析 K8s 集群中的 Warning 事件，并使用 AI 提供解决方案

3. **ChatGPT 交互**
   - `ask chatgpt`: 与 ChatGPT 进行交互，支持：
     - 部署资源
     - 查询资源
     - 删除资源

## 快速开始

### 前置条件

1. Go 1.22 或更高版本
2. 可用的 Kubernetes 集群
3. OpenAI API Key

### 安装步骤

1. 克隆项目：
```bash
git clone <repository-url>
cd aiops_06_ai_ops_client_go_practice
```

2. 安装依赖：
```bash
go mod tidy
```

3. 设置环境变量：
```bash
export OPENAI_API_KEY_2="your-api-key"
```

4. 构建项目：
```bash
go build -o k8scopilot
```

### 使用示例

1. 分析集群事件：
```bash
./k8scopilot analyze event
```

2. 与 ChatGPT 交互：
```bash
./k8scopilot ask chatgpt
```

## 项目结构

```
.
├── cmd/                    # 命令行命令
│   ├── root.go            # 根命令
│   ├── analyze.go         # 分析命令
│   ├── event.go           # 事件分析
│   ├── ask.go             # 询问命令
│   └── chatgpt.go         # ChatGPT 交互
├── utils/                  # 工具函数
│   ├── client_go.go       # K8s 客户端
│   └── openai.go          # OpenAI 客户端
├── main.go                # 主程序入口
├── go.mod                 # Go 模块定义
└── README.md              # 项目文档
```

## 开发步骤

1. 初始化项目：
```bash
# 安装 cobra-cli
go install github.com/spf13/cobra-cli@latest

# 创建新项目
mkdir aiops_06_ai_ops_client_go_practice
cd aiops_06_ai_ops_client_go_practice
go mod init github.com/yourusername/k8scopilot

# 初始化 cobra 项目
cobra-cli init
```

2. 添加命令：
```bash
cobra-cli add hello
cobra-cli add world
cobra-cli add analyze
cobra-cli add event
cobra-cli add ask
cobra-cli add chatgpt
```

3. 实现功能：
   - 在 `utils/` 下实现基础功能
   - 在 `cmd/` 下实现命令行交互
   - 集成 OpenAI API
   - 实现 K8s 客户端功能

## 配置说明

1. Kubernetes 配置：
   - 默认使用 `~/.kube/config`
   - 可通过 `--kubeconfig` 指定配置文件
   - 使用 `--namespace` 指定命名空间

2. OpenAI 配置：
   - 需要设置 `OPENAI_API_KEY_2` 环境变量
   - 默认使用 GPT-4 模型

## 注意事项

1. 确保有足够的 K8s 集群权限
2. 保护好 OpenAI API Key
3. 注意 API 调用频率限制

## 待完成功能

1. 完善资源删除功能
2. 添加更多的事件分析能力
3. 改进 AI 提示词
4. 添加测试用例
