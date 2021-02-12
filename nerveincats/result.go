package nerveincats

import (
	"bytes"
	"image"
	"image/color"
	"log"

	"github.com/ghj1976/games/nerveincats/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

// Result 结果展示类
type Result struct {
	ImageVictory *ebiten.Image
	ImageFailed  *ebiten.Image
}

// NewResult 初始化结果绘图需要的内容
func NewResult() *Result {
	r := &Result{}

	// 成功图片
	img, _, err := image.Decode(bytes.NewReader(resources.Victory_png))
	if err != nil {
		log.Println("Victory_png 资源文件解析错误。")
		log.Fatal(err)
	}
	r.ImageVictory = ebiten.NewImageFromImage(img)

	// 失败图片
	img, _, err = image.Decode(bytes.NewReader(resources.Failed_png))
	if err != nil {
		log.Println("Failed_png 资源文件解析错误。")
		log.Fatal(err)
	}
	r.ImageFailed = ebiten.NewImageFromImage(img)

	return r
}

// Draw 绘图
func (r *Result) Draw(screen *ebiten.Image, status GameStatus, txt string) {
	switch status {
	case Success:
		w, h := r.ImageVictory.Size()
		x := ScreenWidth/2 - float64(w)/2
		y := ScreenHeight/2 - float64(h)*2.0/3.0

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(x, y)
		screen.DrawImage(r.ImageVictory, op)

		// log.Println(txt)
		text.Draw(screen, txt, fontGameResult, int(x)+30, int(y)+240, color.RGBA{0, 0, 205, 255})

	case Failure:
		w, h := r.ImageFailed.Size()
		x := ScreenWidth/2 - float64(w)/2
		y := ScreenHeight/2 - float64(h)*2.0/3.0

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(x, y)
		screen.DrawImage(r.ImageFailed, op)

		// log.Println(txt)
		text.Draw(screen, txt, fontGameResult, int(x)+30, int(y)+240, color.RGBA{218, 165, 32, 255})

	default:
	}

}
