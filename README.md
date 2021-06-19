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
### fibonacci
Fibonacci can be executed with recursion or using a loop and a collection.

`curl http://localhost:8080/fibonacci?type=loop&limit=10`

```
limit is the number used to calculate the fibonacci series
type can be loop or recursive
```

### lorem ipsum
This read a page form internet and write it in a file on temp dir.

`curl http://localhost:8080/lorem-ipsum\?limit\=100`

```
limit is the amount of read and write you want to execute.
```