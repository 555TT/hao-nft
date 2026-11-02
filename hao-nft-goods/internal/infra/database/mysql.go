package database

import (
	"net"
	"strconv"
	"time"

	"hao-nft-goods/internal/config"

	driver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewMySQL(path string) error {
	cfg, err := config.Load(path)
	if err != nil {
		return err
	}
	dsn := driver.NewConfig()
	dsn.User = cfg.MySQL.Username
	dsn.Passwd = cfg.MySQL.Password
	dsn.Net = "tcp"
	dsn.Addr = net.JoinHostPort(cfg.MySQL.Host, strconv.Itoa(cfg.MySQL.Port))
	dsn.DBName = cfg.MySQL.Database
	dsn.ParseTime = true // 解析mysql的时间类型为time.Time
	dsn.Loc = time.Local // 设置时区为本地时区
	dsn.Params = map[string]string{"charset": "utf8mb4"}

	DB, err = gorm.Open(mysql.Open(dsn.FormatDSN()), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}
