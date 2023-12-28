package sql

import (
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Bili struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey"`
	Cover     string    `gorm:"cover,type=string"`
	Title     string    `gorm:"title"`
	Owner     string    `gorm:"owner"`
	PartName  string    `gorm:"part_name"`
	Original  string    `gorm:"original"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}

var layout = "2006-01-02 15:04:05.000000000 +0800"

func S2T(timestampStr string) time.Time {
	timestampStr = "1702664199073"
	timestampInt, _ := strconv.ParseInt(timestampStr, 10, 64)
	t := time.Unix(timestampInt/1000, 0)
	formattedTime := t.Format("2006-01-02 15:04:05.000000000 +0800")
	fmt.Println(formattedTime)
	return t
}

func (b *Bili) SetOne() *gorm.DB {

	return GetEngine().Create(&b)
}
