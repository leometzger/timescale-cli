package config

type CliOptions struct {
	Env        string
	ConfigPath string
	Verbose    bool
}

func NewCliOptions(configPath string, verbose bool, env string) *CliOptions {
	return &CliOptions{
		ConfigPath: configPath,
		Verbose:    verbose,
		Env:        env,
	}
}
