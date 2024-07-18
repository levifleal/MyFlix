package config

import (
	"io"
	"os"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

type Logger = log.Logger

func NewLogger(p string) *log.Logger {
	styles := log.DefaultStyles()
	styles.Levels[log.ErrorLevel] = lipgloss.NewStyle().
		SetString("ERROR").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("204")).
		Foreground(lipgloss.Color("0")).
		Bold(true)

	styles.Levels[log.WarnLevel] = lipgloss.NewStyle().
		SetString("WARNING").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("190")).
		Foreground(lipgloss.Color("0")).
		Bold(true)

	styles.Levels[log.DebugLevel] = lipgloss.NewStyle().
		SetString("WARNING").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("0")).
		Foreground(lipgloss.Color("188")).
		Bold(true)

	styles.Levels[log.InfoLevel] = lipgloss.NewStyle().
		SetString("INFO").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("214")).
		Foreground(lipgloss.Color("0")).
		Bold(true)

	writer := io.Writer(os.Stdout)
	logger := log.NewWithOptions(writer, log.Options{
		Prefix:          p,
		Level:           log.DebugLevel,
		TimeFormat:      time.Kitchen,
		ReportCaller:    true,
		ReportTimestamp: true,
	})
	logger.SetStyles(styles)

	return logger
}
