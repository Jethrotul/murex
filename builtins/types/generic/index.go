package generic

import (
	"bytes"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types/define"
)

func index(p *lang.Process, params []string) error {
	cRecords := make(chan []string, 1)

	go func() {
		err := p.Stdin.ReadLine(func(b []byte) {
			cRecords <- rxWhitespace.Split(string(bytes.TrimSpace(b)), -1)
		})
		if err != nil {
			//ansi.Stderrln(p, ansi.FgRed, err.Error())
			p.Stderr.Writeln([]byte(err.Error()))
		}
		close(cRecords)
	}()

	marshaller := func(s []string) (b []byte) {
		b = []byte(strings.Join(s, "\t"))
		return
	}

	return define.IndexTemplateTable(p, params, cRecords, marshaller)
}
