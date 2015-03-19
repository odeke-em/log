package log

import (
	"os"
	"testing"
)

func TestInitWithNoArgs(t *testing.T) {
	logf := New()
	if logf == nil {
		t.Errorf("Expected non-nil logger")
	}
	if logf.Logf == nil {
		t.Errorf("*.Logf should be non-nil")
	}
	if logf.Logln == nil {
		t.Errorf("*.Logln should be non-nil")
	}
	if logf.LogErrf == nil {
		t.Errorf("*.LogErrf should be non-nil")
	}
	if logf.LogErrln == nil {
		t.Errorf("*.LogErrln should be non-nil")
	}
}

func TestWithOneWriter(t *testing.T) {
	logf := New(os.Stdout)
	if logf == nil {
		t.Errorf("not expecting a nil logger")
	}
	logf.Logf("OutPut: %s %v\n", "trivia", logf)
	logf.Logln("OutPut")

	logf = New(os.Stdout, os.Stderr)
	logf.LogErrf("Errf here: %s calling: %v\n", "bingo", logf)
	logf.LogErrln("Errf here: %s calling: %v\n", "bingo", logf)
}
