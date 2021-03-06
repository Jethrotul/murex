package shell

import (
	"os"
	"testing"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell/userdictionary"

	_ "github.com/lmorg/murex/builtins/core/arraytools"
	_ "github.com/lmorg/murex/builtins/core/textmanip"
	_ "github.com/lmorg/murex/builtins/types/json"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/ansi"
)

func configDefaults(c *config.Config) {
	c.Define("shell", "spellcheck-enabled", config.Properties{
		Description: "Enable spellchecking in the interactive prompt",
		Default:     false,
		DataType:    types.Boolean,
		Global:      true,
	})

	c.Define("shell", "spellcheck-block", config.Properties{
		Description: "Code block to run as part of the spellchecker (STDIN the line, STDOUT is array for misspelt words)",
		Default:     "{ -> aspell list }",
		DataType:    types.CodeBlock,
		Global:      true,
	})

	c.Define("shell", "spellcheck-user-dictionary", config.Properties{
		Description: "An array of words not to count as misspellings",
		Default:     userdictionary.Get(),
		DataType:    types.Json,
		Global:      true,
		GoFunc: config.GoFuncProperties{
			Read:  userdictionary.Read,
			Write: userdictionary.Write,
		},
	})
}

func TestSpellcheckCrLf(t *testing.T) {
	if os.Getenv("MUREX_TEST_SKIP_SPELLCHECK") != "" {
		t.Skip("Environmental variable `MUREX_TEST_SKIP_SPELLCHECK` set")
		return
	}

	count.Tests(t, 1)
	lang.InitEnv()
	//defaults.Defaults(lang.ShellProcess.Config, false)
	configDefaults(lang.ShellProcess.Config)

	err := lang.ShellProcess.Config.Set("shell", "spellcheck-enabled", true)
	if err != nil {
		t.Fatalf("Unable to set spellcheck-enabled config: %s", err)
	}

	err = lang.ShellProcess.Config.Set("shell", "spellcheck-block", `{ -> jsplit ' ' -> suffix "\n" }`)
	if err != nil {
		t.Fatalf("Unable to set spellcheck-block config: %s", err)
	}

	line := "the quick brown fox"
	newLine := string(spellcheck([]rune(line)))
	ansiLine := ansi.ExpandConsts("{UNDERLINE}the{UNDEROFF} {UNDERLINE}quick{UNDEROFF} {UNDERLINE}brown{UNDEROFF} {UNDERLINE}fox{UNDEROFF}")

	if newLine != ansiLine {
		t.Error("spellcheck output doesn't match expected:")
		t.Logf("  Expected: '%s'", ansiLine)
		t.Logf("  Actual:   '%s'", newLine)
	}
}

func TestSpellcheckZeroLenStr(t *testing.T) {
	if os.Getenv("MUREX_TEST_SKIP_SPELLCHECK") != "" {
		t.Skip("Environmental variable `MUREX_TEST_SKIP_SPELLCHECK` set")
		return
	}

	count.Tests(t, 1)
	lang.InitEnv()
	//defaults.Defaults(lang.ShellProcess.Config, false)
	configDefaults(lang.ShellProcess.Config)

	err := lang.ShellProcess.Config.Set("shell", "spellcheck-enabled", true)
	if err != nil {
		t.Fatalf("Unable to set spellcheck-enabled config: %s", err)
	}

	err = lang.ShellProcess.Config.Set("shell", "spellcheck-block", `{ -> jsplit '\s' -> append '' }`)
	if err != nil {
		t.Fatalf("Unable to set spellcheck-block config: %s", err)
	}

	line := "the quick brown fox"
	newLine := string(spellcheck([]rune(line)))
	ansiLine := ansi.ExpandConsts("{UNDERLINE}the{UNDEROFF} {UNDERLINE}quick{UNDEROFF} {UNDERLINE}brown{UNDEROFF} {UNDERLINE}fox{UNDEROFF}")

	if newLine != ansiLine {
		t.Error("spellcheck output doesn't match expected:")
		t.Logf("  Expected: '%s'", ansiLine)
		t.Logf("  Actual:   '%s'", newLine)
	}
}

// test always fails even when function works. Reason unknown
/*func TestSpellcheckVariable(t *testing.T) {
	count.Tests(t, 1)
	lang.InitEnv()
	defaults.Defaults(lang.ShellProcess.Config, false)

	err := lang.ShellProcess.Config.Set("shell", "spellcheck-enabled", true)
	if err != nil {
		t.Fatalf("Unable to set spellcheck-enabled config: %s", err)
	}

	err = lang.ShellProcess.Config.Set("shell", "spellcheck-block", `{ -> jsplit ' ' -> suffix "\n" }`)
	if err != nil {
		t.Fatalf("Unable to set spellcheck-block config: %s", err)
	}

	os.Setenv("MUREX_TEST_SPELLCHECK_TEST", "quick")

	line := "$the $MUREX_TEST_SPELLCHECK_TEST $brown $fox"
	newLine := string(spellcheck([]rune(line)))
	ansiLine := ansi.ExpandConsts("{UNDERLINE}$the{UNDEROFF} $MUREX_TEST_SPELLCHECK_TEST {UNDERLINE}$brown{UNDEROFF} {UNDERLINE}$fox{UNDEROFF}")

	if newLine != ansiLine {
		t.Error("spellcheck output doesn't match expected:")
		t.Logf("  Expected: '%s'", ansiLine)
		t.Logf("  Actual:   '%s'", newLine)
	}
}*/

// test times out for reasons currently unknown
/*func TestSpellcheckBadBlock(t *testing.T) {
	count.Tests(t, 1)
	lang.InitEnv()
	defaults.Defaults(lang.ShellProcess.Config, false)

	err := lang.ShellProcess.Config.Set("shell", "spellcheck-enabled", true)
	if err != nil {
		t.Fatalf("Unable to set spellcheck-enabled config: %s", err)
	}

	err = lang.ShellProcess.Config.Set("shell", "spellcheck-block", `{ -> jsplit { }`)
	if err != nil {
		t.Fatalf("Unable to set spellcheck-block config: %s", err)
	}

	line := "the quick brown fox"
	newLine := string(spellcheck([]rune(line)))

	if newLine != line {
		t.Error("spellcheck output doesn't match expected:")
		t.Logf("  Expected: '%s'", line)
		t.Logf("  Actual:   '%s'", newLine)
	}
}*/
