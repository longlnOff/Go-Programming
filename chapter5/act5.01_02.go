package main

import "fmt"


type Employee struct {
	Id int
	FirstName string
	LastName string
}

type Developer struct {
	Individual Employee
	HourlyRate int
	WorkWeek [7]int
}

type Weekday int

const (
	Monday Weekday = 2
	Tuesday Weekday = 3
	Wednesday Weekday = 4
	Thursday Weekday = 5
	Friday Weekday = 6
	Saturday Weekday = 7
	Sunday Weekday = 8
)

func (d *Developer) LogHours (day Weekday, hours int) {
	d.WorkWeek[day-2] = hours
}

func (d *Developer) HoursWorked() Weekday {
	var total Weekday = 0
	for _, v := range d.WorkWeek {
		total += Weekday(v)
	}

	return total
}

func (d *Developer) PayDay() (int, bool) {
	totalHours := d.HoursWorked()
	if totalHours < 40 {
		return int(totalHours) * d.HourlyRate, false
	}
	return 40 * d.HourlyRate + (int(totalHours) - 40) * d.HourlyRate * 2, true
}

func (d *Developer) PayDetails() {
	for _, v := range d.WorkWeek {
		switch v {
			case int(Monday):
				fmt.Println("Monday hours: ", v)
			case int(Tuesday):
				fmt.Println("Tuesday hours: ", v)
			case int(Wednesday):
				fmt.Println("Wednesday hours: ", v)
			case int(Thursday):
				fmt.Println("Thursday hours: ", v)
			case int(Friday):
				fmt.Println("Friday hours: ", v)
			case int(Saturday):
				fmt.Println("Saturday hours: ", v)
			case int(Sunday):
				fmt.Println("Sunday hours: ", v)
		}
	}
	fmt.Printf("\nHours worked this week:  %d\n", d.HoursWorked())
	pay, overtime := d.PayDay()
	fmt.Println("Pay for the week: $", pay)
	fmt.Println("Is this overtime pay: ", overtime)
	fmt.Println()
}

func nonLoggedHours() func (int) int {
	total := 0
	return func (x int) int {
		total += x
		return total
	}
}
func main() {
	d := Developer{
					Individual: Employee{Id: 1, FirstName: "Tony", LastName: "Stark"}, 
					HourlyRate: 10,
				}
	x := nonLoggedHours()
	fmt.Println("Tracking hours worked thus far today: ", x(2))
	fmt.Println("Tracking hours worked thus far today: ", x(3))
	fmt.Println("Tracking hours worked thus far today: ", x(5))
	fmt.Println()
	d.LogHours(Monday, 8)
	d.LogHours(Tuesday, 10)
	d.LogHours(Wednesday, 10)
	d.LogHours(Thursday, 10)
	d.LogHours(Friday, 6)
	d.LogHours(Saturday, 8)
	d.PayDetails()
}