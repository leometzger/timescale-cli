package commands

import (
	"log/slog"
	"os"
	"path"

	"github.com/leometzger/timescale-cli/internal/config"
	"github.com/leometzger/timescale-cli/internal/container"
	"github.com/spf13/cobra"
)

func newCreateConfigCommand(container *container.CliContainer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create",
		Aliases: []string{},
		Run: func(cmd *cobra.Command, args []string) {
			logger := slog.Default()

			home, err := os.UserHomeDir()
			if err != nil {
				logger.Error("could not identify user home dir", "cause", err)
				os.Exit(1)
			}

			env, err := cmd.Flags().GetString("env")
			if err != nil {
				logger.Error("could not identify environment to create", "cause", err)
				os.Exit(1)
			}

			host, err := cmd.Flags().GetString("host")
			if err != nil {
				logger.Error("could not identify host", "cause", err)
				os.Exit(1)
			}

			port, err := cmd.Flags().GetUint16("port")
			if err != nil {
				logger.Error("could not find port", "cause", err)
				os.Exit(1)
			}

			database, err := cmd.Flags().GetString("database")
			if err != nil {
				logger.Error("could not find database", "cause", err)
				os.Exit(1)
			}

			user, err := cmd.Flags().GetString("user")
			if err != nil {
				logger.Error("could not find user", "cause", err)
				os.Exit(1)
			}

			password, err := cmd.Flags().GetString("password")
			if err != nil {
				logger.Error("could not find password")
				os.Exit(1)
			}

			fullPath := path.Join(home, ".tsctl", config.DefaultConfigFileName)
			err = config.CreateConfig(env, &config.ConfigEnvironment{
				Host:     host,
				Port:     port,
				Database: database,
				User:     user,
				Password: password,
			}, fullPath)
			exitOnError(err)

			logger.Info("✔ setup config on " + fullPath)
		},
	}

	defaultConfig := config.DefaultConfig()

	cmd.Flags().StringP("host", "", defaultConfig.Host, "")
	cmd.Flags().Uint16P("port", "", 5432, "")
	cmd.Flags().StringP("database", "", defaultConfig.Database, "")
	cmd.Flags().StringP("user", "", defaultConfig.User, "")
	cmd.Flags().StringP("password", "", defaultConfig.Password, "")

	return cmd
}
