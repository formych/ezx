package dbx

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/fsm-xyz/ezx/log"
	_ "github.com/go-sql-driver/mysql" // 自动
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 后续动态reload服务增加锁
var (
	sqlMap  = map[string]*sql.DB{}
	sqlxMap = map[string]*sqlx.DB{}
	gormMap = map[string]*gorm.DB{}
)

// Init 初始化db资源
func Init(c []*Config) error {
	// 初始化db
	for _, v := range c {
		d, err := New(v)
		if err != nil {
			return fmt.Errorf("name: %s, err: %s", v.Name, err)
		}
		switch v.Orm {
		case "sqlx":
			sqlxMap[v.Name] = sqlx.NewDb(d, v.Type)

		case "gorm":
			gdb, _ := gorm.Open(mysql.New(mysql.Config{
				Conn: d,
			}), &gorm.Config{
				Logger: logger.Default.LogMode(logger.Silent),
			})
			gormMap[v.Name] = gdb

		default:
			sqlMap[v.Name] = d
		}
	}
	return nil
}

// GetDB ...
func GetDB(name string) *sql.DB {
	if v, ok := sqlMap[name]; ok {
		return v
	}
	return nil
}

// GetSqlxDB ...
func GetSqlxDB(name string) *sqlx.DB {
	if v, ok := sqlxMap[name]; ok {
		return v
	}
	return nil
}

// GetGormDB ...
func GetGormDB(name string) *gorm.DB {
	if v, ok := gormMap[name]; ok {
		return v
	}
	return nil
}
func Close() {
	for k, v := range sqlMap {
		if err := v.Close(); err != nil {
			log.Std.Error("close db failed", zap.String("name", k), zap.Error(err))
		}
	}

	for k, v := range sqlxMap {
		if err := v.Close(); err != nil {
			log.Std.Error("close sqlx db failed", zap.String("name", k), zap.Error(err))
		}
	}

	// 新版本的gorm
	for k, v := range gormMap {
		if db, err := v.DB(); err != nil {
			log.Std.Error("close gorm db failed", zap.String("name", k), zap.Error(err))
		} else {
			if err = db.Close(); err != nil {
				log.Std.Error("close gorm db failed", zap.String("name", k), zap.Error(err))
			}
		}
	}
}

// Config 配置

func (c *Config) validate() {
	if c.MaxIdleConns == 0 {
		c.MaxIdleConns = 10
	}

	if c.MaxOpenConns == 0 {
		c.MaxOpenConns = 100
	}

	if c.ConnMaxIdleTime == durationpb.New(0) {
		c.ConnMaxIdleTime = durationpb.New(5 * time.Minute)
	}

	if c.ConnMaxLifetime == durationpb.New(0) {
		c.ConnMaxLifetime = durationpb.New(time.Hour)
	}

	if c.ConnTimeout == durationpb.New(0) {
		c.ConnTimeout = durationpb.New(200 * time.Millisecond)
	}

	if c.ReadTimeout == durationpb.New(0) {
		c.ReadTimeout = durationpb.New(2 * time.Second)
	}

	if c.WriteTimeout == durationpb.New(0) {
		c.WriteTimeout = c.ReadTimeout
	}

	if c.Type == "mysql" {
		c.Dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s&readTimeout=%s&writeTimeout=%s&interpolateParams=true",
			c.User, c.Password, c.Addr, c.Database, c.ConnTimeout, c.ReadTimeout, c.WriteTimeout)
	} else if c.Type == "postgres" {
		c.Dsn = "postgres://" + c.User + ":" + c.Password + "@" + c.Addr + "/" + c.Database + "?sslmode=disable"
	}
}

// New ...
func New(c *Config) (*sql.DB, error) {
	c.validate()
	db, err := sql.Open(c.Type, c.Dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	// 统一设置db连接参数
	db.SetConnMaxIdleTime(c.ConnMaxIdleTime.AsDuration())
	db.SetConnMaxLifetime(c.ConnMaxLifetime.AsDuration())
	db.SetMaxIdleConns(int(c.MaxIdleConns))
	db.SetMaxOpenConns(int(c.MaxOpenConns))

	return db, nil
}
