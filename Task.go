package main

type Task struct {
	arg1      int
	arg2      int
	operation string
}



func createTask(arg1 int, arg2 int, operation string) Task {
	return Task{arg1, arg2, operation}
}

func operator(i int) (operator string) {
	if i == 0 {
		return "+"
	}
	if i == 1 {
		return "-"
	}
	if i == 2 {
		return "/"
	}
	if i == 3 {
		return "*"
	}
	return "Bad Operator"
}