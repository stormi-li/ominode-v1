# Ominode 节点构建框架
**作者**: stormi-li  
**Email**: 2785782829@qq.com  


## 简介

**Ominode** 是一个帮助初学者快速在 Windows 系统上启动和配置 Redis 的工具。它能够自动生成 Redis 配置文件并启动 Redis 服务，简化了 Redis 在本地开发环境中的配置过程。

## 功能

- **生成 Redis 配置文件**：自动生成 Redis 的配置文件，便于快速启动 Redis 服务。
- **启动 Redis**：在指定的端口上启动 Redis 服务，支持直接运行在本地环境中。
## 教程
### 安装
```shell
go get github.com/stormi-li/ominode-v1
```
### 启动 Redis 
```go
package main

import "github.com/stormi-li/ominode"

func main() {
    // 在 node 目录下生成配置文件并以6379端口启动 Redis
    // 如果存在配置文件则直接启动
	ominode.StartRedis(6379, "node")
}
```