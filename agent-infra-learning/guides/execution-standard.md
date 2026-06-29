# 执行规范

## 编码规范
- 每个 Python 文件顶部有模块用途注释
- 使用 `if __name__ == "__main__":` 入口
- 关键函数有类型注解
- API Key 等敏感信息通过环境变量或 `.env` 读取
- 不提交 `.env` 文件到 git

## 项目结构规范（每个项目）
```
project-name/
├── main.py           # 入口
├── requirements.txt  # 依赖
├── README.md         # 项目说明
├── agent/            # Agent 核心逻辑
│   ├── runtime.py    # 运行时循环
│   ├── tools.py      # 工具注册与执行
│   └── memory.py     # 记忆系统
├── llm/              # LLM 调用封装
│   └── client.py
└── tests/            # 测试
    └── test_*.py
```

## 学习记录规范
- 每天一个 Markdown 文件：`report/weekN/dayN.md`
- 格式：`# Day N：标题` → `## 做了什么` → `## 关键发现` → `## 代码片段` → `## 问题与思考`
- 每周汇总：`report/weekN/weekly-summary.md`

## 提交规范
- 每周一个 commit：`weekN: <本周完成的关键内容>`
- 不包含临时文件（`__pycache__`、`.venv`、`.env`）
