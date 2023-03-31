package dbx

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/formych/ezx/log"
	_ "github.com/go-sql-driver/mysql" // 自动
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
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
func Init(c []Config) error {
	// 初始化db
	for _, v := range c {
		d, err := New(v)
		if err != nil {
			return fmt.Errorf("name: %s, err: %s", v.Name, err)
		}
		switch v.OrmType {
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
type Config struct {
	Name            string        `json:"name" yaml:"name"`
	Type            string        `json:"type" yaml:"type"`
	Addr            string        `json:"addr" yaml:"addr"`
	User            string        `json:"user" yaml:"user"`
	Password        string        `json:"password" yaml:"password"`
	Database        string        `json:"database" yaml:"database"`
	ConnTimeout     time.Duration `json:"conn_timeout" yaml:"conn_timeout"`
	ReadTimeout     time.Duration `json:"read_timeout" yaml:"read_timeout"`
	WriteTimeout    time.Duration `json:"write_timeout" yaml:"write_timeout"`
	ConnMaxIdleTime time.Duration `json:"conn_max_idle_time" yaml:"conn_max_idle_time"`
	ConnMaxLifetime time.Duration `json:"conn_max_lifetime" yaml:"conn_max_lifetime"`
	MaxIdleConns    int           `json:"max_idle_conns" yaml:"max_idle_conns"`
	MaxOpenConns    int           `json:"max_open_conns" yaml:"max_open_conns"`
	OrmType         string        `json:"orm_type" yaml:"orm_type"`

	dsn string
}

func (c *Config) validate() {
	if c.MaxIdleConns == 0 {
		c.MaxIdleConns = 10
	}

	if c.MaxOpenConns == 0 {
		c.MaxOpenConns = 100
	}

	if c.ConnMaxIdleTime == 0 {
		c.ConnMaxIdleTime = 5 * time.Minute
	}

	if c.ConnMaxLifetime == 0 {
		c.ConnMaxLifetime = time.Hour
	}

	if c.ConnTimeout == 0 {
		c.ConnTimeout = 200 * time.Millisecond
	}

	if c.ReadTimeout == 0 {
		c.ReadTimeout = 2 * time.Second
	}

	if c.WriteTimeout == 0 {
		c.WriteTimeout = c.ReadTimeout
	}

	if c.Type == "mysql" {
		c.dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s&readTimeout=%s&writeTimeout=%s&interpolateParams=true",
			c.User, c.Password, c.Addr, c.Database, c.ConnTimeout, c.ReadTimeout, c.WriteTimeout)
	} else if c.Type == "postgres" {
		c.dsn = "postgres://" + c.User + ":" + c.Password + "@" + c.Addr + "/" + c.Database + "?sslmode=disable"
	}
}

// New ...
func New(c Config) (*sql.DB, error) {
	c.validate()
	db, err := sql.Open(c.Type, c.dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	// 统一设置db连接参数
	db.SetConnMaxIdleTime(c.ConnMaxIdleTime)
	db.SetConnMaxLifetime(c.ConnMaxLifetime)
	db.SetMaxIdleConns(c.MaxIdleConns)
	db.SetMaxOpenConns(c.MaxOpenConns)

	return db, nil
}
