package constant

type Lang string

const (
	ZH Lang = "zh-CN"
	EN Lang = "en-US"
)

var Languages = []Lang{
	ZH, EN,
}
