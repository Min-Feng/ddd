package configs

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// configPath 設置多個, 是因為二進位執行檔, 可能在 go module 根目錄執行 或 cmd 目錄執行
func NewLocalRepo(configFileName string) ProjectConfigRepoQ {
	if configFileName == "" {
		log.Fatal().Msg("Not found: configFileName is empty")
	}

	workDir, err := os.Getwd()
	workDir = filepath.ToSlash(workDir) // for window os
	if err != nil {
		log.Fatal().Msgf("Not get work directory: %v", err)
	}

	configPath := ""
	for _, dir := range ModuleDirectory {
		if strings.Contains(workDir, dir) {
			path := strings.Split(workDir, dir)
			configPath = path[0] + dir + "/config"
			break
		}
	}

	vp := viper.New()
	vp.SetConfigType("yaml")
	vp.SetConfigName(configFileName)
	vp.AddConfigPath(configPath)
	if err := vp.ReadInConfig(); err != nil {
		log.Fatal().Msgf("Reading config: %v", err)
	}

	log.Info().Msgf("New local config repository from %v successfully", vp.ConfigFileUsed())
	return &LocalRepo{vp}
}
