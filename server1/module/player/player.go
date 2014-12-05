package player

import "log"
import "github.com/funny/go-project-demo/server1/module"

type PlayerModule struct {
}

// 在初始化的时候将模块注册到module包
func init() {
	module.Player = PlayerModule{}
}

// 扣除铜钱
func (player PlayerModule) DecreaseCoins(num int) {
	log.Printf("DecreaseCoins(%d)", num)
}
