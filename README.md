# go-logger

This project ispired by Java [Logback](http://logback.qos.ch/) implementation, which is implements [slf4j](http://www.slf4j.org/) interface for logger.

go-logger introduces a few simple interfaces in which composition creates a logger.
This library also includes an implementation for a basic logger that stram messages to the `stdout`

***Note: This very begin of a project so it has the basic logger only, a contribution is welcome! 

## Interfaces
* Appenders - Logger output component
* Filters - Logger filters, like log level filter
* Encoders - Output message format

## Installation 
```
go get github.com/yorikya/go-logger
```

## Usage
```
package main

import log "github.com/yorikya/go-logger"

func main(){
  log.Info("hello world")
}

out=>[15:04:05.000][INFO] hello world 
```



