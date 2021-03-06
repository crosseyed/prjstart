package searchcmd

import (
	"path/filepath"
	"testing"

	"github.com/crosseyed/prjstart/internal/services/initialize"
	"github.com/crosseyed/prjstart/internal/settings"
	"github.com/crosseyed/prjstart/internal/settings/iinitialize"
	"github.com/crosseyed/prjstart/internal/utils"
	"github.com/jinzhu/copier"
)

func TestSearch(t *testing.T) {
	args := []string{"search", "keyword"}
	home := filepath.Join(utils.TempDir(), "home")
	s := settings.GetSettings(home)
	i := &initialize.Initialize{}
	copier.Copy(i, iinitialize.Inject(s))
	i.Init()
	Search(args, s)
}
