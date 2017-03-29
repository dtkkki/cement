package config

import (
	"encoding/json"
	"github.com/dtkkki/cement/toolkits/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type configure struct {
	mysql *MysqlConf `yaml:"mysql"`
}

// Configure global configure.
var config = &configure{
	mysql: Mysql,
}

// LoadFile 加载配置文件
func LoadFile(filepath string) error {
	return config.updateWithYamlFile(filepath)
}

// UpdateWithJSONFile updates the global config with json file.
func (c *configure) updateWithJSONFile(filepath string) error {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Errorf("can't read config from %s: %s.", filepath, err.Error())
		return err
	}
	err = json.Unmarshal(b, c)
	if err != nil {
		log.Errorf("can't parse config from %s: %s.", filepath, err.Error())
		return err
	}
	log.Infof("read config from %s", filepath)
	return nil
}

// UpdateWithYamlFile updates the global config with yaml file
func (c *configure) updateWithYamlFile(filepath string) error {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Errorf("can't read config from %s: %s.", filepath, err.Error())
		return err
	}
	err = yaml.Unmarshal(b, c)
	if err != nil {
		log.Errorf("can't parse config from %s: %s.", filepath, err.Error())
		return err
	}
	log.Infof("read config from %s", filepath)
	return nil
}
