package navigation

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

// LinkButton 链接按钮
type LinkButton struct {
	txt      string // 显示的文本信息
	x        int    // 所在父区域的坐标x    左下脚的位置
	y        int    // 所在父区域的坐标y    左下脚的位置
	ltx      int
	lty      int
	width    int             // 宽度
	height   int             // 高度
	linkFace font.Face       // 绘图的字体
	hover    bool            // 是否鼠标移动到按钮的上面了
	clicked  AddClickedEvent // 点击事件
	bgImage  *ebiten.Image   // 选中后的背景
}

// AddClickedEvent 按钮被点击事件
type AddClickedEvent interface {
	ButtonClicked()
}

// NewLinkButton 初始化一个链接按钮
func NewLinkButton(txt string, face font.Face, x, y int, btnClicked func()) *LinkButton {
	lb := &LinkButton{}
	lb.txt = txt
	lb.linkFace = face
	lb.hover = false
	lb.x = x
	lb.y = y
	w, h := lb.MeasureText()
	lb.width = w + 4
	lb.height = h + 4
	lb.ltx = x - 2
	lb.lty = y - h

	lb.bgImage, _ = ebiten.NewImage(lb.width, lb.height, ebiten.FilterDefault)
	lb.bgImage.Fill(color.RGBA{135, 206, 235, 255})

	return lb
}

// Draw 画一个链接按钮
func (lb *LinkButton) Draw(screen *ebiten.Image) {
	if lb.hover {
		// 背景
		opts2 := &ebiten.DrawImageOptions{}
		opts2.GeoM.Translate(float64(lb.ltx), float64(lb.lty))
		screen.DrawImage(lb.bgImage, opts2)

		text.Draw(screen, lb.txt, lb.linkFace, lb.x, lb.y, color.RGBA{218, 112, 214, 255})

	} else {
		text.Draw(screen, lb.txt, lb.linkFace, lb.x, lb.y, color.RGBA{230, 230, 250, 255})
	}
}

// MeasureText 计算指定字体，指定字符串的长宽
// 实现代码 借鉴于 text.Draw( 的内部实现
// https://github.com/hajimehoshi/ebiten/blob/master/text/text.go
func (lb *LinkButton) MeasureText() (width, height int) {

	fx := fixed.I(0)
	fy := lb.linkFace.Metrics().Height
	prevR := rune(-1)
	runes := []rune(lb.txt)
	for _, r := range runes {

		if prevR >= 0 {
			fx += lb.linkFace.Kern(prevR, r)
		}
		if r == '\n' {
			fx = fixed.I(0)
			fy += lb.linkFace.Metrics().Height
			prevR = rune(-1)
			continue
		}

		a, _ := lb.linkFace.GlyphAdvance(r)
		fx += a

		prevR = r
	}
	// 向上取整数
	width = fx.Ceil()
	height = fy.Ceil()
	return width, height
}

// CheckButtonState 检查按钮状态
func (lb *LinkButton) CheckButtonState() {

	// 当 Cursor 划过btn时，btn背景加亮
	// 获得当前Cursor位置
	// https://github.com/hajimehoshi/ebiten/wiki/Tutorial:Handle-user-inputs
	x, y := ebiten.CursorPosition()
	if lb.In(x, y) {
		lb.hover = true

		// 检查是否点击了按钮
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			// 按钮点击事件
		}
	} else {
		lb.hover = false
	}
}

// In 指定的像素点，是不是在这个按钮上
func (lb *LinkButton) In(x, y int) bool {
	// 按钮是一个长方形， 是否在点击的长方形区域，就是是否点击了按钮。
	if x > lb.ltx && x < lb.ltx+lb.width && y > lb.lty && y < lb.lty+lb.height {
		return true
	}
	return false
}
