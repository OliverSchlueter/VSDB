package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"vsdb-server/cache"
)

type jsonResponse struct {
	Status string
	Result map[string]string
}

func apiGet(context *gin.Context) {
	key := context.Request.URL.Query().Get("key")

	contains, _ := dataCache.Contains(key)

	if !contains {
		response := jsonResponse{
			Status: "not found",
		}

		context.IndentedJSON(http.StatusNotFound, response)
		return
	}

	response := jsonResponse{
		Status: "found",
		Result: map[string]string{key: dataCache.Get(key)},
	}

	context.IndentedJSON(http.StatusFound, response)
}

func apiGetAllKeys(context *gin.Context) {
	keys := dataCache.GetAllKeys()
	keyMap := make(map[string]string)

	for _, key := range keys {
		keyMap[key] = ""
	}

	context.IndentedJSON(http.StatusOK, jsonResponse{
		Status: "success",
		Result: keyMap,
	})
}

func apiGetAllEntries(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, jsonResponse{
		Status: "success",
		Result: dataCache.GetAllEntries(),
	})
}

func apiInsert(context *gin.Context) {
	key := context.Request.URL.Query().Get("key")
	value := context.Request.URL.Query().Get("value")

	dataCache.InsertOrUpdate(key, value)

	response := jsonResponse{
		Status: "inserted",
		Result: map[string]string{key: value},
	}

	context.IndentedJSON(http.StatusCreated, response)
}

func apiDelete(context *gin.Context) {
	key := context.Request.URL.Query().Get("key")

	contains, _ := dataCache.Contains(key)

	if !contains {
		response := jsonResponse{
			Status: "not found",
		}

		context.IndentedJSON(http.StatusNotFound, response)
		return
	}

	dataCache.Delete(key)

	response := jsonResponse{
		Status: "deleted",
		Result: map[string]string{key: ""},
	}

	context.IndentedJSON(http.StatusOK, response)
}

var dataCache = cache.CreateCache()

func main() {

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

	dataCache.InsertOrUpdate("hello", "world")
	dataCache.InsertOrUpdate("hello2", "world")

	router := gin.Default()
	router.GET("/get", apiGet)
	router.GET("/getAllKeys", apiGetAllKeys)
	router.GET("/getAllEntries", apiGetAllEntries)
	router.GET("/insert", apiInsert)
	router.GET("/delete", apiDelete)
	err := router.Run("localhost:" + strconv.FormatInt(port, 10))

	if err != nil {
		return
	}
}
