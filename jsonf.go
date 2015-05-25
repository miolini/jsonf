package main

import (
	"encoding/json"
	"fmt"
	"os"
	"log"
	"flag"
	"regexp"
	"github.com/robertkrimen/otto"
)

type Rule struct {
	Expr    string
	Replace string
}

var Rules = []Rule {{"",""},
	{`(?m)(\"[^\"]+\"):`,                     "\033[36m$1\033[39m:"},
	{`(?m)(^\s*[{}\]]{1}[,]*$|: [{\[])`, "\033[33m$1\033[39m"},
	{`: (\"[^\"]+\")`,                     ": \033[31m$1\033[39m$2"},
	{`(?m): ([\d][\d\.e+]*)([,]*)$`,                       ": \033[33m$1\033[39m"},
	{`(?m)(^\s+(?:[\d][\d\.e+]*|true|false|null))([,]*)$`, "\033[35m$1\033[39m$2"},
	{`(?::) ((?:true|false|null))`,                        ": \033[31m$1\033[39m"},
	{`(?m)^(true|false|null)$`,                              "\033[31m$1\033[39m"},
}

var (
	flSyntaxHighlight = flag.Bool("s", true, "syntax hightlight")
	flQuery           = flag.String("q", "", "json query")
	flFormat          = flag.Bool("f", true, "format output json to stdout")
	flVerbose         = flag.Bool("v", false, "verbose output to stderr")
	flDebug           = flag.Bool("d", false, "debug output to stderr")
	flColorize        = flag.Bool("c", true, "colorize output")
)

func main() {
	var (
		err  error
		data []byte
		v    interface{}
	)
	
	// Flags parsing
	flag.Parse()
	
	// Json Decoding
	err = json.NewDecoder(os.Stdin).Decode(&v)
	if err != nil {
		log.Fatalf("json decode error: %s", err)
	}
	
	// Query
	if *flQuery != "" {
		if v, err = query(*flQuery, v); err != nil {
			log.Fatalf("query err: %s", err)
		}
	}
	
	// Json Encode
	if *flFormat {
		data, err = json.MarshalIndent(v, "", "  ")
	} else {
		data, err = json.Marshal(v)
	}
	if err != nil {
		log.Fatalf("json encode error: %s", err)
	}
	
	// Syntax Highlight
	if *flFormat && *flSyntaxHighlight {
		if data, err = highlight(data); err != nil {
			log.Fatalf("highlight error: %s", err)
		}
	}
	
	// Print result
	fmt.Println(string(data))
}

func query(query string, v interface{}) (interface{}, error) {
	vm := otto.New()
	vm.Set("value", v)
	vmResult, err := vm.Run(query)
	if err == nil {
		v, err = vmResult.Export()
	}
	return v, err
}

func highlight(data []byte) ([]byte, error) {
	var (
		re *regexp.Regexp
		err  error
	)
	for _, rule := range Rules {
		if rule.Expr == "" {
			continue
		}
		if *flDebug {
			log.Printf("compile highlight re: %#v", rule)
		}
		re, err = regexp.Compile(rule.Expr)
		if err != nil {
			return nil, err
		}
		data = re.ReplaceAll(data, []byte(rule.Replace))
	}
	if *flVerbose {
		if *flColorize {
			log.Printf("colorize enabled")
		} else {
			log.Printf("colorize disabled")
		}
	}
	if !*flColorize {
		re, err = regexp.Compile(`\033\[[0-9;]*m`)
		if err != nil {
			return nil, err
		}
		data = re.ReplaceAll(data, []byte(""))
	}
	return data, nil
}
