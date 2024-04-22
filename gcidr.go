package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"sort"
)

type Subnet struct {
	CIDR  string   `json:"cidr"`
	Count int      `json:"count"`
	IPs   []string `json:"ips"`
}

func main() {
	flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	listPtr := flags.String("l", "", "The path to the IP list file")
	jsonPtr := flags.Bool("j", false, "Output the result as JSON")
	sortAscPtr := flags.Bool("sa", false, "Sort the result in ascending order by count")
	sortDescPtr := flags.Bool("sd", false, "Sort the result in descending order by count")
	outputPtr := flags.String("o", "", "The path to save the result")

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

	if *sortAscPtr {
		sort.Slice(subnetList, func(i, j int) bool {
			return subnetList[i].Count < subnetList[j].Count
		})
	}

	if *sortDescPtr {
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
		if *outputPtr != "" {
			err := os.WriteFile(*outputPtr, data, 0644)
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			fmt.Println(string(data))
		}
	} else {
		for _, subnet := range subnetList {
			fmt.Printf("CIDR: %s, Count: %d, IPs: %v\n", subnet.CIDR, subnet.Count, subnet.IPs)
		}
	}
}

func generateCIDR(ip net.IP) string {
	ip = ip.To4()
	ip[3] = 0
	return fmt.Sprintf("%s/24", ip.String())
}
