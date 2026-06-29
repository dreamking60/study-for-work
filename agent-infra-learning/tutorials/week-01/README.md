# Week 1：ReAct 深入 + 工具调用

**目标**：用现有工具快速搭建完整 Agent，搞清"一个 Agent 运行起来究竟需要哪些部件"。同时从源码层面理解 ReAct 循环和 Function Calling 协议。

**环境准备**
```bash
mkdir -p ~/agent-infra-learning/week1 && cd ~/agent-infra-learning/week1
python -m venv venv && source venv/bin/activate
pip install langchain langchain-openai openai python-dotenv httpx duckduckgo-search
```

---

## Day 1：裸调 LLM API，理解 Function Calling 底层

不依赖框架，用 `openai` SDK 手写一次 Function Calling。

- 读 [OpenAI Function Calling 文档](https://platform.openai.com/docs/guides/function-calling)（45min）
- 手写脚本：定义 `get_weather(city)` 工具，构造 `tools` 参数发给 ChatCompletion，处理 `tool_calls` 响应（1.5h）
- 打印完整请求/响应 JSON，理解 `role: tool`、`tool_call_id` 对应关系（30min）

**产出**：`day1_function_calling.py`

---

## Day 2：LangChain Agent + ReAct 循环

用 LangChain 搭建第一个完整 Agent，跑通 ReAct 循环。

- 读 [LangChain Agent 快速入门](https://python.langchain.com/docs/modules/agents/quick_start/)（30min）
- 实现：`create_react_agent` + `AgentExecutor`，注册 DuckDuckGo Search 和 Calculator 工具（1.5h）
- 设置 `verbose=True`，观察 Thought/Action/Observation 日志；设 `max_iterations=3` 观察截断（1h）

**产出**：`day2_langchain_agent.py`

---

## Day 3：读 ReAct 论文 + 手动实现

读原始论文，用代码把 ReAct 循环重新实现一遍。

- 读 [ReAct 论文](https://arxiv.org/abs/2210.03629) 核心 4-5 页（1h）
- 画 ReAct 循环状态机图（30min）
- 手动实现 ReAct：构造 system prompt，循环调用 LLM，解析 `Action:`，执行工具，填入 `Observation:`，直到 `Final Answer:`（1.5h）

**产出**：`day3_react_from_scratch.py`

---

## Day 4：拆解 AgentExecutor 源码

读 LangChain AgentExecutor 源码，理解框架抽象。

- 找到源码位置，定位 `agent.py` 和 `agent_executor.py`（30min）
- 读 `AgentExecutor._take_next_step()`，理解 AgentFinish vs AgentAction 处理（1.5h）
- 读 `AgentExecutor._call()`，理解循环整体设计（1h）

**产出**：核心调用链路图 + 3 个源码设计发现

---

## Day 5：自定义工具 + 错误处理

掌握工具开发的高级模式与异常处理。

- 用 `@tool` 装饰器写 3 个自定义工具（1h）
- 实现 `handle_tool_error=True`，测试异常恢复（45min）
- 读 `BaseTool` 的 `_run` / `_arun` 双接口设计（1h）

**产出**：`day5_custom_tools.py` + 源码笔记

---

## Day 6：系统架构图 + 知识沉淀

把一周所学提炼成可复用的知识体系。

- 画 Agent 系统架构图：从用户输入到最终输出的完整链路（1.5h）
- 列出核心抽象（至少 5 个）：LLM、PromptTemplate、Tool、Memory、Parser…（1h）
- 写 Markdown 笔记总结 ReAct、Function Calling、LangChain 抽象（1h）

**产出**：`architecture.md`

---

## Day 7：查漏补缺 + 下周预习

- 回顾 Day 1-6 代码，确保均可运行（1h）
- 读 [Lilian Weng Agent 综述](https://lilianweng.github.io/posts/2023-06-23-agent/) 前三节（1.5h）
- 搭建 Week 2 环境：`pip install chromadb langchain-chroma langchain-community tiktoken`（30min）

**产出**：代码整理完毕，Week 2 环境就绪
