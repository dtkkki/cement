package config

// Mysql 配置项
var Mysql = new(MysqlConf)

// MysqlConf  mysql的配置项
type MysqlConf struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	MaxIdle  int    `yaml:"maxidle"`
	MaxOpen  int    `yaml:"maxopen"`
}
