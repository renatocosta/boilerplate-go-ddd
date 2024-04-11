package logfile

import (
	"testing"

	"github.com/ddd/pkg/support"
	"github.com/google/uuid"
)

func BenchmarkSelectFile(b *testing.B) {

	logFile := NewLogFile(uuid.New(), support.NewString("c:/nonexistent-file.csv"))

	for i := 0; i < b.N; i++ {
		logFile.Select([]string{"a", "b", "c"})
	}

}
