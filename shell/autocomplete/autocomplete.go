package autocomplete

import (
	"sort"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/utils/parser"
	"github.com/lmorg/readline"
)

type AutoCompleteT struct {
	Items             []string
	Definitions       map[string]string
	TabDisplayType    readline.TabDisplayType
	ErrCallback       func(error)
	DelayedTabContext readline.DelayedTabContext
	ParsedTokens      parser.ParsedTokens
}

func (act *AutoCompleteT) append(items ...string) {
	// Dedup
	for _, item := range items {
		for i := range act.Items {
			if act.Items[i] == item {
				goto next
			}
		}

		act.Items = append(act.Items, item)
	next:
	}
}

func (act *AutoCompleteT) appendDef(item, def string) {
	act.Definitions[item+" "] = def
	act.append(item)
}

func (act *AutoCompleteT) disposable() *AutoCompleteT {
	return &AutoCompleteT{
		Items:             []string{},
		Definitions:       make(map[string]string),
		ErrCallback:       act.ErrCallback,
		DelayedTabContext: act.DelayedTabContext,
		ParsedTokens:      act.ParsedTokens,
	}
}

// MatchFunction returns autocomplete suggestions for functions / executables
// based on a partial string
func MatchFunction(partial string, act *AutoCompleteT) (items []string) {
	switch {
	case pathIsLocal(partial):
		items = matchLocal(partial, true)
		items = append(items, matchDirs(partial, act)...)
	default:
		exes := allExecutables(true)
		items = matchExes(partial, exes, true)
	}
	return
}

// MatchVars returns autocomplete suggestions for variables based on a partial
// string
func MatchVars(partial string) (items []string) {
	vars := lang.ShellProcess.Variables.DumpMap()

	for name := range vars {
		if strings.HasPrefix(name, partial[1:]) {
			items = append(items, name[len(partial)-1:])
		}
	}

	sort.Strings(items)
	return
}

// MatchFlags is the entry point for murex's complex system of flag matching
func MatchFlags(flags []Flags, partial, exe string, params []string, pIndex *int, act *AutoCompleteT) int {
	args := dynamicArgs{
		exe:    exe,
		params: params,
	}

	return matchFlags(flags, partial, exe, params, pIndex, args, act)
}
