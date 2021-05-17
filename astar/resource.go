package main

import (
	"image/color"
	_ "image/png"
	"log"

	"github.com/blizzy78/ebitenui/image"
	"github.com/blizzy78/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/font"
)

const (
	backgroundColor = "131a22"

	textIdleColor     = "dff4ff"
	textDisabledColor = "5a7a91"

	labelIdleColor     = textIdleColor
	labelDisabledColor = textDisabledColor

	buttonIdleColor     = textIdleColor
	buttonDisabledColor = labelDisabledColor

	listSelectedBackground         = "4b687a"
	listDisabledSelectedBackground = "2a3944"

	toolTipColor = backgroundColor
)

type uiResources struct {
	fonts       *fonts
	background  *image.NineSlice
	text        *textResources
	label       *labelResources
	comboButton *comboButtonResources
	list        *listResources
}

type textResources struct {
	idleColor     color.Color
	disabledColor color.Color
	face          font.Face
	titleFace     font.Face
	bigTitleFace  font.Face
	smallFace     font.Face
}

type labelResources struct {
	text *widget.LabelColor
	face font.Face
}

type listResources struct {
	image        *widget.ScrollContainerImage
	track        *widget.SliderTrackImage
	trackPadding widget.Insets
	handle       *widget.ButtonImage
	handleSize   int
	face         font.Face
	entry        *widget.ListEntryColor
	entryPadding widget.Insets
}

func newUIResources() (*uiResources, error) {
	background := image.NewNineSliceColor(hexToColor(backgroundColor))
	log.Println("finish background")
	fonts, err := loadFonts()
	if err != nil {
		return nil, err
	}
	log.Println("finish fonts")

	comboButton, err := newComboButtonResources(fonts)
	if err != nil {
		return nil, err
	}
	log.Println("finish comboButton")

	list, err := newListResources(fonts)
	if err != nil {
		return nil, err
	}
	log.Println("finish newListResources")

	return &uiResources{
		fonts: fonts,

		background: background,

		text: &textResources{
			idleColor:     hexToColor(textIdleColor),
			disabledColor: hexToColor(textDisabledColor),
			face:          fonts.face,
			titleFace:     fonts.titleFace,
			bigTitleFace:  fonts.bigTitleFace,
			smallFace:     fonts.toolTipFace,
		},

		label:       newLabelResources(fonts),
		comboButton: comboButton,
		list:        list,
	}, nil
}

func newLabelResources(fonts *fonts) *labelResources {
	return &labelResources{
		text: &widget.LabelColor{
			Idle:     hexToColor(labelIdleColor),
			Disabled: hexToColor(labelDisabledColor),
		},

		face: fonts.face,
	}
}

func newListResources(fonts *fonts) (*listResources, error) {
	idle, _, err := ebitenutil.NewImageFromFile("resource/list-idle.png")
	if err != nil {
		return nil, err
	}

	disabled, _, err := ebitenutil.NewImageFromFile("resource/list-disabled.png")
	if err != nil {
		return nil, err
	}

	mask, _, err := ebitenutil.NewImageFromFile("resource/list-mask.png")
	if err != nil {
		return nil, err
	}

	trackIdle, _, err := ebitenutil.NewImageFromFile("resource/list-track-idle.png")
	if err != nil {
		return nil, err
	}

	trackDisabled, _, err := ebitenutil.NewImageFromFile("resource/list-track-disabled.png")
	if err != nil {
		return nil, err
	}

	handleIdle, _, err := ebitenutil.NewImageFromFile("resource/slider-handle-idle.png")
	if err != nil {
		return nil, err
	}

	handleHover, _, err := ebitenutil.NewImageFromFile("resource/slider-handle-hover.png")
	if err != nil {
		return nil, err
	}

	return &listResources{
		image: &widget.ScrollContainerImage{
			Idle:     image.NewNineSlice(idle, [3]int{25, 12, 22}, [3]int{25, 12, 25}),
			Disabled: image.NewNineSlice(disabled, [3]int{25, 12, 22}, [3]int{25, 12, 25}),
			Mask:     image.NewNineSlice(mask, [3]int{26, 10, 23}, [3]int{26, 10, 26}),
		},

		track: &widget.SliderTrackImage{
			Idle:     image.NewNineSlice(trackIdle, [3]int{5, 0, 0}, [3]int{25, 12, 25}),
			Hover:    image.NewNineSlice(trackIdle, [3]int{5, 0, 0}, [3]int{25, 12, 25}),
			Disabled: image.NewNineSlice(trackDisabled, [3]int{0, 5, 0}, [3]int{25, 12, 25}),
		},

		trackPadding: widget.Insets{
			Top:    5,
			Bottom: 24,
		},

		handle: &widget.ButtonImage{
			Idle:     image.NewNineSliceSimple(handleIdle, 0, 5),
			Hover:    image.NewNineSliceSimple(handleHover, 0, 5),
			Pressed:  image.NewNineSliceSimple(handleHover, 0, 5),
			Disabled: image.NewNineSliceSimple(handleIdle, 0, 5),
		},

		handleSize: 5,
		face:       fonts.face,

		entry: &widget.ListEntryColor{
			Unselected:         hexToColor(textIdleColor),
			DisabledUnselected: hexToColor(textDisabledColor),

			Selected:         hexToColor(textIdleColor),
			DisabledSelected: hexToColor(textDisabledColor),

			SelectedBackground:         hexToColor(listSelectedBackground),
			DisabledSelectedBackground: hexToColor(listDisabledSelectedBackground),
		},

		entryPadding: widget.Insets{
			Left:   30,
			Right:  30,
			Top:    10,
			Bottom: 10,
		},
	}, nil
}

func (u *uiResources) close() {
	u.fonts.close()
}
