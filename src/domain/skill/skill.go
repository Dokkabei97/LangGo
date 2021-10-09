package skill

import (
	"common"
	"config"
	"gorm.io/gorm"
)

var skill Skill

const Skill_PREFIX = "ski_"

type Skill struct {
	gorm.Model        // 기본적으로 id, create_at, updated_at, deleted_at 제공
	SkillToken string `gorm:"not null"` // 스킬 토큰
	Lang       string `gorm:"not null"` // 보유 스킬
	Tier       uint   `gorm:"not null"` // 스킬 숙련도
	UserID     uint   // 유저 fk
}

func FindByLang(lang string) *gorm.DB {
	return config.DB.
		Where("lang = ?", lang).
		Find(&skill)
}

func FindByTier(tier uint) *gorm.DB {
	return config.DB.
		Where("tier = ?", tier).
		Find(&skill)
}

func CreateSkill(skill *Skill) (err error) {
	skill.SkillToken = common.RandomCharacterWithPrefix(Skill_PREFIX)
	if err = config.DB.Create(skill).Error; err != nil {
		return err
	}
	return nil
}

func UpdateSkill(tier uint) (err error) {
	if err = config.DB.Model(&skill).Update("tier", tier).Error; err != nil {
		return err
	}

	return nil
}

func DeleteSkill(skillToken string) (err error) {
	config.DB.Where("skill_token = ?", skillToken).Delete(&skill)
	return nil
}
