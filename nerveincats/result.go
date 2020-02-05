package nerveincats

import (
	"bytes"
	"image"
	"image/color"
	"log"

	"github.com/ghj1976/games/nerveincats/resources"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

// Result 结果展示类
type Result struct {
	ImageVictory *ebiten.Image
	ImageFailed  *ebiten.Image
	ktFont       font.Face
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
	r.ImageVictory, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	// 失败图片
	img, _, err = image.Decode(bytes.NewReader(resources.Failed_png))
	if err != nil {
		log.Println("Failed_png 资源文件解析错误。")
		log.Fatal(err)
	}
	r.ImageFailed, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	// 加载游戏需要的字体
	tt, err := truetype.Parse(resources.FontHuaKangWaWaTi_ttc)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	r.ktFont = truetype.NewFace(tt, &truetype.Options{
		Size:    30,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

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

		log.Println(txt)
		text.Draw(screen, txt, r.ktFont, int(x)+30, int(y)+240, color.RGBA{0, 0, 205, 255})

	case Failure:
		w, h := r.ImageFailed.Size()
		x := ScreenWidth/2 - float64(w)/2
		y := ScreenHeight/2 - float64(h)*2.0/3.0

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(x, y)
		screen.DrawImage(r.ImageFailed, op)

		log.Println(txt)
		text.Draw(screen, txt, r.ktFont, int(x)+30, int(y)+240, color.RGBA{218, 165, 32, 255})

	default:
	}

}
