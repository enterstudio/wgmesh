
package main

import (
	"os"
	"fmt"
	"github.com/jessevdk/go-flags"
)

type Options struct {
	Verbose	[]bool `short:"v"`
	Mode	string `short:"m" long:"mode"`
	Dev	string `short:"i" long:"interface"`
	Port	int `short:"p" long:"port"`
	
	/* Master Options */
	ListenOn string `short:"l" long:"listen-on"`

	/* Slave Options */
	MasterPubkies	[]string `short:"k" long:"key"`
	MasterEndPoints	[]string `short:"e" long:"endpoint"`
	MasterDistances	[]int `short:"d" long:"distance"`
	MasterIPAddrs	[]string `short:"a" long:"addr"`
	
	Groups		[]string `short:"g" long:"group"`
}


func main() {

	var opts Options

	_, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		panic(err)
	}

	switch opts.Mode {
	case "master"	: Master(opts)
	case "slave"	: Slave(opts)
	case "test"	: Test(opts)
	default:
		fmt.Fprintf(os.Stderr, "invalid mode '%s'\n", opts.Mode)
		return
	}
}
