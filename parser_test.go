package gkin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	const name = "./test/.gkin.yml"
	actual, err := Parse(name)
	if err != nil {
		t.Errorf("error test func: %v", err)
		return
	}

	expected := Gkin{
		Pipeline: []Pipe{
			{
				Image: "golang:1.10.2",
				Name:  "build",
				Commands: []string{
					"go get",
					"go build",
				},
			},
			{
				Image: "golang:1.10.2",
				Name:  "test",
				Commands: []string{
					"go test",
				},
			},
		},
	}

	assert.Equal(t, expected, actual, "not equal parsed Gkin struct")
}
