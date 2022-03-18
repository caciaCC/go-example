package main

import (
	"errors"
	"fmt"
)

func changeBefore() error {
	if true {
		if true {
			if true {
				if true {

				} else {
					return errors.New("...")
				}
			} else {
				return errors.New("...")
			}
		} else {
			return errors.New("...")
		}
	} else {
		return errors.New("...")
	}
	return nil
}

func changeAfter() error {
	for {
		if !true {
			return errors.New("...")
		}
		if !true {
			return errors.New("...")
		}
		if !true {
			return errors.New("...")
		}
		if !true {
			return errors.New("...")
		}
		break
	}
	return nil
}

func main() {
	fmt.Println(changeBefore())
	//	使用卫语句代替嵌套if
	fmt.Println(changeAfter())
}
