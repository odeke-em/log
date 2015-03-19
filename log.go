package log

import (
	"fmt"
	"io"
	"os"
)

type Loggerf func(string, ...interface{}) (int, error)
type Loggerln func(...interface{}) (int, error)

type logy struct {
	printf  Loggerf
	println Loggerln
}

type Logger struct {
	Logf     Loggerf
	Logln    Loggerln
	LogErrf  Loggerf
	LogErrln Loggerln
}

func newLogger(f io.Writer) *logy {
	if f == nil {
		f = os.Stdout
	}
	fl := func(format string, args ...interface{}) (int, error) {
		return fmt.Fprintf(f, format, args...)
	}

	fln := func(args ...interface{}) (int, error) {
		return fmt.Fprintln(f, args...)
	}

	return &logy{
		printf:  fl,
		println: fln,
	}
}

func New(writers ...io.Writer) *Logger {
	var stdout, stderr io.Writer

	wLen := len(writers)
	if wLen >= 1 {
		stdout = writers[0]
	}
	if wLen >= 2 {
		stderr = writers[1]
	}

	stdouter := newLogger(stdout)
	stderrer := newLogger(stderr)

	return &Logger{
		Logf:     stdouter.printf,
		Logln:    stdouter.println,
		LogErrf:  stderrer.printf,
		LogErrln: stderrer.println,
	}
}
