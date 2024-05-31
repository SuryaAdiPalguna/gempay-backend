package main

import (
	"database/sql"
	"gempay/account"
	"gempay/balance"
	"gempay/history"
	"gempay/payment"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Koneksi Database
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/gempay")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Router
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	router.POST("/login", account.Login(db))
	router.POST("/register", account.Register(db))

	router.POST("/account/display_profile", account.DisplayProfile(db))
	router.POST("/account/edit_profile", account.EditProfile(db))

	router.POST("/balance", balance.Balance(db))
	router.POST("/balance/balance_list", balance.BalanceList(db))
	router.POST("/balance/top_up_balance", balance.TopUpBalance(db))
	router.POST("/balance/check_top_up_balance", balance.CheckTopUpBalance(db))
	router.POST("/balance/top_up_receipt", balance.TopUpReceipt(db))

	router.POST("/payment/electric", payment.Electric(db))
	router.POST("/payment/internet", payment.Internet(db))
	router.POST("/payment/water", payment.Water(db))
	router.POST("/payment/pulse", payment.Pulse(db))
	router.POST("/payment/ewallet", payment.Ewallet(db))
	router.POST("/payment/data", payment.Data(db))
	router.POST("/payment/fee_check", payment.FeeCheck(db))

	router.POST("/history", history.History(db))
	router.POST("/history/payment_receipt", history.PaymentReceipt(db))

	router.Run(":8000")
}

// connect, row, stmt
