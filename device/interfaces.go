package device

import (
	"net"

	"github.com/TheHippo/mnc/util"
)

// Interface contains basic information about a network interface
type Interface struct {
	Name string
	IP   string
}

// GetInterfaces return all useable network interfaces for this computer
func GetInterfaces(listOnlyIP4 bool) ([]*Interface, error) {
	ifaces, err := net.Interfaces()
	res := make([]*Interface, 0)
	if err != nil {
		return res, err
	}
	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			return res, err
		}
		for _, addr := range addrs {
			switch v := addr.(type) {
			// case *net.IPAddr:
			// 	fmt.Printf("%v : %s (%s)\n", iface.Name, v, v.IP.DefaultMask())

			case *net.IPNet:
				// fmt.Printf("%v : %s [%v/%v]\n", iface.Name, v, v.IP, v.Mask)
				if (v.IP.To4() != nil && listOnlyIP4) || !listOnlyIP4 {
					res = append(res, &Interface{
						Name: iface.Name,
						IP:   v.IP.String(),
					})
				}
			}
		}

	}
	return res, nil
}

// MakePrettyIter make iterator for prettier output
func MakePrettyIter(i []*Interface) *InterfaceIter {
	return &InterfaceIter{
		ifaces: i,
	}
}

// Preface returns first part for pretty print output
func (i *Interface) Preface() string {
	return i.Name
}

// Remainder returns the remaining part of the pretty print output
func (i *Interface) Remainder() string {
	return i.IP
}

// InterfaceIter iterator for interface output
type InterfaceIter struct {
	ifaces  []*Interface
	current int
}

// HasNext wether another interface is available
func (i *InterfaceIter) HasNext() bool {
	return i.current < len(i.ifaces)
}

// Next returns the next interface
func (i *InterfaceIter) Next() util.Prettifyer {
	i.current++
	return i.ifaces[i.current-1]
}

// Reset the iterator
func (i *InterfaceIter) Reset() {
	i.current = 0
}
