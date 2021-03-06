package cmdtest

import "github.com/lmorg/murex/lang"

func testState(p *lang.Process) error {
	name, err := p.Parameters.String(1)
	if err != nil {
		return errUsage("", err)
	}

	block, err := p.Parameters.Block(2)
	if err != nil {
		return errUsage("", err)
	}

	return p.Tests.State(name, block)
}
