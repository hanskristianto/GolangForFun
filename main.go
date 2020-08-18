package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"time"
	"strconv"
)

var data = [][]string{{"ID", "Name", "Age", "Amount", "Date_time", "Status"}}

func main() {
	fmt.Println("")
	fmt.Println("#Example format")
	fmt.Println("-- add no_id name age amount")
	fmt.Println("-- installment no_id month")
	fmt.Println("-- status no_id")
	fmt.Println("-- find_by_amount_accepted amount")
	fmt.Println("-- find_by_amount_rejected amount")
	fmt.Println("")
	
	reader := bufio.NewReader(os.Stdin)
	
	getInput, _ := reader.ReadString('\n')
	
	input := strings.Fields(getInput)

	switch input[0]{
		case "add": 
			add(input[1:5])
		case "installment":
			installment(data[1:len(data)], input[1:3])
		case "status":
			checkStatus(input[1])
		case "find_by_amount_accepted":
			findAmountAccepted(input[1])
		case "find_by_amount_rejected":
			findAmountRejected(input[1])
		case "exit":
			defer fmt.Println("Goodbye")
		default:
			fmt.Printf("Wrong format, please follow example format!")
			main()
		
	}	
}

func add(row []string){
	// get date
	dt := time.Now()
	
	age_to_int, _ := strconv.Atoi(row[2])
	// custom_row := append(row, dt.Format("2006-01-02"), age_to_int > 17?"Accepted":"Rejected")
	if age_to_int > 17{
		// append date and status
		custom_row := append(row, dt.Format("2006-01-02"), "Accepted")
		data = append(data, custom_row)
	}else{

		custom_row := append(row, dt.Format("2006-01-02"), "Rejected")
		data = append(data, custom_row)
	}
	
	for i:= 0; i<len(data); i++{
		fmt.Println(data[i])
	}
	fmt.Println("=============== [Add success max request",len(data)-1,"of 50]=================")
	
	main()

}

func installment(row [][]string, input []string){

	// just for header but still in array
	var dataInstallment = [][]string{{"Month", "DueDate", "AdministrationFee", "Capital", "Total"}}
	fmt.Println(dataInstallment[0])

	for i:=0; i<len(row); i++{
		// selection if no id 12345 == input
		if row[i][0] == input[0]{
			// convert from string to int
			month, _ := strconv.Atoi(input[1])
			amount, _ := strconv.Atoi(row[i][3])
			// looping how much month devided
			for j:=0; j<month; j++{
				fee := 100
				date,_ := time.Parse("2006-01-02", row[i][4])
				dueDate := date.AddDate(0,j+1,0)

				data := []string{strconv.Itoa(j+1), dueDate.Format("2006-01-02"), strconv.Itoa(fee), strconv.Itoa(amount/month), strconv.Itoa(amount/month+fee)}
				// dataInstallment = append(dataInstallment, data)
				fmt.Println(data)
			}
		}
	}
	
	main()
}


func checkStatus(value string){

	for i:=0; i<len(data); i++{
		if data[i][0] == value{
			// data[i][5] status
			fmt.Println("Loan ID : ", data[i][0], "is", data[i][5])
		}
	}
	main()
}

func findAmountAccepted(value string){
	data_exist := true

	for i:=0; i<len(data); i++{

		// data[i][3]  amount
		// data[i][5] status
		amount, _ := strconv.Atoi(data[i][3])
		input_amount, _ := strconv.Atoi(value)

		if amount == input_amount && data[i][5] == "Accepted"{
			fmt.Println("ID : ", data[i][0])
			data_exist = false
			
		}
	}
	if data_exist{
		fmt.Println("Sorry, data not found!")
	}
	main()
}

func findAmountRejected(value string){
	data_exist := true

	for i:=0; i<len(data); i++{

		// data[i][3]  amount
		// data[i][5] status
		amount, _ := strconv.Atoi(data[i][3])
		input_amount, _ := strconv.Atoi(value)

		if amount == input_amount && data[i][5] == "Rejected"{
			fmt.Println("ID : ", data[i][0])
			data_exist = false
			
		}
	}
	if data_exist{
		fmt.Println("Sorry, data not found!")
	}
	main()
}