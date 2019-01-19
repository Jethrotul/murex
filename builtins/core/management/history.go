package management

import (
	"errors"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell"
	"github.com/lmorg/murex/utils/json"
)

func init() {
	lang.GoFunctions["history"] = cmdHistory
	//lang.GoFunctions["^"] = cmdHistCmd
	//lang.GoFunctions["history-set-write-pipe"] = cmdHistPipe
}

func cmdHistory(p *lang.Process) (err error) {
	p.Stdout.SetDataType(types.Json)
	if !shell.Interactive {
		return errors.New("This is only designed to be run when the shell is in interactive mode")
	}

	list := shell.Prompt.History.Dump()

	b, err := json.Marshal(list, p.Stdout.IsTTY())
	if err != nil {
		return err
	}

	_, err = p.Stdout.Writeln(b)
	return err
}

func cmdHistCmd(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)
	return errors.New("Invalid usage of history variable")
}

/*func cmdHistPipe(p *lang.Process) error {
	if !shell.Interactive {
		return errors.New("This is only designed to be run when the shell is in interactive mode.")
	}

	p.Stdout.SetDataType(types.Null)

	name, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	if lang.GlobalPipes.Dump()[name] == "" {
		return errors.New("No pipe exists named: " + name)
	}

	pipe, err := lang.GlobalPipes.Get(name)
	if err != nil {
		return err
	}

	shell.History.Writer = pipe

	return nil
}*/
