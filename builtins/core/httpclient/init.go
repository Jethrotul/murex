package httpclient

import (
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"net/http"
)

func init() {
	proc.GoFunctions["get"] = cmdGet
	proc.GoFunctions["post"] = cmdPost
	proc.GoFunctions["getfile"] = cmdGetFile

	proc.GlobalConf.Define("http", "user-agent", config.Properties{
		Description: "User agent string for `get` and `getfile`.",
		Default:     config.AppName + "/" + config.Version,
		DataType:    types.String,
	})

	proc.GlobalConf.Define("http", "timeout", config.Properties{
		Description: "Timeout in seconds for `get` and `getfile`.",
		Default:     10,
		DataType:    types.Integer,
	})

	proc.GlobalConf.Define("http", "insecure", config.Properties{
		Description: "Ignore certificate errors.",
		Default:     false,
		DataType:    types.Boolean,
	})

	proc.GlobalConf.Define("http", "redirect", config.Properties{
		Description: "Automatically follow redirects.",
		Default:     true,
		DataType:    types.Boolean,
	})

	proc.GlobalConf.Define("http", "default-https", config.Properties{
		Description: "If true then when no protocol is specified (`http://` not `https://`) then default to `https://`.",
		Default:     false,
		DataType:    types.Boolean,
	})
}

type httpStatus struct {
	Code    int
	Message string
}

type jsonHttp struct {
	Status  httpStatus
	Headers http.Header
	Body    string
}