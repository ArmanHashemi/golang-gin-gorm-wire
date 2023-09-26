package repository

import (
	"application/src/model"
	"errors"
	"time"

	"application/config"

	_ "github.com/mattn/go-sqlite3"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Erors
var (
	ErrorRepoNotInitialized = errors.New("repo not initialized")
)

// DataSouce is the struct that holds all the repository sources
type DataSource struct {
	logger   *zap.Logger
	cfg      *config.ViperConfig
	redis    *redis.Client
	sqliteDB *gorm.DB
	mysqlDB  *gorm.DB
	// pgsqlDb *gorm.DB
}

// NewDataSource creates a new DataSource
func NewDataSource(logger *zap.Logger, cfg *config.ViperConfig) (*DataSource, error) {
	ds := &DataSource{
		logger: logger.With(zap.String("type", "datasource")),
		cfg:    cfg,
		redis:  nil,
	}
	err := ds.Init()
	if err != nil {
		return nil, err
	}
	return ds, nil
}

func (ds *DataSource) Init() error {
	if ds.cfg.DatasourceConfig.Redis.Enabled {
		err := ds.initRedis()
		if err != nil {
			ds.logger.Error("redis init failed", zap.Error(err))
			return ErrorRepoNotInitialized
		}
	}
	if ds.cfg.DatasourceConfig.Sqlite.Enabled {
		ds.logger.Debug("env", zap.Bool("sqlite enabled", ds.cfg.DatasourceConfig.Sqlite.Enabled))
		err := ds.initSqlite()
		if err != nil {
			ds.logger.Error("sqlite init failed", zap.Error(err))
			return ErrorRepoNotInitialized
		}
	}

	if ds.cfg.DatasourceConfig.Mysql.Enabled {
		err := ds.InitMysql()
		if err != nil {
			ds.logger.Error("mysql init failed", zap.Error(err))
			return ErrorRepoNotInitialized
		}
	}
	return nil
}

func (ds *DataSource) Close() error {
	if ds.cfg.DatasourceConfig.Redis.Enabled {
		err := ds.redis.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func (ds *DataSource) initSqlite() error {
	db, err := gorm.Open(sqlite.Open(ds.cfg.DatasourceConfig.Sqlite.Dns), &gorm.Config{})
	if err != nil {
		ds.logger.Error("sqlite open failed", zap.Error(err))
		return err
	}
	sqlDB, err := db.DB()
	if err != nil {
		ds.logger.Error("sqlite open failed", zap.Error(err))
		return err
	}
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	ds.sqliteDB = db
	return nil
}

func (ds *DataSource) initRedis() error {
	return nil
}

func (ds *DataSource) InitMysql() error {
	ds.logger.Debug("start mysqldb")
	dns := ds.cfg.DatasourceConfig.Mysql.Dns

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		ds.logger.Error("mysql open failed", zap.Error(err))
		return err
	}

	// auto migrate
	err = db.AutoMigrate(&model.User{})

	if err != nil {
		panic("failed to migrate the database schema")
	}

	sqlDB, err := db.DB()
	if err != nil {
		ds.logger.Error("mysql open failed", zap.Error(err))
		return err
	}
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(time.Hour)
	// ping db
	err = sqlDB.Ping()
	if err != nil {
		ds.logger.Panic("mysql ping failed", zap.Error(err))
	}
	ds.mysqlDB = db
	return nil

}

// // init pgsql
