package tank90

import (
	"container/list"
	"log"
	"math/rand"

	"github.com/ghj1976/games/tank90/images"
)

// EnemyBot 带人工智能的敌人坦克
type EnemyBot struct {
	Tank                     // 派生坦克类的所有功能
	actions       *list.List // 要依次执行的命令
	currBotAction *BotAction // 当前执行的命令
	aStarPP       *PathPoint // 调试用的，当前A*算法算出来的路径， 正式使用时不用管它。
}

// NewEnemyBot 构造一个敌人机器人
func NewEnemyBot(name, typeName string, level int) *EnemyBot {
	bot := &EnemyBot{}

	toward := ""
	mapx, mapy := 0, 0
	// 敌人坦克随机出现在三个位置
	switch r := rand.Intn(3); r {
	case 0:
		mapx = 0
		toward = "right"
	case 1:
		mapx = 12
		toward = "right"
	case 2:
		mapx = 24
		toward = "left"
	default:
		log.Println(r)
	}
	log.Printf("NewEnemyBot x:%d y:%d toward:%s", mapx, mapy, toward)

	bot.Tank = *NewTank(name, typeName, level, toward, mapx, mapy)
	log.Printf("NewEnemyBot x:%d y:%d toward:%s", mapx, mapy, toward)
	log.Printf("NewEnemyBot cx:%d cy:%d ", bot.cx, bot.cy)

	bot.actions = list.New()

	return bot
}

// Update 每帧更新 敌人AI坦克
// 返回值 true 是要重新算新的目标
func (bot *EnemyBot) Update(count int) bool {
	if bot.currBotAction == nil {
		log.Println("无命令了！")
		// 无命令，等待接受新的命令
		return true
	}

	switch bot.currBotAction.Action {
	case "finished":
		bot.Stop()
		// 之前的命令执行完成了， 取下一个命令
		// 取下个命令不用等时间， 在一帧内完成
		bot.currBotAction = bot.PopNewCommand()
	case "moveto":
		distance := bot.Tank.Update(count)
		bot.UpdateMoveTo(distance)
		bot.Move()
	case "turn":
		// 转向不用等时间， 在一帧内完成即可
		bot.Tank.Turn(bot.currBotAction.StrParameter)
		log.Printf("bot turn %s", bot.currBotAction.StrParameter)
		bot.currBotAction.Action = "finished"
	case "parse": // 暂停
		bot.Parse()
	default:
		log.Fatalf("异常命令:%s", bot.currBotAction.Action)
	}
	return false

}

// PopNewCommand 得到最新的一条命令
func (bot *EnemyBot) PopNewCommand() *BotAction {
	// 从队列的最前面哪一个命令
	f1 := bot.actions.Front()
	if f1 != nil {
		cmd := f1.Value.(*BotAction)
		if cmd != nil {
			// 从队列里把这个命令删除
			bot.actions.Remove(f1)
		}
		return cmd
	}
	return nil
}

// PushFrontCommand 插队在最前面增加一条命令
func (bot *EnemyBot) PushFrontCommand(act *BotAction) {
	bot.actions.PushFront(act)
}

// PushFrontCommandList 插队在最前面增加一批命令
func (bot *EnemyBot) PushFrontCommandList(acts []*BotAction) {
	for _, act := range acts {
		bot.actions.PushFront(act)
	}
}

// PrintCommand 把目前的命令集合打印出来
func (bot *EnemyBot) PrintCommand() {
	log.Println("命令集合")
	log.Printf("出发点: x:%d, y:%d", bot.MapX, bot.MapY)

	for ff := bot.actions.Front(); ff != nil; ff = ff.Next() {
		cmd := ff.Value.(*BotAction)
		log.Printf("命令: %s %d %s", cmd.Action, cmd.IntParameter, cmd.StrParameter)
	}
}

// PushBackCommand 在命令列表最后加入一条命令
func (bot *EnemyBot) PushBackCommand(act *BotAction) {
	bot.actions.PushBack(act)
}

// PushBackCommandList 在命令列表最后加入一批命令
func (bot *EnemyBot) PushBackCommandList(acts []*BotAction) {
	for _, act := range acts {
		bot.actions.PushBack(act)
	}
}

// CommandIsEmpty 是否命令空了。
func (bot *EnemyBot) CommandIsEmpty() bool {
	log.Printf("bot %s actions.Len() %d", bot.GetName(), bot.actions.Len())
	if bot.actions.Len() <= 0 {
		return true
	}
	return false
}

// UpdateMoveTo 更新坦克目前距离终点的距离，
// 同时判断是完成了这个距离的移动。
func (bot *EnemyBot) UpdateMoveTo(distance int) {
	if distance <= 0 {
		return
	}
	bot.currBotAction.IntParameter = bot.currBotAction.IntParameter - distance
	log.Printf("bot %s move %d, 目标: %d", bot.GetName(), distance, bot.currBotAction.IntParameter)
	if bot.currBotAction.IntParameter <= 0 {
		bot.currBotAction.Action = "finished"
	}
}

// CommandParseFrames 给机器人一个暂停的插队命令
func (bot *EnemyBot) CommandParseFrames(num int) {
	act := &BotAction{}
	act.Action = "parse"
	act.IntParameter = num // 暂停1秒，传入的参数应该是 60
	log.Printf("bot %s parse %d", bot.GetName(), bot.currBotAction.IntParameter)

	curr := bot.currBotAction
	bot.currBotAction = act
	bot.PushFrontCommand(curr)
}

// Parse 暂停指定长时间
func (bot *EnemyBot) Parse() {
	bot.Tank.Stop()
	bot.currBotAction.IntParameter = bot.currBotAction.IntParameter - 1
	log.Printf("bot %s parse %d", bot.GetName(), bot.currBotAction.IntParameter)
	if bot.currBotAction.IntParameter <= 0 {
		bot.Tank.Move()
	}
}

// ResetWanderTargetActionList 重新来一个目标
func (bot *EnemyBot) ResetWanderTargetActionList(gmap [26][26]int) {

	// 删除之前所有命令
	bot.ClearAllCommand()
	log.Printf("清除%s的所有命令！", bot.GetName())

	// 使用的新的地图重算一个目标路径
	bot.SetWanderTargetActionList(gmap)
}

// ClearAllCommand 把所有命令清空
func (bot *EnemyBot) ClearAllCommand() {
	var n *list.Element
	f1 := bot.actions.Front()
	for ; f1 != nil; f1 = n {
		// 之前清除踩雷了
		// https://www.cnblogs.com/ziyouchutuwenwu/p/3780800.html
		cmd := f1.Value.(*BotAction)
		n = f1.Next()
		if cmd != nil {
			// 从队列里把这个命令删除
			bot.actions.Remove(f1)
		}
	}
	bot.currBotAction = nil
}

// SetWanderTargetActionList 新找一个瞎逛的目标，生成过去的命令集合
func (bot *EnemyBot) SetWanderTargetActionList(gmap [26][26]int) *PathPoint {
	if !bot.CommandIsEmpty() {
		log.Println("bot 命令还没执行完，不需要新的命令产生！")
		return nil
	}
	// 得到一个随机目标
	target := GetRandomTarget(gmap, bot.Shipable)
	log.Printf("新的目标: x:%d y:%d", target.Col, target.Row)
	source := Point{Col: bot.MapX, Row: bot.MapY}
	pp := FindPath(gmap, source, target, bot.Shipable)
	if pp == nil {
		log.Printf("A*算法找不到路径: source: %v target:%v shipable %v", source, target, bot.Shipable)
		return nil
	}

	bot.aStarPP = pp
	// 	m.aStarPP = pp
	// 把寻路结果路径变成命令队列集合
	targetPP := pp
	sourcePP := pp.Parent
	var currBotCommand *BotAction

	var lastToward string
	for {
		if targetPP == nil || sourcePP == nil {
			// 处理到最后时
			if currBotCommand != nil {
				bot.PushFrontCommand(currBotCommand)
			}
			// 无论如何，先给一个转向命令，避免默认坦克朝向不对。
			turnBotCommand := &BotAction{}
			turnBotCommand.Action = "turn"
			turnBotCommand.StrParameter = lastToward
			bot.PushFrontCommand(turnBotCommand)

			break
		}
		currToward := getNeighborsTowward(sourcePP.Point, targetPP.Point)
		if len(currToward) <= 0 {
			break
		}
		if len(lastToward) <= 0 { // 第一次进入这个逻辑
			currBotCommand = &BotAction{}
			currBotCommand.Action = "moveto"
			currBotCommand.IntParameter = images.TileSize
		} else if lastToward == currToward {
			currBotCommand.IntParameter += images.TileSize
		} else { // 朝向不一样，需要增加一个转向命令
			if currBotCommand != nil {
				bot.PushFrontCommand(currBotCommand)
			}

			turnBotCommand := &BotAction{}
			turnBotCommand.Action = "turn"
			turnBotCommand.StrParameter = lastToward
			bot.PushFrontCommand(turnBotCommand)

			// 新的移动命令
			currBotCommand = &BotAction{}
			currBotCommand.Action = "moveto"
			currBotCommand.IntParameter = images.TileSize

		}

		// 处理下一组
		targetPP = sourcePP
		sourcePP = targetPP.Parent
		lastToward = currToward
	}

	bot.PrintCommand()
	bot.currBotAction = bot.PopNewCommand()
	return pp
}
