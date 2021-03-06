package iinitialize

import (
	"github.com/crosseyed/prjstart/internal/settings"
)

// Inject creates settings for initialize.New
func Inject(s *settings.Settings) (opts struct {
	ConfigPath         string // Path to configuration file
	ConfigTemplatePath string // Path to template configuration file
	DBDriver           string // SQL Driver to use
	DSN                string // SQL DSN
	HomeDir            string // Path to home directory
	MetadataDir        string // Path to metadata directory
	SQLiteFile         string // Path to DB file
	TemplateDir        string // Path to template directory
}) {
	opts.ConfigPath = s.PathUserConf
	opts.ConfigTemplatePath = s.PathTemplateConf
	opts.DBDriver = s.DBDriver
	opts.DSN = s.DBDsn
	opts.HomeDir = s.Home
	opts.MetadataDir = s.PathMetadataDir
	opts.SQLiteFile = s.SqliteDB
	opts.TemplateDir = s.PathTemplateDir
	return opts
}
