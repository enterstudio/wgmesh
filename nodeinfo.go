/* nodeinfo.go */

package main

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"strconv"
)


type EndPoint struct {
	addr	net.IP
	port	int
}

func (ep EndPoint) String() string {
	return fmt.Sprintf("%s:%d", ep.addr, ep.port)
}

func ParseEndPoint(epstr string) *EndPoint {

	s := strings.Split(epstr, ":")
	if len(s) > 2 {
		return nil
	}

	addrstr := s[0]
	portstr := s[1]

	addr := net.ParseIP(addrstr); 
	if addr == nil {
		return nil
	}
	
	port, err := strconv.Atoi(portstr)
	if err != nil {
		return nil
	}

	ep := &EndPoint{
		addr:	addr,
		port:	port,
	}

	return ep
}




type NodeInfo struct {
	/* common parameters */
	pubkey		string		// Wireguard Publickey
	addr		net.IP		// IP Addr the wg device has
	endpoint	EndPoint	// EndPoint of the wg device
	groups		[]string	// Groups this node belongs
	prefixes	[]net.IPNet	// Prefixes this node accommodates

	/* parameters Master node uses in Slave mode */
	distance	int
}



func (ni NodeInfo) String() string {

	var strprefixes []string

	for _, prefix := range ni.prefixes {
		strprefixes = append(strprefixes, prefix.String())
	}
	return fmt.Sprintf("<%s %s Group:[%s] Prefix:[%s]>",
		ni.pubkey,
		ni.addr,
		strings.Join(ni.groups, " "),
		strings.Join(strprefixes, " "))
}


func NewNodeInfo(pubkey string, addr_s string, ep_s string)(*NodeInfo, error) {

	/* validate params */
	addr := net.ParseIP(addr_s); 
	if addr == nil {
		return nil, fmt.Errorf("Invalid IP address '%s'", addr_s)
	}

	ep := ParseEndPoint(ep_s); 
	if ep == nil {
		return nil, fmt.Errorf("invalid EndPoint '%s'", ep_s)
	}

	ni := &NodeInfo{
		pubkey:	pubkey,
		addr:	addr,
		endpoint: *ep,
	}

	return ni, nil
}



type NodeInformationBase struct {
	table	map[string]*NodeInfo
	mutex	*sync.RWMutex
}

func NewNodeInformationBase() *NodeInformationBase{
	nib := &NodeInformationBase {
		table:	make(map[string]*NodeInfo),
		mutex:	new(sync.RWMutex),
	}
	return nib
}

func (nib *NodeInformationBase) AddNodeInfo(ni *NodeInfo) error {

	nib.mutex.Lock()
	defer nib.mutex.Unlock()

	
	if _, exist := nib.table[ni.pubkey]; exist {
		return fmt.Errorf("NodeInfo %s exists", ni.pubkey)
	}

	nib.table[ni.pubkey] = ni

	return nil
}

func (nib *NodeInformationBase) DelNodeInfo(pubkey string) error {
	nib.mutex.Lock()
	defer nib.mutex.Unlock()

	if _, exist := nib.table[pubkey]; exist {
		delete(nib.table, pubkey)
	} else {
		return fmt.Errorf("NodeInfo %s does not exist", pubkey)
	}

	return nil
}

func (nib *NodeInformationBase) FindNodeInfo(pubkey string) *NodeInfo {
	nib.mutex.RLock()
	defer nib.mutex.RUnlock()

	
	if ni, exist := nib.table[pubkey]; exist {
		return ni
	}

	return nil
}

func (nib *NodeInformationBase) ForeachNodeInfo() chan NodeInfo {

	iter := make(chan NodeInfo)

	go func() {
		nib.mutex.RLock()
		defer nib.mutex.RUnlock()
		for _, ni := range nib.table {
			iter <- *ni
		}
		close(iter)
	}()
	
	return iter
}

