package main

import (
	"errors"
	"fmt"
	"os"
)

type Employee struct {
	Id int
	FirstName string
	LastName string
}

const (
	Excelent = 5
	Good = 4
	Fair = 3
	Poor = 2
	Unsatisfactory = 1
)

type Developer struct {
	Individual Employee
	HourlyRate int
	HoursWorkedInYear int
	Review map[string]interface{}
}

type Manager struct {
	Individual Employee
	Salary float64
	CommissionRate float64
}

type Payer interface {
	Pay() (string, float64)
}



func main() {
	employeeReview := make(map[string]interface{})
	employeeReview["WorkQuality"] = 5
	employeeReview["TeamWork"] = 2
	employeeReview["Communication"] = "Poor"
	employeeReview["Problem-solving"] = 4
	employeeReview["Dependability"] = "Unsatisfactory"
	d := Developer{Individual: Employee{Id: 1, FirstName: "Eric", LastName: "Davis"}, HourlyRate: 35, HoursWorkedInYear: 2400, Review: employeeReview}
	m := Manager{Individual: Employee{Id: 2, FirstName: "Mr.", LastName: "Boss"}, Salary: 150000, CommissionRate: .07}
	err := d.ReviewRating()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	payDetails(d)
	payDetails(m)
}



func (d Developer) Pay() (string, float64) {
	name := d.Individual.FirstName + " " + d.Individual.LastName
	yearPay := float64(d.HourlyRate) * float64(d.HoursWorkedInYear)

	return name, yearPay
}

func (m Manager) Pay() (string, float64) {
	name := m.Individual.FirstName + " " + m.Individual.LastName
	yearPay := m.Salary + (m.Salary * m.CommissionRate)

	return name, yearPay
}

func payDetails(p Payer) {
	name, pay := p.Pay()
	fmt.Println("Name: ", name)
	fmt.Println("Pay: ", pay)
}

func (d Developer) ReviewRating() error {
	total := 0

	for _, v := range d.Review {
		rating, err := OverallReview(v)
		if err != nil {
			return err
		}
		total += rating
	}

	averageReview := float64(total) / float64(len(d.Review))
	fmt.Printf("%s got a review rating of %.2f\n", d.Individual.FirstName, averageReview)

	return nil
}

func OverallReview(i interface{}) (int, error) {
	switch v := i.(type) {
	case int:
		return v, nil
	case string:
		rating, err := convertReviewToInt(v)
		if err != nil {
			return 0, err
		}		
		return rating, nil
	default:
		return 0, errors.New("invalid type")
	}
}


func convertReviewToInt(str string) (int, error) {
	switch  str {
	case "Excelent":
		return Excelent, nil
	case "Good":
		return Good, nil
	case "Fair":
		return Fair, nil
	case "Poor":
		return Poor, nil
	case "Unsatisfactory":
		return Unsatisfactory, nil
	default:
		return 0, errors.New("invalid review: " + str)
	}
}