package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

type ExploitApiJson []struct {
	ID   string `json:"id"`
	Date string `json:"date"`
	Name string `json:"name"`
}

func main() {
	argparserPage := flag.Int("range", 20, "goploit -range [Number]")
	argparserName := flag.String("name", "wordpress", "goploit -name [Framework Name]")
	flag.Parse()

	Banner()

	request, err := http.Get("https://www.exploitalert.com/api/search-exploit?name=" + *argparserName)
	defer request.Body.Close()

	if err != nil {
		fmt.Printf("An error has occurred: %s", err)
		return
	}

	var jsn ExploitApiJson

	json.NewDecoder(request.Body).Decode(&jsn)

	if len(jsn) <= 0 {
		fmt.Println("\n\033[1;31m—\033[0;0m I didn't find any exploit with that name.\n")
		return
	}

	fmt.Printf("I found the following Exploit's: \033[1;35m%d\033[0;0m\n", len(jsn))

	for i := 0; i < len(jsn); i++ {
		if i < *argparserPage {
			fmt.Printf("— \033[1;35mID:\033[0;0m %s \033[1;35mDate:\033[0;0m %s \033[1;35mName:\033[0;0m %s\n", jsn[i].ID, jsn[i].Date, jsn[i].Name)
		} else {
			break
		}
	}
}

func Banner() {
	fmt.Println("\033[1;96m", `———————————————————————————————————————————————————————————————————————
	██████╗  ██████╗ ██████╗ ██╗      ██████╗ ██╗████████╗
	██╔════╝ ██╔═══██╗██╔══██╗██║     ██╔═══██╗██║╚══██╔══╝
	██║  ███╗██║   ██║██████╔╝██║     ██║   ██║██║   ██║   
	██║   ██║██║   ██║██╔═══╝ ██║     ██║   ██║██║   ██║   
	╚██████╔╝╚██████╔╝██║     ███████╗╚██████╔╝██║   ██║   
	 ╚═════╝  ╚═════╝ ╚═╝     ╚══════╝ ╚═════╝ ╚═╝   ╚═╝   
		By: www.github.com/blkzy — Version: 1.1
———————————————————————————————————————————————————————————————————————`, "\033[0;0m")
}
