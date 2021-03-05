/**
 * @Author: lzw5399
 * @Date: 2021/1/14 23:11
 * @Desc:
 */
package model

import "time"

type EntityBase struct {
	Id uint `gorm:"primarykey"`
}

type AuditableBase struct {
	EntityBase
	CreateTime time.Time `gorm:"default:now()"`
	UpdateTime time.Time `gorm:"default:now()"`
	CreateBy   uint
	UpdateBy   uint
}
