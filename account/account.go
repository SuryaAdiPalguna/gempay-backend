package account

import (
	"database/sql"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"

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

func Login(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ambil input dari menu login
		var data Account
		data.Username = c.DefaultPostForm("username", "")
		data.Password = c.DefaultPostForm("password", "")

		// // fmt.Println(data.Username)
		// // fmt.Println(data.Password)
		// if err := c.BindJSON(&data); err != nil {
		// 	// fmt.Println("Error di input data")
		// 	c.JSON(400, gin.H{"error": "Bad Request"})
		// 	return
		// }

		// cek username ada atau nggak
		var user Account
		err := db.QueryRow("SELECT * FROM account WHERE username = ?", data.Username).Scan(&user.IdAccount, &user.Username, &user.Password, &user.Email, &user.Phone, &user.Name, &user.Gender, &user.Address, &user.Balance)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error1": err.Error()})
			return
		}

		// cek password bener atau salah
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error2": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": user, "message": "OK"})
	}
}
func Register(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ambil input dari menu registrasi
		var data Account
		// if err := c.BindJSON(&data); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err})
		// 	return
		// }
		data.Username = c.DefaultPostForm("username", "")
		data.Password = c.DefaultPostForm("password", "")

		// cek username ada atau nggak
		var count int
		err := db.QueryRow("SELECT COUNT(*) FROM account WHERE username = ?", data.Username).Scan(&count)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Akun Sudah Ada"})
			return
		}

		// enkripsi password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
		if err != nil {
			panic(err.Error())
		}

		// tambahkan akun
		_, err = db.Exec("INSERT INTO account(username, password) VALUES (?, ?)", data.Username, string(hashedPassword))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": data, "message": "Registration successful"})
	}
}
func DisplayProfile(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ambil input
		var data Account
		// if err := c.BindJSON(&data); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		// 	return
		// }
		data.Username = c.DefaultPostForm("username", "")

		// ambil data
		err := db.QueryRow("SELECT * FROM account WHERE username = ?", data.Username).Scan(&data.IdAccount, &data.Username, &data.Password, &data.Email, &data.Phone, &data.Name, &data.Gender, &data.Address, &data.Balance)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": data, "message": "OK"})
	}
}
func EditProfile(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ambil input
		var data Account
		// if err := c.BindJSON(&data); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		// 	return
		// }
		data.Username = c.DefaultPostForm("username", "")
		data.Email = c.DefaultPostForm("password", "")
		data.Phone = c.DefaultPostForm("phone", "")
		data.Name = c.DefaultPostForm("name", "")

		var genderString = c.DefaultPostForm("gender", "")
		gender, err := strconv.Atoi(genderString)
		if err != nil {
			// Handle error: invalid integer format
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gender format"})
			return
		}

		data.Gender = gender
		data.Address = c.DefaultPostForm("address", "")

		// update data
		_, err = db.Exec("UPDATE users SET email = ?, phone = ?, name = ?, gender = ?, address = ? WHERE username = ?", data.Email, data.Phone, data.Name, data.Gender, data.Address, data.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		// stmt, err := db.Prepare("UPDATE users SET email = ?, phone = ?, name = ?, gender = ?, address = ? WHERE username = ?")
		// if err != nil {
		// 	// Handle prepare statement error
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		// 	return
		// }
		// defer stmt.Close()

		// _, err = stmt.Exec(data.Email, data.Phone, data.Name, data.Gender, data.Address, data.Username)
		// if err != nil {
		// 	// Handle update error (e.g., no user found)
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "No User Found"})
		// 	return
		// }

		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}
