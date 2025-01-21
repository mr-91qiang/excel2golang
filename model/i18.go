package strategy

import "excel2golang/global"

type GameLevelI18 struct {
	global.GVA_MODEL
	GameLevelID uint
	Name        string `json:"name" form:"name" gorm:"column:name;comment:;size:64;" `                    //名称
	Main        string `json:"main" form:"main" gorm:"column:main;comment:;size:128;" binding:"required"` //主关卡
	Sub         string `json:"sub" form:"sub" gorm:"column:sub;comment:;size:128;"`                       //副关卡
	Lang        string `json:"lang"`
}

// TableName 装备 Equipment自定义表名 equipment
func (GameLevelI18) TableName() string {
	return "gamelevel_i18"
}

type HeroI18 struct {
	global.GVA_MODEL
	HeroID     uint
	Name       string `json:"name" form:"name" gorm:"column:name;comment:;size:64;" `                    //英雄名称'
	Profession string `json:"profession" form:"profession" gorm:"column:profession;comment:;size:128;" ` //职业
	Lang       string `json:"lang"`
}

func (HeroI18) TableName() string {
	return "hero_i18"
}

type EquipmentI18 struct {
	global.GVA_MODEL
	EquipmentID uint
	Profession  string `json:"profession" form:"profession" binding:"required"`        //职业
	Type        string `json:"type" form:"type" gorm:"column:type;comment:;size:128;"` //类型
	Name        string `json:"name" form:"name" binding:"required"`                    //名称
	Lang        string `json:"lang"`
}

// TableName 装备 Equipment自定义表名 equipment
func (EquipmentI18) TableName() string {
	return "equipment_i18"
}

type ArtifactI18 struct {
	global.GVA_MODEL
	Name       string `json:"name" form:"name" gorm:"column:name;comment:;size:64;"` //名称'
	ArtifactID uint   `json:"artifactID"`
	Content    string `json:"content" form:"content" gorm:"column:content;comment:;size:512;"`
	Profession string `json:"profession" form:"profession" gorm:"column:profession;comment:;size:128;" binding:"required"` //职业
	Lang       string `json:"lang"`
}

func (ArtifactI18) TableName() string {
	return "artifact_i18"
}
