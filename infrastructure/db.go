package infrastructure

import (
	"fmt"

	"github.com/NuttayotSukkum/user-profile/internal/core/domain"
	logrus "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

type DBCfg struct {
	Database     string `mapstructure:"database"`
	Host         string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

func InitDb() {
	var cfg DBCfg

	if err := viper.UnmarshalKey("db.user-profile", &cfg); err != nil {
		return
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:          logger.Default.LogMode(logger.Info),
		CreateBatchSize: 10000,
	})

	Log.Info("connect to database", zap.String("host:", cfg.Host), zap.String("port", cfg.Port), zap.String("database:", cfg.Database), zap.String("user:", cfg.Username))

	// if merr := db.AutoMigration()

	if err != nil {
		logrus.Errorf("ERROR:%s", err)
		panic(err)
	}

	cfgDB, err := db.DB()
	if err != nil {
		Log.Fatal("failed to connect db:", zap.Error(err))
	}

	cfgDB.SetMaxOpenConns(cfg.MaxOpenConns)
	cfgDB.SetMaxIdleConns(cfg.MaxIdleConns)

	DB = db
	if err := DB.AutoMigrate(&domain.UserProfile{}); err != nil {
		Log.Fatal("auto migration failed", zap.Error(err))
	}

	Log.Info("database connected successfully",
		zap.String("database", cfg.Database),
		zap.Int("maxOpenConns", cfg.MaxOpenConns),
		zap.Int("maxIdleConns", cfg.MaxIdleConns),
	)

}
