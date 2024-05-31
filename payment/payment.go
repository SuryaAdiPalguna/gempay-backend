package payment

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

func Electric(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ambil input
		var data Transaction
		// if err := c.BindJSON(&data); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err})
		// 	return
		// }
		var IdAccountString = c.DefaultPostForm("id_account", "")
		IdAccount, err := strconv.Atoi(IdAccountString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.IdAccount = IdAccount

		var IdServiceString = c.DefaultPostForm("id_service", "")
		IdService, err := strconv.Atoi(IdServiceString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.IdService = IdService

		data.Date = c.DefaultPostForm("date", "")

		var NumberCostumerString = c.DefaultPostForm("number_costumer", "")
		NumberCostumer, err := strconv.Atoi(NumberCostumerString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.NumberCostumer = NumberCostumer

		data.Reference = c.DefaultPostForm("reference", "")
		data.Description = c.DefaultPostForm("description", "")

		var PriceSystemString = c.DefaultPostForm("price_system", "")
		PriceSystem, err := strconv.Atoi(PriceSystemString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.PriceSystem = PriceSystem

		var PriceMutationString = c.DefaultPostForm("price_mutation", "")
		PriceMutation, err := strconv.Atoi(PriceMutationString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.PriceMutation = PriceMutation

		var StatusDepositeString = c.DefaultPostForm("status_deposite", "")
		StatusDeposite, err := strconv.Atoi(StatusDepositeString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.StatusDeposite = StatusDeposite

		var StatusMutationString = c.DefaultPostForm("status_mutation", "")
		StatusMutation, err := strconv.Atoi(StatusMutationString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.StatusMutation = StatusMutation

		// insert data
		_, err = db.Exec("INSERT INTO balance VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", data.IdAccount, data.IdService, data.Date, data.NumberCostumer, data.Reference, data.Description, data.PriceSystem, data.PriceMutation, data.StatusDeposite, data.StatusMutation)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}
func Internet(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ambil input
		var data Transaction
		// if err := c.BindJSON(&data); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err})
		// 	return
		// }
		var IdAccountString = c.DefaultPostForm("id_account", "")
		IdAccount, err := strconv.Atoi(IdAccountString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.IdAccount = IdAccount

		var IdServiceString = c.DefaultPostForm("id_service", "")
		IdService, err := strconv.Atoi(IdServiceString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.IdService = IdService

		data.Date = c.DefaultPostForm("date", "")

		var NumberCostumerString = c.DefaultPostForm("number_costumer", "")
		NumberCostumer, err := strconv.Atoi(NumberCostumerString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.NumberCostumer = NumberCostumer

		data.Reference = c.DefaultPostForm("reference", "")
		data.Description = c.DefaultPostForm("description", "")

		var PriceSystemString = c.DefaultPostForm("price_system", "")
		PriceSystem, err := strconv.Atoi(PriceSystemString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.PriceSystem = PriceSystem

		var PriceMutationString = c.DefaultPostForm("price_mutation", "")
		PriceMutation, err := strconv.Atoi(PriceMutationString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.PriceMutation = PriceMutation

		var StatusDepositeString = c.DefaultPostForm("status_deposite", "")
		StatusDeposite, err := strconv.Atoi(StatusDepositeString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.StatusDeposite = StatusDeposite

		var StatusMutationString = c.DefaultPostForm("status_mutation", "")
		StatusMutation, err := strconv.Atoi(StatusMutationString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.StatusMutation = StatusMutation

		// insert data
		_, err = db.Exec("INSERT INTO balance VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", data.IdAccount, data.IdService, data.Date, data.NumberCostumer, data.Reference, data.Description, data.PriceSystem, data.PriceMutation, data.StatusDeposite, data.StatusMutation)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}
func Water(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ambil input
		var data Transaction
		// if err := c.BindJSON(&data); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err})
		// 	return
		// }
		var IdAccountString = c.DefaultPostForm("id_account", "")
		IdAccount, err := strconv.Atoi(IdAccountString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.IdAccount = IdAccount

		var IdServiceString = c.DefaultPostForm("id_service", "")
		IdService, err := strconv.Atoi(IdServiceString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.IdService = IdService

		data.Date = c.DefaultPostForm("date", "")

		var NumberCostumerString = c.DefaultPostForm("number_costumer", "")
		NumberCostumer, err := strconv.Atoi(NumberCostumerString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.NumberCostumer = NumberCostumer

		data.Reference = c.DefaultPostForm("reference", "")
		data.Description = c.DefaultPostForm("description", "")

		var PriceSystemString = c.DefaultPostForm("price_system", "")
		PriceSystem, err := strconv.Atoi(PriceSystemString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.PriceSystem = PriceSystem

		var PriceMutationString = c.DefaultPostForm("price_mutation", "")
		PriceMutation, err := strconv.Atoi(PriceMutationString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.PriceMutation = PriceMutation

		var StatusDepositeString = c.DefaultPostForm("status_deposite", "")
		StatusDeposite, err := strconv.Atoi(StatusDepositeString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.StatusDeposite = StatusDeposite

		var StatusMutationString = c.DefaultPostForm("status_mutation", "")
		StatusMutation, err := strconv.Atoi(StatusMutationString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.StatusMutation = StatusMutation

		// insert data
		_, err = db.Exec("INSERT INTO balance VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", data.IdAccount, data.IdService, data.Date, data.NumberCostumer, data.Reference, data.Description, data.PriceSystem, data.PriceMutation, data.StatusDeposite, data.StatusMutation)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}
func Pulse(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ambil input
		var data Transaction
		// if err := c.BindJSON(&data); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err})
		// 	return
		// }
		var IdAccountString = c.DefaultPostForm("id_account", "")
		IdAccount, err := strconv.Atoi(IdAccountString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.IdAccount = IdAccount

		var IdServiceString = c.DefaultPostForm("id_service", "")
		IdService, err := strconv.Atoi(IdServiceString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.IdService = IdService

		data.Date = c.DefaultPostForm("date", "")

		var NumberCostumerString = c.DefaultPostForm("number_costumer", "")
		NumberCostumer, err := strconv.Atoi(NumberCostumerString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.NumberCostumer = NumberCostumer

		data.Reference = c.DefaultPostForm("reference", "")
		data.Description = c.DefaultPostForm("description", "")

		var PriceSystemString = c.DefaultPostForm("price_system", "")
		PriceSystem, err := strconv.Atoi(PriceSystemString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.PriceSystem = PriceSystem

		var PriceMutationString = c.DefaultPostForm("price_mutation", "")
		PriceMutation, err := strconv.Atoi(PriceMutationString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.PriceMutation = PriceMutation

		var StatusDepositeString = c.DefaultPostForm("status_deposite", "")
		StatusDeposite, err := strconv.Atoi(StatusDepositeString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.StatusDeposite = StatusDeposite

		var StatusMutationString = c.DefaultPostForm("status_mutation", "")
		StatusMutation, err := strconv.Atoi(StatusMutationString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.StatusMutation = StatusMutation

		// insert data
		_, err = db.Exec("INSERT INTO balance VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", data.IdAccount, data.IdService, data.Date, data.NumberCostumer, data.Reference, data.Description, data.PriceSystem, data.PriceMutation, data.StatusDeposite, data.StatusMutation)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}
func Ewallet(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ambil input
		var data Transaction
		// if err := c.BindJSON(&data); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err})
		// 	return
		// }
		var IdAccountString = c.DefaultPostForm("id_account", "")
		IdAccount, err := strconv.Atoi(IdAccountString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.IdAccount = IdAccount

		var IdServiceString = c.DefaultPostForm("id_service", "")
		IdService, err := strconv.Atoi(IdServiceString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.IdService = IdService

		data.Date = c.DefaultPostForm("date", "")

		var NumberCostumerString = c.DefaultPostForm("number_costumer", "")
		NumberCostumer, err := strconv.Atoi(NumberCostumerString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.NumberCostumer = NumberCostumer

		data.Reference = c.DefaultPostForm("reference", "")
		data.Description = c.DefaultPostForm("description", "")

		var PriceSystemString = c.DefaultPostForm("price_system", "")
		PriceSystem, err := strconv.Atoi(PriceSystemString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.PriceSystem = PriceSystem

		var PriceMutationString = c.DefaultPostForm("price_mutation", "")
		PriceMutation, err := strconv.Atoi(PriceMutationString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.PriceMutation = PriceMutation

		var StatusDepositeString = c.DefaultPostForm("status_deposite", "")
		StatusDeposite, err := strconv.Atoi(StatusDepositeString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.StatusDeposite = StatusDeposite

		var StatusMutationString = c.DefaultPostForm("status_mutation", "")
		StatusMutation, err := strconv.Atoi(StatusMutationString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.StatusMutation = StatusMutation

		// insert data
		_, err = db.Exec("INSERT INTO balance VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", data.IdAccount, data.IdService, data.Date, data.NumberCostumer, data.Reference, data.Description, data.PriceSystem, data.PriceMutation, data.StatusDeposite, data.StatusMutation)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}
func Data(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ambil input
		var data Transaction
		// if err := c.BindJSON(&data); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err})
		// 	return
		// }
		var IdAccountString = c.DefaultPostForm("id_account", "")
		IdAccount, err := strconv.Atoi(IdAccountString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.IdAccount = IdAccount

		var IdServiceString = c.DefaultPostForm("id_service", "")
		IdService, err := strconv.Atoi(IdServiceString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.IdService = IdService

		data.Date = c.DefaultPostForm("date", "")

		var NumberCostumerString = c.DefaultPostForm("number_costumer", "")
		NumberCostumer, err := strconv.Atoi(NumberCostumerString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.NumberCostumer = NumberCostumer

		data.Reference = c.DefaultPostForm("reference", "")
		data.Description = c.DefaultPostForm("description", "")

		var PriceSystemString = c.DefaultPostForm("price_system", "")
		PriceSystem, err := strconv.Atoi(PriceSystemString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.PriceSystem = PriceSystem

		var PriceMutationString = c.DefaultPostForm("price_mutation", "")
		PriceMutation, err := strconv.Atoi(PriceMutationString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.PriceMutation = PriceMutation

		var StatusDepositeString = c.DefaultPostForm("status_deposite", "")
		StatusDeposite, err := strconv.Atoi(StatusDepositeString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.StatusDeposite = StatusDeposite

		var StatusMutationString = c.DefaultPostForm("status_mutation", "")
		StatusMutation, err := strconv.Atoi(StatusMutationString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.StatusMutation = StatusMutation

		// insert data
		_, err = db.Exec("INSERT INTO balance VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", data.IdAccount, data.IdService, data.Date, data.NumberCostumer, data.Reference, data.Description, data.PriceSystem, data.PriceMutation, data.StatusDeposite, data.StatusMutation)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}
func FeeCheck(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ambil input
		var data Transaction
		// if err := c.BindJSON(&data); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err})
		// 	return
		// }
		var IdTransactionString = c.DefaultPostForm("id_transaction", "")
		IdTransaction, err := strconv.Atoi(IdTransactionString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.IdTransaction = IdTransaction

		// ambil data
		var transaction Transaction
		var service Service
		err = db.QueryRow("SELECT * FROM transaction INNER JOIN service ON transaction.id_service = service.id_service WHERE id_transaction = ?", data.IdTransaction).Scan(&transaction.IdTransaction, &transaction.IdAccount, &transaction.IdService, &transaction.Date, &transaction.NumberCostumer, &transaction.Reference, &transaction.Description, &transaction.PriceSystem, &transaction.PriceMutation, &transaction.StatusDeposite, &transaction.StatusMutation, &service.IdService, &service.NameService, &service.SubnameService, &service.TypeService, &service.SubtypeService, &service.Description, &service.PriceService)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"dataTransaction": transaction, "dataService": service, "message": "OK"})
	}
}
