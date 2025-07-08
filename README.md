# 🧠 AI Agent Tool Hub (Go)

## 🔍 Overview

This project is a **Go-based dynamic tool hub** designed to run in coordination with an **AI agent**. It provides a runtime where tools (executables, `.wasm`, and `.so` in the future) can be registered, discovered, and executed based on a high-level description.

---

## 📌 Key Features

- 🔧 Execute  tools via subprocess (CLI)
- 📦 Isolated tool execution using JSON input/output
- 📁 Dynamic tool registry based on description + name
- 🔒 Future support for `.wasm` and `.so`
- 🌐 Designed to scale in Kubernetes via NFS/shared volume

---

## 🗂️ Folder Structure

```text
.
├── bin/         # Built tool binaries (add, hash, etc.), for local testing
├── build/       # Compiled final application binary
├── tools/       # Source files for tools (Go, C, etc.), for local testing
├── internal/
│   └── runner/  # Tool runner logic (e.g., CLIRunTool, to run cli tool)
├── PythonAIAgent #python code for ai agent to test tool 
├── cmd/api      # project entry point
└── Makefile    # Makefile


```

## 🧠 How It Works

### 🔁 Current Flow(for local test)

- Tools are compiled and placed in the `bin/` directory.
- Each tool must:
    - Accept **JSON input** via `stdin`
    - Return **JSON output** via `stdout`

  ```go
  CLIRunTool(path, input)

### 🔁 Normal Current FLow
- Tools are uploaded to system with description
- Each tool must:
  - Accept **JSON input** via `stdin`
  - Return **JSON output** via `stdout`