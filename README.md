# apiai-go-client

[![CircleCI](https://circleci.com/gh/mlabouardy/apiai-go-client.svg?style=svg&circle-token=1f1f702ba1cb3cb0ee1385992d5fdf2dba02ea3f)](https://circleci.com/gh/mlabouardy/apiai-go-client) [![License](https://img.shields.io/badge/License-Apache%202.0-yellowgreen.svg)](https://opensource.org/licenses/Apache-2.0)

This library allows integrating agents from the [Api.ai](http://api.ai) natural language processing service with your Golang application.

* [Prerequsites](#prerequsites)
* [Installation](#installation)
* [Features](#features)
* [Usage](#usage)
* [Documentation](#documentation)

# Prerequsites

Create an [API.AI account](http://api.ai).

# Installation

```shell
go get github.com/mlabouardy/apiai-go-client
```
# Features

* Queries
* Contexts
* Intents
* UserIntents
* Entities

# Usage

* Create `main.go` file with the following code:
```golang
package main

import (
	"fmt"
	"github.com/mlabouardy/apiai-go-client"
	"github.com/mlabouardy/apiai-go-client/models"
	"log"
)

func main() {
	err, client := NewApiAiClient(Options{
		AccessToken: "<API.AI TOKEN GOES HERE>",
	})
	if err != nil {
		log.Fatal(err)
	}

	entities, err := client.EntitiesFindAllRequest()
	if err != nil {
		log.Fatal(err)
	}
	for _, entity := range entities {
		fmt.Println(entity.Name)
	}
}
```
* Run following command.
```shell
go run main.go
```
* Your can find more examples in [`examples`](examples) directory.

# Documentation

Documentation is available at http://api.ai.
