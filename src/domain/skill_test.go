package domain

import (
	"config"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func dbConnection() {
	config.DB, _ = gorm.Open(mysql.Open(config.DBUrl(config.BuildDBConfigTest())), &gorm.Config{})
	_ = config.DB.AutoMigrate(&Skill{})
}

func TestCreateSkill(t *testing.T) {
	dbConnection()
	assert := assert.New(t)

	skillEx := Skill{Lang: "GO", Tier: 1}
	if err := CreateSkill(&skillEx); err != nil {
		t.Error("not created")
	}

	var skill Skill
	config.DB.First(&skill, 1)
	assert.Equal("GO", skillEx.Lang)
	config.DB.Migrator().DropTable("skills")
}
