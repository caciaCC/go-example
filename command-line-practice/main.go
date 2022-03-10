package main

import (
	"fmt"
	"github.com/alecthomas/kong"
)

func main() {
	var Cli struct {
		Play         bool  `optional:"" help:"运行play函数，适用于测试和开发"`
		Debug        bool  `optional:"" help:"启动debug模式"`
		SetAdmin     int64 `optional:"" xor:"c" help:"设置admin权限"`
		Version      bool  `optional:"" xor:"c" short:"v" help:"打印版本信息"`
		SyncBilibili bool  `optional:"" xor:"c" help:"同步b站帐号的关注，适用于更换或迁移b站帐号的时候"`
	}
	kong.Parse(&Cli)
	fmt.Println(Cli)
}
