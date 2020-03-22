package tank90

// BotAction 机器人的动作
type BotAction struct {
	Action       string // 动作 moveto、turn 、finished 移动、转弯 暂停等待新的命令中
	StrParameter string // 命令的字符串参数，用于转弯 left right top bottom
	IntParameter int    // 命令的int参数，用于移动距离，单位像素
}
