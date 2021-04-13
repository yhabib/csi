package metrics

import (
	"encoding/csv"
	"log"
	"math"
	"os"
	"strconv"
)

type (
	UserId  int
	UserAge uint8
)

var (
	ages     = []UserAge{}
	payments = []int{}

	// Leveraging that we know these values and compiler doesnt'
	numUsers        = 100_000
	numPayments     = 1_000_000
	maxPaymentCents = 100_000_000
)

// Worst case scenario: NumUsers * MaxAge = 12_000_000 32 bits are enough to hold it
func AverageAge() float64 {
	var sum uint32

	// By writing the loop in this way rather than the traditional for i=0;i<numUsers; i++ it gets 6k faster
	// More crazy. range is more performant than while
	for _, age := range ages {
		sum += uint32(age)
	}
	return float64(sum) / float64(numUsers)
}

// Worst case scenario: NumPayments * MaxPaymentCents < int64
func AveragePaymentAmount() float64 {
	var sum int64

	for _, cents := range payments {
		sum += int64(cents)
	}
	return float64(sum) / float64(100*numPayments)
}

// Compute the standard deviation of payment amounts
func StdDevPaymentAmount() float64 {
	mean := AveragePaymentAmount()
	squaredDiffs := 0.0
	count := 0

	for _, cents := range payments {
		count++
		diff := float64(cents/100) - mean
		squaredDiffs += diff * diff
	}
	return math.Sqrt(squaredDiffs / float64(count))
}

func LoadData() {
	f, err := os.Open("users.csv")
	if err != nil {
		log.Fatalln("Unable to read users.csv", err)
	}
	reader := csv.NewReader(f)
	userLines, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("Unable to parse users.csv as csv", err)
	}

	for _, line := range userLines {
		age, _ := strconv.Atoi(line[2])
		ages = append(ages, UserAge(age))
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
		paymentCents, _ := strconv.Atoi(line[0])
		payments = append(payments, paymentCents)
	}
}
