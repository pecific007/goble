package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	connect := flag.Bool("c", false, "Connect to bluetooth device")
	disconnect := flag.Bool("d", false, "Disconnect from bluetooth device")
	sister := flag.Bool("s", false, "Connect to sister's earphones")

	flag.Parse()

	if len(os.Args) == 1 {
		flag.Usage()
		os.Exit(1)
	}

	boat := os.Getenv("boat")
	realme := os.Getenv("realme")

	var deviceMacAddr string
	if *sister {
		deviceMacAddr = realme
	} else {
		deviceMacAddr = boat
	}

	if *connect == *disconnect {
		fmt.Fprintln(os.Stderr, "Cannot pass -c (connect) and -d (disconnect) both.")
		flag.Usage()
		os.Exit(1)
	}

	if *connect {
		if err := runBluetoothCtl("connect", deviceMacAddr); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to connect %s\n", err)
			os.Exit(1)
		}
		if err := runBluetoothCtl("trust", deviceMacAddr); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to trust %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("Connected and trusted: %v", deviceMacAddr)
	}
	if *disconnect {
		if err := runBluetoothCtl("disconnect", deviceMacAddr); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to disconnect %s\n", err)
			os.Exit(1)
		}
		if err := runBluetoothCtl("untrust", deviceMacAddr); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to untrust %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("Disconnected and untrusted: %v", deviceMacAddr)
	}
}

func runBluetoothCtl(action, macAddr string) error {
	cmd := exec.Command("bluetoothctl", action, macAddr)
	out, err := cmd.Output()
	output := strings.TrimSpace(string(out))
	if output != "" {
		fmt.Println(output)
	}
	if err != nil {
		return fmt.Errorf("%v (output: %s)", err, output)
	}

	if strings.Contains(output, "failed") {
		return fmt.Errorf("bluetoothctl reported failure: %s", output)
	}
	return nil
}
