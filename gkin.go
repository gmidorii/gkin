package gkin

import "fmt"

// Argument is cli arg definition
type Argument struct {
	Gkin string
}

// Run is gkin strating
func Run(arg Argument) error {
	gkin, err := Parse(arg.Gkin)
	if err != nil {
		return err
	}
	name, err := Build(gkin)
	if err != nil {
		return err
	}
	fmt.Println(name)
	return nil
}
