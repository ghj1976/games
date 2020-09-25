package resources

import (
	"log"

	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/vorbis"
)

var (
	audioContext *audio.Context
	movePlayer   *audio.Player
)

// GetAudioPlayer 得到音乐播放器
func GetAudioPlayer(playerType string) *audio.Player {

	if playerType == "tankmove" {
		return movePlayer
	}
	return nil
}

// InitAudio 初始化声音
func InitAudio() {

	audioContext, _ = audio.NewContext(44100)

	moveD, err := vorbis.Decode(audioContext, audio.BytesReadSeekCloser(Move_ogg))
	if err != nil {
		log.Fatal(err)
	}
	movePlayer, err = audio.NewPlayer(audioContext, moveD)
	if err != nil {
		log.Fatal(err)
	}
}
