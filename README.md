# gin-k8s-api

## 【初始化】

### 拉取代码

```shell
git clone https://github.com/chingzio/gin-rdmg-service.git
```

### 创建数据库

```sql
CREATE DATABASE `gin-rdmg-service` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
```

### 修改配置文件
config/app.ini

| 配置分组     | 配置名称   | 配置描述           | 示例参数           | 备注          |
| ------------ | ---------- | ------------------ | ------------------ | ------------- |
| [app]        | AppName    | 项目名称           | gin-rdmg-service    |               |
| [app]        | Port       | 项目监听端口       | :9000              | 冒号不能省略  |
| [database]   | Db         | 数据库类型         | mysql              | 现只支持mysql |
| [database]   | DbHost     | 数据库地址         | 127.0.0.1          |               |
| [database]   | DbPort     | 数据库端口         | 3306               |               |
| [database]   | DbUser     | 数据库用户名       |                    |               |
| [database]   | DbPassWord | 数据库密码         |                    |               |
| [database]   | DbName     | 数据库库名         |                    |               |

### 下载依赖

```shell
go mod download
```

### 启动项目

```shell
go run main.go
```

