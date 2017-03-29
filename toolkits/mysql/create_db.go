package mysql

import (
	"fmt"
	"strings"
	"time"

	"github.com/dtkkki/cement/toolkits/config"
	"github.com/dtkkki/cement/toolkits/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// DBLogger 用以记录日子的Logger
type DBLogger struct{}

// Print 日志记录接口，默认忽略数据库重复的错误
func (logger DBLogger) Print(values ...interface{}) {
	s := fmt.Sprint(values...)
	if !IsDuplicateError(s) {
		log.Info(s)
	} else {
		log.Debug(s)
	}
}

// IsDuplicateError 判断是否为数据库记录重复的错误
func IsDuplicateError(err string) bool {
	return strings.Contains(err, "Error 1062")
}

// Clean 清理工作
func Clean() {
	log.Info("Doing clean work for MYSQL...")
	db.Close()
}

// CreateDB 初始化MYSQL实例
func CreateDB() *WrappedDB {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		dbName,
	)

	log.Warnf("try to connect to MYSQL %s:%d", host, port)
	database, err := gorm.Open("mysql", connStr)
	if err != nil {
		log("failed to connect MYSQL %s:%d/%s: %s", host, port, dbName, err.Error())
		return nil
	}
	log.Infof("connected to MYSQL %s:%d/%s", host, port, dbName)

	database.SetLogger(DBLogger{})
	database.DB().SetMaxIdleConns(maxIdle)
	database.DB().SetMaxOpenConns(maxOpen)
	// database.DB().SetConnMaxLifetime(d)

	return &WrappedDB{database}
}

// WrappedDB gorm.DB的包装体
type WrappedDB struct {
	database *gorm.DB
}

// Session 返回一个Transation
// Must remember to commit or rollback
func (db *WrappedDB) Session() *gorm.DB {
	return db.database.Begin()
}

// Conn 返回一个MYSQL的空闲连接
func (db *WrappedDB) Conn() *gorm.DB {
	return db.database
}

// TimeMixin mixin
type TimeMixin struct {
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP" json:"updatedAt"`
}
