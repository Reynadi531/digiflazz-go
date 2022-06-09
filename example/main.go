package main

import (
	"fmt"
	"os"

	"github.com/Reynadi531/digiflazz-go"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Faild load env")
		return
	}
	username := os.Getenv("DIGIFLAZZ_USERNAME")
	key := os.Getenv("DIGIFLAZZ_KEY")

	client := digiflazz.New(username, key)
	saldo, err := client.Balance.CekSaldo()
	if err != nil {
		fmt.Println(err)
		return
	}

	// pricelist, err := client.Pricelist.CekHarga()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// deposit, err := client.Deposit.BuatTiket(200000, digiflazz.BankMandiri, "FooBarz")
	// if err != nil {
	// fmt.Println(err)
	// return
	// }

	topup, err := client.Transaction.Topup("1", "1", "1", true, "")
	if err != nil {
		fmt.Println(err)
		return
	}

	checkTopup, err := client.Transaction.CekTopup(topup.RefrenceId)

	fmt.Printf("Saldo: %f\n\n", saldo)
	// fmt.Printf("%+v\n", pricelist)
	fmt.Printf("%+v\n\n", topup)
	fmt.Printf("%+v\n\n", checkTopup)
	// fmt.Printf("%+v\n", deposit)
}
