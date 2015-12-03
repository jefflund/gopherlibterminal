package glt

import "testing"
import "math/rand"

func TestRain(t *testing.T) {
	cols, lines := 0, 0
	delay := 50

	xpos, ypos := make([]int, 5), make([]int, 5)
	Open()
	defer Close()

	cols = State(TkWidth) - 4
	lines = State(TkHeight) - 4

	for j := 4; j >= 0; j-- {
		xpos[j] = rand.Intn(cols) + 2
		ypos[j] = rand.Intn(lines) + 2
	}

	j := 0
	for {
		x := rand.Intn(cols) + 2
		y := rand.Intn(lines) + 2
		Color(0xffffffff)
		Print(x, y, ".")

		Print(xpos[j], ypos[j], "o")
		if j == 0 {
			j = 4
		} else {
			j--
		}
		Print(xpos[j], ypos[j], "O")
		if j == 0 {
			j = 4
		} else {
			j--
		}
		Color(0xff00007f)
		Print(xpos[j], ypos[j]-1, "-")
		Print(xpos[j]-1, ypos[j], "|.|")
		Print(xpos[j], ypos[j]+1, "-")
		if j == 0 {
			j = 4
		} else {
			j--
		}
		Color(0xff0000ff)
		Print(xpos[j], ypos[j]-2, "-")
		Print(xpos[j]-1, ypos[j]-1, "/ \\")
		Print(xpos[j]-2, ypos[j], "| O |")
		Print(xpos[j]-1, ypos[j]+1, "\\ /")
		Print(xpos[j], ypos[j]+2, "-")
		if j == 0 {
			j = 4
		} else {
			j--
		}
		Print(xpos[j], ypos[j]-2, " ")
		Print(xpos[j]-1, ypos[j]-1, "   ")
		Print(xpos[j]-2, ypos[j], "     ")
		Print(xpos[j]-1, ypos[j]+1, "   ")
		Print(xpos[j], ypos[j]+2, " ")

		xpos[j] = x
		ypos[j] = y
		Refresh()

		Delay(delay)

		if HasInput() {
			inp := Read()
			if inp == TkQ && Check(TkShift) || inp == TkClose {
				break
			}
		}
	}
}
