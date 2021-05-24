package main

import (
	"fmt"
	"github.com/go-ping/ping"
)

func main() {
	pinger, err := ping.NewPinger("129.254.75.33")
	if err != nil {
		panic(err)
	}
	pinger.OnRecv = func(pkt *ping.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt)
	}

	pinger.Count = 10
	pinger.Run() // blocks until finished

	stats := pinger.Statistics() // get send/receive/rtt stats
	fmt.Printf("round-trip min/avg/max/stddev/dupl_recv = %v/%v/%v/%v/%v bytes\n",
		stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt, stats.PacketsRecv * 24)

}