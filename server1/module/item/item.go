package item

import "github.com/funny/go-project-demo/server1/module"

type ItemModule struct {
}

// 在初始化的时候将模块注册到module包
func init() {
	module.Item = ItemModule{}
}

func (item ItemModule) BuyItem() {
	module.Player.DecreaseCoins(100)
}
