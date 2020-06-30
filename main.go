package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	hostIndex, portIndex, userIndex, passIndex, dbIndex := CommandArgs()
	host := os.Args[hostIndex]
	port := os.Args[portIndex]
	user := os.Args[userIndex]
	pass := os.Args[passIndex]
	db := os.Args[dbIndex]

	// seda zadane func'i ke db misaze.
	DbCreate(db, host, port, user, pass)
}

// DbCreate = baraye sakhte db bedune hich table'i.
func DbCreate(name, host, port, user, pass string) {
	// baz kardane connection be db.
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable",
		host, port, user, pass)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// ersale dasture create.
	_, err = db.Exec("CREATE DATABASE " + name)
	if err != nil {
		log.Fatalln("\nSomething happened!\n", err)
	} else {
		fmt.Printf("Database '%v' created.\n", name)
	}
}

//CommandArgs = maqadire arg ro migireo meqdar dehi mikone.
func CommandArgs() (int, int, int, int, int) {
	var hostIndex, portIndex, userIndex, passIndex, dbIndex int
	var hostTester, portTester, userTester, passTester, dbTester bool = false, false, false, false, false
	if len(os.Args) == 11 { // yeki esme khode file ejraei mishe.
		for i := 1; i < len(os.Args); i++ {
			if os.Args[i] == "-host" {
				hostIndex = i + 1
				hostTester = true
			} else if os.Args[i] == "-port" {
				portIndex = i + 1
				portTester = true
			} else if os.Args[i] == "-user" {
				userIndex = i + 1
				userTester = true
			} else if os.Args[i] == "-pass" {
				passIndex = i + 1
				passTester = true
			} else if os.Args[i] == "-db" {
				dbIndex = i + 1
				dbTester = true
			}
		}
		if hostTester == true && portTester == true && userTester == true && passTester == true && dbTester == true {
			fmt.Println("Waiting ...")
		} else {
			ArgsUnknown()
		}
	} else {
		ArgsUnknown()
	}
	return hostIndex, portIndex, userIndex, passIndex, dbIndex
}

//ArgsUnknown = agar hich valuei ashna nabud.
func ArgsUnknown() {
	fmt.Println("\nSomething went wrong. Please run it as the following Example:" +
		"\n\t-host [hostname] -port [port number] -user [username] -pass [password] -db [db name]\n")

}
