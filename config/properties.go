package config

import (
	"fmt"
	"log"
	"os"

	"github.com/arielsrv/go-archaius"
	"github.com/ugurcsen/gods-generic/lists/arraylist"
)

type Config struct {
	File   string
	Folder string
	Err    error
}

type Builder struct {
	filename string
	folder   string
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) WithFile(filename string) *Builder {
	b.filename = filename
	return b
}

func (b *Builder) WithFolder(folder string) *Builder {
	b.folder = folder
	return b
}

func (b *Builder) Build() *Config {
	propertiesPath, environment, scope :=
		b.folder,
		GetEnv(),
		GetScope()

	compositeConfig := arraylist.New[string]()

	scopeConfig := fmt.Sprintf("%s/%s/%s.%s", propertiesPath, environment, scope, b.filename)
	if pathExist(scopeConfig) {
		compositeConfig.Add(scopeConfig)
	}

	envConfig := fmt.Sprintf("%s/%s/%s", propertiesPath, environment, b.filename)
	if pathExist(envConfig) {
		compositeConfig.Add(envConfig)
	}

	sharedConfig := fmt.Sprintf("%s/%s", propertiesPath, b.filename)
	if pathExist(fmt.Sprintf("%s/%s", propertiesPath, b.filename)) {
		compositeConfig.Add(sharedConfig)
	}

	err := archaius.Init(
		archaius.WithENVSource(),
		archaius.WithRequiredFiles(compositeConfig.Values()),
	)

	config := new(Config)
	config.File = b.filename
	config.Folder = b.folder

	if err != nil {
		config.Err = err
		return config
	}

	log.Printf("ENV: %s, SCOPE: %s\n", environment, scope)

	return config
}

func String(key string) string {
	value, err := archaius.GetValue(key).ToString()
	if err != nil {
		fallback := ""
		return fallback
	}
	return value
}

func TryBool(key string, defaultValue bool) bool {
	value := archaius.Exist(key)
	if !value {
		log.Printf("warn: config %s not found, fallback to %t", key, defaultValue)
		return defaultValue
	}
	return archaius.GetBool(key, defaultValue)
}

func TryInt(key string, defaultValue int) int {
	value, err := archaius.GetValue(key).ToInt()
	if err != nil {
		log.Printf("warn: config %s not found, fallback to %d", key, defaultValue)
		return defaultValue
	}
	return value
}

func pathExist(path string) bool {
	fileInfo, err := os.Stat(path)

	if err != nil {
		return false
	}

	if fileInfo.IsDir() {
		return false
	}

	return true
}
