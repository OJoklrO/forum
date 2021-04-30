package setting

import (
	"github.com/spf13/viper"
	"time"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("config/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}

type ServerSettings struct {
	RunMode      string
	Url          string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettings struct {
	StaticPagePath   string
	LogSavePath      string
	LogFileName      string
	LogFileExt       string
	InviteCode       string
	UploadSavePath   string
	UploadMaxSize    int
	UploadApi        string
	UploadExtensions []string
}

type DatabaseSettings struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type JWTSettings struct {
	Secret string
	Issuer string
	Expire time.Duration
}
