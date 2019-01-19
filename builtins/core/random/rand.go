package random

import (
	"errors"
	"math/rand"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.GoFunctions["rand"] = cmdRand
}

func cmdRand(p *lang.Process) error {
	dt, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	p.Stdout.SetDataType(dt)
	var v interface{}

	switch dt {
	case types.Integer, types.Number:
		max, _ := p.Parameters.Int(1)
		if max > 0 {
			v = rand.Intn(max + 1)
		} else {
			v = rand.Int()
		}

	case types.Float:
		v = rand.Float64()

	default:
		return errors.New("I don't know how to generate random data for the data type `" + dt + "`")
	}

	s, err := types.ConvertGoType(v, types.String)
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write([]byte(s.(string)))
	return err
}
