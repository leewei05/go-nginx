package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func main() {
	m := make(map[interface{}]interface{})

	yamlFile, err := ioutil.ReadFile("nginx.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Println(m)
	fmt.Println(m["default"])
}
