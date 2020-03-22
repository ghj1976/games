package tank90

import (
	"log"

	"github.com/ghj1976/games/tank90/resources"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/vorbis"
)

var (
	audioContext *audio.Context
	movePlayer   *audio.Player
)

// InitAudio 初始化声音
func InitAudio() {

	audioContext, _ = audio.NewContext(44100)

	moveD, err := vorbis.Decode(audioContext, audio.BytesReadSeekCloser(resources.Move_ogg))
	if err != nil {
		log.Fatal(err)
	}
	movePlayer, err = audio.NewPlayer(audioContext, moveD)
	if err != nil {
		log.Fatal(err)
	}
}
