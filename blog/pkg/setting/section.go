package setting

import "time"

type ServerSetting struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSetting struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
	UploadSavePath	string
	UploadServerUrl string
	UploadImageMaxSize int
	UploadImageAllowExts []string
}

type DataBaseSetting struct {
	DBType      string
	UserName    string
	Password    string
	Host        string
	DBName      string
	TablePrefix string
	Charset     string
	ParseTime   bool
	MaxIdleConn int
	MaxOpnConn  int
}

type JWTSetting struct {
	Secret string
	Issuer string
	Expire time.Duration
}

// 序列化成对象
func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}