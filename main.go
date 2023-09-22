package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"os/user"

	"github.com/zcalusic/sysinfo"
)

func main() {

	outCommand, err := exec.Command("ls", "-ahl").Output()

	if err != nil {
		fmt.Printf("%s", err)
	}

	output := string(outCommand[:])
	fmt.Println(output)

	current, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	if current.Uid != "0" {
		// log.Fatal("requires superuser privilege")
		fmt.Println("Probably should try doing this as Root")
	}

	var si sysinfo.SysInfo

	si.GetSysInfo()

	data, err := json.MarshalIndent(&si, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
