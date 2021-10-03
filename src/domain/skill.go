package domain

import (
	"config"
	"gorm.io/gorm"
)

type Skill struct {
	gorm.Model        // 기본적으로 id, create_at, updated_at, deleted_at 제공
	Lang       string // 보유 스킬
	Tier       uint   // 스킬 숙련도
	UserID     uint   // 유저 fk
}

func findByLang(lang string) {
	// todo 스킬 언어별 검색
}

func findByTier(tier uint) {
	// todo 스킬 티어별 검색
}

func CreateSkill(skill *Skill) (err error) {
	if err = config.DB.Create(skill).Error; err != nil {
		return err
	}
	return nil
}

func UpdateSkill(skill *Skill) (err error) {
	// todo 스킬 업데이트
	return nil
}

func DeleteSkill(skill *Skill) (err error) {
	// todo 스킬 삭제
	return nil
}
