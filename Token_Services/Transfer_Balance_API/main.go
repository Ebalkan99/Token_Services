package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"math/big"
	"net/http"
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/core/types"
	"database/sql"
    _ "github.com/lib/pq"
)

var (
	avaxClient *rpc.Client
	privateKey *ecdsa.PrivateKey
	 db *sql.DB // PostgreSQL veritabanı bağlantısı
)

func init() {
	var err error
	// Bağlantı kurmak istediğiniz Avalanche Fuji istemcisini belirtilmesi
	avaxClient, err = rpc.Dial("https://api.avax-test.network/ext/bc/C/rpc")
	if err != nil {
		log.Fatal(err)
	}

	privateKeyHex := "d1e884006d15fce05a6e0c610abf89645859f94e6d65ad106105ebe742c026f9"

	privateKey, err = crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal(err)
	}
		// PostgreSQL veritabanı bağlantısı
	const (
		host     = "127.0.0.1" //localhost
		port     = 5432
		user     = "postgres"
		password = "180299"
		dbname   = "postgres"
	)
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	 // Tabloyu oluşturma
    createTableSQL := `
    CREATE TABLE IF NOT EXISTS logs (
        id SERIAL PRIMARY KEY,
        message TEXT,
        timestamp TIMESTAMP
    );
    `
    _, err = db.Exec(createTableSQL)
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
	http.HandleFunc("/balance", getBalance)
	http.HandleFunc("/transfer", transferTokens)

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}

func getBalance(w http.ResponseWriter, r *http.Request) {
	// Adres parametresini al
	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "Address parameter is missing", http.StatusBadRequest)
		return
	}

	// AVAX bakiye sorgusu
	var avaxResult *big.Int
	err := avaxClient.CallContext(context.Background(), &avaxResult, "eth_getBalance", address, "latest")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Yanıtı gönderme işlemi
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"AVAXBalance": "%s"}`, avaxResult.String())
	// Günlük kayıt icin
	insertLogSQL := `
    INSERT INTO logs (message, timestamp)
    VALUES ($1, $2);
    `
    _, err = db.Exec(insertLogSQL, "API işlemi gerçekleştirildi: getBalance", time.Now())
    if err != nil {
        log.Println("Günlük kaydedilemedi:", err)
    }
}

func transferTokens(w http.ResponseWriter, r *http.Request) {
	// Transfer işlemi parametreleri
	nonce := uint64(0) // Transfer işlemi için nonce değeri
	gasPrice := big.NewInt(225000000000) // Avalanche Fuji'de örnek bir gas price
	toAddress := common.HexToAddress("0x4E98fF06B1257C43599d024F6E62720a6078BD5F") // Gönderilecek adres
	value := big.NewInt(1000000000000000000) // 1 AVAX in nAVAX

	// İşlem verilerini oluşturma
	tx := types.NewTransaction(nonce, toAddress, value, 21000, gasPrice, nil)

	// İşlemi imzalama
	signer := types.NewEIP155Signer(nil)
	signedTx, err := types.SignTx(tx, signer, privateKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// İşlemi ağa gönderme
	err = avaxClient.CallContext(context.Background(), nil, "eth_sendRawTransaction", signedTx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Yanıtı gönder
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"Message": "Transfer completed"}`)
	// Günlük kaydını veritabanına ekle
    insertLogSQL := `
    INSERT INTO logs (message, timestamp)
    VALUES ($1, $2);
    `
    _, err = db.Exec(insertLogSQL, "API işlemi gerçekleştirildi: transferTokens", time.Now())
    if err != nil {
        log.Println("Günlük kaydedilemedi:", err)
    }
}
