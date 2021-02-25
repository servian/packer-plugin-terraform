package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hashicorp/packer-plugin-sdk/plugin"
)

// Version number constant.
const Version = "0.0.6"

var (
	versDisp = flag.Bool("version", false, "Display version")
)

func main() {
	flag.Parse()

	if *versDisp {
		fmt.Printf("Version: v%s\n", Version)
		os.Exit(0)
	}

	pps := plugin.NewSet()
	pps.RegisterProvisioner(plugin.DEFAULT_NAME, new(Provisioner))
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
