package resources

import (
	"log"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

var (
	settingFace font.Face // 游戏提示信息的字体
)

// GetFont 获得指定字体
func GetFont(fonttype string) font.Face {
	if fonttype == "setting" {
		return settingFace
	}
	return nil
}

// InitFontFace 加载字体
// 考虑到对象会多次创建，这个方案不是类的方案
func InitFontFace() {

	// 加载游戏需要的字体
	tt, err := truetype.Parse(Setting_ttf)
	if err != nil {
		log.Fatal(err)
	}
	settingFace = truetype.NewFace(tt, &truetype.Options{
		Size:    12,
		DPI:     72,
		Hinting: font.HintingFull,
	})
}
