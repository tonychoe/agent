package staticconvert_test

import (
	"runtime"
	"testing"

	"github.com/grafana/agent/converter/internal/staticconvert"
	"github.com/grafana/agent/converter/internal/test_common"
	_ "github.com/grafana/agent/pkg/metrics/instance" // Imported to override default values via the init function.
)

func TestConvert(t *testing.T) {
	test_common.TestDirectory(t, "testdata", ".yaml", true, staticconvert.Convert)

	// This test has a race condition due to downstream code so skip loading the config
	test_common.TestDirectory(t, "testdata-race", ".yaml", false, staticconvert.Convert)

	if runtime.GOOS == "windows" {
		test_common.TestDirectory(t, "testdata_windows", ".yaml", true, staticconvert.Convert)
	}
}
