package handler

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
)

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func GenerateOrderSn(userId int32) string {
	u := strings.Replace(uuid.New().String(), "-", "", -1)[:16]
	return fmt.Sprintf("d%s", u)
}
