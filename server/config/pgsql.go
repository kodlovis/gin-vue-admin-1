package config

type Pgsql struct {
	Host         string `mapstructure:"host" json:"host" yaml:"host"` // 服务器地址:端口
	TimeZone     string `mapstructure:"timeZone" json:"timeZone" yaml:"timeZone"`
	Dbname       string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`     // 数据库名
	Username     string `mapstructure:"username" json:"username" yaml:"username"` // 数据库用户名
	Password     string `mapstructure:"password" json:"password" yaml:"password"` // 数据库密码
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	LogMode      bool   `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
	LogZap       string `mapstructure:"log-zap" json:"logZap" yaml:"log-zap"`
	Sslmode      string `mapstructure:"sslmode" json:"sslmode" yaml:"sslmode"`
	Port         string `mapstructure:"port" json:"port" yaml:"port"`
}

func (m *Pgsql) Dsn() string {
	return "host=" + m.Host + "user=" + m.Username + "password=" + m.Password + "dbname=" + m.Dbname + "port=" + m.Port + "sslmode=" + m.Sslmode + "TimeZone=" + m.TimeZone
}
