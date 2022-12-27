package version

import (
	"fmt"
	"runtime"
)

const versionTemplate = `financial-period-api
Version: %s
Build: %s %s
GitSha: %s
Go: %s
`

var (
	Version = ""

	BuildTime = ""

	BuildNumber = ""

	GitSha = ""

	GoVersion = runtime.Version()
)

// DetailedVersion retuns detailed version
func DetailedVersion() string {
	return fmt.Sprintf(
		versionTemplate,
		Version, BuildNumber, BuildTime, GitSha, GoVersion)

}
