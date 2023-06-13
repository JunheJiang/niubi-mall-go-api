package cfg

type Mysql struct {
	//支持格式
	Path string `mapstructure:"path" json:"path" yaml:"path"`

	Port string `mapstructure:"port" json:"port" yaml:"port"`

	Config string `mapstructure:"config" json:"config" yaml:"config"`

	Dbname string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`

	Username string `mapstructure:"username" json:"username" yaml:"username"`

	Password string `mapstructure:"password" json:"password" yaml:"password"`

	MaxIdleConns int `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`

	MaxOpenConns int `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`

	LogMode string `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"` //Gorm全局日志

	LogZap string `mapstructure:"log-map" json:"logZap" yaml:"log-map"`
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}
