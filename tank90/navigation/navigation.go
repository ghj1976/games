package navigation

import (
	"github.com/ghj1976/games/tank90/resources"
	"github.com/hajimehoshi/ebiten"
)

// Navigation 功能导航
type Navigation struct {
	btnSinglePlayer         *LinkButton
	btnSinglePlayerNewGame  *LinkButton
	btnSinglePlayerLoadGame *LinkButton
	btnSinglePlayerCampaign *LinkButton
	btnMultiplayer          *LinkButton
	btnConstruction         *LinkButton
	btnGallery              *LinkButton
}

// NewNavigation 新建导航
func NewNavigation() *Navigation {
	nav := &Navigation{}
	nav.btnSinglePlayer = NewLinkButton("Single Player", resources.GetFont("setting"), 130, 90, func() {})

	nav.btnMultiplayer = NewLinkButton("Multiplayer", resources.GetFont("setting"), 130, 110, func() {})

	nav.btnConstruction = NewLinkButton("Construction", resources.GetFont("setting"), 130, 130, func() {})

	nav.btnGallery = NewLinkButton("Gallery", resources.GetFont("setting"), 130, 150, func() {})

	return nav
}

// Update updates the current game state.
func (nav *Navigation) Update() error {

	if nav.btnSinglePlayer != nil {
		nav.btnSinglePlayer.CheckButtonState()
	}

	if nav.btnMultiplayer != nil {
		nav.btnMultiplayer.CheckButtonState()
	}

	if nav.btnConstruction != nil {
		nav.btnConstruction.CheckButtonState()
	}

	if nav.btnGallery != nil {
		nav.btnGallery.CheckButtonState()
	}
	return nil
}

// Draw 绘图
func (nav *Navigation) Draw(screen *ebiten.Image) {
	if nav.btnSinglePlayer != nil {
		nav.btnSinglePlayer.Draw(screen)
	}

	if nav.btnMultiplayer != nil {
		nav.btnMultiplayer.Draw(screen)
	}

	if nav.btnConstruction != nil {
		nav.btnConstruction.Draw(screen)
	}

	if nav.btnGallery != nil {
		nav.btnGallery.Draw(screen)
	}
}
