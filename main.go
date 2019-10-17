package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"net/smtp"
	"os"
	"runtime"
	"strings"
	"sync"
)

var (
	username string
	password string
	name     string
	address  string
)

var intro []string
var info []string
var outro []string

var wg = sync.WaitGroup{}

func init() {
	info = []string{
		" I wanted to get as much information as possible, would you care to send me a tee-shirt and some additional information?",
		" it's has become a goal of mine to enroll here. I was wondering if you could send me some information and maybe even a tee-shirt :)",
		" I wanted some more information, potentially even some merch if that's not too much trouble.",
		" I was wondering if you could send me some information along with a t-shirt",
	}
	outro = []string{
		" I always loved the colors over at ",
		" I always found the campus beautiful at ",
		" Like the saying goes, the grass is greener at ",
	}
	intro = []string{
		"I am currently a junior in high school and was genuinely interested in enrolling into one of the programs at ",
		"I am in the college search process and came across ",
		"I am new in the college search process and found  ",
		"As someone who is new in the college search process, I was overwhelmed before I came across ",
	}

}

func main() {
	runtime.GOMAXPROCS(100)
	// username, password, name, email, address
	if len(os.Args) != 5 {
		log.Fatal("USAGE [google username] [google password] [first last name] [address]")
	}
	username, password, name, address = os.Args[1], os.Args[2], os.Args[3], os.Args[4]

	file, err := os.Open("college_emails_names.csv")
	if err != nil {
		log.Fatal("could not open")
	}

	lines, err := csv.NewReader(file).ReadAll()

	for index, line := range lines {
		if index == 0 {
			continue
		}
		temp := line[0]
		collegeName := strings.Split(temp, ",")[0]
		email := line[1]
		fmt.Println(collegeName)

		wg.Add(1)
		go sendMail(name, email, collegeName, address)

	}

	wg.Wait()
	log.Println("finished sending all emails")
}

// Email is the your email!
func sendMail(name, email, college, address string) {
	auth := smtp.PlainAuth("", username, password, "smtp.gmail.com")
	to := email
	msg := "From: " + email + "\n" +
		"To: " + to + "\n" +
		"Subject:" + name + " Interest\n\n" +
		"Hello, " + "\n\n" + intro[rand.Intn(len(info))] + college + info[rand.Intn(len(info))] + outro[rand.Intn(len(outro))] + college + "\n\n" +
		"My address is " + "\n" + address + "\n\n" +

		"Thank you" + "\n" + name + "\n"
	err := smtp.SendMail("smtp.gmail.com:587", auth, username, []string{to}, []byte(msg))
	if err != nil {
		log.Printf("failed %v", err)
	}

	wg.Done()
}

// Function for testing the email responses
func testMail(name, email, college, address string) {
	to := email
	msg := "From: " + email + "\n" +
		"To: " + to + "\n" +
		"Subject:" + name + " " + college + " Interest\n\n" +
		"Hello, " + "\n\n" + intro[rand.Intn(len(info))] + college + info[rand.Intn(len(info))] + outro[rand.Intn(len(outro))] + college + "\n\n" +
		"My address is " + "\n" + address + "\n\n" +

		"Thank you" + "\n" + name + "\n"

	fmt.Println(msg)
	wg.Done()

}
