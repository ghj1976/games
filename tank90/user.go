package tank90

import (
	"log"

	"github.com/ghj1976/games/tank90/resources"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// User 玩家坦克
type User struct {
	Tank
}

// NewUser 新建一个玩家
// 比坦克类多控制键的配置
func NewUser(typeName string) *User {
	u := &User{}

	switch typeName {
	case "user1":
		u.Tank = *NewTank("user1", typeName, 1, "top", 8, 24)
	case "user2":
		u.Tank = *NewTank("user2", typeName, 1, "top", 16, 24)
	default:
		log.Fatalf("错误的玩家 %s", typeName)
	}

	return u
}

// Update 处理按键逻辑
func (u *User) Update() {
	// oldTowards := u.Towards

	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		// u.Move()
		//
		u.IsMove = true

		u.Turn("top")
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		u.IsMove = true
		u.Turn("left")

	}
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		u.IsMove = true
		u.Turn("right")

	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		u.IsMove = true
		u.Turn("bottom")

	}

	if inpututil.IsKeyJustReleased(ebiten.KeyUp) {
		u.Stop()
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyLeft) {
		u.Stop()
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyRight) {
		u.Stop()
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyDown) {
		u.Stop()
	}
}

// Move 玩家移动
// 音乐播放停止的逻辑在玩家这里
func (u *User) Move(count int) {
	// log.Printf("max tps %d", ebiten.MaxTPS())
	movePlayer := resources.GetAudioPlayer("tankmove")
	if !u.IsMove {
		// 不移动了，停止播放声音
		if movePlayer.IsPlaying() {
			movePlayer.Pause()
		}

		return
	}

	u.Tank.Update(count)

	if !movePlayer.IsPlaying() {
		movePlayer.Rewind()
		movePlayer.Play()
	}
}
