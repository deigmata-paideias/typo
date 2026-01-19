package version

import (
	"encoding/json"
	"fmt"
	"io"
	"runtime"

	"gopkg.in/yaml.v3"
)

type Info struct {
	TypoVersion   string `json:"envoyGatewayVersion"`
	GitCommitID   string `json:"gitCommitID"`
	GolangVersion string `json:"golangVersion"`
}

func Get() Info {

	return Info{
		TypoVersion:   typoVersion,
		GitCommitID:   gitCommitID,
		GolangVersion: runtime.Version(),
	}
}

var (
	typoVersion string
	gitCommitID string
)

// Print shows the versions of the typo.
func Print(w io.Writer, format string) error {

	v := Get()

	switch format {
	case "json":
		if marshalled, err := json.MarshalIndent(v, "", "  "); err == nil {
			_, _ = fmt.Fprintln(w, string(marshalled))
		}
	case "yaml":
		if marshalled, err := yaml.Marshal(v); err == nil {
			_, _ = fmt.Fprintln(w, string(marshalled))
		}
	default:
		_, _ = fmt.Fprintf(w, "TYPO_GO_VERSION: %s\n", v.TypoVersion)
		_, _ = fmt.Fprintf(w, "GIT_COMMIT_ID: %s\n", v.GitCommitID)
		_, _ = fmt.Fprintf(w, "GOLANG_VERSION: %s\n", v.GolangVersion)
	}

	return nil
}
