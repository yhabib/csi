package metrics

import (
	"encoding/csv"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

type (
	UserId  int
	UserMap map[UserId]*User
	UserAge uint8
)

type Address struct {
	fullAddress string
	zip         int
}

type DollarAmount struct {
	dollars, cents uint64
}

type Payment struct {
	amount DollarAmount
	time   time.Time
}

type User struct {
	id       UserId
	name     string
	age      uint8
	address  Address
	payments []Payment
}

var (
	userAge       = make(map[UserId]UserAge)
	userIdToCents = make(map[UserId][]int)
)

func AverageAge(users UserMap) float64 {
	var average float32
	count := 0
	// for _, u := range users {
	// 	count += 1
	// 	average += (float64(u.age) - average) / count
	// }

	for _, age := range userAge {
		count++
		average += (float32(age) - average) / float32(count)
	}

	return float64(average)
}

func AveragePaymentAmount(users UserMap) float64 {
	var average float64
	count := 0
	// for _, u := range users {
	// 	for _, p := range u.payments {
	// 		count += 1
	// 		amount := float64(p.amount.dollars) + float64(p.amount.cents)/100
	// 		average += (amount - average) / count
	// 	}
	// }
	for _, payments := range userIdToCents {
		for _, cents := range payments {
			count++
			average += (float64(cents) - average) / float64(count)
		}
	}
	return average / 100
}

// Compute the standard deviation of payment amounts
func StdDevPaymentAmount(users UserMap) float64 {
	mean := AveragePaymentAmount(users)
	squaredDiffs := 0.0
	count := 0
	// for _, u := range users {
	// 	for _, p := range u.payments {
	// 		count += 1
	// 		amount := float64(p.amount.dollars) + float64(p.amount.cents)/100
	// 		diff := amount - mean
	// 		squaredDiffs += diff * diff
	// 	}
	// }

	for _, payments := range userIdToCents {
		for _, cents := range payments {
			count++
			diff := float64(cents/100) - mean
			squaredDiffs += diff * diff
		}
	}
	return math.Sqrt(squaredDiffs / float64(count))
}

func LoadData() UserMap {
	f, err := os.Open("users.csv")
	if err != nil {
		log.Fatalln("Unable to read users.csv", err)
	}
	reader := csv.NewReader(f)
	userLines, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("Unable to parse users.csv as csv", err)
	}

	users := make(UserMap, len(userLines))
	for _, line := range userLines {
		id, _ := strconv.Atoi(line[0])
		name := line[1]
		age, _ := strconv.Atoi(line[2])
		address := line[3]
		zip, _ := strconv.Atoi(line[3])
		users[UserId(id)] = &User{UserId(id), name, uint8(age), Address{address, zip}, []Payment{}}
		userAge[UserId(id)] = UserAge(age)

	}

	f, err = os.Open("payments.csv")
	if err != nil {
		log.Fatalln("Unable to read payments.csv", err)
	}
	reader = csv.NewReader(f)
	paymentLines, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("Unable to parse payments.csv as csv", err)
	}

	for _, line := range paymentLines {
		userId, _ := strconv.Atoi(line[2])
		paymentCents, _ := strconv.Atoi(line[0])
		datetime, _ := time.Parse(time.RFC3339, line[1])
		users[UserId(userId)].payments = append(users[UserId(userId)].payments, Payment{
			DollarAmount{uint64(paymentCents / 100), uint64(paymentCents % 100)},
			datetime,
		})
		userIdToCents[UserId(userId)] = append(userIdToCents[UserId(userId)], paymentCents)
	}

	return users
}
