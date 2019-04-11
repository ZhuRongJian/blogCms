package conf

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/fsnotify/fsnotify"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

//初始化全局DB
var DB *gorm.DB

func SetUp() {
	// 初始化配置
	cerr := initConfig()
	if cerr != nil {
		fmt.Println(cerr)
	}

	//初始化数据库
	dberr := initDb()
	if dberr != nil {
		fmt.Println(dberr)
	}
}

func initConfig() error {
	viper.SetConfigFile("./app.yaml") // 如果指定了配置文件，则解析指定的配置文件
	viper.SetConfigType("yaml")       // 设置配置文件格式为YAML
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改:", in.Name)
	})
	return nil
}

func initDb() error {
	user := viper.GetString("mysql.user")
	password := viper.GetString("mysql.password")
	dbName := viper.GetString("mysql.db")
	local := viper.GetString("mysql.host") + ":" + viper.GetString("mysql.port")
	DB, err := gorm.Open("mysql", ""+user+":"+password+"@("+local+")/"+dbName+"?charset=utf8&parseTime=True&loc=")
	defer DB.Close()
	return err
}
