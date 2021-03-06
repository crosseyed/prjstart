package isync

import (
	"database/sql"

	"github.com/crosseyed/prjstart/internal/resources/config"
	"github.com/crosseyed/prjstart/internal/resources/gitclient/plumbing"
	"github.com/crosseyed/prjstart/internal/settings"
	"github.com/crosseyed/prjstart/internal/settings/iplumbing"
	"github.com/jinzhu/copier"
)

// Inject creates settings for tablesync.New
func Inject(s *settings.Settings) (opts struct {
	DB                 *sql.DB
	Config             *config.File
	ConfigTemplatePath string
	Plumb              *plumbing.Plumbing
}) {
	plumb := &plumbing.Plumbing{}
	copier.Copy(plumb, iplumbing.Inject(s))
	opts.DB = s.GetDB()
	opts.Config = s.ConfigFile()
	opts.ConfigTemplatePath = s.PathTemplateConf
	opts.Plumb = plumb
	return
}
