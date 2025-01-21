// 自动生成模板GameLevel
package strategy

import (
	"encoding/json"
	"excel2golang/global"
)

// 游戏关卡 结构体  GameLevel
type GameLevel struct {
	global.GVA_MODEL
	Name          string `json:"name" form:"name" gorm:"column:name;comment:;size:64;" binding:"required"`                //名称
	GameLevelId   int    `json:"gameLevelId" form:"gameLevelId" gorm:"column:game_level_id;comment:;" binding:"required"` //关卡id
	Main          string `json:"main" form:"main" gorm:"column:main;comment:;size:128;" binding:"required"`               //主关卡
	Sub           string `json:"sub" form:"sub" gorm:"column:sub;comment:;size:128;"`                                     //副关卡
	State         int    `json:"state" form:"state" gorm:"column:state;comment:状态 1发布 2下架;"`                              //状态
	GameLevelI18s []GameLevelI18
}

// TableName 游戏关卡 GameLevel自定义表名 gameLevel
func (GameLevel) TableName() string {
	return "gamelevel"
}

// 序列化
func (m *GameLevel) MarshalBinary() (data []byte, err error) {
	return json.Marshal(m)
}

// 反序列化
func (m *GameLevel) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)

}
