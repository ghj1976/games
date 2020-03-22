package tank90

import "github.com/ghj1976/games/tank90/tools"

// Bullet 子弹类
type Bullet struct {
	tools.RectangleCollisionObject
	Power        int    // 子弹的威力，值只能是1，2
	DestoryStone bool   // 是否具备摧毁石墙的能力
	Towards      string // 子弹飞行方向， 只能是 top bottom left right 四种
	IsFinish     bool   // 是否已经完成攻击了， true 就不用再做其他判断了
}
