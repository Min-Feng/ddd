package helpertest

import (
	"regexp"

	"github.com/rs/zerolog/log"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/json"
)

func FormatToRawJSON(prettyJSON []byte) []byte {
	m := minify.New()
	m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), json.Minify)
	rawJSON, err := m.Bytes("application/json", prettyJSON)
	if err != nil {
		log.Fatal().Msgf("FormatToRawJSON: %v", err)
	}
	return rawJSON
}
