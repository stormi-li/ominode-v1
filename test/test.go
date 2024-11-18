package main

import "github.com/stormi-li/ominode-v1"

func main() {
	ominode.StartRedis(6379, "node")
}
