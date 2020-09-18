package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/ema/qdisc"

	mp "github.com/mackerelio/go-mackerel-plugin"
)

type QdiscPlugin struct {
	Prefix     string
	Interfaces []string
}

func (q QdiscPlugin) GraphDefinition() map[string]mp.Graphs {
	return map[string]mp.Graphs{
		"#": {
			Label: "Qdisc data Summary",
			Unit:  "integer",
			Metrics: []mp.Metrics{
				{Name: "tx_bytes", Label: "tx_bytes", Diff: true},
				{Name: "tx_packets", Label: "tx_packets", Diff: true},
				{Name: "tx_drops", Label: "tx_drops", Diff: false},
				{Name: "tx_overlimits", Label: "tx_overlimits", Diff: false},
				{Name: "tx_requeues", Label: "tx_requeues", Diff: false},
				{Name: "qlen", Label: "qlen", Diff: false},
				{Name: "backlog", Label: "backlog", Diff: false},
			},
		},
	}
}

func (q QdiscPlugin) FetchMetrics() (map[string]float64, error) {
	info, err := qdisc.Get()
	if err != nil {
		return nil, err
	}
	metrics := map[string]float64{}
	for _, msg := range info {
		for _, i := range q.Interfaces {
			if msg.IfaceName != i {
				continue
			}
			metrics[i+".tx_bytes"] = float64(msg.Bytes)
			metrics[i+".tx_packets"] = float64(msg.Packets)
			metrics[i+".tx_drops"] = float64(msg.Drops)
			metrics[i+".tx_overlimits"] = float64(msg.Overlimits)
			metrics[i+".tx_requeues"] = float64(msg.Requeues)
			metrics[i+".qlen"] = float64(msg.Qlen)
			metrics[i+".backlog"] = float64(msg.Backlog)
		}
	}

	if len(metrics) == 0 {
		return nil, fmt.Errorf("interface is not found")
	}

	return metrics, nil
}

func (q QdiscPlugin) MetricKeyPrefix() string {
	if q.Prefix == "" {
		q.Prefix = "qdisc"
	}
	return q.Prefix
}

// main
func main() {
	optPrefix := flag.String("metric-key-prefix", "qdisc", "Metric key prefix")
	optTempfile := flag.String("tempfile", "", "Temp file name")
	optInterface := flag.String("interface", "eth0", "network interface name")
	flag.Parse()

	interfaces := strings.Split(*optInterface, ",")
	q := QdiscPlugin{
		Prefix:     *optPrefix,
		Interfaces: interfaces,
	}

	plugin := mp.NewMackerelPlugin(q)
	plugin.Tempfile = *optTempfile
	plugin.Run()
}
