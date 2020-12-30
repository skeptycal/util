package main

import (
	"github.com/BurntSushi/toml"
	"log"
)

var tomlData = `title = "config"
[feature1]
enable = true
userids = [
  "12345", "67890"
]

[feature2]
enable = false`

type feature1 struct {
	Enable  bool
	Userids []string
}

type feature2 struct {
	Enable bool
}

type tomlConfig struct {
	Title string
	F1    feature1 `toml:"feature1"`
	F2    feature2 `toml:"feature2"`
}

func main() {
	var conf tomlConfig
	if _, err := toml.Decode(tomlData, &conf); err != nil {
		log.Fatal(err)
	}
	log.Printf("title: %s", conf.Title)
	log.Printf("Feature 1: %#v", conf.F1)
	log.Printf("Feature 2: %#v", conf.F2)
}
