# mackerel-plugin-qdisc
Linux qdisc(Queueing discipline) custom metrics plugin for mackerel.io agent.

# Install
```
$ mkr plugin install livesense-inc/mackerel-plugin-qdisc
```

# Synopsis
```
mackerel-plugin-qdisc [-metric-key-prefix=<prefix>] [-tempfile=<tempfile>] [-interface=<prefix>]
```

# Example of mackerel-agent.conf
```
[plugin.metrics.qdisc]
command = ["/path/to/mackerel-plugin-qdisc", "-interface=en0,en1"]
```

# Example
mackerel-plugin-qdisc gets queuing discipline information via netlink and outputs result as a mackerel plugin metric format.
```
$ /path/to/mackerel-plugin-qdisc -interface=en0
qdisc.en0.tx_bytes     39242.857143    1600254416
qdisc.en0.tx_packetbs  242.857143      1600254416
qdisc.en0.tx_drops     0       1600254416
qdisc.en0.tx_overlimits        0       1600254416
qdisc.en0.tx_requeues  0       1600254416
qdisc.en0.qlen 0       1600254416
qdisc.en0.backlog      0       1600254416
```