package player

import "github.com/Icedroid/go-project-demo/server1/module"

type PlayerModule struct {
}

// 在初始化的时候将模块注册到module包
func init() {
	module.Player = PlayerModule{}
}

// 扣除铜钱
func (player PlayerModule) DecreaseCoins(num int) {
	println("DecreaseCoins(", num, ")")
}
