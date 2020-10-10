package searchcmd

import (
	"github.com/crosseyed/prjstart/internal/resources/db/tablesync"
	"github.com/crosseyed/prjstart/internal/services/search"
	"github.com/crosseyed/prjstart/internal/settings"
	"github.com/crosseyed/prjstart/internal/settings/isearch"
	"github.com/crosseyed/prjstart/internal/settings/itablesync"
	"github.com/crosseyed/prjstart/internal/utils/options"
)

var usageDoc = `Search for templates using a keyword

Usage:
    prjstart search [--long] <term>

Options:
    -h --help     Print help.
    -l --long     Long output.
`

// OptSearch bindings for docopts
type OptSearch struct {
	Search bool   `docopt:"search"`
	Term   string `docopt:"<term>"`
	Long   bool   `docopt:"--long"`
}

// Search for templates
func Search(args []string, s *settings.Settings) int {
	opts := &OptSearch{}
	options.Bind(usageDoc, args, opts)

	// Sync DB table "installed" with configuration file
	sync := tablesync.New(itablesync.Inject(s))
	sync.SyncInstalled()
	srch := search.New(isearch.Inject(s))
	return srch.Search2Output(opts.Term)
}
