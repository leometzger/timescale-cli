package config

type CliOptions struct {
	ConfigPath string
	Env        string
	Verbose    bool
}

func NewCliOptions(configPath string, verbose bool, env string) *CliOptions {
	return &CliOptions{
		ConfigPath: configPath,
		Verbose:    verbose,
		Env:        env,
	}
}
