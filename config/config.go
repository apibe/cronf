package config

type Config struct {
	FTP  FTP  `json:"ftp" yaml:"ftp"`
	SCP  SCP  `json:"scp" yaml:"scp"`
	Log  Log  `json:"log" yaml:"log"`
	HTTP HTTP `json:"http" yaml:"http"`
}
