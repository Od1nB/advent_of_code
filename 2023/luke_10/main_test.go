package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var grid = []string{"-L|F7", "7S-7|", "L|7||", "-L-J|", "L|-JF"}
var pm = fillMaze(grid)

func TestTypParse(t *testing.T) {
	assert.Equal(t, pm.pipes["00"].typ, "-")
	assert.Equal(t, pm.pipes["01"].typ, "7")
	assert.Equal(t, pm.pipes["44"].typ, "F")
	assert.Equal(t, pm.pipes["10"].typ, "L")
	assert.Equal(t, pm.pipes["13"].typ, "L")
	assert.Equal(t, pm.pipes["43"].typ, "|")
	assert.Equal(t, pm.pipes["42"].typ, "|")
}

func TestSLoop(t *testing.T) {
	assert.Equal(t, pm.pipes["11"].typ, "S")
	assert.Equal(t, pm.pipes["11"].inLoop, true)
	assert.Equal(t, (&pipePart{}).n, pm.pipes["11"].n)
	assert.Equal(t, pm.pipes["12"], pm.pipes["11"].s)
	assert.Equal(t, true, pm.pipes["12"].inLoop)
	assert.Equal(t, "-", pm.pipes["11"].e.typ)
}
