package main

import (
	"fmt"

	"github.com/Reynadi531/digiflazz-go"
)

func main() {
	client := digiflazz.New("", "")
	saldo, err := client.Balance.CekSaldo()
	if err != nil {
		fmt.Println(err)
		return
	}

	pricelist, err := client.Pricelist.CekHarga()
	if err != nil {
		fmt.Println(err)
		return
	}

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

	fmt.Printf("Saldo: %f\n", saldo)
	fmt.Printf("%+v\n", pricelist)
	fmt.Printf("%+v\n", topup)
	// fmt.Printf("%+v\n", deposit)
}
