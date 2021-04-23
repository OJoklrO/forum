package setting

import "time"

type ServerSettings struct {
	RunMode         string
	HttpPort        string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
}

type AppSettings struct {
	DefaultDescLen      int
	LogSavePath         string
	LogFileName         string
	LogFileExt          string
	StaticPagePath      string
}

type DatabaseSettings struct {
	DBType          string
	UserName        string
	Password        string
	Host            string
	DBName          string
	Charset         string
	ParseTime       bool
	MaxIdleConns    int
	MaxOpenConns    int
}

type JWTSettings struct {
	Secret string
	Issuer string
	Expire time.Duration
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}