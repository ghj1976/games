package tank90

import (
	"fmt"
	"log"
	"strconv"

	"github.com/ghj1976/games/tank90/images"
	"github.com/ghj1976/games/tank90/tools"
	"github.com/hajimehoshi/ebiten"
)

/*
1、碰撞检测 跟地图没有关系。 而是根据 坦克、子弹、墙 矩形 是否碰撞来判断的。
2、坦克可运行区域跟地图有关系。
3、坦克碰到坦克 也属于碰撞检测， 会触发新的行走方向选择。
*/

// Map 游戏地图
type Map struct {
	mapImageNeedUpdate bool // 地图需要更新

	mapArr   [26][26]int      // 地图， 一个26*26 的二维数组
	WallMap  map[string]*Tile // 所有墙的集合， 碰撞判断、绘图是这里做循环。包括砖墙、石墙
	GrassMap map[string]*Tile // 森林的集合， 森林要画在坦克上面
	SandMap  map[string]*Tile // 沙漠的集合， 这里坦克停止将随机滑行一段距离
	WaterMap map[string]*Tile // 水域的集合， 没有船，坦克无法过河
	MyBird   *Bird
}

// NewMap 构建地图信息
// 传入的是地图信息
func NewMap(id int) *Map {
	mt := getMapByName(id)
	if len(mt) <= 0 {
		log.Fatalf("没有编号为%d的地图！", id)
	}
	tm := &Map{}
	tm.WallMap = make(map[string]*Tile)
	tm.GrassMap = make(map[string]*Tile)
	tm.SandMap = make(map[string]*Tile)
	tm.WaterMap = make(map[string]*Tile)

	tm.string2Map(mt)
	tm.buildTileMap()

	return tm
}

// string2Map 把地图字符串配置信息，变成二维数组数据
// 非数字的字符如何组织都不影响加载
func (m *Map) string2Map(mt string) {
	// 把字符串的地图转换成 [26][26]int 对象
	x := 0
	y := 0
	for _, rv := range mt {

		v, err := strconv.Atoi(string(rv))
		if err != nil {
			continue
		}

		if x >= 26 || y >= 26 {
			//log.Println(v)
			break
		} else {
			m.mapArr[y][x] = v
		}

		// 下一轮的 x y 计算
		if x < 25 {
			x++
		} else {
			x = 0
			y++
		}
	}
	// 打印出来，看地图数据
	for _, ja := range m.mapArr {
		ss := ""
		for _, v := range ja {
			ss += fmt.Sprintf("%d", v)
		}
		log.Println(ss)
	}
}

// buildTileMap 基于地图，生成地图中各个元素的map对象
func (m *Map) buildTileMap() {

	// 鹰巢只能在一个位置，而且必须有
	m.MyBird = NewBird(12, 24)

	// 遍历地图map
	for i := 0; i < 26; i++ { // y 轴
		for j := 0; j < 26; j++ { // x 轴
			node := m.mapArr[i][j] // m.mapArr[y][x] 这样的存储
			if node == 9 {
				// 鹰巢  已经处理了
				continue
			}
			if node == 0 { // 道路不用处理
				continue
			}

			switch node {
			case 1, 5: // 砖墙、石墙
				m.WallMap[fmt.Sprintf("%d-%d", j, i)] = NewTile(node, j, i)
			case 2: // 森林
				m.GrassMap[fmt.Sprintf("%d-%d", j, i)] = NewTile(node, j, i)
			case 3: // 水域
				m.WaterMap[fmt.Sprintf("%d-%d", j, i)] = NewTile(node, j, i)
			case 4: // 沙漠
				m.SandMap[fmt.Sprintf("%d-%d", j, i)] = NewTile(node, j, i)
			default:
				log.Println(fmt.Sprintf("位置%d-%d没处理%d逻辑。", j, i, node))
			}

		}
	}

}

// Draw 在地图上绘制猫的图像
func (m *Map) Draw(screen *ebiten.Image) {

	if ebiten.IsDrawingSkipped() {
		return
	}

	m.MyBird.Draw(screen)

	for _, wall := range m.WallMap {
		wall.Draw(screen)
	}

	for _, water := range m.WaterMap {
		water.Draw(screen)
	}

	for _, sand := range m.SandMap {
		sand.Draw(screen)
	}

}

// GrassDraw 更新森林绘图
// 森林要画在坦克上面，所以后面画
func (m *Map) GrassDraw(screen *ebiten.Image) {

	for _, grass := range m.GrassMap {
		grass.Draw(screen)
	}
}

// CanMove 坦克是否可以移动
// 地图边缘 + 碰撞判断
func (m *Map) CanMove(t ITankMeeting) bool {

	cx, cy, w, h := t.GetCentorPositionAndSize()
	mx, my := t.GetMapXY()
	// 地图边缘判断
	if cx < w/2 || cy < h/2 || cx > images.TileSize*26-w/2 || cy > images.TileSize*26-h/2 {
		return false
	}

	// 提高性能考虑， 只判断附近的地图元素
	for _, c := range m.NearByTiles(mx, my, t.GetTowards()) {

		// 碰撞判断
		if tools.CheckCollision(c, t) {
			// 程序调试提示
			// switch ty := c.(type) {
			// case *Tile:
			// 	log.Printf("与 x:%d,y:%d 位置的发生碰撞 ", ty.MapX, ty.MapY)
			// 	log.Printf("之前位置: x:%d,y:%d ", t.PreCX, t.PreCY)
			// default:
			// 	log.Println(reflect.TypeOf(c))
			// }
			return false
		}
	}

	// 出现碰撞后，要回退，所以要记录之前的坐标点
	t.SetPreCXY(cx, cy)
	return true // 可以移动

}

// NearByTiles 返回附近的具备检查碰撞的地图对象
// 不会阻碍前进的  道路不算， 森林不算
// 1 砖墙 3 水域 4 沙漠 5 石墙  9 司令部
// xy 是坦克左上角的坐标， 所以附近的要返回下面几个位置
// 考虑到还有移动的偏差， 还有一个外围
//   * * * *     y-2   x-2 到 x+3
// * * * * * *   y-1
// * * x m * *   y
// * * m m * *   y+1
// * * * * * *   y+2
//   * * * *     y+3
// top      x-1 ~ x+2  y-2 ~ y+0
// bottom   x-1 ~ x+2  y+1 ~ y+3
// left     x-2 ~ x+0  y-1 ~ y+2
// right    x+1 ~ x+3  y-1 ~ y+2
func (m *Map) NearByTiles(x, y int, toward string) []tools.ICollision {
	// log.Printf("x,y,t=%d,%d,%s\r\n", x, y, toward)
	arr := []tools.ICollision{}

	var arrPoints []tools.Point
	switch toward {
	case "top":
		arrPoints = getICollision(x, y, -1, 2, -2, 0)
	case "bottom":
		arrPoints = getICollision(x, y, -1, 2, 1, 3)
	case "left":
		arrPoints = getICollision(x, y, -2, 0, -1, 2)
	case "right":
		arrPoints = getICollision(x, y, 1, 3, -1, 2)
	default:
		return arr
	}
	// log.Println(len(arrPoints))

	for _, p := range arrPoints {
		wall, b1 := m.WallMap[fmt.Sprintf("%d-%d", p.X, p.Y)]
		if b1 {
			arr = append(arr, wall)
		}
		sand, b2 := m.SandMap[fmt.Sprintf("%d-%d", p.X, p.Y)]
		if b2 {
			arr = append(arr, sand)
		}
		water, b3 := m.WaterMap[fmt.Sprintf("%d-%d", p.X, p.Y)]
		if b3 {
			arr = append(arr, water)
		}
		if m.mapArr[p.Y][p.X] == 9 {
			arr = append(arr, m.MyBird)
		}
	}

	// log.Printf("point %d, mapcheck %d\r\n", len(arrPoints), len(arr))

	return arr
}

// getICollision 取 x，y 在 minx, maxx, miny, maxy 范围内的所有点
func getICollision(x, y, minx, maxx, miny, maxy int) []tools.Point {
	arr := []tools.Point{}
	for m := miny; m <= maxy; m++ {
		ny := y + m
		if ny <= 0 || ny > 25 {
			continue
		}

		for n := minx; n <= maxx; n++ {
			nx := x + n
			if nx <= 0 || nx > 25 {
				continue
			}
			arr = append(arr, tools.Point{X: nx, Y: ny})
			// log.Printf("x,y=%d,%d\r\n", nx, ny)
		}
	}
	// log.Printf("len %d\r\n", len(arr))
	return arr
}

// CloneMap 克隆一份当面map的地图数组
// 注意，克隆出来的要做修改，这里必须完整的克隆，避免影响线上使用的
func (m *Map) CloneMap() [26][26]int {
	var mmap [26][26]int // 这是一个数组，不是切片， 值传递

	for y := 0; y < 26; y++ {
		for x := 0; x < 26; x++ {
			mmap[y][x] = m.mapArr[y][x]
		}
	}

	return mmap
}
