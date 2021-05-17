package main

import (
	"log"

	"github.com/blizzy78/ebitenui"
	"github.com/blizzy78/ebitenui/widget"
)

type Item struct {
	Key   int    // 主键
	Value string // 显示信息
}

func createUI() (*ebitenui.UI, func(), error) {

	res, err := newUIResources()
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}
	log.Println("020")

	rootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Padding(widget.Insets{
				Top:    10,
				Left:   20,
				Right:  30,
				Bottom: 40,
			}),
			widget.RowLayoutOpts.Spacing(4))))

	rootContainer.AddChild(widget.NewText(
		widget.TextOpts.Text("地图类型：", res.text.face, res.text.idleColor)))

	rootContainer.AddChild(newListComboButton(
		// []interface{}{"50*50随机地图", "26*26坦克大战地图", "26*26 U型地图"},
		[]interface{}{Item{Key: 1, Value: "50*50随机地图"}, Item{Key: 2, Value: "26*26坦克大战地图"}, Item{Key: 3, Value: "26*26 U型地图"}},
		func(e interface{}) string {
			log.Printf("111-%s", e.(Item).Value)
			return e.(Item).Value
		},
		func(e interface{}) string {
			log.Printf("222-%s", e.(Item).Value)
			return e.(Item).Value
		},
		func(args *widget.ListComboButtonEntrySelectedEventArgs) {

			rootContainer.RequestRelayout()
			selected := args.Entry.(Item)
			log.Printf("old %d-%d", CurrDemoMapType, CurrHScoreType)

			CurrDemoMapType = selected.Key
			log.Printf("new %d-%d", CurrDemoMapType, CurrHScoreType)
			CurrGame.aMap.Reset(CurrDemoMapType, CurrHScoreType)
		},
		res))

	rootContainer.AddChild(widget.NewText(
		widget.TextOpts.Text("H计算方式：", res.text.face, res.text.idleColor)))

	rootContainer.AddChild(newListComboButton(
		// []interface{}{"x + y", "x*x + y*y", "2 * (x + y)"},
		[]interface{}{Item{Key: 1, Value: "x + y"}, Item{Key: 2, Value: "x*x + y*y"}, Item{Key: 3, Value: "2 * (x + y)"}},
		func(e interface{}) string {
			return e.(Item).Value
		},
		func(e interface{}) string {
			return e.(Item).Value
		},
		func(args *widget.ListComboButtonEntrySelectedEventArgs) {

			rootContainer.RequestRelayout()
			selected := args.Entry.(Item)
			log.Printf("old %d-%d", CurrDemoMapType, CurrHScoreType)

			CurrHScoreType = selected.Key
			log.Printf("new %d-%d", CurrDemoMapType, CurrHScoreType)
			CurrGame.aMap.Reset(CurrDemoMapType, CurrHScoreType)
		},
		res))

	ui := &ebitenui.UI{
		Container: rootContainer,

		ToolTip: nil,

		DragAndDrop: nil,
	}

	return ui, func() {
		res.close()
	}, nil
}

func newListComboButton(entries []interface{}, buttonLabel widget.SelectComboButtonEntryLabelFunc, entryLabel widget.ListEntryLabelFunc,
	entrySelectedHandler widget.ListComboButtonEntrySelectedHandlerFunc, res *uiResources) *widget.ListComboButton {

	return widget.NewListComboButton(
		widget.ListComboButtonOpts.SelectComboButtonOpts(
			widget.SelectComboButtonOpts.ComboButtonOpts(
				widget.ComboButtonOpts.ButtonOpts(
					widget.ButtonOpts.Image(res.comboButton.image),
					widget.ButtonOpts.TextPadding(res.comboButton.padding),
				),
			),
		),
		widget.ListComboButtonOpts.Text(res.comboButton.face, res.comboButton.graphic, res.comboButton.text),
		widget.ListComboButtonOpts.ListOpts(
			widget.ListOpts.Entries(entries),
			widget.ListOpts.ScrollContainerOpts(
				widget.ScrollContainerOpts.Image(res.list.image),
			),
			widget.ListOpts.SliderOpts(
				widget.SliderOpts.Images(res.list.track, res.list.handle),
				widget.SliderOpts.HandleSize(res.list.handleSize),
				widget.SliderOpts.TrackPadding(res.list.trackPadding)),
			widget.ListOpts.EntryFontFace(res.list.face),
			widget.ListOpts.EntryColor(res.list.entry),
			widget.ListOpts.EntryTextPadding(res.list.entryPadding),
		),
		widget.ListComboButtonOpts.EntryLabelFunc(buttonLabel, entryLabel),
		widget.ListComboButtonOpts.EntrySelectedHandler(entrySelectedHandler))
}
