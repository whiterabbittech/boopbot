package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// These values encode the names of the environment
// variables used for configuring this program.
const (
	Environment      = "ENV"
	LogLevel         = "VERBOSITY"
	GivebutterApiKey = "GIVEBUTTER_API_KEY"
)

// init will load the environment variables used
// to configure this app into Viper for later access.
// init runs when the package is first imported.
func init() {
	viper.BindEnv(LogLevel)
	viper.BindEnv(Environment)
	viper.BindEnv(GivebutterApiKey)
	viper.SetDefault(LogLevel, "info")
	viper.SetDefault(Environment, "")
}

// Config represents the configuration for the application.
type Config struct {
	environment      Env
	logLevel         logrus.Level
	givebutterApiKey string
}

// NewFromEnv creates a new Config object from the environment.
// Some rudimentary validation is performed to ensure Config
// values are valid.
func NewFromEnv() *Config {
	// Fetch all of the values from the environment.
	// Strings and ints can be captured directly.
	// For the environment and log level, which both
	// come from enums, ensure that the string
	// provided in the environment maps to an enum variant.
	return &Config{
		logLevel:         parseLogLevel(),
		environment:      parseEnv(),
		givebutterApiKey: viper.GetString(GivebutterApiKey),
	}
}

// parseLogLevel will return the strongly-typed log level,
// as configured in the environment. If the log level is
// improperly specified, the program will terminate.
func parseLogLevel() logrus.Level {
	var levelStr = viper.GetString(LogLevel)
	var level, err = logrus.ParseLevel(levelStr)
	if err != nil {
		logrus.Fatalf("Unable to parse log level: %v", err)
	}
	return level
}

// parseEnv returns the Env as configured by the environment.
// The Env is an enum representing one of OtterTune's possible
// deployment environments.
func parseEnv() Env {
	var env = viper.GetString(Environment)
	return NewEnvFromString(env)
}

// Env returns the env this service is running in.
func (config *Config) Env() Env {
	return config.environment
}

// LogLevel returns the log level for this application.
func (config *Config) LogLevel() logrus.Level {
	return config.logLevel
}

func (config *Config) GivebutterAPIKey() string {
	return config.givebutterApiKey
}
