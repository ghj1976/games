package tank90

import (
	"log"
	"math"

	"github.com/ghj1976/games/tank90/images"
	"github.com/ghj1976/games/tank90/tools"
	"github.com/hajimehoshi/ebiten/v2"
)

// Tank 不论是敌人，还是玩家通用的坦克处理类
type Tank struct {
	name               string  // 坦克唯一编号
	TypeName           string  // 有下面七种  user1、 user2 、 enemy1 、 enemy2、 enemy3 enemy4 enemy5
	Level              int     // 坦克的级别，所有坦克都有4个级别 1，2，3，4
	towards            string  // 朝向 只能是 top bottom left right 四种
	MapX               int     // 在地图上的坐标位置
	MapY               int     // 在地图上的坐标位置
	NextIMG            bool    // 是否要使用第二张图
	speed              float64 // 坦克速度
	BulletSpeed        float64 // 炮弹速度
	BulletPower        int     // 炮弹威力， 1或2
	BulletDestoryStone bool    // 炮弹是否能击毁石墙
	Point              int     // 敌方坦克的积分
	PreCX              int     //  上次行动时的屏幕坐标， 用于碰撞后的回退
	PreCY              int     //  上次行动时的屏幕坐标， 用于碰撞后的回退
	IsMove             bool    //  是否在运动 。 true 在移动， false 停止
	Shipable           bool    // 是否可以过水域
	countdownFrames    float64 // 倒计时帧数， 当帧数小于1的时候，触发一次移动
	cx                 int     // 图像的中心点，屏幕坐标点，用于碰撞检测用
	cy                 int     // 图像的中心点，屏幕坐标点，用于碰撞检测用
	width              int     // 宽度，用于碰撞检测用
	height             int     // 高度，用于碰撞检测用
}

// NewTank 初始化一个坦克
// mapx, mapy 地图左上角坐标，不是像素坐标,
func NewTank(name, typeName string, level int, toward string, mapx, mapy int) *Tank {
	t := &Tank{}
	t.name = name
	t.Shipable = false
	t.MapX = mapx
	t.MapY = mapy
	if typeName == "user1" || typeName == "user2" || typeName == "enemy1" || typeName == "enemy2" || typeName == "enemy3" || typeName == "enemy4" || typeName == "enemy5" {
		t.TypeName = typeName
	} else {
		log.Fatalf("Tank TypeName error %s\r\n", typeName)
	}

	if level > 0 && level <= 4 {
		t.Level = level
	} else {
		log.Fatalf("Tank Level error %d\r\n", level)
	}

	t.SetTowards(toward)

	// 碰撞判断需要的数据
	t.cx = (mapx + 1) * images.TileSize
	t.cy = (mapy + 1) * images.TileSize
	t.width = images.TankSize
	t.height = images.TankSize
	t.PreCX, t.PreCY = t.cx, t.cy

	// 玩家坦克的速度和炮弹的威力设置
	if t.TypeName == "user1" || t.TypeName == "user2" {
		t.Point = 0
		switch t.Level {
		case 1:
			t.speed = 3.2
			t.BulletSpeed = 8.8
			t.BulletPower = 1
			t.BulletDestoryStone = false
		case 2:
			t.speed = 3.8
			t.BulletSpeed = 12.8
			t.BulletPower = 1
			t.BulletDestoryStone = false
		case 3:
			t.speed = 4.2
			t.BulletSpeed = 12.8
			t.BulletPower = 2
			t.BulletDestoryStone = false
		case 4:
			t.speed = 3.4
			t.BulletSpeed = 11.2
			t.BulletPower = 2
			t.BulletDestoryStone = true
		}
	} else {
		switch t.TypeName {
		case "enemy1":
			t.speed = 2.6
			t.BulletSpeed = 8.0
			t.BulletPower = 1
			t.BulletDestoryStone = false
			t.Point = 100
		case "enemy2":
			t.speed = 4.6
			t.BulletSpeed = 10.0
			t.BulletPower = 1
			t.BulletDestoryStone = false
			t.Point = 200
		case "enemy3":
			t.speed = 3.2
			t.BulletSpeed = 11.2
			t.BulletPower = 2
			t.BulletDestoryStone = false
			t.Point = 300
		case "enemy4":
			t.speed = 3.6
			t.BulletSpeed = 11.6
			t.BulletPower = 1
			t.BulletDestoryStone = false
			t.Point = 400
		case "enemy5":
			t.speed = 3.0
			t.BulletSpeed = 9.6
			t.BulletPower = 1
			t.BulletDestoryStone = true
			t.Point = 500
		}
	}
	t.countdownFrames = 0.0

	return t
}

// GetName 获得坦克的唯一编号
func (t *Tank) GetName() string {
	return t.name
}

// GetSpeed 获得坦克的速度
func (t *Tank) GetSpeed() float64 {
	return t.speed
}

// GetTowards 获得坦克当前朝向
func (t *Tank) GetTowards() string {
	return t.towards
}

// GetMapXY 获得坦克当前在地图的位置
func (t *Tank) GetMapXY() (mapx, mapy int) {
	return t.MapX, t.MapY
}

// SetPreCXY 缓存坦克的位置
func (t *Tank) SetPreCXY(x, y int) {
	t.PreCX = x
	t.PreCY = y
}

// SetTowards 设置坦克朝向
func (t *Tank) SetTowards(to string) {
	if to == "top" || to == "bottom" || to == "left" || to == "right" {
		t.towards = to
	} else {
		log.Fatalf("Tank Towards error %s\r\n", to)
	}
}

// GetCentorPositionAndSize 获得中心点位置及长宽, 碰撞判断用
func (t *Tank) GetCentorPositionAndSize() (x, y, w, h int) {
	return t.cx, t.cy, t.width, t.height
}

// 加载坦克图片
func (t *Tank) getTankImage() *ebiten.Image {

	enemyColor := ""
	switch t.Level {
	case 1:
		enemyColor = "silver"
	case 2:
		enemyColor = "green"
	case 3:
		enemyColor = "golden"
	case 4:
		enemyColor = "red"
	}

	switch t.TypeName {
	case "user1":
		return images.GetTankImage(t.Level, "golden", t.towards, t.NextIMG)
	case "user2":
		return images.GetTankImage(t.Level, "green", t.towards, t.NextIMG)
	case "enemy1":
		return images.GetTankImage(5, enemyColor, t.towards, t.NextIMG)
	case "enemy2":
		return images.GetTankImage(6, enemyColor, t.towards, t.NextIMG)
	case "enemy3":
		return images.GetTankImage(7, enemyColor, t.towards, t.NextIMG)
	case "enemy4":
		return images.GetTankImage(8, enemyColor, t.towards, t.NextIMG)
	default:
		return nil
	}
}

// Draw 画坦克
func (t *Tank) Draw(mapImage *ebiten.Image) {

	// if t.TypeName == "enemy1" {
	// 	log.Printf("enemy1 cx %d,cy %d, mx %d, my %d, toward %s", t.CX, t.CY, t.MapX, t.MapY, t.Towards)
	// }
	opts2 := &ebiten.DrawImageOptions{}
	opts2.GeoM.Translate(float64(t.cx-t.width/2), float64(t.cy-t.height/2))
	mapImage.DrawImage(t.getTankImage(), opts2) // 游戏地图区域绘制

}

// Update 定时更新坦克数据
// count 定时器参数
// 返回这次坦克的移动距离
func (t *Tank) Update(count int) int {
	distance := 0

	if t.IsMove {
		// 换坦克图
		if count%(ebiten.MaxTPS()/12) == 0 {
			// log.Printf("mx:%d,my:%d", u.MapX, u.MapY)
			// 更新动画图片
			t.NextIMG = !t.NextIMG
		}

		t.countdownFrames = t.countdownFrames - 1.0

		if t.countdownFrames < 1.0 { // 该触发移动了
			// 计算下次应该多少桢后触发移动
			// 正常是一秒60桢的节奏update
			t.countdownFrames += float64(ebiten.MaxTPS()) / (t.speed * images.TileSize)

			t.PreCX, t.PreCY = t.cx, t.cy
			// 每次移动距离是1像素
			distance = 1
			// 发生移动后的位置
			switch t.towards {
			case "top":
				t.cy -= distance
			case "bottom":
				t.cy += distance
			case "left":
				t.cx -= distance
			case "right":
				t.cx += distance
			}

			// golang向上取整、向下取整和四舍五入
			// https://studygolang.com/articles/12965
			// math.Floor()  向下取整
			// math.Round()  四舍五入
			// math.Ceil() 向上取整
			t.MapX = int(math.Floor(float64(t.cx-images.TileSize) / images.TileSize)) // cx 是中心位置，需要变成左上角位置
			t.MapY = int(math.Floor(float64(t.cy-images.TileSize) / images.TileSize)) // cy 是中心位置，需要变成左上角位置
		}

	}
	return distance
}

// Move 坦克 开始移动
func (t *Tank) Move() {
	t.IsMove = true
}

// Stop 坦克停止移动
func (t *Tank) Stop() {
	t.IsMove = false
}

// Turn 坦克转向那个方向
// 转向需要修正位置，避免出现被几像素卡住的问题
func (t *Tank) Turn(toward string) {

	if t.towards == toward || len(toward) == 0 {
		// 没有变化方向
		return
	}

	log.Printf("old %s new %s", t.towards, toward)
	t.towards = toward

	// 从之间撞上物体的坐标返回没有撞上之前的。
	t.ResetPosition()

	// 误差在10之内的，做修正, 避免不好控制拐弯
	modX := t.cx % images.TileSize
	if modX != 0 {
		if modX <= 3 {
			t.cx = int(t.cx/images.TileSize) * images.TileSize
		} else if modX >= 5 {
			t.cx = (int(t.cx/images.TileSize) + 1) * images.TileSize
		}
	}

	modY := t.cy % images.TileSize
	if modY != 0 {
		if modY <= 3 {
			t.cy = int(t.cy/images.TileSize) * images.TileSize
		} else if modY >= 5 {
			t.cy = (int(t.cy/images.TileSize) + 1) * images.TileSize
		}
	}
}

// ResetPosition 返回上一步的位置，用于碰撞后的处理
func (t *Tank) ResetPosition() {
	t.cx = t.PreCX
	t.cy = t.PreCY
}

// ITankMeeting 要做坦克相遇判断必须实现的方法接口
type ITankMeeting interface {
	GetName() string
	GetSpeed() float64
	GetTowards() string
	GetMapXY() (mapx, mapy int)
	ResetPosition()
	SetPreCXY(x, y int)
	tools.ICollision
}
