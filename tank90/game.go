package tank90

import (
	"image/color"
	"log"
	"math/rand"
	"reflect"

	"github.com/ghj1976/games/tank90/images"
	"github.com/ghj1976/games/tank90/resources"
	"github.com/ghj1976/games/tank90/tools"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	// ScreenWidth 屏幕大小，宽度
	ScreenWidth = 350
	// ScreenHeight 屏幕大小，高度
	ScreenHeight = 270
)

// Game 游戏逻辑
type Game struct {
	tmap     *Map
	mapImage *ebiten.Image // 地图图片缓存。 提高性能考虑，不用每次都绘制图片
	user1    *User         // 玩家一
	user2    *User         // 玩家二
	count    int           // 控制刷新速度的值
	enemyArr []*EnemyBot   // 所有敌人
}

// NewGame generates a new Game object.
func NewGame() (*Game, error) {
	// seed := time.Now().Unix()
	seed := int64(1583566915)
	log.Printf("当前随机种子: %d", seed)
	rand.Seed(seed)

	images.InitGameIMG()

	g := &Game{}
	g.count = 0
	g.tmap = NewMap(1)
	g.user1 = NewUser("user1")
	g.enemyArr = []*EnemyBot{NewEnemyBot("e1", "enemy1", 1), NewEnemyBot("e2", "enemy2", 1)}

	return g, nil
}

// Draw 绘图
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{132, 132, 132, 255})

	// 建立一個 728x728 的新画布,作为map区域
	square := ebiten.NewImage(208, 208)
	square.Fill(color.Black)

	// map 区域的元素绘制
	g.tmap.Draw(square)

	// 我方坦克绘制
	g.user1.Draw(square)

	// 敌人绘制
	i := 0
	for _, enemy := range g.enemyArr {

		if enemy.aStarPP != nil {
			colour := color.RGBA{34, 49, 255, 255}
			if i%2 == 0 {
				colour = color.RGBA{34, 139, 34, 255}
			}
			log.Println("画去目标路径。")
			// 调试时绘出路径
			for pp := enemy.aStarPP; pp != nil; pp = pp.Parent {
				DrawDebugInfo(square, pp.Point.Col, pp.Point.Row, colour)
			}
		}
		i++

		enemy.Draw(square)

	}

	// 炮弹绘制

	// 森林绘制
	g.tmap.GrassDraw(square)

	// 把map区域画在地图上。
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(ScreenWidth/2-104, ScreenHeight/2-104)
	screen.DrawImage(square, opts) // 游戏地图区域绘制

	drawReferenceLine(screen)
}

// Update updates the current game state.
func (g *Game) Update() error {

	// int 最大  2147483647
	if g.count > 2147483640 {
		g.count = 0
	}

	// 玩家情况更新
	g.user1.Update()

	if g.CanMove(g.user1) {
		g.user1.Move(g.count)
	} else {
		// 不能移动时，要暂停
		g.user1.IsMove = false
		g.user1.Move(g.count)
	}

	// 敌人 AI 坦克相关逻辑更新
	for _, enemy := range g.enemyArr {
		if g.CanMove(enemy) {
			b := enemy.Update(g.count)
			if b { // 需要重算一个命令集合
				enemy.ResetWanderTargetActionList(g.tmap.CloneMap())
				// g.tmap.SetWanderTargetActionList(enemy)
			}
		}

	}

	g.count++
	return nil
}

// drawReferenceLine  绘制调试用的坐标参考线
func drawReferenceLine(screen *ebiten.Image) {
	ebitenutil.DrawLine(screen, ScreenWidth/2-112, ScreenHeight/2-88, ScreenWidth/2+112, ScreenHeight/2-88, color.White)
	ebitenutil.DrawLine(screen, ScreenWidth/2-112, ScreenHeight/2-72, ScreenWidth/2+112, ScreenHeight/2-72, color.White)
	ebitenutil.DrawLine(screen, ScreenWidth/2-112, ScreenHeight/2-56, ScreenWidth/2+112, ScreenHeight/2-56, color.White)
	ebitenutil.DrawLine(screen, ScreenWidth/2-112, ScreenHeight/2-40, ScreenWidth/2+112, ScreenHeight/2-40, color.White)
	ebitenutil.DrawLine(screen, ScreenWidth/2-112, ScreenHeight/2-24, ScreenWidth/2+112, ScreenHeight/2-24, color.White)
	ebitenutil.DrawLine(screen, ScreenWidth/2-112, ScreenHeight/2-8, ScreenWidth/2+112, ScreenHeight/2-8, color.White)
	ebitenutil.DrawLine(screen, ScreenWidth/2-112, ScreenHeight/2+8, ScreenWidth/2+112, ScreenHeight/2+8, color.White)
	ebitenutil.DrawLine(screen, ScreenWidth/2-112, ScreenHeight/2+24, ScreenWidth/2+112, ScreenHeight/2+24, color.White)
	ebitenutil.DrawLine(screen, ScreenWidth/2-112, ScreenHeight/2+40, ScreenWidth/2+112, ScreenHeight/2+40, color.White)
	ebitenutil.DrawLine(screen, ScreenWidth/2-112, ScreenHeight/2+56, ScreenWidth/2+112, ScreenHeight/2+56, color.White)
	ebitenutil.DrawLine(screen, ScreenWidth/2-112, ScreenHeight/2+72, ScreenWidth/2+112, ScreenHeight/2+72, color.White)
	ebitenutil.DrawLine(screen, ScreenWidth/2-112, ScreenHeight/2+88, ScreenWidth/2+112, ScreenHeight/2+88, color.White)

	text.Draw(screen, "y0", resources.GetFont("setting"), ScreenWidth/2+112, ScreenHeight/2-96, color.RGBA{0, 0, 205, 255})
	text.Draw(screen, "y2", resources.GetFont("setting"), ScreenWidth/2+112, ScreenHeight/2-80, color.RGBA{0, 0, 205, 255})
	text.Draw(screen, "y4", resources.GetFont("setting"), ScreenWidth/2+112, ScreenHeight/2-64, color.RGBA{0, 0, 205, 255})
	text.Draw(screen, "y6", resources.GetFont("setting"), ScreenWidth/2+112, ScreenHeight/2-48, color.RGBA{0, 0, 205, 255})
	text.Draw(screen, "y8", resources.GetFont("setting"), ScreenWidth/2+112, ScreenHeight/2-32, color.RGBA{0, 0, 205, 255})
	text.Draw(screen, "y10", resources.GetFont("setting"), ScreenWidth/2+112, ScreenHeight/2-16, color.RGBA{0, 0, 205, 255})
	text.Draw(screen, "y12", resources.GetFont("setting"), ScreenWidth/2+112, ScreenHeight/2, color.RGBA{0, 0, 205, 255})
	text.Draw(screen, "y14", resources.GetFont("setting"), ScreenWidth/2+112, ScreenHeight/2+16, color.RGBA{0, 0, 205, 255})
	text.Draw(screen, "y16", resources.GetFont("setting"), ScreenWidth/2+112, ScreenHeight/2+32, color.RGBA{0, 0, 205, 255})
	text.Draw(screen, "y18", resources.GetFont("setting"), ScreenWidth/2+112, ScreenHeight/2+48, color.RGBA{0, 0, 205, 255})
	text.Draw(screen, "y20", resources.GetFont("setting"), ScreenWidth/2+112, ScreenHeight/2+64, color.RGBA{0, 0, 205, 255})
	text.Draw(screen, "y22", resources.GetFont("setting"), ScreenWidth/2+112, ScreenHeight/2+80, color.RGBA{0, 0, 205, 255})
	text.Draw(screen, "y24", resources.GetFont("setting"), ScreenWidth/2+112, ScreenHeight/2+96, color.RGBA{0, 0, 205, 255})

	ebitenutil.DrawLine(screen, ScreenWidth/2-88, ScreenHeight/2-112, ScreenWidth/2-88, ScreenHeight/2+112, color.White)
	ebitenutil.DrawLine(screen, ScreenWidth/2-72, ScreenHeight/2-112, ScreenWidth/2-72, ScreenHeight/2+112, color.White)
	ebitenutil.DrawLine(screen, ScreenWidth/2-56, ScreenHeight/2-112, ScreenWidth/2-56, ScreenHeight/2+112, color.White)
	ebitenutil.DrawLine(screen, ScreenWidth/2-40, ScreenHeight/2-112, ScreenWidth/2-40, ScreenHeight/2+112, color.White)
	ebitenutil.DrawLine(screen, ScreenWidth/2-24, ScreenHeight/2-112, ScreenWidth/2-24, ScreenHeight/2+112, color.White)
	ebitenutil.DrawLine(screen, ScreenWidth/2-8, ScreenHeight/2-112, ScreenWidth/2-8, ScreenHeight/2+112, color.White)
	ebitenutil.DrawLine(screen, ScreenWidth/2+8, ScreenHeight/2-112, ScreenWidth/2+8, ScreenHeight/2+112, color.White)
	ebitenutil.DrawLine(screen, ScreenWidth/2+24, ScreenHeight/2-112, ScreenWidth/2+24, ScreenHeight/2+112, color.White)
	ebitenutil.DrawLine(screen, ScreenWidth/2+40, ScreenHeight/2-112, ScreenWidth/2+40, ScreenHeight/2+112, color.White)
	ebitenutil.DrawLine(screen, ScreenWidth/2+56, ScreenHeight/2-112, ScreenWidth/2+56, ScreenHeight/2+112, color.White)
	ebitenutil.DrawLine(screen, ScreenWidth/2+72, ScreenHeight/2-112, ScreenWidth/2+72, ScreenHeight/2+112, color.White)
	ebitenutil.DrawLine(screen, ScreenWidth/2+88, ScreenHeight/2-112, ScreenWidth/2+88, ScreenHeight/2+112, color.White)

	text.Draw(screen, "x0", resources.GetFont("setting"), ScreenWidth/2-104, ScreenHeight/2+122, color.RGBA{0, 0, 205, 255})
	text.Draw(screen, "x4", resources.GetFont("setting"), ScreenWidth/2-72, ScreenHeight/2+122, color.RGBA{0, 0, 205, 255})
	text.Draw(screen, "x8", resources.GetFont("setting"), ScreenWidth/2-40, ScreenHeight/2+122, color.RGBA{0, 0, 205, 255})
	text.Draw(screen, "x12", resources.GetFont("setting"), ScreenWidth/2-10, ScreenHeight/2+122, color.RGBA{0, 0, 205, 255})
	text.Draw(screen, "x16", resources.GetFont("setting"), ScreenWidth/2+20, ScreenHeight/2+122, color.RGBA{0, 0, 205, 255})
	text.Draw(screen, "x20", resources.GetFont("setting"), ScreenWidth/2+52, ScreenHeight/2+122, color.RGBA{0, 0, 205, 255})
	text.Draw(screen, "x24", resources.GetFont("setting"), ScreenWidth/2+88, ScreenHeight/2+122, color.RGBA{0, 0, 205, 255})

}

// GetCurrMapCopy 获得当前地图+ 敌人坦克+ 玩家坦克的地图镜像copy
// 用于静态寻路分析用。
func (g *Game) GetCurrMapCopy() [26][26]int {
	mmap := g.tmap.CloneMap()

	if g.user1 != nil {
		mmap[g.user1.MapY][g.user1.MapX] = 10
		mmap[g.user1.MapY+1][g.user1.MapX] = 10
		mmap[g.user1.MapY][g.user1.MapX+1] = 10
		mmap[g.user1.MapY+1][g.user1.MapX+1] = 10
	}
	if g.user2 != nil {
		mmap[g.user2.MapY][g.user2.MapX] = 10
		mmap[g.user2.MapY+1][g.user2.MapX] = 10
		mmap[g.user2.MapY][g.user2.MapX+1] = 10
		mmap[g.user2.MapY+1][g.user2.MapX+1] = 10
	}
	for _, enemy := range g.enemyArr {
		mmap[enemy.MapY][enemy.MapX] = 20
		mmap[enemy.MapY+1][enemy.MapX] = 20
		mmap[enemy.MapY][enemy.MapX+1] = 20
		mmap[enemy.MapY+1][enemy.MapX+1] = 20
	}

	return mmap
}

// CanMove 判断一辆坦克是否能移动
func (g *Game) CanMove(t ITankMeeting) bool {
	// 地图检查
	if !g.tmap.CanMove(t) {
		return false
	}

	// 跟其他坦克检查是否碰撞
	if t.GetName() != g.user1.GetName() {
		if tools.CheckCollision(t, g.user1) {
			log.Printf("%s 相撞 %s", t.GetName(), g.user1.GetName())
			g.TankMeeting(t, g.user1)
			return false
		}
	}

	if g.user2 != nil && t.GetName() != g.user2.GetName() {
		if tools.CheckCollision(t, g.user2) {
			log.Printf("%s 相撞 %s", t.GetName(), g.user1.GetName())
			g.TankMeeting(t, g.user2)
			return false
		}

	}

	for _, en := range g.enemyArr {
		if en.GetName() != t.GetName() {
			if tools.CheckCollision(t, en) {
				log.Printf("%s 相撞 %s", t.GetName(), en.GetName())
				g.TankMeeting(t, en)
				return false
			}
		}
	}

	return true
}

// TankMeeting 两辆坦克相遇的处理逻辑
func (g *Game) TankMeeting(t1, t2 ITankMeeting) {
	// 复位到没有碰撞前的位置
	t1.ResetPosition()
	t2.ResetPosition()

	// 判断两辆坦克的走向，看是否可以通过一台暂停来解决
	conflict, selectTank := towardsConflictSelectTank(t1, t2)
	if conflict {
		log.Printf("selecttank:%v", selectTank)

		// 方向冲突, selectTank 重新计算新目标
		eb, b := checkEnemyBot(selectTank)
		log.Println(b)
		log.Println(eb)
		if b {
			log.Printf("%s 重算路径", eb.GetName())
			log.Println("开始重算路径")
			//  重新算一条路径， 这个map需要有对应坦克的信息。
			gmap := g.GetCurrMapCopy()           // 获得当前地图
			eb.ResetWanderTargetActionList(gmap) // 重算路径
		}

	} else {
		// 方向不冲突， selectTank 暂停1秒
		eb, b := checkEnemyBot(selectTank)
		if b {
			log.Printf("%s 让路", eb.GetName())
			log.Println("开始让路等待")
			eb.CommandParseFrames(60)
		}

	}

}

func checkEnemyBot(v interface{}) (eb *EnemyBot, b bool) {
	log.Println(reflect.TypeOf(v))
	eb, b = v.(*EnemyBot)

	return eb, b
	// return nil, false
}

// towardsConflict 两个走势是否冲突， 只有相对的方向才算冲突
func towardsConflictSelectTank(t1, t2 ITankMeeting) (conflict bool, selectTank ITankMeeting) {
	// 机器人和 玩家相遇， 机器人规避
	_, b := t1.(*User)
	if b {
		return true, t2
	}
	_, b = t2.(*User)
	if b {
		return true, t1
	}

	if t1.GetTowards() == "top" && t2.GetTowards() == "bottom" {
		// 如果方向冲突，速度快的重算目标，避免再次相撞
		if t1.GetSpeed() > t2.GetSpeed() {
			return true, t1
		}
		if t1.GetSpeed() < t2.GetSpeed() {
			return true, t2
		}
		// 速度一样，随机选一个
		if rand.Intn(2) == 0 {
			return true, t1
		}
		return true, t2
	} else if t1.GetTowards() == "bottom" && t2.GetTowards() == "top" {

		// 如果方向冲突，速度快的重算目标，避免再次相撞
		if t1.GetSpeed() > t2.GetSpeed() {
			return true, t1
		}
		if t1.GetSpeed() < t2.GetSpeed() {
			return true, t2
		}
		// 速度一样，随机选一个
		if rand.Intn(2) == 0 {
			return true, t1
		}
		return true, t2
	} else if t1.GetTowards() == "left" && t2.GetTowards() == "right" {
		// 如果方向冲突，速度快的重算目标，避免再次相撞
		if t1.GetSpeed() > t2.GetSpeed() {
			return true, t1
		}
		if t1.GetSpeed() < t2.GetSpeed() {
			return true, t2
		}
		// 速度一样，随机选一个
		if rand.Intn(2) == 0 {
			return true, t1
		}
		return true, t2
	} else if t1.GetTowards() == "right" && t2.GetTowards() == "left" {
		// 如果方向冲突，速度快的重算目标，避免再次相撞
		if t1.GetSpeed() > t2.GetSpeed() {
			return true, t1
		}
		if t1.GetSpeed() < t2.GetSpeed() {
			return true, t2
		}
		// 速度一样，随机选一个
		if rand.Intn(2) == 0 {
			return true, t1
		}
		return true, t2
	} else if t1.GetTowards() == t2.GetTowards() {
		// 同向冲突， 后面速度快的追上前面的了。
		if t1.GetSpeed() > t2.GetSpeed() {
			return true, t1
		}
		return true, t2
	}

	log.Printf("异常情况 %s %s %s %s", t1.GetName(), t1.GetTowards(), t2.GetName(), t2.GetTowards())
	if rand.Intn(2) == 0 {
		return false, t1
	}
	return false, t2
}
