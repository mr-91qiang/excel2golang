package constant

// 普通玩家的角色id
const (
	PLAYER_RULE      = 1314
	PLAYER_RULE_PLUS = 520
	ENABLE           = 1
	// 未启用
	DISABLED = 2
)

type MapItem string

const (
	Spawn   MapItem = "Spawn"
	Walls   MapItem = "Walls"
	Evelate MapItem = "Evelate"
	Ground  MapItem = "Ground"
	Core    MapItem = "Core"
	Arrows  MapItem = "Arrows"
)

const EmailRegisterKey = `"[EmailRegisterKey][%s][%s][%s]",`
const EmailLoginKey = `"[EmailLoginKey][%s][%s],`

const TokenPre = "[H5]Token"

const RetryKey = `[Retry]`
