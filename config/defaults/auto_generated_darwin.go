// +build darwin

package defaults

/*
   WARNING:
   --------

   This Go source file has been automatically generated from
   profile_osx.mx using docgen.

   Please do not manually edit this file because it will be automatically
   overwritten by the build pipeline. Instead please edit the aforementioned
   profile_osx.mx file located in the same directory.
*/

func init() {
	murexProfile = append(murexProfile, "autocomplete set brew { [\n\t{\n\t\t\"Flags\": [\n            # Built-in commands\n            \"--cache\", \"--env\", \"--version\", \"cleanup\", \"deps\", \"fetch\", \"home\", \"leaves\", \"log\", \"options\", \"postinstall\", \"search\", \"tap-info\", \"unlink\", \"update-report\", \"upgrade\", \"--caskroom\", \"--prefix\", \"analytics\", \"commands\", \"desc\", \"gist-logs\", \"info\", \"link\", \"migrate\", \"outdated\", \"readall\", \"shellenv\", \"tap\", \"unpin\", \"update-reset\", \"uses\", \"--cellar\", \"--repository\", \"cask\", \"config\", \"doctor\", \"help\", \"install\", \"list\", \"missing\", \"pin\", \"reinstall\", \"switch\", \"uninstall\", \"untap\", \"update\", \"vendor-install\",\n            \n            # Built-in developer commands\n            \"audit\", \"bump-revision\", \"create\", \"extract\", \"linkage\", \"pr-automerge\", \"prof\", \"sh\", \"test\", \"update-license-data\", \"bottle\", \"bump\", \"dispatch-build-bottle\", \"formula\", \"livecheck\", \"pr-publish\", \"pull\", \"sponsors\", \"tests\", \"update-python-resources\", \"bump-cask-pr\", \"cat\", \"diy\", \"install-bundler-gems\", \"man\", \"pr-pull\", \"release-notes\", \"style\", \"typecheck\", \"update-test\", \"bump-formula-pr\", \"command\", \"edit\", \"irb\", \"mirror\", \"pr-upload\", \"ruby\", \"tap-new\", \"unpack\", \"vendor-gems\",\n            \n            # External commands\n            \"aspell-dictionaries\", \"postgresql-upgrade-database\"\n        ],\n        \"FlagValues\": {\n            \"cask\": [{\n                \"Flags\": [\n                    # Cask commands\n                    \"--cache\", \"_help\", \"_stanza\", \"audit\", \"cat\", \"create\", \"doctor\", \"edit\", \"fetch\", \"help\", \"home\", \"info\", \"install\", \"list\", \"outdated\", \"reinstall\", \"style\", \"uninstall\", \"upgrade\", \"zap\",\n\n                    # External cask commands\n                    \"ci\"\n                ]\n            }]\n        }\n\t},\n    {\n        \"DynamicDesc\": ({ \n            brew help @{ $ARGS->@[1..] } -> tabulate: --map --split-comma --column-wraps\n        }),\n        \"ListView\": true\n    }\n] }\n\nnull {\n# brew commands -> grep -v '==>' -> cast str\nmkautocomplete {\n    # Compiles a persistant autocomplete file optimised for executables with a slow launch time\n\n    set usage = (Example usage:\n$ARGS[0] cmd {\n    \"VersionPin\":      ({ cmd -v }),      # --- required\n    \"topDynamic\":      ({ code-block }),  # \\__ one or\n    \"topDynamicDesc\":  ({ code-block }),  # /     both\n    \"topInclude\":      { json },          # --- optional\n    \"flagDynamic\":     ({ code-block }),  # \\__ one or\n    \"flagDynamicDesc\": ({ code-block }),  # /     other\n    \"flagInclude\":     { json }           # --- optional\n})\n\n    mkdir -p ~/.murex_modules/.mkautocomplete\n\n    !if { $ARGS[2] } then {\n        err \"Invalid usage!\"\n        err $usage\n    }\n\n    !if { $ARGS[2]->[VersionPin] } then {\n        err \"Invalid usage!\"\n        err $usage\n    }\n\n    set cmd=$ARGS[1]\n    open <!null> ~/.murex_modules/.mkautocomplete/$cmd.version -> set oldVersion\n    open <!null> ~/.murex_modules/.mkautocomplete/$cmd.autocomplete -> set oldAutocomplete\n\n    set f=${$ARGS[2]} # lazier flag notation\n\n    trypipe {\n        source $f[VersionPin] -> set newVersion\n        if { \"$oldVersion\" == \"$newVersion\" } then {\n            autocomplete $cmd $oldAutocomplete\n            false\n        }\n\n        if { $f[topDynamic] } then {\n            source $f[topDynamic] -> set topDynamic\n        } else {\n            set json topDynamic=\"[]\"\n        }\n        \n        if { $f[topDynamicDesc] } then {\n            source $f[topDynamicDesc] -> set topDynamicDesc\n        } else {\n            set json topDynamicDesc=\"{}\"\n        }\n\n        #if { $f[flagDynamic] } then {\n        #    source $f[flagDynamic] -> foreach\n        #}  \n\n        tout json ({\n\n})\n    }\n}\n}")
}
