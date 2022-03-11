package zpl

import "fmt"

func Parese(input []byte) {
	var buffCommands = make(map[int]Command)

	for i, item := range input {
		//	var c Command
		var commandToken string

		if "^" == string(item) {
			commandToken = fmt.Sprintf("^%s%s", string(input[i+1]), string(input[i+2]))
		}

		comm, err := CreateCommand(commandToken)
		if nil == err {
			buffCommands[i] = comm
		}
	}
	/*
		var commadIndex int
		for i, item := range input {
			c, ok := buffCommands[i]
			if ok {
				commadIndex = 0
			}

		}
	*/
	//fmt.Print(buffCommands)

}
