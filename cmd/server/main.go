package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

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

func main() {
	//checkParams(os.Args)

	e := echo.New()
	e.GET("/fibonacci", func(c echo.Context) error {
		type_ := c.QueryParam("type")

		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		fibonacci := 0
		if type_ == "loop" {
			for i := 0; i <= limit; i++ {
				fibonacci = FibonacciLoop(i)
			}
		} else if type_ == "recursive" {
			for i := 0; i <= limit; i++ {
				fibonacci = FibonacciRecursion(i)
			}
		} else {
			return c.String(http.StatusBadRequest, "Usage:\n\n\tcurl http://host:1323/fibonacci?limit=41&type=recursive\n\n\tYou can use two types \"loop\" or \"recursive\"\n")
		}

		return c.String(http.StatusOK, strconv.Itoa(fibonacci)+"\n")
	})
	e.Logger.Fatal(e.Start(":8080"))
}