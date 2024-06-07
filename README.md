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
   go.sum:
       github.com/DIN-center/din-sc/apps/din-go v0.0.0-20240607023748-1c5f62c5f131 h1:6dSyF5bsvx99Omf3NMiWnR7NV0CbaVij7O1ut3zdKVU=
       github.com/DIN-center/din-sc/apps/din-go v0.0.0-20240607023748-1c5f62c5f131/go.mod h1:lfYyZ8F8CIEcyIRyE+Hwso1xMfC+OvGoDgVJ8E0YIWk=

   Yay... go mod download - did its thing!
   ```


