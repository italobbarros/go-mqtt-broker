package logger

import (
	"fmt"
	"os"

	"github.com/cihub/seelog"
)

// Logger represents the log levels.
type Logger struct {
	logger seelog.LoggerInterface
}

var myLevelToString = map[seelog.LogLevel]string{
	seelog.TraceLvl:    "PROTOCOL",
	seelog.DebugLvl:    "DEBUG",
	seelog.InfoLvl:     "INFO",
	seelog.WarnLvl:     "WARN",
	seelog.ErrorLvl:    "ERROR",
	seelog.CriticalLvl: "CRITICAL",
	seelog.Off:         "OFF",
}

func createMyLevelFormatter(params string) seelog.FormatterFunc {
	return func(message string, level seelog.LogLevel, context seelog.LogContextInterface) interface{} {
		levelStr, ok := myLevelToString[level]
		if !ok {
			return "Broken level!"
		}
		return levelStr
	}
}

func InitCustomFormatter() {
	err := seelog.RegisterCustomFormatter("MYLEVEL", createMyLevelFormatter)
	if err != nil {
		fmt.Println("Error creating register custom formatter")
	}

}

// NewLogger creates a new Logger instance with the specified name.
func NewLogger(name string) *Logger {
	var configFile string
	if os.Getenv("LOG_CONFIG_PROD") == "true" {
		configFile = "configs/log-config-prod.xml" // Substitua pelo caminho correto
	} else {
		configFile = "configs/log-config.xml" // Substitua pelo caminho correto
	}

	// Carrega o arquivo de configuração
	configData, err := os.ReadFile(configFile)
	if err != nil {
		fmt.Printf("ERRO ao criar o Logger name: %s\n", name)
	}
	// Substitui a marca de lugar [%NAME] no arquivo XML pelo nome fornecido
	configString := fmt.Sprintf(string(configData), name, name, name, name, name, name)
	// Inicializa o logger a partir da string de configuração
	logger, err := seelog.LoggerFromConfigAsString(configString)
	if err != nil {
		fmt.Printf("ERRO ao criar o Logger name: %s\n", name)
	}

	return &Logger{logger: logger}
}

func (ll *Logger) reloadConfig(name string) {
	var configFile string
	if os.Getenv("LOG_CONFIG_PROD") == "true" {
		configFile = "configs/log-config-prod.xml" // Substitua pelo caminho correto
	} else {
		configFile = "configs/log-config.xml" // Substitua pelo caminho correto
	}
	//Carrega o arquivo de configuração
	configData, err := os.ReadFile(configFile)
	if err != nil {
		fmt.Printf("ERRO ao criar o Logger name: %s\n", name)
	}

	// Substitui a marca de lugar [%NAME] no arquivo XML pelo nome fornecido
	configString := fmt.Sprintf(string(configData), name)
	// Inicializa o logger a partir da string de configuração
	logger, err := seelog.LoggerFromConfigAsString(configString)
	if err != nil {
		fmt.Printf("ERRO ao criar o Logger name: %s\n", name)
	}
	ll.logger = logger
}

// Debug logs a debug message.
// Debug logs a debug message.
func (ll *Logger) Debug(format string, params ...interface{}) {
	ll.logger.Debugf(format, params...)
	ll.logger.Flush()
}

// Info logs an info message.
func (ll *Logger) Info(format string, params ...interface{}) {
	ll.logger.Infof(format, params...)
	ll.logger.Flush()
}

// Warning logs a warning message.
func (ll *Logger) Warning(format string, params ...interface{}) {
	ll.logger.Warnf(format, params...)
	ll.logger.Flush()
}

// Error logs an error message.
func (ll *Logger) Error(format string, params ...interface{}) {
	ll.logger.Errorf(format, params...)
	ll.logger.Flush()
}

func (ll *Logger) Ask(format string, params ...interface{}) {
	ll.logger.Tracef(format, params...)
	ll.logger.Flush()
}

// Answer logs a message with the ANSWER level.
func (ll *Logger) Answer(format string, params ...interface{}) {
	ll.logger.Tracef(format, params...)
	ll.logger.Flush()
}
