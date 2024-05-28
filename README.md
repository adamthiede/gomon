# gomon

Monitoring some basic stuff with go

Goals:

- Monitor certificates, TCP ports, and ping hosts (DONE)
- Read config from a config file (currently using toml) (DONE)
- Display a basic webpage with an ascii graph
- Send emails if host goes down

Mostly I'm trying to replicate a hacky shell script I have written, and instead of relying on external tools, do more of a 'from scratch' approach.

## Dependencies:

(pro-bing)[github.com/prometheus-community/pro-bing] - to ping. Otherwise ICMP requires root access, and I'm not going to just shell out to the `ping` command.
(toml)[github.com/BurntSushi/toml] - to parse toml config file.
