// Package version is the exported surface of mock-api that other mock repos
// (mock-agent) import, mirroring how real CCF repos depend on api's packages.
package version

// Version is the library version of this module. Bump it when cutting a
// release so importers can see which version they were built against.
const Version = "0.1.0"

// Hello returns a greeting that names this module and its version.
func Hello() string {
	return "hello from mock-api " + Version
}
