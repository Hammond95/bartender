package version

/* Go linker will set BuildTime, Commit and Release variables during compilation.

env GOOS=${OS} GOARCH=${ARCH} go build \
		-ldflags "-s -w -X ${PROJECT}/version.Release=${RELEASE} \
		-X ${PROJECT}/version.Commit=${COMMIT} -X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
		-o ../bin/gintonic-${RELEASE}-${OS}-${ARCH}
*/

var (
	// BuildTime is a time label of the moment when the binary was built
	BuildTime = "unset"
	// Commit is a last commit hash at the moment when the binary was built
	Commit = "unset"
	// Release is a semantic version of current build
	Release = "unset"
)
