# gin-k8s-api

## 【初始化】

### 拉取代码

```shell
git clone https://github.com/chingzio/gin-k8s-api.git
```

### 创建数据库

```sql
CREATE DATABASE `gin-k8s-db-test` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
```

### 获取kubernetes集群config文件

```shell
# kubernetes集群master
cat .kube/config
# 复制配置文件至本地，在项目配置文件中会使用此文件保存在本地的路径
```

### 修改配置文件

| 配置分组     | 配置名称   | 配置描述           | 示例参数           | 备注          |
| ------------ | ---------- | ------------------ | ------------------ | ------------- |
| [app]        | AppName    | 项目名称           | gin-k8s-api        |               |
| [app]        | Port       | 项目监听端口       | :9000              | 冒号不能省略  |
| [database]   | Db         | 数据库类型         | mysql              | 现只支持mysql |
| [database]   | DbHost     | 数据库地址         | 127.0.0.1          |               |
| [database]   | DbPort     | 数据库端口         | 3306               |               |
| [database]   | DbUser     | 数据库用户名       |                    |               |
| [database]   | DbPassWord | 数据库密码         |                    |               |
| [database]   | DbName     | 数据库库名         |                    |               |
| [kubernetes] | ConfigPath | k8s config文件路径 | /root/.kube/config | 绝对路径      |

### 下载依赖

```shell
go mod download
```

### 启动项目

```shell
go run main.go
```

