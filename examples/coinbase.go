package main

import (
	"fmt"

	"github.com/jyap808/go-coinbase"
)

const (
	API_KEY    = ""
	API_SECRET = ""
)

func main() {
	// Coinbase client
	coinbase := coinbase.New(API_KEY, API_SECRET)

	// Get Ticker (btcusd)
	ticker, err := coinbase.GetTicker("btcusd")
	fmt.Println(err, ticker)

	// Get Distribution (JBS)
	/* distribution, err := coinbase.GetDistribution("SPHR")
	    fmt.Println(err)
		for _, balance := range distribution.Distribution {
			fmt.Println(balance.BalanceD)
		} */

	// Get market summaries
	/*
		marketSummaries, err := coinbase.GetMarketSummaries()
		fmt.Println(err, marketSummaries)
	*/

	// Get orders book
	/*
		orderBook, err := coinbase.GetOrderBook("BTC-DRK", "both", 100)
		fmt.Println(err, orderBook)
	*/

	// Market history
	/*
		marketHistory, err := coinbase.GetMarketHistory("BTC-DRK", 100)
		for _, trade := range marketHistory {
			fmt.Println(err, trade.Timestamp.String(), trade.Quantity, trade.Price)
		}
	*/

	// Market

	// BuyLimit
	/*
		uuid, err := coinbase.BuyLimit("BTC-DOGE", 1000, 0.00000102)
		fmt.Println(err, uuid)
	*/

	// BuyMarket
	/*
		uuid, err := coinbase.BuyLimit("BTC-DOGE", 1000)
		fmt.Println(err, uuid)
	*/

	// Sell limit
	/*
		uuid, err := coinbase.SellLimit("BTC-DOGE", 1000, 0.00000115)
		fmt.Println(err, uuid)
	*/

	// Cancel Order
	/*
		err := coinbase.CancelOrder("e3b4b704-2aca-4b8c-8272-50fada7de474")
		fmt.Println(err)
	*/

	// Get open orders
	/*
		orders, err := coinbase.GetOpenOrders("BTC-DOGE")
		fmt.Println(err, orders)
	*/

	// Account
	// Get balances
	/*
		balances, err := coinbase.GetBalances()
		fmt.Println(err, balances)
	*/

	// Get balance
	/*
		balance, err := coinbase.GetBalance("DOGE")
		fmt.Println(err, balance)
	*/

	// Get address
	/*
		address, err := coinbase.GetDepositAddress("QBC")
		fmt.Println(err, address)
	*/

	// WithDraw
	/*
		whitdrawUuid, err := coinbase.Withdraw("QYQeWgSnxwtTuW744z7Bs1xsgszWaFueQc", "QBC", 1.1)
		fmt.Println(err, whitdrawUuid)
	*/

	// Get order history
	/*
		orderHistory, err := coinbase.GetOrderHistory("BTC-DOGE", 10)
		fmt.Println(err, orderHistory)
	*/

	// Get getwithdrawal history
	/*
		withdrawalHistory, err := coinbase.GetWithdrawalHistory("all", 0)
		fmt.Println(err, withdrawalHistory)
	*/

	// Get deposit history
	/*
		deposits, err := coinbase.GetDepositHistory("all", 0)
		fmt.Println(err, deposits)
	*/

}
