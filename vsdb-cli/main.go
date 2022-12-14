package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type jsonResponse struct {
	Status string
	Result map[string]string
}

func getResponse(url string) (jsonResponse, error) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Server error")
		fmt.Println(err)
		return jsonResponse{}, err
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	respObj := jsonResponse{}

	_ = json.Unmarshal(body, &respObj)

	return respObj, nil
}

func printTitle() {
	fmt.Println("+--------------------+")
	fmt.Println("+ Very Slow Database +")
	fmt.Println("+--------------------+")
	fmt.Println("")
}

func main() {
	printTitle()

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

	fmt.Println("Using port " + strconv.FormatInt(port, 10))

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

		switch strings.ToLower(args[0]) {
		case "help":
			fmt.Println("All commands are documented here: https://github.com/OliverSchlueter/VSDB/blob/main/vsdb-cli/README.md#commands")

		case "cls", "clearscreen":
			cmd := exec.Command("cmd", "/c", "cls")
			cmd.Stdout = os.Stdout
			cmd.Run()

			printTitle()

		case "exit", "stop":
			os.Exit(0)

		case "get":
			if len(args) < 2 {
				fmt.Println("Syntax: get <key>")
				continue
			}

			key := args[1]

			respObj, err := getResponse("http://localhost:" + strconv.FormatInt(port, 10) + "/get?key=" + key)

			if err != nil {
				fmt.Println("Error")
				fmt.Println(err)
				continue
			}

			if respObj.Status == "not found" {
				fmt.Println("Entry not found")
				continue
			}

			if respObj.Status == "found" {
				fmt.Println(respObj.Result[key])
				continue
			}

		case "getallkeys":
			respObj, err := getResponse("http://localhost:" + strconv.FormatInt(port, 10) + "/getAllKeys")

			if err != nil {
				fmt.Println("Error")
				fmt.Println(err)
				continue
			}

			for k, _ := range respObj.Result {
				fmt.Println(k)
			}

		case "getallentries":
			respObj, err := getResponse("http://localhost:" + strconv.FormatInt(port, 10) + "/getAllEntries")

			if err != nil {
				fmt.Println("Error")
				fmt.Println(err)
				continue
			}

			for k, v := range respObj.Result {
				fmt.Println(k + " : " + v)
			}

		case "insert":
			if len(args) < 3 {
				fmt.Println("Syntax: insert <key> <value>")
				continue
			}

			key := args[1]
			value := args[2]

			respObj, err := getResponse("http://localhost:" + strconv.FormatInt(port, 10) + "/insert?key=" + key + "&value=" + value)

			if err != nil {
				fmt.Println("Error")
				fmt.Println(err)
				continue
			}

			if respObj.Status == "inserted" && respObj.Result[key] == value {
				fmt.Println("Inserted (" + key + " : " + value + ")")
				continue
			}

		case "delete":
			if len(args) < 2 {
				fmt.Println("Syntax: get <key>")
				continue
			}

			key := args[1]

			respObj, err := getResponse("http://localhost:" + strconv.FormatInt(port, 10) + "/delete?key=" + key)

			if err != nil {
				fmt.Println("Error")
				fmt.Println(err)
				continue
			}

			if respObj.Status == "not found" {
				fmt.Println("Entry not found")
				continue
			}

			if respObj.Status == "deleted" {
				fmt.Println("Deleted " + key)
				continue
			}
		}

	}

}
