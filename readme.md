## Introduction to concurrency?

#### Definition?

Concurrency is the ability of a software system to execute multiple tasks or operations at the same time, often on different processors or cores. Software that can do such things are known as _concurrent_ software.

## Requirements

> brew install go 

> go version 

> brew install 

## Start a go server 
``` go run main.go ``` 

## Start the load test script

This will generate multiple requests to the counter endpoints

```k6 run scripts/loadtest.js```

## Check the value of the Counter

```
curl -X GET http://localhost:8080/valueWithOptimisticVersioning
```
