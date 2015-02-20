package gopherlibterminal

import "testing"

func TestEverything(t *testing.T) {
	x, y := 0, 0
	nx, ny := 0, 0
	Open()

	for {
		Color(0xff0000ff)
		Print(x, y, "@")
		x, y = nx, ny
		Color(0xffff0000)
		Print(x, y, "@")
		Refresh()

		inp := Read()
		if inp == TkQ && Check(TkShift) || inp == TkClose {
			break
		} else if inp == TkLeft {
			nx--
		} else if inp == TkRight {
			nx++
		} else if inp == TkUp {
			ny--
		} else if inp == TkDown {
			ny++
		}
	}

	Close()
}