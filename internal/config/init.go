package config

import (
	"os"
	"path/filepath"

	"github.com/derhabicht/capa5/internal/logging"
)

func init() {
	cwd, _ := os.Getwd()
	logging.Trace().Str("cwd", cwd).Msg("")

	cfgDir, err := ConfigDir()
	if err != nil {
		logging.Error().Err(err).Msg("could not find user config dir")
	} else {
		logging.Debug().Str("usrCfgDir", cfgDir).Msg("user config dir found")
	}

	SetConfigType("yaml")
	SetConfigName("config")
	AddConfigPath(filepath.Join(cfgDir, "cfg"))
	err = ReadInConfig()
	if err != nil {
		path := filepath.Join(cfgDir, "cfg", "config.yaml")
		logging.Warn().Err(err).Str("path", path).Msg("config file not found, creating a default config")
		err = WriteConfigAs(path)
		if err != nil {
			logging.Error().Err(err).Msg("failed to create default config file")
		}
	}
}
