package source

type Level struct {
	GameLevelId int    `json:"game_level_id" xlsx:"关卡IVd" index:"1" `
	Sub         string `json:"sub" xlsx:"关卡序号" index:"2" `
	Cn          string `json:"cn" xlsx:"关卡名称（中文）" index:"3" `
	En          string `json:"en" xlsx:"关卡名称（H5）" index:"4" `
	Ru          string `json:"ru" xlsx:"RU" index:"5" `
	De          string `json:"de" xlsx:"DE" index:"6" `
	Fr          string `json:"fr" xlsx:"FR" index:"7" `
}
