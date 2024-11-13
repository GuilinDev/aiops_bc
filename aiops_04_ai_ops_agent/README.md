# RAGFlow - 智能运维知识库系统

RAGFlow 是一个基于 RAG (Retrieval Augmented Generation) 的智能运维知识库系统，它能够帮助运维团队更好地管理和利用内部运维知识。

## 核心特性

1. **智能文档检索**
   - 基于向量数据库的相似度搜索
   - 支持多种文档格式（Markdown、Text等）
   - 智能文档分块和向量化

2. **上下文感知回答**
   - 基于检索到的相关文档生成答案
   - 保持答案的准确性和可溯源性
   - 提供参考文档来源

3. **运维知识库管理**
   - 文档自动加载和更新
   - 支持多级目录结构
   - 版本控制和变更追踪

## 快速开始

1. 安装依赖：
```bash
pip install langchain openai chromadb
```

2. 设置环境变量：
```bash
export OPENAI_API_KEY="your-api-key"
```

3. 使用示例：
```python
from ragflow.knowledge_base import KnowledgeBase
from ragflow.agent import RAGFlowAgent

# 初始化知识库
kb = KnowledgeBase(docs_dir="./docs")
kb.load_documents()

# 创建Agent
agent = RAGFlowAgent(kb)

# 查询示例
response = agent.answer("如何排查K8s Pod启动失败的问题？")
print(response["answer"])
print("\n参考文档:")
for ref in response["references"]:
    print(f"- {ref['source']} (相关度: {ref['score']})")
```

## 文档结构

建议的文档组织结构：

```
docs/
├── kubernetes/
│   ├── troubleshooting.md
│   └── best-practices.md
├── monitoring/
│   ├── prometheus.md
│   └── grafana.md
└── deployment/
    ├── helm.md
    └── gitlab-ci.md
```

## 最佳实践

1. **文档编写规范**
   - 使用清晰的标题和层级结构
   - 包含具体的命令和配置示例
   - 添加故障排查步骤和解决方案

2. **知识库维护**
   - 定期更新和审查文档
   - 记录实际案例和解决方案
   - 保持文档的一致性和准确性

3. **使用建议**
   - 优先搜索常见问题
   - 参考相似案例的解决方案
   - 及时补充新的知识点

## 示例文档

创建一个示例故障排查文档：
