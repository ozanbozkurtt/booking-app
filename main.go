package main

import (
	"fmt"
	"project/helper"
	"sync"
	"time"
)

var conferenceName = "Go Conference" //var ile değişken tanımlayabiliriz. var yerine := de kullanılabilir.
const conferenceTickets int = 50     //const ile sabit değişken tanımlayabiliriz.
var remaningTickets uint = 50        //uint ile yalnızca pozitif tamsayı değişken tanımlayabiliriz.
var bookings = make([]UserData, 0)   //[]UserData ile diziye ekleme yaptık tanımladık.
//var bookings = make([]map[string]string, 0)  aldığımız mapleri stringe çevirip diziye ekleyebiliriz.

//var bookings =[]string{}  []string ile slice tanımlayabiliriz.

type UserData struct { // struct ile mapde olduğu gibi degişkenleri tanımlayabiliriz. struct yapısında hepsinin aynı tip olması gerekmiyor.
	FirstName string
	LastName  string
	Email     string
	Tickets   uint
}

var wg = sync.WaitGroup{} //sync.WaitGroup ile goroutine sayısını belirleyebiliriz.

func main() {

	greetUser() //fonksiyonu çağırdık.

	// fmt.Printf("Welcome to %v booking app\n", conferenceName) , fmt.Printf ile değişkenleri %v'den sonra ekrana yazdırabiliriz. \n ile alt satıra geçebiliriz.
	//fmt.Println("There are", conferenceTickets, "total tickets and", remaningTickets, "remaining") // , ile değişkenleri birleştirebiliriz.
	//fmt.Println("Get your ticket now!")                                                            //println yazarsak bir alt satıra geçer

	//for {  remaningTickets > 0 && len(bookings) < 50  loopu statement ile yazabiliriz.
	//[]string{} ile dizi tanımlayabiliriz. var bookings = []string{} şeklinde de tanımlanabilir.

	firstName, lastName, email, userTickets := getUserInput() //getUserInput ile kullanıcıdan değer alırız.
	isValidName, isValidTicketNumber, isValidEmail := helper.ValidateUserInput(firstName, lastName, email, userTickets, remaningTickets)

	if isValidEmail && isValidName && isValidTicketNumber {
		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)                                              //wg.Add ile goroutine sayısını bir arttırırız.
		go sendTicket(userTickets, firstName, lastName, email) //go ile goroutine ile fonksiyonu çağırırız. Bir fonksiyonu beklemek zorunda kalmadan diğerine geçer.

		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remaningTickets == 0 {
			fmt.Println("Sorry we are sold out")

		}
	} else {
		if !isValidEmail {
			fmt.Println("Please enter a valid email")
		}
		if !isValidName {
			fmt.Println("Please enter a valid name")
		}
		if !isValidTicketNumber {
			fmt.Println("Please enter a valid ticket number")
		}
	}
	wg.Wait() //goroutine sayısını beklemek zorunda kalmadan diğerine geçer.

}
func greetUser() { // func ile fonksiyon tanımladık ve yukarıda bu fonksunu cağırdık.
	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Println("There are", conferenceTickets, "total tickets and", remaningTickets, "remaining")
	fmt.Println("Get your ticket now!")
}
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		//var names = strings.Fields(booking)
		firstNames = append(firstNames, booking.FirstName)
	}
	return firstNames //bookings içinden her bir first name in alınmasını sağladık.

}
func getUserInput() (string, string, string, uint) {
	var firstName string //daha sonradan değer tanımlayacağımız için data type'ını belirtmemiz gerekiyor.
	var email string
	var userTickets uint
	var lastName string

	fmt.Println("Please enter your first name:")
	fmt.Scan(&firstName) //fmt.Scan ile kullanıcıdan değer alırız. & ile değişkenin memorydeki adresini alırız.

	fmt.Println("Please enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email:")
	fmt.Scan(&email)

	fmt.Println("Please enter how many tickets you want:")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}
func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remaningTickets = remaningTickets - userTickets
	var userData = UserData{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Tickets:   userTickets,
	}

	bookings = append(bookings, userData) //append ile diziye değer ekleyebiliriz.
	fmt.Printf("Thank you %v %v, your email is %v and you have %v tickets\n", firstName, lastName, email, userTickets)
	fmt.Printf("There are %v tickets remaining\n", remaningTickets)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second) //10 saniye bekler. time.Sleep ile fonksiyonu bekleriz.
	var ticket = fmt.Sprintf("%v tickets for %v %v\n", userTickets, firstName, lastName)
	fmt.Println("####################################")
	fmt.Printf("Sending ticket: \n %v \n to email adress %v\n", ticket, email)
	fmt.Println("####################################")
	wg.Done() //wg.Done ile goroutine sayısını azaltırız.
}
