module github.com/cds-amal/go-mono

go 1.21.6

toolchain go1.21.11

require github.com/DIN-center/din-sc/apps/din-go v0.0.0-20240605161942-b2e562f2da9c // pseudo-version to satisfy go mod requirements

require (
	github.com/golang/mock v1.6.0 // indirect
	golang.org/x/mod v0.14.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/tools v0.15.0 // indirect
)

// specify the version with the test branch, this should change when we merge the PR
replace github.com/DIN-center/din-sc/apps/din-go => github.com/DIN-center/din-sc/apps/din-go v0.0.0-20240607023748-1c5f62c5f131
