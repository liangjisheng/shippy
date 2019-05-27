package go_micro_srv_user

import (
	"github.com/labstack/gommon/log"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// BeforeCreate 我们将 ORM 创建的 UUID 字符串修改为一个整数，用来作为表的主键
// 或 ID 是比较安全的。MongoDB 使用了类似的技术，但是 Postgres 需要我们使用
// 第三方库手动来生成
// 函数 BeforeCreate() 指定了 GORM 库使用 uuid 作为 ID 列值
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("created uuid error: %v\n", err)
	}
	return scope.SetColumn("Id", uuid.String())
}
