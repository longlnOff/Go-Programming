package main
import "fmt"

// s1 s2 s3 point to same hidden array
func linked() (int, int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1
	s3 := s1[:]
	s1[3] = 99

	return s1[3], s2[3], s3[3]
}

// append create new hidden array for s1
// so s1 and s2 point to different hidden array
func noLink() (int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1
	s1 = append(s1, 6)
	s1[3] = 99

	return s1[3], s2[3]
}

// current length s1 is 5
// append one value to s1 but total length is smaller than capacity (10)
// so s1 and s2 point to the same hidden array
func capLinked() (int, int) {
	s1 := make([]int, 5, 10)
	s1[0], s1[1], s1[2], s1[3], s1[4] = 1, 2, 3, 4, 5
	s2 := s1
	s1 = append(s1, 6)
	s1[3] = 99

	return s1[3], s2[3]
}

// same as above but we the total length surpass capacity
// so it need to create an other hidden array for s1
func capNoLink() (int, int) {
	s1 := make([]int, 5, 10)
	s1[0], s1[1], s1[2], s1[3], s1[4] = 1, 2, 3, 4, 5
	s2 := s1
	s1 = append(s1, []int{10: 11}...)
	s1[3] = 99
	return s1[3], s2[3]
}

// copy to an other hidden array
func copyNoLink() (int, int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, len(s1))
	copied := copy(s2, s1)
	s1[3] = 99
	return s1[3], s2[3], copied
}

func appendNoLink() (int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := append([]int{}, s1...)
	s1[3] = 99
	return s1[3], s2[3]
}

func main() {
	l1, l2, l3 := linked()
	fmt.Println("Linked   :", l1, l2, l3)
	nl1, nl2 := noLink()
	fmt.Println("No Link   :", nl1, nl2)
	cl1, cl2 := capLinked()
	fmt.Println("Cap Link  :", cl1, cl2)
	cnl1, cnl2 := capNoLink()
	fmt.Println("Cap No Link :", cnl1, cnl2)
	copy1, copy2, copied := copyNoLink()
	fmt.Print("Copy No Link: ", copy1, copy2)
	fmt.Printf(" (Number of elements copied %v)\n", copied)
	a1, a2 := appendNoLink()
	fmt.Println("Append No Link:", a1, a2)
}