# acsm_db_convert

Convert Asetto Corsa Server Manager database from BoltDB to JSON

Note, this is very crufy and my very first attempt at any Go code - it is written to get the job done, not particularly do anything right.

---

## How to Use

Install Go - I'll publish binaries if/when I get around to learning how.

Copy your server_manager.db file to the current directory.

run
`go run main.go`

It should create a shared_store.json directory. Copy this to your server and set the appropriate permissions.
Then edit `config.yml` in your acsm server to use this instead of boltdb.

### For multi servers, 
run the assetto-multiserver-manager once and stop it (so it creates the 'servers' directory structure.
Copy the `shared_store.json/server_options.json` file to `servers/SERVER_00/store.json/server_options.json`

Restart servers/SERVER_00/store.json/server_options.json
