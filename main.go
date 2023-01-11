package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var HOST string
var PASSWORD string
var USERNAME string
var NODE string

func printHelp() {
	fmt.Println("Usage: [cmd] VMID")
	fmt.Println("  start     start VMID")
	fmt.Println("  stop      stop  VMID")
	fmt.Println("  reset     reset VMID")
	fmt.Println("  shutdown  shutdown VMID")
	fmt.Println("  reboot    reboot VMID")
	fmt.Println("  suspend   suspend VMID")
	fmt.Println("")
	fmt.Println("running with no arguments prints out VM status")
}

func doCommand() bool {
	switch os.Args[1] {
	case "start":
		toggleVM("start", os.Args[2])
	case "stop":
		toggleVM("stop", os.Args[2])
	case "reset":
		toggleVM("reset", os.Args[2])
	case "shutdown":
		toggleVM("shutdown", os.Args[2])
	case "reboot":
		toggleVM("reboot", os.Args[2])
	case "suspend":
		toggleVM("suspend", os.Args[2])
	default:
		return false
	}
	return true
}

func validateInput() {
	if len(HOST) == 0 {
		fmt.Println("Error reading config file. 'Host' field invalid")
		os.Exit(1)
	}
	if len(USERNAME) == 0 {
		fmt.Println("Error reading config file. 'User' field invalid")
		os.Exit(1)
	}
	if len(PASSWORD) == 0 {
		fmt.Println("Error reading config file. 'Password' field invalid")
		os.Exit(1)
	}
}

func main() {
	viper.SetConfigName("remoteproxmoxcli")
	viper.AddConfigPath("$HOME/.config/")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file. Please modify example.cfg and copy it to ~/.config/remoteproxmoxcli.yml")
		return
	}
	HOST = viper.GetString("Host")
	USERNAME = viper.GetString("User")
	PASSWORD = viper.GetString("Pass")
	validateInput()

	getTicket()
	nodeList := getNodes()
	NODE = nodeList[0].Node
	vmList := listVMs()

	switch len(os.Args) {
	case 2:
		printHelp()
	case 3:
		if doCommand() == false {
			printVmStatus(nodeList[0], vmList)
		}
	default:
		printVmStatus(nodeList[0], vmList)

	}
}
