package sql

import (
	"gorm.io/gorm"
	"time"
)

type Ytdlp struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	URL       string `gorm:"url,type=string"`
	Status    string `gorm:"status"`
	ErrorMsg  string `gorm:"error_msg"`
	ErrorCode string `gorm:"error_code"`
	Request   string `gorm:"request"`
	CreatedAt time.Time
}

func (y *Ytdlp) FindOneByURL() *gorm.DB {
	return GetEngine().First(&y, "url = ?", y.URL)
}
func (y *Ytdlp) SetOne() *gorm.DB {
	return GetEngine().Create(&y)
}
func (y *Ytdlp) UpdateStatusByURL() *gorm.DB {
	return GetEngine().Model(&Ytdlp{}).Where("url = ?", y.URL).Update("status = ?", y.Status)
}
