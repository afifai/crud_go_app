package handler

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/rysmaadit/go-template/app"
	"github.com/rysmaadit/go-template/external/gorm_client"
	"gorm.io/gorm"
)

func InitMigration() (db *gorm.DB, err error) {
	db, err = gorm_client.Connect(app.Init())
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(gorm_client.Employee{})
	return db, nil
}

func CreateEmployee(db *gorm.DB, name, email, address string, age int) {
	new_emp := gorm_client.Employee{Name: name, Email: email, Address: address, Age: age}
	db.Create(&new_emp)
	fmt.Println("Data Berhasil Masuk !")
}

func ReadEmployee(db *gorm.DB) {
	var emp []gorm_client.Employee
	db.Find(&emp)

	for _, e := range emp {
		fmt.Println("Name\t:", e.Name, "\nE-Mail\t:", e.Email, "\nAddress\t:", e.Address, "\nAge\t:", e.Age)
	}
}

func UpdateEmployee(db *gorm.DB, email, ent string, sel int) {
	var emp gorm_client.Employee
	selection := [4]string{"name", "email", "age", "address"}
	db.Model(&emp).Where("email=?", email).Update(selection[sel-1], ent)
	info := fmt.Sprintf("%s berhasil diupdate", selection[sel-1])
	fmt.Println(info)
}

func DeleteEmployee(db *gorm.DB, email string) {
	var emp gorm_client.Employee
	db.Model(&emp).Where("email=?", email).Delete(&emp)
	fmt.Println("Data Berhasil Dihapus")
}

func getInputData() (name, email, address string, age int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan Nama : ")
	name, _ = reader.ReadString('\n')
	fmt.Print("Masukkan Email : ")
	email, _ = reader.ReadString('\n')
	fmt.Print("Masukkan Alamat : ")
	address, _ = reader.ReadString('\n')
	fmt.Print("Masukkan Usia : ")
	fmt.Scanf("%d", &age)
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)
	address = strings.TrimSpace(address)
	return name, email, address, age
}

func MainCLI() {
	var pilihan int
	reader := bufio.NewReader(os.Stdin)

	db, err := InitMigration()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Menu : ")
	fmt.Println("1. Masukkan Data\n2. Lihat Data\n3. Update Data\n4. Hapus Data")
	fmt.Print("Masukkan pilihan anda : ")
	fmt.Scanf("%d", &pilihan)
	if pilihan == 1 {
		name, email, address, age := getInputData()
		CreateEmployee(db, name, email, address, age)
	} else if pilihan == 2 {
		ReadEmployee(db)
	} else if pilihan == 3 {
		fmt.Print("Masukkan email yang ingin diupdate : ")
		email, _ := reader.ReadString('\n')
		email = strings.TrimSpace(email)
		fmt.Println("Data yang ingin diubah")
		fmt.Println("1. Nama\n2. Email\n3. Usia\n4. Alamat")
		fmt.Scanf("%d", &pilihan)
		fmt.Print("Masukkan perubahan : ")
		input_mod, _ := reader.ReadString('\n')
		input_mod = strings.TrimSpace(input_mod)
		UpdateEmployee(db, email, input_mod, pilihan)
	} else if pilihan == 4 {
		fmt.Print("Masukkan email yang ingin dihapus : ")
		email, _ := reader.ReadString('\n')
		email = strings.TrimSpace(email)
		DeleteEmployee(db, email)
	}
}
