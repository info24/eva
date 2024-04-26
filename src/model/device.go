package model

import (
	"github.com/info24/eva/common"
	"gorm.io/gorm"
)

type Device struct {
	gorm.Model
	//Id          int    `orm:"auto"`
	CreateAt    int64  `gorm:"type:int(11);default:0" json:"create_at"`
	Name        string `gorm:"type:varchar(255);default:''" json:"name"`
	Ip          string `gorm:"type:varchar(15);default:''" json:"ip"`
	Username    string `gorm:"type:varchar(255);default:''" json:"username"`
	Password    string `gorm:"type:varchar(255);default:''" json:"password"`
	Pty         string `gorm:"type:varchar(255);default:''" json:"pty"`
	Description string `gorm:"type:varchar(255);default:''" json:"description"`
}

func (d *Device) ToMap(row, col int) map[string]interface{} {
	return map[string]interface{}{
		common.ID:       d.ID,
		common.IP:       d.Ip,
		common.USERNAME: d.Username,
		common.PTY:      d.Pty,
		common.PASSWORD: d.Password,
		common.ROW:      row,
		common.COL:      col,
	}
}
