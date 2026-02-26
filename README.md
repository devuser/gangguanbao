# 钢管宝查询系统 (gangguanbao)

前后端分离的订单数据查询系统，用于统计本公司往年订单，支持按材质、规格、厂家等条件查询、统计分析及 Excel 导出。

**主要功能：** 订单查询（详情/编辑）、统计分析、材质/规格/厂家管理（详情/编辑）、新增/删除、导出 Excel

---

## 一、环境要求

| 软件 | 最低版本 | 用途 |
|------|----------|------|
| **Go** | 1.21+ | 后端运行环境 |
| **Node.js** | 18+ | 前端构建与运行 |
| **MySQL** | 5.0+ | 数据库 |
| **Make** | - | 可选，简化命令 |

---

## 二、环境安装（初学者）

### 2.1 Go

**macOS：** `brew install go`  
**Windows：** 从 [go.dev/dl](https://go.dev/dl/) 下载安装包  
**Linux：** `wget` 下载 tar.gz 解压至 `/usr/local`，并配置 PATH

### 2.2 Node.js

**推荐 nvm：**
```bash
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.0/install.sh | bash
nvm install 18 && nvm use 18
```

或从 [nodejs.org](https://nodejs.org/) 下载 LTS 版本。

### 2.3 MySQL

**macOS：** `brew install mysql@5.5 && brew services start mysql@5.5`  
**Docker（跨平台）：**
```bash
docker run -d --name mysql-gangguanbao -p 13306:3306 \
  -e MYSQL_ROOT_PASSWORD=root123 \
  -e MYSQL_DATABASE=gangguanbao_production \
  mysql:5.5
```

### 2.4 Make（可选）

**macOS：** 已预装  
**Windows：** `choco install make` 或使用 Git Bash

---

## 三、部署步骤

### 3.1 初始化数据库

**全新安装：**
```bash
mysql -h 127.0.0.1 -P 3306 -u root -p < sql/schema.sql
# Docker 端口 13306：-P 13306
```

**已有旧表（materials/specs/vendors/orders）需迁移到 biz_ 前缀：**
```bash
mysql -h 127.0.0.1 -P 3306 -u root -p < sql/migrate_biz_prefix.sql
```

### 3.2 配置后端

```bash
cd backend
cp config.example.yaml config.yaml
# 编辑 config.yaml，修改 database.host/port/user/password
```

### 3.3 安装依赖

```bash
# 后端
cd backend && go mod tidy
# 若慢：go env -w GOPROXY=https://goproxy.cn,direct

# 前端（支持 npm / pnpm）
cd frontend && npm install
# 若慢：npm config set registry https://registry.npmmirror.com
```

### 3.4 启动服务

**终端 1 - 后端：**
```bash
make run-backend
# 或：cd backend && go build -buildvcs=false -o bin/server ./cmd/server && ./bin/server
```

**终端 2 - 前端：**
```bash
make dev-frontend
# 或：cd frontend && npm run dev
```

### 3.5 访问

浏览器打开：**http://localhost:5173**

---

## 四、生产部署

1. **构建前端：** `make frontend` 或 `cd frontend && npm run build` → 输出在 `frontend/dist/`
2. **Nginx：** 托管 `dist/`，反向代理 `/api` 到 `http://127.0.0.1:8080`
3. **后端服务：** 用 systemd/supervisor 运行 `./bin/server`，`config.yaml` 中 `mode: release`

---

## 五、Makefile 命令

| 命令 | 说明 |
|------|------|
| `make backend` | 编译后端 |
| `make run-backend` | 运行后端 |
| `make frontend` | 构建前端（生产） |
| `make dev-frontend` | 格式化并启动前端开发服务器 |
| `make deps-backend` | 后端依赖（go mod tidy） |
| `make deps-frontend` | 前端依赖（npm/pnpm install） |
| `make test` | 运行测试 |
| `make fmt` | 格式化代码 |
| `make clean` | 清理构建产物 |

---

## 六、项目结构

```
gangguanbao/
├── Makefile
├── README.md              # 本文档（开发/部署）
├── README-SOP.md          # 用户操作手册
├── sql/
│   ├── schema.sql         # 建表及示例数据（biz_ 前缀）
│   └── migrate_biz_prefix.sql   # 旧表迁移脚本
├── backend/               # Go + Gin + GORM
│   ├── cmd/server/
│   ├── internal/{config,handler,model,repository}
│   └── config.yaml
└── frontend/              # Vue 3 + Vite + TypeScript
    ├── src/{api,components,views,router,assets}
    └── package.json
```

---

## 七、常见问题

**Q：`sql: unknown driver "mysql"`**  
A：`go get github.com/go-sql-driver/mysql` 后重新编译。

**Q：接口 404 或跨域**  
A：确认后端在 8080 端口，`vite.config.ts` 已配置 `/api` 代理。

**Q：数据库连接失败**  
A：检查 `config.yaml` 中 host/port/user/password，确认 MySQL 已启动。

**Q：表不存在或 Unknown table**  
A：确认已执行 `sql/schema.sql`（新库）或 `sql/migrate_biz_prefix.sql`（旧库迁移）。
