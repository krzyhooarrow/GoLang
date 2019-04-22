package main

type Task struct {
	arg1      int
	arg2      int
	operation string
	score int
}



func createTask(arg1 int, arg2 int, operation string) Task {
	var pointer int
	return Task{arg1, arg2, operation,pointer}
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