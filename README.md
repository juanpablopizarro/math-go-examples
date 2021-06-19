# math-go-examples


This project holds some math operations just to test how the implementations consumes host computer resources.

## Prerequisites
* `docker`
* `golang`

## How to build it

`docker build -t math .`

## How to run it
`docker run -p 8080:8080 --rm math .`

## How to test it
`curl http://localhost:8080/fibonacci?type=loop&limit=10`


```
limit is the number used to calculate the fibonacci series
type can be loop or recursive
```

