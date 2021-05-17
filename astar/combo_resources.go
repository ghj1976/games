package main

import (
	"image/color"
	_ "image/png"
	"log"
	"strconv"

	"github.com/blizzy78/ebitenui/image"
	"github.com/blizzy78/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/font"
)

type comboButtonResources struct {
	image   *widget.ButtonImage
	text    *widget.ButtonTextColor
	face    font.Face
	graphic *widget.ButtonImageImage
	padding widget.Insets
}

func newComboButtonResources(fonts *fonts) (*comboButtonResources, error) {
	log.Println("begin newComboButtonResources")
	idle, err := loadImageNineSlice("resource/combo-button-idle.png", 12, 0)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("finish Combo idle")
	hover, err := loadImageNineSlice("resource/combo-button-hover.png", 12, 0)
	if err != nil {
		return nil, err
	}
	log.Println("finish Combo hover")

	pressed, err := loadImageNineSlice("resource/combo-button-pressed.png", 12, 0)
	if err != nil {
		return nil, err
	}
	log.Println("finish Combo pressed")

	disabled, err := loadImageNineSlice("resource/combo-button-disabled.png", 12, 0)
	if err != nil {
		return nil, err
	}
	log.Println("finish Combo disabled")

	i := &widget.ButtonImage{
		Idle:     idle,
		Hover:    hover,
		Pressed:  pressed,
		Disabled: disabled,
	}

	arrowDown, err := loadGraphicImages("resource/arrow-down-idle.png", "resource/arrow-down-disabled.png")
	if err != nil {
		return nil, err
	}
	log.Println("finish Combo arrowDown")

	return &comboButtonResources{
		image: i,

		text: &widget.ButtonTextColor{
			Idle:     hexToColor(buttonIdleColor),
			Disabled: hexToColor(buttonDisabledColor),
		},

		face:    fonts.face,
		graphic: arrowDown,

		padding: widget.Insets{
			Left:   30,
			Right:  30,
			Top:    10,
			Bottom: 10,
		},
	}, nil

}

func loadGraphicImages(idle string, disabled string) (*widget.ButtonImageImage, error) {
	idleImage, _, err := ebitenutil.NewImageFromFile(idle)
	if err != nil {
		return nil, err
	}

	var disabledImage *ebiten.Image
	if disabled != "" {
		disabledImage, _, err = ebitenutil.NewImageFromFile(disabled)
		if err != nil {
			return nil, err
		}
	}

	return &widget.ButtonImageImage{
		Idle:     idleImage,
		Disabled: disabledImage,
	}, nil
}

func loadImageNineSlice(path string, centerWidth int, centerHeight int) (*image.NineSlice, error) {
	i, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Println("ebitenutil.NewImageFromFile 错误！")
		log.Println(err)
		return nil, err
	}

	w, h := i.Size()
	return image.NewNineSlice(i,
			[3]int{(w - centerWidth) / 2, centerWidth, w - (w-centerWidth)/2 - centerWidth},
			[3]int{(h - centerHeight) / 2, centerHeight, h - (h-centerHeight)/2 - centerHeight}),
		nil
}

func hexToColor(h string) color.Color {
	u, err := strconv.ParseUint(h, 16, 0)
	if err != nil {
		panic(err)
	}

	return color.RGBA{
		R: uint8(u & 0xff0000 >> 16),
		G: uint8(u & 0xff00 >> 8),
		B: uint8(u & 0xff),
		A: 255,
	}
}
