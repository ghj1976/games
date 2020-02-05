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
		Size:    36,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	return r
}

// Draw 绘图
func (r *Result) Draw(screen *ebiten.Image, status GameStatus, txt string) {
	switch status {
	case Success:
		text.Draw(r.ImageVictory, txt, r.ktFont, 30, 240, color.RGBA{0, 0, 205, 255})

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(ScreenWidth/2, ScreenHeight/2)
		w, h := r.ImageVictory.Size()
		op.GeoM.Translate(-float64(w)/2, -float64(h)*2.0/3.0)
		screen.DrawImage(r.ImageVictory, op)

	case Failure:
		text.Draw(r.ImageFailed, txt, r.ktFont, 30, 240, color.RGBA{238, 232, 170, 255})

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(ScreenWidth/2, ScreenHeight/2)
		w, h := r.ImageFailed.Size()
		op.GeoM.Translate(-float64(w)/2, -float64(h)*2.0/3.0)
		screen.DrawImage(r.ImageFailed, op)

	default:
	}

}
