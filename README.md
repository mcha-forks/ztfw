
# ztfw

Userspace libzt wrapper port forwarder.

```shell
usage: ztfw [<flags>]

Flags:
      --help                   Show context-sensitive help (also try --help-long and --help-man).
  -n, --network="8056c2e21c000001"  
                               zerotier network id
  -f, --forward-port="22"      port to forward (in listen mode)
  -a, --accept-port="2222"     port to accept (in connect mode)
  -u, --use-udp                UDP instead of TCP (TCP default)
  -c, --connect-to=CONNECT-TO  server (zerotier) ip to connect
      --version                Show application version.
```
