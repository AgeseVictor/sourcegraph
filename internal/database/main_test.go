package database

import (
	"flag"
	"os"
	"testing"

	"github.com/sourcegraph/sourcegraph/lib/log"
	"github.com/sourcegraph/sourcegraph/lib/log/logtest"
)

func TestMain(m *testing.M) {
	flag.Parse()
	if !testing.Verbose() {
		logtest.InitWithLevel(m, log.LevelNone)
	}
	os.Exit(m.Run())
}
