package port

import (
	"net"
	"fmt"
)

type ConnStatus struct {
	port int
	status string
}

func ScanPorts(protocol, host string) []ConnStatus {
	ports := []ConnStatus{}

	for i := 0; i < 1025; i++ {
		address := fmt.Sprintf("%s:%d", host, i)
		_, err := net.Dial(protocol, address)
	
		if err != nil {
			connStatus := ConnStatus{
				port: i,
				status: "Closed",
			}
			ports = append(ports, connStatus)
			result := fmt.Sprintf("Port %d: %s", i, connStatus.status)
			fmt.Println(result)
		} else {
			connStatus := ConnStatus{
				port: i,
				status: "Open",
			}
			ports = append(ports, connStatus)
			result := fmt.Sprintf("Port %d: %s", i, connStatus.status)
			fmt.Println(result)
		}
	}

	return ports
}