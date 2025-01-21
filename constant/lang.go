package constant

type Lang string

const (
	ZH Lang = "zh-CN"
	EN Lang = "en-US"
	FR Lang = "fr-FR"
	DE Lang = "de-DE"
	RU Lang = "ru-RU"
)

var Languages = []Lang{
	ZH,
	EN,
	FR,
	DE,
	RU,
}
