package models

import (
	"database/sql"

	"github.com/dtkkki/cement/toolkits/mysql"
)

//DBUser just user
//should be moved to dtkkki
type DBUser struct {
	ID           int64          `gorm:"primary_key" json:"id"`
	Email        sql.NullString `gorm:"type:varchar(200);unique_index;default:null" json:"email"`
	Mobile       sql.NullString `gorm:"type:varchar(50);unique_index;default:null" json:"mobile"`
	DisplayName  string         `gorm:"type:varchar(100);default:null" json:"displayname"`
	Password     string         `gorm:"type:varchar(255);default:null"`
	Image        string         `gorm:"type:varchar(255);default:null" json:"image"`
	SignupMethod string         `gorm:"type:varchar(255);default:null"`
	Roles        string         `gorm:"type:varchar(2048);index;default:null"`
	AccessToken  string         `gorm:"type:varchar(255);default:null" json:"accesstoken"`
	Setting      string         `gorm:"type:longtext"`
	Info         string         `gorm:"type:longtext"`
	mysql.TimeMixin
}

//DBCemetUser is dbtable of cemet's user
type DBCemetUser struct {
	DBUser
}
