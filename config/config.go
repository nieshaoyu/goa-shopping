package config

import (
	"os"

	"github.com/jinzhu/configor"
	"go.uber.org/zap"
)

type Config struct {
	Debug       bool             `json:"debug,omitempty" default:"false" `
	Logger      LoggerConfig     `yaml:"logger,omitempty"`
	Database    DatabaseConfig   `yaml:"database,omitempty"`
	Server      ServerConfig     `yaml:"server,omitempty"`
	Jwt         JwtConfig        `yaml:"jwt,omitempty"`
	VerifyCode  VerifyCodeConfig `yaml:"verify_code,omitempty"`
	EmailServer EmailServerConf  `yaml:"email,omitempty"`
}

type EmailServerConf struct {
	// 发送邮件地址
	Username string `yaml:"username"`
	// smtp服务器地址
	Host string `yaml:"host"`
	// 端口 默认465
	Port int `yaml:"port" default:"465"`
	// 邮箱密码或授权码
	Password string `yaml:"password"`
	// 发件人
	Sender string `yaml:"sender"`
}

type DatabaseConfig struct {
	// 仅支持 mysql
	DSN             string `yaml:"dsn"`
	MaxIdleConns    int    `yaml:"max_idle_conns" default:"10"`
	MaxOpenConns    int    `yaml:"max_open_conns" default:"100"`
	LogMode         bool   `yaml:"log_mode" default:"false"`
	ConnMaxLifetime string `yaml:"conn_max_lifetime" default:"1h"`
}

type LoggerConfig struct {
	Level string `yaml:"level,omitempty" default:"debug"`
	// json or text
	Format string `yaml:"format,omitempty" default:"json"`
	// file
	Output string `yaml:"output,omitempty" default:""`
}

type ServerConfig struct {
	Host     string `yaml:"host,omitempty" default:"0.0.0.0"`
	HTTPPort string `yaml:"http_port,omitempty" default:"8080"`
}

type VerifyCodeConfig struct {
	TplId        int    `yaml:"tpl_id"`
	ExpireMinute int    `yaml:"expire_minute" default:"5"`
	CodeLength   int    `yaml:"code_length" default:"6"`
	MockCode     string `yaml:"mock_code" default:""`
}

type JwtConfig struct {
	Secret   string `json:"secret,omitempty"`
	ExpireIn int    `json:"expire_in,omitempty" default:"86400"`
}

func initLogger(debug bool, level, output string) {
	var conf zap.Config
	if debug {
		conf = zap.NewDevelopmentConfig()
	} else {
		conf = zap.NewProductionConfig()
	}

	var zapLevel = zap.NewAtomicLevel()
	if err := zapLevel.UnmarshalText([]byte(level)); err != nil {
		zap.L().Panic("set logger level fail",
			zap.Strings("only", []string{"debug", "info", "warn", "error", "dpanic", "panic", "fatal"}),
			zap.Error(err),
		)
	}

	conf.Level = zapLevel
	conf.Encoding = "json"

	if output != "" {
		conf.OutputPaths = []string{output}
		conf.ErrorOutputPaths = []string{output}
	}

	logger, _ := conf.Build()

	zap.RedirectStdLog(logger)
	zap.ReplaceGlobals(logger)
}

var C *Config = &Config{}

func Init(cfgFile string) {
	_ = os.Setenv("CONFIGOR_ENV_PREFIX", "-")

	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)

	if cfgFile != "" {
		if err := configor.New(&configor.Config{AutoReload: true}).Load(C, cfgFile); err != nil {
			zap.L().Panic("init config fail", zap.Error(err))
		}
	} else {
		if err := configor.New(&configor.Config{AutoReload: true}).Load(C); err != nil {
			zap.L().Panic("init config fail", zap.Error(err))
		}
	}

	initLogger(C.Debug, C.Logger.Level, C.Logger.Output)

	zap.L().Debug("loaded config", zap.Any("config", C))
}

func init() {
	C = &Config{}
}
