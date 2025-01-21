package source

import (
	"fmt"
	"strings"
)

type I18 struct {
	Num int    `json:"num,omitempty" xlsx:"序号"  index:"1" `
	Img string `json:"img,omitempty" xlsx:"图片"  index:"2" `
	ZH  string `json:"zh,omitempty" xlsx:"文本"  index:"3" `
	En  string `json:"en,omitempty" xlsx:"EN"  index:"4" `
	Ru  string `json:"ru,omitempty" xlsx:"RU"  index:"5" `
	DE  string `json:"de,omitempty" xlsx:"DE"  index:"6" `
	FR  string `json:"fr,omitempty" xlsx:"FR"  index:"7" `
}

func (receiver I18) FileName() string {
	return "本地化文案.xlsx"
}
func Target(list []*I18) string {
	end := `package target

import "strategy/constant"
var (
		I18Map  = map[int]I18{}
		I18List = []I18{
			%s
		}


		
	)
const (
	%s
)
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


%s

`
	line := "{\n\t\t\t\tNum: %d,\n\t\t\t\tImg: `%s`,\n\t\t\t\tZH: `%s`,\n\t\t\t\tEn: `%s`,\n\t\t\t\tRu: `%s`,\n\t\t\t\tDE: `%s`,\n\t\t\t\tFR: `%s`,\n\t\t\t},"
	Num := "/*%s*/ \n\t\t\t\tNum%d = %d\n"

	var ss string
	for _, i18 := range list {
		if i18.Num <= 0 {
			continue
		}
		i18.Img = strings.ReplaceAll(i18.Img, `"`, "\"")
		i18.ZH = strings.ReplaceAll(i18.ZH, `"`, "\"")
		i18.En = strings.ReplaceAll(i18.En, `"`, "\"")
		i18.Ru = strings.ReplaceAll(i18.Ru, `"`, "\"")
		i18.DE = strings.ReplaceAll(i18.DE, `"`, "\"")
		i18.FR = strings.ReplaceAll(i18.FR, `"`, "\"")
		ss += fmt.Sprintf(line, i18.Num, i18.Img, i18.ZH, i18.En, i18.Ru, i18.DE, i18.FR)
	}
	var nums string
	for _, i18 := range list {
		if i18.Num <= 0 {
			continue
		}
		nums += fmt.Sprintf(Num, i18.ZH, i18.Num, i18.Num)
	}

	var s = "type I18 struct {\n\tNum int    `json:\"num,omitempty\" xlsx:\"序号\"  index:\"1\" `\n\tImg string `json:\"img,omitempty\" xlsx:\"图片\"  index:\"2\" `\n\tZH  string `json:\"zh,omitempty\" xlsx:\"文本\"  index:\"3\" `\n\tEn  string `json:\"en,omitempty\" xlsx:\"EN\"  index:\"4\" `\n\tRu  string `json:\"ru,omitempty\" xlsx:\"RU\"  index:\"5\" `\n\tDE  string `json:\"de,omitempty\" xlsx:\"DE\"  index:\"6\" `\n\tFR  string `json:\"fr,omitempty\" xlsx:\"FR\"  index:\"7\" `\n}\n"

	return fmt.Sprintf(end, ss, nums, s)
}

type C interface {
	FileName() string
}
