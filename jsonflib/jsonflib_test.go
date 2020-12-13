package jsonflib_test

import (
	"testing"

	"github.com/miolini/jsonf/jsonflib"
	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	inputJSON := []byte(
		"\n{\n" +
			"  \"aptly\": [\n" +
			"    \"develop\"\n" +
			"  ]\n" +
			"}\n")
	expectedJSON := []byte(
		"\033[33m\n{\033[39m\n" +
			"  \033[36m\"aptly\"\033[39m\033[33m: [\033[39m\n" +
			"    \"develop\"\n" +
			"\033[33m  ]\033[39m\n" +
			"\033[33m}\033[39m\n")

	outputJSON, err := jsonflib.Highlight(
		inputJSON,
		jsonflib.HighlightFlags{Colorize: true, Verbose: false, Debug: false},
	)
	assert.Nil(t, err)
	assert.Equal(t, outputJSON, expectedJSON)
}

func TestNumberComa(t *testing.T) {
	inputJSON := []byte(
		"\n{\n" +
			"	\"id\": 1,\n" +
			"	\"num\": 2\n" +
			"}")
	expectedJSON := []byte(
		"\033[33m\n{\033[39m\n" +
			"	\033[36m\"id\"\033[39m: \033[33m1\033[39m,\n" +
			"	\033[36m\"num\"\033[39m: \033[33m2\033[39m\n" +
			"\033[33m}\033[39m")
	outputJSON, err := jsonflib.Highlight(
		inputJSON,
		jsonflib.HighlightFlags{Colorize: true, Verbose: false, Debug: false},
	)
	assert.Nil(t, err)
	assert.Equal(t, expectedJSON, outputJSON)
}
