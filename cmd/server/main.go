package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	//checkParams(os.Args)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/fibonacci", fibonacciHandler)
	e.GET("/lorem-ipsum", loremIpsumHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func FibonacciLoop(n int) int {
	f := make([]int, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

func FibonacciRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}

func loremIpsumHandler(c echo.Context) error {
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 1
	}

	bytes := 0
	for i := 0; i < limit; i++ {
		resp, err := http.Get("https://www.lipsum.com/")
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)

		err = ioutil.WriteFile("/tmp/lorem-ipsum-data", body, 0644)
		check(err)

		bytes = bytes + len(body)
	}

	return c.String(http.StatusOK, "bytes read: "+strconv.Itoa(bytes)+"\n")
}

func fibonacciHandler(c echo.Context) error {
	type_ := c.QueryParam("type")

	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	fibonacci := "[ "
	if type_ == "loop" {
		for i := 0; i <= limit; i++ {
			fibonacci = fibonacci + strconv.Itoa(FibonacciLoop(i)) + " "
		}
	} else if type_ == "recursive" {
		for i := 0; i <= limit; i++ {
			fibonacci = fibonacci + strconv.Itoa(FibonacciRecursion(i)) + " "
		}
	} else {
		return c.String(http.StatusBadRequest, "Usage:\n\n\tcurl http://host:1323/fibonacci?limit=41&type=recursive\n\n\tYou can use two types \"loop\" or \"recursive\"\n")
	}
	fibonacci = fibonacci + "]"
	return c.String(http.StatusOK, fibonacci+"\n")
}
