package main

import (
	"fmt"
	"github.com/fanxiaoping/http-mux/mux"
	"log"
	"net/http"
)

func main(){
	engine := mux.NewEngine()
	engine.Use(func(c *mux.Context) {
		fmt.Println("begin")
		fmt.Println("end")
	})

	engine.Use(func(c *mux.Context) {
		fmt.Println("begin 2")
		c.Next()
		fmt.Println("end 2")
	})
	engine.Use(func(c *mux.Context) {
		fmt.Println("begin 3")
		c.Next()
		fmt.Println("end 3")
	})

	engine.AddRoute("/path1", func(c *mux.Context) {
		c.Writer.Write([]byte("path1"))
	})
	engine.AddRoute("/path2", func(c *mux.Context) {
		c.Writer.Write([]byte("path2"))
	})

	if err := http.ListenAndServe(":9990",engine);err != nil{
		log.Fatal(err)
	}
}
