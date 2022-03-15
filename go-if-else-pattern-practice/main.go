package main

func changeBefore() {
	if true {

	}
	if true {

	}
	if true {

	}
	if true {

	}
	if true {

	}
}
func changeAfter() {
	if check1() {

	}
	if check2() {

	}
}

func check1() bool {
	return true && true && true
}
func check2() bool {
	return true && true
}

func main() {
	changeBefore()
	//多个单个if
	changeAfter()
	//合并成多个if并列条件的函数
}
