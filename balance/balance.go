package balance

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
type Deposite struct {
	IdDeposite     int    `json:"id_deposite"`
	IdAccount      int    `json:"id_account"`
	IdPayment      int    `json:"id_payment"`
	Date           string `json:"date"`
	AmountDeposite int    `json:"amount_deposite"`
	Description    string `json:"description"`
	PriceUnique    int    `json:"price_unique"`
	PriceMutation  int    `json:"price_mutation"`
	StatusDeposite int    `json:"status_deposite"`
	StatusMutation int    `json:"status_mutation"`
}
type Payment struct {
	IdPayment     int    `json:"id_payment"`
	MethodPayment string `json:"method_payment"`
	PricePayment  int    `json:"price_payment"`
}

// function balance digunakan untuk mengecek value saldo saat ini
func Balance(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ambil input
		var data Account
		// if err := c.BindJSON(&data); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err})
		// 	return
		// }
		data.Username = c.DefaultPostForm("username", "")

		// ambil data
		err := db.QueryRow("SELECT * FROM account WHERE username = ?", data.Username).Scan(&data.IdAccount, &data.Username, &data.Password, &data.Email, &data.Phone, &data.Name, &data.Gender, &data.Address, &data.Balance)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": data, "message": "OK"})
	}
}

// function balancelist digunakan untuk melihat riwayat topup/pembelian saldo
func BalanceList(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ambil input
		var data Account
		// if err := c.BindJSON(&data); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err})
		// 	return
		// }
		var idAccountString = c.DefaultPostForm("id_account", "")
		IdAccount, err := strconv.Atoi(idAccountString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.IdAccount = IdAccount

		// ambil data
		rows, err := db.Query("SELECT * FROM deposite INNER JOIN payment ON deposite.id_payment = payment.id_payment WHERE transaction.id_account = ?", data.IdAccount)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		defer rows.Close()

		var resultDeposite []Deposite
		var resultPayment []Payment
		for rows.Next() {
			var tempDeposite Deposite
			var tempPayment Payment
			err = rows.Scan(&tempDeposite.IdDeposite, &tempDeposite.IdAccount, &tempDeposite.IdPayment, &tempDeposite.Date, &tempDeposite.AmountDeposite, &tempDeposite.Description, &tempDeposite.PriceUnique, &tempDeposite.PriceMutation, &tempDeposite.StatusDeposite, &tempDeposite.StatusMutation, &tempPayment.IdPayment, &tempPayment.MethodPayment, &tempPayment.PricePayment)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err})
			}

			resultDeposite = append(resultDeposite, tempDeposite)
			resultPayment = append(resultPayment, tempPayment)
		}

		c.JSON(http.StatusOK, gin.H{"dataDeposite": resultDeposite, "dataPayment": resultPayment, "message": "OK"})
	}
}

// function topupbalance digunakan ketika mau nginput mau topup brp rupiah dan methodnya apa
func TopUpBalance(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ambil input
		var data Deposite
		// if err := c.BindJSON(&data); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err}) // error client
		// 	return
		// }
		var idAccountString = c.DefaultPostForm("id_deposite", "")
		IdAccount, err := strconv.Atoi(idAccountString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.IdAccount = IdAccount

		var idPaymentString = c.DefaultPostForm("id_payment", "")
		IdPayment, err := strconv.Atoi(idPaymentString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.IdPayment = IdPayment

		data.Date = c.DefaultPostForm("date", "")

		var AmountDepositeString = c.DefaultPostForm("amount_deposite", "")
		AmountDeposite, err := strconv.Atoi(AmountDepositeString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.AmountDeposite = AmountDeposite 

		data.Description = c.DefaultPostForm("description", "")

		var PriceUniqueString = c.DefaultPostForm("price_unique", "")
		PriceUnique, err := strconv.Atoi(PriceUniqueString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.PriceUnique = PriceUnique 

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
		_, err = db.Exec("INSERT INTO deposite VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", data.IdAccount, data.IdPayment, data.Date, data.AmountDeposite, data.Description, data.PriceUnique, data.PriceMutation, data.StatusDeposite, data.StatusMutation)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err}) // error server
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

// riwayat topupbalance yang belum lunas payment nya
func CheckTopUpBalance(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ambil input
		var data Deposite
		// if err := c.BindJSON(&data); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err})
		// 	return
		// }
		var idDepositeString = c.DefaultPostForm("id_deposite", "")
		IdDeposite, err := strconv.Atoi(idDepositeString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.IdDeposite = IdDeposite

		// ambil data
		var deposite Deposite
		var payment Payment
		err = db.QueryRow("SELECT * FROM deposite INNER JOIN payment ON deposite.id_payment = payment.id_payment WHERE deposite.id_deposite = ?", data.IdDeposite).Scan(&deposite.IdDeposite, &deposite.IdAccount, &deposite.IdPayment, &deposite.Date, &deposite.AmountDeposite, &deposite.Description, &deposite.PriceUnique, &deposite.PriceMutation, &deposite.StatusDeposite, &deposite.StatusMutation, &payment.IdPayment, &payment.MethodPayment, &payment.PricePayment)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"dataDeposite": deposite, "dataPayment": payment, "message": "OK"})
	}
}

// riwayat topupbalance yang sudah lunas payment nya
func TopUpReceipt(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ambil input
		var data Deposite
		// if err := c.BindJSON(&data); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err})
		// 	return
		// }
		var idDepositeString = c.DefaultPostForm("id_deposite", "")
		IdDeposite, err := strconv.Atoi(idDepositeString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data.IdDeposite = IdDeposite

		// ambil data
		var deposite Deposite
		var payment Payment
		err = db.QueryRow("SELECT * FROM deposite INNER JOIN payment ON deposite.id_payment = payment.id_payment WHERE deposite.id_deposite = ?", data.IdDeposite).Scan(&deposite.IdDeposite, &deposite.IdAccount, &deposite.IdPayment, &deposite.Date, &deposite.AmountDeposite, &deposite.Description, &deposite.PriceUnique, &deposite.PriceMutation, &deposite.StatusDeposite, &deposite.StatusMutation, &payment.IdPayment, &payment.MethodPayment, &payment.PricePayment)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"dataDeposite": deposite, "dataPayment": payment, "message": "OK"})
	}
}
