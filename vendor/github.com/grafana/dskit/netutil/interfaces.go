package netutil

import (
	"fmt"
	"net"
	"os"
)

func ipForAddr(addr net.Addr) (net.IP, error) {
	switch a := addr.(type) {
	case *net.IPAddr:
		return a.IP, nil
	case *net.IPNet:
		return a.IP, nil
	default:
		return net.IP{}, fmt.Errorf("no valid ip for address '%s'", addr.String())
	}
}

func PrivateNetworkInterfaces() []string {
	ifaces := []string{}

	all, err := net.Interfaces()
	if err != nil {
		return ifaces
	}

	for _, iface := range all {
		if iface.Flags&net.FlagLoopback == 0 && iface.Flags&net.FlagUp != 0 {
			addrs, err := iface.Addrs()
			if err != nil {
				continue
			}
			for _, addr := range addrs {
				ip, err := ipForAddr(addr)
				if err != nil {
					continue
				}
				if !ip.IsPrivate() {
					continue
				}
			}
			ifaces = append(ifaces, iface.Name)
		}
	}
	fmt.Fprintln(os.Stderr, ifaces)
	return ifaces
}
