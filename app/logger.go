package app

import (
	"io"
	"os"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

func Logger() *loggerUtil {
	if logger == nil {
		logger = &loggerUtil{}
		logger.configure()
	}
	return logger
}

var logger *loggerUtil

type loggerUtil struct {
	zerolog.Logger
}

// configure sets up the logging framework
//
// Even if you’re shipping your logs to a main platform, we recommend writing them to a file on your local
// machine first. You will want to make sure your logs are always available locally and not lostin the network.
// In addition, writing to a file means that you can decouple the task of writing your logs from the task of
// sending them to a main platform. Your applications themselves will not need to establish connections or
// stream your logs, and you can leave these jobs to specialized software like the Datadog Agent. If you’re
// running your Go applications within a containerized infrastructure that does not already include persistent
// storage (e.g., containers running on AWS Fargate) you may want to configure your log management tool to
// collect logs directly from your containers’ STDOUT and STDERR streams (this is handled differently in Docker
// and Kubernetes).
//
// The output log file will be located at LOG_FILE_FILENAME and will be rolled according to configuration set.
func (l *loggerUtil) configure() {
	var writers []io.Writer

	if LOG_CONSOLE_ENABLED {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
	}
	if LOG_FILE_ENABLED && ENV_FILE == "" {
		writers = append(writers, &lumberjack.Logger{
			Filename:   LOG_FILE_FILENAME,
			MaxSize:    LOG_FILE_MAX_SIZE,
			MaxAge:     LOG_FILE_MAX_AGE,
			MaxBackups: LOG_FILE_MAX_BACKUPS,
		})
	}
	l.Logger = zerolog.New(io.MultiWriter(writers...)).
		With().
		Timestamp().
		Logger()

	l.Info().
		Bool("LOG_CONSOLE_ENABLED", LOG_CONSOLE_ENABLED).
		Bool("LOG_FILE_ENABLED", LOG_FILE_ENABLED).
		Str("LOG_FILE_FILENAME", LOG_FILE_FILENAME).
		Int("LOG_FILE_MAX_SIZE", LOG_FILE_MAX_SIZE).
		Int("LOG_FILE_MAX_AGE", LOG_FILE_MAX_AGE).
		Int("LOG_FILE_MAX_BACKUPS", LOG_FILE_MAX_BACKUPS).
		Msg("Logging configured")
}
