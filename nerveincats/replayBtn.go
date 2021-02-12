package nerveincats

import (
	"bytes"
	"image"
	"log"
	"math"

	"github.com/ghj1976/games/nerveincats/resources"
	"github.com/hajimehoshi/ebiten/v2"
)

// ReplayBtn 重玩按钮处理逻辑
type ReplayBtn struct {
	ImageReplay *ebiten.Image
	x           int
	y           int
	w           int
	h           int
}

// NewReplayBtn 初始化重玩按钮
func NewReplayBtn() *ReplayBtn {
	btn := &ReplayBtn{}
	// 再玩一次图片
	img, _, err := image.Decode(bytes.NewReader(resources.Replay_png))
	if err != nil {
		log.Println("Replay_png 资源文件解析错误。")
		log.Fatal(err)
	}
	btn.ImageReplay = ebiten.NewImageFromImage(img)

	btn.w, btn.h = btn.ImageReplay.Size()
	btn.x = int(math.Floor(float64(ScreenWidth/2) - float64(btn.w)/2))
	btn.y = int(math.Floor(float64(ScreenHeight*4.0/5.0) - float64(btn.h)*2.0/3.0))
	return btn
}

// Draw 绘图
func (btn *ReplayBtn) Draw(screen *ebiten.Image) {

	op2 := &ebiten.DrawImageOptions{}
	op2.GeoM.Translate(float64(btn.x), float64(btn.y))
	screen.DrawImage(btn.ImageReplay, op2)

	// log.Println(btn.ImageReplay.)
}

// In 指定的像素点，是不是在这个按钮上
func (btn *ReplayBtn) In(x, y int) bool {

	// btn是一个长方形， 是否在点击的长方形区域，就是是否点击了按钮。
	if x > btn.x && x < btn.x+btn.w && y > btn.y && y < btn.y+btn.h {
		return true
	}
	return false
}
