from typing import List, Dict
from langchain.embeddings import OpenAIEmbeddings
from langchain.vectorstores import Chroma
from langchain.text_splitter import CharacterTextSplitter
from langchain.document_loaders import DirectoryLoader, TextLoader
import os

class KnowledgeBase:
    def __init__(self, docs_dir: str = "docs", embedding_model: str = "text-embedding-ada-002"):
        """
        初始化知识库
        
        Args:
            docs_dir: 文档目录
            embedding_model: OpenAI embedding模型名称
        """
        self.docs_dir = docs_dir
        self.embeddings = OpenAIEmbeddings(model=embedding_model)
        self.vector_store = None
        
    def load_documents(self) -> None:
        """加载文档并创建向量存储"""
        # 加载文档
        loader = DirectoryLoader(self.docs_dir, glob="**/*.md", loader_cls=TextLoader)
        documents = loader.load()
        
        # 分割文档
        text_splitter = CharacterTextSplitter(
            chunk_size=1000,
            chunk_overlap=200
        )
        texts = text_splitter.split_documents(documents)
        
        # 创建向量存储
        self.vector_store = Chroma.from_documents(
            documents=texts,
            embedding=self.embeddings,
            persist_directory="./data/chroma_db"
        )
        
    def search(self, query: str, k: int = 3) -> List[Dict]:
        """
        搜索相关文档
        
        Args:
            query: 查询文本
            k: 返回结果数量
            
        Returns:
            相关文档列表
        """
        if not self.vector_store:
            raise ValueError("Knowledge base not initialized. Call load_documents() first.")
            
        results = self.vector_store.similarity_search_with_score(query, k=k)
        return [
            {
                "content": doc.page_content,
                "source": doc.metadata.get("source", ""),
                "score": score
            }
            for doc, score in results
        ] 
