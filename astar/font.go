package main

import (
	"io/ioutil"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

const (
	fontFaceRegular = "resource/FZ-Regular.ttf"
	fontFaceBold    = "resource/FZ-Bold.ttf"
)

// fonts  加载字体的类
type fonts struct {
	face         font.Face
	titleFace    font.Face
	bigTitleFace font.Face
	toolTipFace  font.Face
}

func loadFonts() (*fonts, error) {
	fontFace, err := loadFont(fontFaceRegular, 12)
	if err != nil {
		return nil, err
	}

	titleFontFace, err := loadFont(fontFaceBold, 14)
	if err != nil {
		return nil, err
	}

	bigTitleFontFace, err := loadFont(fontFaceBold, 16)
	if err != nil {
		return nil, err
	}

	toolTipFace, err := loadFont(fontFaceRegular, 12)
	if err != nil {
		return nil, err
	}

	return &fonts{
		face:         fontFace,
		titleFace:    titleFontFace,
		bigTitleFace: bigTitleFontFace,
		toolTipFace:  toolTipFace,
	}, nil
}

func (f *fonts) close() {
	if f.face != nil {
		_ = f.face.Close()
	}

	if f.titleFace != nil {
		_ = f.titleFace.Close()
	}

	if f.bigTitleFace != nil {
		_ = f.bigTitleFace.Close()
	}
}

func loadFont(path string, size float64) (font.Face, error) {
	fontData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	ttfFont, err := truetype.Parse(fontData)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(ttfFont, &truetype.Options{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	}), nil
}
