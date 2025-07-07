# Gin 博客 API 项目

这是一个基于 Go 语言 Gin 框架开发的博客 API 系统，支持用户认证、文章管理和评论功能。


## 运行环境要求

- **Go 版本**: 1.24.4 或更高版本
- **数据库**: MySQL 5.7


## 依赖安装步骤

### 1. 安装 Go 环境

确保已安装 Go 1.24.4 或更高版本：

```bash
go version
```

### 2. 克隆项目

```bash
git clone <repository-url>
cd blog
```

### 3. 安装项目依赖

```bash
go mod download
```

或者

```bash
go mod tidy
```

### 4. 安装 MySQL 数据库

确保 MySQL 已安装并运行，默认配置：
- 主机：127.0.0.1:3306
- 用户名：root
- 密码：root
- 数据库：test_blog

## 配置说明

### 数据库配置

编辑 `config/config.yaml` 文件，根据你的实际环境修改数据库连接信息：

```yaml
server:
  port: 8080                  # 服务器端口

database:
  driver: mysql
  dsn: root:root@tcp(127.0.0.1:3306)/test_blog?charset=utf8mb4&parseTime=True

jwt:
  secret: dfuhsudhfysgd        # JWT 密钥（生产环境请修改）
  expire_hours: 168            # Token 过期时间（小时）
```

### 创建数据库

在 MySQL 中创建数据库：

```sql
CREATE DATABASE test_blog CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

## 启动方式

### 1. 开发环境启动

```bash
go run main.go
```

### 2. 编译后启动

```bash
# 编译
go build -o blog.exe main.go

# 运行
./blog.exe
```

### 3. 启动成功

当看到以下输出时，表示启动成功：

```
配置加载成功：config/config.yaml
服务器启动在端口 :8080
```

## 测试说明

### 使用 Postman 测试

项目在 `result/` 文件夹中提供了完整的 API 测试集合：

1. **测试文件位置**: `result/测试.postman_collection.json`
2. **导入方式**: 
   - 打开 Postman
   - 点击 "Import"
   - 选择 `result/测试.postman_collection.json` 文件
   - 导入完成后即可看到所有 API 测试用例

3. **测试用例包含**:
   - 用户注册和登录测试
   - 文章的增删改查测试
   - 评论功能测试
   - 认证权限测试