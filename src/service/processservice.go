/**
 * @Author: lzw5399
 * @Date: 2021/1/15 23:35
 * @Desc:
 */
package service

import (
	"errors"
	"gorm.io/gorm"
	"workflow/src/global"
	. "workflow/src/model"
	"workflow/src/model/request"
	"workflow/src/util"
)

// 创建新的process流程
func CreateProcess(r *request.ProcessRequest, originXml string) error {
	// 检查流程是否已存在
	var c int64
	global.BankDb.Model(&Process{}).Where("id=?", r.ID).Count(&c)
	if c != 0 {
		return errors.New("当前流程标识已经在，请检查后重试")
	}

	// 校验
	if err := validate(r); err != nil {
		return err
	}

	// 开始事务
	err := global.BankDb.Transaction(func(tx *gorm.DB) error {
		events := r.ToEvents()
		for _, event := range events {
			if err := tx.Create(&event).Error; err != nil {
				return err
			}
		}

		process := r.ToProcess(originXml)
		tx.Create(&process)

		// 返回nil提交事务
		return nil
	})

	return err
}

// 校验
func validate(r *request.ProcessRequest) error {
	if r.StartEvent == nil || len(r.StartEvent) == 0 {
		return errors.New(util.PropertyNotFound("StartEvent"))
	}

	if r.EndEvent == nil || len(r.EndEvent) == 0 {
		return errors.New(util.PropertyNotFound("EndEvent"))
	}

	return nil
}