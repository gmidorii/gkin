package gkin

import "fmt"

// Argument is cli arg definition
type Argument struct {
	Gkin Gkin
}

// Run is gkin strating
func Run(arg Argument) error {
	name, err := Build(arg.Gkin.Pipeline[0])
	if err != nil {
		return err
	}
	fmt.Println(name)
	return nil
}
