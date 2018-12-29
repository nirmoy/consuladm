package netutils

import (
	"fmt"
	"net"

	"golang.org/x/sys/unix"

	"github.com/vishvananda/netlink"
)

//taken from https://github.com/cilium/cilium/blob/36cdd98c7278c87c4034c7a8fbdc57b74271ecbc/pkg/node/node_address_linux.go#L28

func FirstGlobalV4Addr(intf string) (net.IP, error) {
	var link netlink.Link
	var err error

	addr, err := netlink.AddrList(link, netlink.FAMILY_V4)
	if err != nil {
		return nil, err
	}

	for _, a := range addr {
		if a.Scope == unix.RT_SCOPE_UNIVERSE {
			if len(a.IP) >= 4 {
				return a.IP, nil
			}
		}
	}

	return nil, fmt.Errorf("No address found")
}
