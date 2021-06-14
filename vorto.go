package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// "host.docker.internal"
const (
	host     = "host.docker.internal"
	port     = 5432
	user     = "postgres"
	password = "Newuser123"
	dbname   = "vorto_ai"
)

func main() {
	fmt.Println("Successful!")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// if we consider that the supplier id 8 can give only bean_type 2 to carrier id 2 then the below is the correct approach
	fmt.Println("Invalid transactions if consider supplied id 8 can only give bean_type 2 to the carrier id 2 : ")
	view1, err := db.Query(`SELECT	Distinct(D.id) from delivery D where NOT EXISTS
	(select distinct(C.driver, C.supplier) from (select S.id as Supplier, SBT.bean_type_id as Bean, CBT.carrier_id as carrier, D.id as Driver from carrier_bean_type CBT, carrier C, driver D, supplier S, supplier_bean_type SBT, bean_type B where S.id = SBT.supplier_id AND SBT.bean_type_id = B.id AND B.id = CBT.bean_type_id AND CBT.carrier_id = D.carrier_id order by S.id desc) C where C.driver = D.driver_id AND C.supplier= D.supplier_id)`)
	if err != nil {
		// handle this error
		panic(err)

	}
	defer view1.Close()
	for view1.Next() {
		//fmt.Println("inside for loop")
		var id int

		err := view1.Scan(&id)
		if err != nil {
			// handle this error
			panic(err)

		}
		fmt.Println("the two transactions that are invalid are ", id)
	}
	fmt.Println("Invalid transactions if consider supplied id 8 can both bean_types 1, 2 to the carrier id 2 : ")
	//else we can dig a little more deeper and show that for delivery 2 can be invalid for bean type 1.
	view2, err := db.Query(`Select DISTINCT(c1.id),c1.bean_type_id from
	(Select  D.id, D.supplier_id, D.driver_id, SBT.bean_type_id from supplier_bean_type SBT inner join delivery D on SBT.supplier_id = D.supplier_id) as c1
	where NOT EXISTS
	(select distinct(c2.driver, c2.bean_id) from (Select C.carrier_id as carrier , C.bean_type_id as bean_id , D.id as driver from carrier_bean_type C  inner join driver D on C.carrier_id = D.carrier_id) as c2  where c2.driver = c1.driver_id AND c1.bean_type_id= c2.bean_id) order by c1.id`)
	if err != nil {
		// handle this error
		panic(err)

	}
	defer view2.Close()
	for view2.Next() {
		//fmt.Println("inside for loop")
		var id int
		var bean_id int
		err := view2.Scan(&id, &bean_id)
		if err != nil {
			// handle this error
			panic(err)

		}
		fmt.Println("the  transactions that are invalid are ", "ID", id, "for bean", bean_id)
	}
	fmt.Println("done")
}
