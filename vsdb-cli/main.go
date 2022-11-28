package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type jsonResponse struct {
	Status string
	Result string
}

func main() {
	fmt.Println("+--------------------+")
	fmt.Println("+ Very Slow Database +")
	fmt.Println("+--------------------+")
	fmt.Println("")

	var port int64 = 80

	args := os.Args[1:]

	for i, arg := range args {
		if arg == "-p" {
			portArg, err := strconv.ParseInt(args[i+1], 0, 64)

			if err == nil {
				port = portArg
			} else {
				fmt.Println("Invalid port")
				return
			}
		}
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)
		input = strings.Replace(input, "\r", "", -1)

		args := strings.Split(input, " ")

		if len(args) == 0 {
			continue
		}

		if strings.Compare(args[0], "get") == 0 {
			if len(args) < 2 {
				fmt.Println("Syntax: get <key>")
				continue
			}

			key := args[1]

			resp, err := http.Get("http://localhost:" + strconv.FormatInt(port, 10) + "/get?key=" + key)

			if err != nil {
				fmt.Println("Server error")
				fmt.Println(err)
				continue
			}

			defer resp.Body.Close()

			body, _ := io.ReadAll(resp.Body)

			respObj := jsonResponse{}

			_ = json.Unmarshal(body, &respObj)

			if respObj.Status == "not found" {
				fmt.Println("Entry not found")
				continue
			}

			if respObj.Status == "found" {
				fmt.Println(respObj.Result)
				continue
			}

		}

	}

}
