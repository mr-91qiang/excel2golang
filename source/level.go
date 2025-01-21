package source

type Level struct {
	GameLevelId int    `json:"game_level_id" xlsx:"关卡ID" index:"1" `
	CNSub       string `json:"CNSub" xlsx:"CN副关卡" index:"2" ` //
	ENSub       string `json:"ENSub" xlsx:"EN副关卡" index:"3" ` //
	RUSub       string `json:"RUSub" xlsx:"RU副关卡" index:"4" ` //
	DESub       string `json:"DESub" xlsx:"DE副关卡" index:"5" ` //
	FRSub       string `json:"FRSub" xlsx:"FR副关卡" index:"6" ` //
	CnMain      string `json:"cnMain" xlsx:"CN主关卡" index:"7" `
	EnMain      string `json:"enMain" xlsx:"EN主关卡" index:"8" `
	RuMain      string `json:"ruMain" xlsx:"RU主关卡" index:"9" `
	DeMain      string `json:"deMain" xlsx:"DE主关卡" index:"10" `
	FrMain      string `json:"frMain" xlsx:"FR主关卡" index:"11" `
}
