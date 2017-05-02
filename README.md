go-coinbase
==========

go-coinbase is an implementation of the Coinbase API (public and private) in Golang.

Based off of https://github.com/toorop/go-bittrex/

## Import
	import "github.com/jyap808/go-coinbase"
	
## Usage
~~~ go
package main

import (
	"fmt"
	"github.com/jyap808/go-coinbase"
)

const (
	API_KEY    = "YOUR_API_KEY"
	API_SECRET = "YOUR_API_SECRET"
)

func main() {
	// Coinbase client
	coinbase := coinbase.New(API_KEY, API_SECRET)

	// Get markets
	markets, err := coinbase.GetMarkets()
	fmt.Println(err, markets)
}
~~~	

See ["Examples" folder for more... examples](https://github.com/jyap808/go-coinbase/blob/master/examples/coinbase.go)

