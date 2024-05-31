# gomon

Monitoring some basic stuff with go

Goals:

- Monitor certificates, TCP ports, and ping hosts (DONE)
- Read config from a config file (currently using toml) (DONE)
- Display a basic webpage with an ascii graph (DONE but does not persist through restarts. Doesn't bother me much though.)
- Send emails if host goes down (DONE)

Mostly I'm trying to replicate a hacky shell script I have written, and instead of relying on external tools, do more of a 'from scratch' approach.

Other features to add in the future:

- tests, mostly for config parsing
- save webpage to HTML file in addition to/instead of a webserver

## Dependencies:

- [pro-bing](github.com/prometheus-community/pro-bing) - to ping. Otherwise ICMP requires root access, and I'm not going to just shell out to the `ping` command.
    - requires setting sysctl value: `sudo sysctl -w net.ipv4.ping_group_range="0 2147483647"` otherwise weird behavior occurs
- [toml](github.com/BurntSushi/toml) - to parse toml config file.
- [email](github.com/jordan-wright/email) - to send emails. There is a way to do it with just net/smtp, but this simplifies things.


## Screenshot

Since it's just text, it'll look like this in your browser:

```
Time:2024-05-28 12:34:24.297703837 -0500 CDT m=+3.395556712
Cert checks:

Tcp checks:

192.168.7.1:22: +

192.168.7.11:53: +

192.168.7.11:80: +

192.168.7.11:22: +

192.168.7.1:80: +

192.168.7.1:443: +

Ping checks:

alpinelinux.org: +

postmarketos.org: +
```
