package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hashicorp/packer-plugin-sdk/plugin"
)

// Version number constant.
const Version = "0.0.5"

var (
	versDisp = flag.Bool("version", false, "Display version")
)

func main() {
	flag.Parse()

	if *versDisp {
		fmt.Printf("Version: v%s\n", Version)
		os.Exit(0)
	}

	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterProvisioner(&Provisioner{})
	server.Serve()
}
