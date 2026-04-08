package futuapi

// Version information / 版本信息
const (
	Version      = "0.4.0-dev"
	VersionMajor = 0
	VersionMinor = 4
	VersionPatch = 0
	VersionPre   = "dev"
)

// Build information (set via ldflags during release builds)
var (
	BuildTime    = "unknown"
	BuildCommit  = "unknown"
	BuildGoVer   = "unknown"
)

// VersionInfo returns detailed version and build information.
func VersionInfo() string {
	return Version
}
