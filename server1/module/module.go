package module

// 玩家模块
type PlayerModule interface {
	// 扣除铜钱
	DecreaseCoins(num int)
}

// 物品模块
type ItemModule interface {
	// 购买物品
	BuyItem()
}

// 这些是接口的具体实现，等待外部主动注册进来，
// 这样module包永远是被引用的，不会出现递归引用问题。
var (
	Player PlayerModule
	Item   ItemModule
)
