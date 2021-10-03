package domain

import "gorm.io/gorm"

type Skill struct {
	gorm.Model        // 기본적으로 id, create_at, updated_at, deleted_at 제공
	Lang       string // 보유 스킬
	Tier       uint   // 스킬 숙련도
	UserID     uint   // 유저 fk
}
