package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/miolini/jsonf/jsonflib"
	"os"
	"log"
)

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
		if v, err = jsonflib.Query(*flQuery, v); err != nil {
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
		data, err = jsonflib.Highlight(
			data,
			jsonflib.HighlightFlags{Colorize: *flColorize, Verbose: *flVerbose, Debug: *flDebug},
		)
		if err != nil {
			log.Fatalf("highlight error: %s", err)
		}
	}
	
	// Print result
	fmt.Println(string(data))
}
