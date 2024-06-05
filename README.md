## TLDR

1. Set `GOPRIVATE`
   ```console
   export GOPRIVATE=github.com/DIN-center
   ```
1. Run go mod download
   ```console
   go mod download github.com/DIN-center/din-sc/apps/din-go
   ```

## Test it

1. Run the test script
   ```console
   ➜ ./test.sh
   rm: ./go.sum: No such file or directory
   ======================================================================================
   Listing contents of /Users/amal/go/pkg/mod/github.com/\!d\!i\!n-center/din-sc/apps
   ======================================================================================

   /Users/amal/go/pkg/mod/github.com/!d!i!n-center/din-sc/apps
   └── din-go@v0.0.0-20240607023748-1c5f62c5f131
       ├── Makefile
       ├── README.md
       ├── cmd
       │   └── main.go
       ├── go.mod
       ├── go.sum
       ├── lib
       │   └── din
       │       ├── client.go
       │       └── helpers.go
       ├── pkg
       │   └── dinregistry
       │       ├── contract.go
       │       ├── din_registry.go
       │       ├── din_registry_test.go
       │       ├── endpoint_collection.go
       │       ├── endpoint_collection_test.go
       │       ├── interface.go
       │       ├── interface_mock.go
       │       ├── provider.go
       │       ├── provider_test.go
       │       ├── service.go
       │       ├── service_test.go
       │       └── types.go
       ├── project.json
       └── tools.go

   7 directories, 21 files
   go.sum:
       github.com/DIN-center/din-sc/apps/din-go v0.0.0-20240607023748-1c5f62c5f131 h1:6dSyF5bsvx99Omf3NMiWnR7NV0CbaVij7O1ut3zdKVU=
       github.com/DIN-center/din-sc/apps/din-go v0.0.0-20240607023748-1c5f62c5f131/go.mod h1:lfYyZ8F8CIEcyIRyE+Hwso1xMfC+OvGoDgVJ8E0YIWk=

   Yay... go mod download - did its thing!
   ```


