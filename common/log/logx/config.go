package logx

// A LogConf is a logging config.
type LogConf struct {
	ServiceName         string `yaml:"ServiceName" json:",optional"`
	Mode                string `yaml:"Mode" json:",default=console,options=console|file|volume"`
	TimeFormat          string `yaml:"TimeFormat" json:",optional"`
	Path                string `yaml:"Path" json:",default=logs"`
	Level               string `yaml:"Level" json:",default=info,options=info|error|severe"`
	Compress            bool   `yaml:"Compress" json:",optional"`
	KeepDays            int    `yaml:"KeepDays" json:",optional"`
	StackCooldownMillis int    `yaml:"StackCooldownMillis" json:",default=100"`
}
