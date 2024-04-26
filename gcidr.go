package main

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"sort"
)

func ip2int(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

type Subnet struct {
	CIDR  string   `json:"cidr"`
	Count int      `json:"count"`
	IPs   []string `json:"ips"`
}

func cidr2int(cidr string) uint32 {
	ip, _, _ := net.ParseCIDR(cidr)
	return ip2int(ip)
}

func main() {
	flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	listPtr := flags.String("l", "", "The path to the IP list file")
	jsonPtr := flags.Bool("json", false, "Output the result as JSON")
	sortAscCIDRPtr := flags.Bool("sa-cidr", false, "Sort the result in ascending order by CIDR")
	sortDescCIDRPtr := flags.Bool("sd-cidr", false, "Sort the result in descending order by CIDR")
	sortAscCountPtr := flags.Bool("sa-count", false, "Sort the result in ascending order by count")
	sortDescCountPtr := flags.Bool("sd-count", false, "Sort the result in descending order by count")

	flags.Usage = func() {
		fmt.Printf("USAGE: %s [OPTIONS]\n\nOptions:\n", os.Args[0])
		flags.PrintDefaults()
	}

	flags.Parse(os.Args[1:])

	var input io.Reader
	if *listPtr != "" {
		file, err := os.Open(*listPtr)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		input = file
	} else {
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) != 0 {
			flags.Usage()
			return
		}
		input = os.Stdin
	}

	subnets := make(map[string][]string)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		ip := net.ParseIP(scanner.Text())
		if ip != nil {
			cidr := generateCIDR(ip)
			subnets[cidr] = append(subnets[cidr], scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	subnetList := make([]Subnet, 0, len(subnets))
	for cidr, ips := range subnets {
		subnetList = append(subnetList, Subnet{CIDR: cidr, Count: len(ips), IPs: ips})
	}

	if *sortAscCIDRPtr {
		sort.Slice(subnetList, func(i, j int) bool {
			return cidr2int(subnetList[i].CIDR) < cidr2int(subnetList[j].CIDR)
		})
	}

	if *sortDescCIDRPtr {
		sort.Slice(subnetList, func(i, j int) bool {
			return cidr2int(subnetList[i].CIDR) > cidr2int(subnetList[j].CIDR)
		})
	}

	if *sortAscCountPtr {
		sort.Slice(subnetList, func(i, j int) bool {
			return subnetList[i].Count < subnetList[j].Count
		})
	}

	if *sortDescCountPtr {
		sort.Slice(subnetList, func(i, j int) bool {
			return subnetList[i].Count > subnetList[j].Count
		})
	}

	if *jsonPtr {
		data, err := json.MarshalIndent(subnetList, "", "  ")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(data))
	} else {
		for _, subnet := range subnetList {
			fmt.Printf("CIDR: %s\tCount: %d\tIPs: %v\n", subnet.CIDR, subnet.Count, subnet.IPs)
		}
	}
}

func generateCIDR(ip net.IP) string {
	ip = ip.To4()
	ip[3] = 0
	return fmt.Sprintf("%s/24", ip.String())
}
