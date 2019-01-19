package typemgmt

import (
	"errors"
	"io"
	"os"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.GoFunctions["datatype"] = cmdSetDt
	lang.GoFunctions["exec"] = lang.External
	lang.GoFunctions["die"] = cmdDie
	lang.GoFunctions["exit"] = cmdExit
	lang.GoFunctions["null"] = cmdNull
	lang.GoFunctions["true"] = cmdTrue
	lang.GoFunctions["false"] = cmdFalse
	lang.GoFunctions["!"] = cmdNot
	lang.GoFunctions["cast"] = cmdCast
}

func cmdSetDt(p *lang.Process) error {
	dt := p.Parameters.StringAll()
	//p.Scope.Stdout.SetDataType(dt)
	//p.Parent.Stdout.SetDataType(dt)
	p.Stdout.SetDataType(dt)
	return nil
}

func cmdNull(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)
	p.Stdin.ReadAll()
	return nil
}

func cmdTrue(p *lang.Process) error {
	if s, _ := p.Parameters.String(0); s != "-s" {
		p.Stdout.SetDataType(types.Boolean)
		p.Stdout.Writeln(types.TrueByte)
	} else {
		p.Stdout.SetDataType(types.Null)
	}

	return nil
}

func cmdFalse(p *lang.Process) error {
	if s, _ := p.Parameters.String(0); s != "-s" {
		p.Stdout.SetDataType(types.Boolean)
		p.Stdout.Writeln(types.FalseByte)
	} else {
		p.Stdout.SetDataType(types.Null)
	}

	p.ExitNum = 1
	return nil
}

func cmdNot(p *lang.Process) error {
	p.Stdout.SetDataType(types.Boolean)

	b, err := p.Stdin.ReadAll()
	if err != nil {
		return err
	}

	val := !types.IsTrue(b, p.Previous.ExitNum)
	if val {
		p.Stdout.Writeln(types.TrueByte)
	} else {
		p.Stdout.Writeln(types.FalseByte)
	}
	return nil
}

func cmdDie(p *lang.Process) error {
	p.Stdout.SetDataType(types.Die)

	os.Exit(1)
	return nil
}

func cmdExit(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)

	i, _ := p.Parameters.Int(0)

	os.Exit(i)
	return nil
}

func cmdCast(p *lang.Process) error {
	s, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	// Data types are lower case. So lets help people out a little here.
	dt := strings.ToLower(s)

	// Technically you could use the following values as data types, but it's unlikely anyone would intend to do so,
	// so lets just disable them with a helpful error to ease debugging.
	switch dt {
	case "string":
		return errors.New("`" + s + "` is an invalid data type. Presumably you meant `" + types.String + "`?")
	case "number":
		return errors.New("`" + s + "` is an invalid data type. Presumably you meant `" + types.Number + "`?")
	case "integer":
		return errors.New("`" + s + "` is an invalid data type. Presumably you meant `" + types.Integer + "`?")
	case "boolean":
		return errors.New("`" + s + "` is an invalid data type. Presumably you meant `" + types.Boolean + "`?")
	case "code", "codeblock":
		return errors.New("`" + s + "` is an invalid data type. Presumably you meant `" + types.CodeBlock + "`?")
	case "generic":
		return errors.New("`" + s + "` is an invalid data type. Presumably you meant `" + types.Generic + "`?")
	}

	p.Stdout.SetDataType(dt)
	_, err = io.Copy(p.Stdout, p.Stdin)
	return err
}
