package history

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Account struct {
	IdAccount int    `json:"id_account"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Name      string `json:"name"`
	Gender    int    `json:"gender"`
	Address   string `json:"address"`
	Balance   int    `json:"balance"`
}
type Transaction struct {
	IdTransaction  int    `json:"id_transaction"`
	IdAccount      int    `json:"id_account"`
	IdService      int    `json:"id_service"`
	Date           string `json:"date"`
	NumberCostumer int    `json:"number_costumer"`
	Reference      string `json:"reference"`
	Description    string `json:"description"`
	PriceSystem    int    `json:"price_system"`
	PriceMutation  int    `json:"price_mutation"`
	StatusDeposite int    `json:"status_deposite"`
	StatusMutation int    `json:"status_mutation"`
}
type Service struct {
	IdService      int    `json:"id_service"`
	NameService    string `json:"name_service"`
	SubnameService string `json:"subname_service"`
	TypeService    string `json:"type_service"`
	SubtypeService string `json:"subtype_service"`
	Description    string `json:"description"`
	PriceService   int    `json:"price_service"`
}

// untuk func historynya
func History(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ambil input
		var data Account
		// if err := c.BindJSON(&data); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		// 	return
		// }
		var idAccountString = c.DefaultPostForm("id_account", "")
		idAccount, err := strconv.Atoi(idAccountString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gender format"})
			return
		}

		data.IdAccount = idAccount

		// ambil data
		rows, err := db.Query("SELECT * FROM transaction INNER JOIN service ON transaction.id_service = service.id_service WHERE transaction.id_account = ?", data.IdAccount)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		defer rows.Close()

		var resultTransaction []Transaction
		var resultService []Service
		for rows.Next() {
			var tempTransaction Transaction
			var tempService Service
			err = rows.Scan(&tempTransaction.IdTransaction, &tempTransaction.IdAccount, &tempTransaction.IdService, &tempTransaction.Date, &tempTransaction.NumberCostumer, &tempTransaction.Reference, &tempTransaction.Description, &tempTransaction.PriceSystem, &tempTransaction.PriceMutation, &tempTransaction.StatusDeposite, &tempTransaction.StatusMutation, &tempService.IdService, &tempService.NameService, &tempService.SubnameService, &tempService.TypeService, &tempService.SubtypeService, &tempService.Description, &tempService.PriceService)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			}

			resultTransaction = append(resultTransaction, tempTransaction)
			resultService = append(resultService, tempService)
		}

		c.JSON(http.StatusOK, gin.H{"dataTransaction": resultTransaction, "dataService": resultService, "message": "OK"})
	}
}
func PaymentReceipt(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ambil input
		var data Transaction
		// if err := c.BindJSON(&data); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		// 	return
		// }
		var idTransactionString = c.DefaultPostForm("id_transaction", "")
		IdTransaction, err := strconv.Atoi(idTransactionString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gender format"})
			return
		}
		data.IdTransaction = IdTransaction

		// ambil data
		var transaction Transaction
		var service Service
		err = db.QueryRow("SELECT * FROM transaction INNER JOIN service ON transaction.id_service = service.id_service WHERE transaction.id_transaction = ?", data.IdTransaction).Scan(&transaction.IdTransaction, &transaction.IdAccount, &transaction.IdService, &transaction.Date, &transaction.NumberCostumer, &transaction.Reference, &transaction.Description, &transaction.PriceSystem, &transaction.PriceMutation, &transaction.StatusDeposite, &transaction.StatusMutation, &service.IdService, &service.NameService, &service.SubnameService, &service.TypeService, &service.SubtypeService, &service.Description, &service.PriceService)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"dataTransaction": transaction, "dataService": service, "message": "OK"})
	}
}
