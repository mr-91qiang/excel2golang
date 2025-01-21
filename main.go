package main

import (
	"excel2golang/source"
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"reflect"
	"strconv"
	"strings"
)

/*// 定义结构体 TableColumn, “xlsx"是自定义标签
type TableColumn struct {
	Classification string `json:"classification" xlsx:"分类"`
	TagNumber      string `json:"tagNumber" xlsx:"标签序号"`
	FirstlyLabel   string `json:"firstlyLabel" xlsx:"一级标签"`
	SecondaryLabel string `json:"secondaryLabel" xlsx:"二级标签"`
	TertiaryLabel  string `json:"tertiaryLabel" xlsx:"三级标签"`
}
*/
// 按行读取数据
func getRows(fileName string, sheet int) [][]string {
	file, err := excelize.OpenFile(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	rows, err := file.GetRows(file.GetSheetList()[sheet], excelize.Options{})
	if err != nil {
		panic(err)
	}
	return rows
}

// 初始化 标签序列 对应的 结构体字段 顺序
func initTag2FieldIdx(v interface{}, tagKey string) map[string]int {
	u := reflect.TypeOf(v)
	numField := u.NumField()
	tag2fieldIndex := map[string]int{}
	for i := 0; i < numField; i++ {
		f := u.Field(i)
		tagValue, ok := f.Tag.Lookup(tagKey)
		if ok {
			tag2fieldIndex[tagValue] = i
		} else {
			continue
		}
	}
	return tag2fieldIndex
}

// 将 行数据 按照反射的方式转为结构体 TableColumn
func rowsToTableColumns[T any](rows [][]string, tag2fieldIndex map[string]int) []*T {
	var data []*T
	// 默认第一行对应tag
	head := rows[0]
	for _, row := range rows[1:] {
		tbCol := new(T)
		rv := reflect.ValueOf(tbCol).Elem()
		for i := 0; i < len(head); i++ {
			colCell := row[i]
			// 通过 tag 取到结构体字段下标
			fieldIndex, ok := tag2fieldIndex[head[i]]
			if !ok {
				continue
			}

			colCell = strings.Trim(colCell, " ")
			// 通过字段下标找到字段放射对象
			v := rv.Field(fieldIndex)
			// 根据字段的类型，选择适合的赋值方法
			switch v.Kind() {
			case reflect.String:
				value := colCell
				v.SetString(value)
			case reflect.Int64, reflect.Int32:
				value, err := strconv.Atoi(colCell)
				if err != nil {
					panic(err)
				}
				v.SetInt(int64(value))
			case reflect.Int:
				if colCell == "" {
					continue
				}
				value, err := strconv.Atoi(colCell)
				if err != nil {
					panic(err)
				}
				v.SetInt(int64(value))
			case reflect.Float64:
				value, err := strconv.ParseFloat(colCell, 64)
				if err != nil {
					panic(err)
				}
				v.SetFloat(value)
			}
		}

		data = append(data, tbCol)
	}
	return data
}
func I18Excel() {
	sheet := 0
	args := os.Args
	for _, arg := range args {
		if strings.Contains(arg, ".xlsx") {
			initTag2FieldIdx := initTag2FieldIdx(source.I18{}, "xlsx")
			rows := getRows(arg, sheet)
			tbCols := rowsToTableColumns[source.I18](rows, initTag2FieldIdx)
			//c := Map[arg]

			var file = source.Target(tbCols)
			os.Mkdir("./target", os.ModePerm)
			err := os.WriteFile("./target/i18.go", []byte(file), os.ModePerm)
			if err != nil {
				panic(err)
			}
		}
	}
}

func main() {
	LevelExcel()
}

var fmtStr = `{ GameLevelId: %d, Sub: "%s", Main: "%s", Name: "%s",State:1,
				GameLevelI18s: []strategy.GameLevelI18{
					{Lang: string(constant.ZH), Main: "%s", Name: "%s", Sub: "%s"},
					{Lang: string(constant.EN), Main: "%s", Name: "%s", Sub: "%s"},
					{Lang: string(constant.RU), Main: "%s", Name: "%s", Sub: "%s"},
					{Lang: string(constant.DE), Main: "%s", Name: "%s", Sub: "%s"},
					{Lang: string(constant.FR), Main: "%s", Name: "%s", Sub: "%s"},
				},
			},
`

func LevelExcel() {
	sheet := 5
	//id := 312
	args := os.Args
	for _, arg := range args {
		if strings.Contains(arg, ".xlsx") {
			initTag2FieldIdx := initTag2FieldIdx(source.Level{}, "xlsx")
			rows := getRows(arg, sheet)
			tbCols := rowsToTableColumns[source.Level](rows, initTag2FieldIdx)
			//c := Map[arg]
			//var list []strategy.GameLevel
			fileContent := ""

			for _, level := range tbCols {
				val := fmt.Sprintf(fmtStr, level.GameLevelId, level.Sub, level.Cn, level.Cn,
					level.Cn, level.Cn, level.Sub,
					level.En, level.En, level.Sub,
					level.Ru, level.Ru, level.Sub,
					level.De, level.De, level.Sub,
					level.Fr, level.Fr, level.Sub,
				)
				fileContent += val
			}
			os.WriteFile("level.go", []byte(fileContent), os.ModePerm)
		}
	}
}

// getTags 从结构体中获取指定标签名的所有字段标签值
func getTags(structInstance source.C, tagName string) map[string]string {
	// 将interface{}转换为reflect.Type
	t := reflect.TypeOf(structInstance)
	// 确保传入的是结构体类型
	if t.Kind() != reflect.Struct {
		panic("getTags expects a struct instance")
	}

	// 创建一个map来存储标签名和对应的标签值
	tags := make(map[string]string)

	// 遍历结构体的所有字段
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		// 获取字段的标签
		tag := field.Tag.Get(tagName)
		// 如果标签存在，则添加到map中
		if tag != "" {
			tags[field.Name] = tag
		}
	}

	return tags
}
