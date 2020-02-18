package main

import (
	"fmt"
	"strconv"
	"net"
	"time"
	"net/icmp"
)

func main() {
	host := "baidu.com"
	ips, err := net.LookupIP(host)
	if err != nil {
		fmt.Println(err)
	}
	for _, ip := range ips {
		fmt.Println(ip.String())
	}
	i := icmp.Echo{}
}
