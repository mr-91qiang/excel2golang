package target

import "strategy/constant"

var (
	I18Map  = map[int]I18{}
	I18List = []I18{}
)

const ()

func init() {
	for _, i18 := range I18List {
		I18Map[i18.Num] = i18
	}
}
func GetEN(lang constant.Lang, num int) string {
	i18, ok := I18Map[num]
	if !ok {
		return ""
	}
	switch lang {
	case constant.EN:
		return i18.En
	case constant.DE:
		return i18.DE
	case constant.RU:
		return i18.Ru
	case constant.ZH:
		return i18.ZH
	case constant.FR:
		return i18.FR
	default:
		return i18.ZH
	}
}

type I18 struct {
	Num int    `json:"num,omitempty" xlsx:"序号"  index:"1" `
	Img string `json:"img,omitempty" xlsx:"图片"  index:"2" `
	ZH  string `json:"zh,omitempty" xlsx:"文本"  index:"3" `
	En  string `json:"en,omitempty" xlsx:"EN"  index:"4" `
	Ru  string `json:"ru,omitempty" xlsx:"RU"  index:"5" `
	DE  string `json:"de,omitempty" xlsx:"DE"  index:"6" `
	FR  string `json:"fr,omitempty" xlsx:"FR"  index:"7" `
}
