package calc

func Add(x, y int) int {
	return x + y
}

func Subtract(x, y int) int {
	return x - y
}

func VoteEligibility(age int) string {
	if age >= 18 {
		return "You are eligible to vote"
	} else {
		return "You are not eligible to vote"
	}
}

func GetById(id string) (string, error) {
	if id == "105" {
		return "Task 105", nil
	}
	return "", nil
}
