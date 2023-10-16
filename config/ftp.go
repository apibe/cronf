package config

type FTP struct {
	Auth struct {
		Addr     string `json:"addr" yaml:"addr"`
		User     string `json:"user" yaml:"user"`
		Password string `json:"password" yaml:"password"`
	} `json:"auth" yaml:"auth"`
	Download []FtpUD `json:"download" yaml:"download"`
	Upload   []FtpUD `json:"upload" yaml:"upload"`
}

type FtpUD struct {
	Name      string `json:"name" yaml:"name"`
	Exec      bool   `json:"exec" yaml:"exec"`
	JoinCron  bool   `json:"join_cron" yaml:"join-cron"`
	Cron      string `json:"cron" yaml:"cron"`
	FtpPath   string `json:"ftp-path" yaml:"ftp-path"`
	LocalPath string `json:"local-path" yaml:"local-path"`
	Retry     struct {
		Interval int `json:"interval"`
		Times    int `json:"times"`
	} `json:"retry"`
}
