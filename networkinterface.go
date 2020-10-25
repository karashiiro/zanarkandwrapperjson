package main

import (
	"net"
	"regexp"

	"github.com/ayyaruq/zanarkand/devices"
)

// GetNetworkInterface gets the network interface name corresponding to the provided IP address.
// If the input value is nil, the default network device is returned.
func GetNetworkInterface(networkDevice net.IP) (string, error) {
	// Get the default network device
	// Looking for ranges:
	// 24-bit private block: 10.*.*.*
	// 20-bit private block: 172.16.*.*-172.31.*.*
	// 16-bit private block: 192.168.*.*
	// 16-bit APIPA block: 169.254.*.*
	_, subnet24, _ := net.ParseCIDR("10.0.0.0/8")
	_, subnet20, _ := net.ParseCIDR("172.16.0.0/12")
	_, subnet16, _ := net.ParseCIDR("192.168.0.0/16")
	_, APIPA16, _ := net.ParseCIDR("169.254.0.0/16")

	rNetIfaces, err := devices.ListDeviceNames(false, true)
	if err != nil {
		return "", err
	}
	netIfaces, _ := devices.ListDeviceNames(false, false) // It seems safe to assume that if the call that includes this and more doesn't error, this will not error if called immediately after either
	netIfaceIdx := len(netIfaces) - 1
	for i, nif := range rNetIfaces {
		ip := net.ParseIP(regexp.MustCompile("\\d+\\.\\d+\\.\\d+\\.\\d+").FindString(nif)) // Lazy but it works and it only runs once
		if networkDevice != nil && networkDevice.Equal(ip) {
			netIfaceIdx = i
			break
		}
		if subnet24.Contains(ip) || subnet20.Contains(ip) || subnet16.Contains(ip) || APIPA16.Contains(ip) {
			netIfaceIdx = i
		}
	}

	return netIfaces[netIfaceIdx], nil
}
