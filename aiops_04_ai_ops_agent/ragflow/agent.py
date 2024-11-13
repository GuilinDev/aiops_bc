from typing import List, Dict
from langchain.chat_models import ChatOpenAI
from langchain.prompts import ChatPromptTemplate
from knowledge_base import KnowledgeBase


class RAGFlowAgent:
    def __init__(self, knowledge_base: KnowledgeBase, model_name: str = "gpt-4o-mini"):
        """
        初始化RAGFlow Agent
        
        Args:
            knowledge_base: 知识库实例
            model_name: OpenAI模型名称
        """
        self.knowledge_base = knowledge_base
        self.llm = ChatOpenAI(model_name=model_name)

        self.prompt_template = ChatPromptTemplate.from_template("""
            你是一个专业的运维工程师助手。基于以下运维知识库的内容，请回答用户的问题。
            如果知识库中没有相关信息，请明确说明。

            知识库内容:
            {context}

            用户问题: {question}
            
            请用专业、清晰的方式回答，并在适当的时候提供具体的命令或配置示例。
        """)

    def answer(self, question: str) -> Dict:
        """
        回答用户问题
        
        Args:
            question: 用户问题
            
        Returns:
            包含答案和参考文档的字典
        """
        # 搜索相关文档
        docs = self.knowledge_base.search(question)

        # 构建上下文
        context = "\n\n".join([doc["content"] for doc in docs])

        # 生成回答
        prompt = self.prompt_template.format(context=context, question=question)
        response = self.llm.predict(prompt)

        return {
            "answer": response,
            "references": docs
        }
