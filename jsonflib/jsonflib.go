package jsonflib

import (
	"log"
	"regexp"
	"github.com/robertkrimen/otto"
)

type Rule struct {
	Expr    string
	Replace string
}

var Rules = []Rule {{"",""},
	{`(?m)(\"[^\"]+\"):`,                     "\033[36m$1\033[39m:"},
	{`(?m)(^\s*[{}\]]{1}[,]*$|: [{\[][}\]]?)`, "\033[33m$1\033[39m"},
	{`: (\"[^\"]+\")`,                     ": \033[31m$1\033[39m$2"},
	{`(?m): ([\d][\d\.e+]*)([,]*)$`,                       ": \033[33m$1\033[39m"},
	{`(?m)(^\s+(?:[\d][\d\.e+]*|true|false|null))([,]*)$`, "\033[35m$1\033[39m$2"},
	{`(?::) ((?:true|false|null))`,                        ": \033[31m$1\033[39m"},
	{`(?m)^(true|false|null)$`,                              "\033[31m$1\033[39m"},
}

func Query(query string, v interface{}) (interface{}, error) {
	vm := otto.New()
	vm.Set("value", v)
	vmResult, err := vm.Run(query)
	if err == nil {
		v, err = vmResult.Export()
	}
	return v, err
}

type HighlightFlags struct {
	Colorize, Verbose, Debug bool
}

func Highlight(data []byte, highlightFlags HighlightFlags) ([]byte, error) {
	var (
		re *regexp.Regexp
		err  error
	)
	for _, rule := range Rules {
		if rule.Expr == "" {
			continue
		}
		if highlightFlags.Debug {
			log.Printf("compile highlight re: %#v", rule)
		}
		re, err = regexp.Compile(rule.Expr)
		if err != nil {
			return nil, err
		}
		data = re.ReplaceAll(data, []byte(rule.Replace))
	}
	if highlightFlags.Verbose {
		if highlightFlags.Colorize {
			log.Printf("colorize enabled")
		} else {
			log.Printf("colorize disabled")
		}
	}
	if !highlightFlags.Colorize {
		re, err = regexp.Compile(`\033\[[0-9;]*m`)
		if err != nil {
			return nil, err
		}
		data = re.ReplaceAll(data, []byte(""))
	}
	return data, nil
}
