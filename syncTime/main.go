package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

var syncURI = []string{
	"5time.nist.gov",
	"time-nw.nist.gov",
	"time-a.nist.gov",
	"time-b.nist.gov",
	"tick.mit.edu",
	"time.windows.com",
	"clock.sgi.com",
	"13.65.245.138",
}

func main() {
	var listen net.Listener
	var err error
	for _, url := range syncURI {
		remoteIPs := getRemoteIP(url)
		if remoteIPs != nil {
			for _, ip := range remoteIPs {
				listen, err = net.Listen("tcp", ip+":13")
				if err != nil {
					fmt.Println(err)
					continue
				}
				break
			}
		}
	}
	if listen == nil {
		log.Println("none tcp begin listen")
		os.Exit(1)
	}

	defer listen.Close()
	log.Println("tcp begin listen")

	for {
		con, err := listen.Accept()
		if err != nil {
			continue
		}

		log.Println(con.RemoteAddr().String(), " tcp connect succeed")
	}
}

func getIP() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}
		}
	}
}

func getRemoteIP(url string) []string {
	ns, err := net.LookupHost(url)
	if err != nil {
		//fmt.Fprintf(os.Stderr, "Err: %s", err.Error())
		return nil
	}

	// fmt.Fprintf(os.Stdout, "----%s\n", url)
	// for _, n := range ns {
	// 	fmt.Fprintf(os.Stdout, "----%s\n", n)
	// }

	return ns
}
