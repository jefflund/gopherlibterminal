package gopherlibterminal

// #cgo CFLAGS:
// #cgo amd64 386 CFLAGS: -DX86=1
// #cgo LDFLAGS: -lBearLibTerminal -Wl,-rpath .
// #include <stdlib.h>
// #include <BearLibTerminal.h>
import "C"
import "errors"
import "unsafe"

type Input int

const (
	TkNone Input =		0x00
	TkA =			0x04
	TkB =			0x05
	TkC =			0x06
	TkD =			0x07
	TkE =			0x08
	TkF =			0x09
	TkG =			0x0A
	TkH =			0x0B
	TkI =			0x0C
	TkJ =			0x0D
	TkK =			0x0E
	TkL =			0x0F
	TkM =			0x10
	TkN =			0x11
	TkO =			0x12
	TkP =			0x13
	TkQ =			0x14
	TkR =			0x15
	TkS =			0x16
	TkT =			0x17
	TkU =			0x18
	TkV =			0x19
	TkW =			0x1A
	TkX =			0x1B
	TkY =			0x1C
	TkZ =			0x1D
	Tk1 =			0x1E
	Tk2 =			0x1F
	Tk3 =			0x20
	Tk4 =			0x21
	Tk5 =			0x22
	Tk6 =			0x23
	Tk7 =			0x24
	Tk8 =			0x25
	Tk9 =			0x26
	Tk0 =			0x27
	TkReturn =		0x28
	TkEnter =		0x28
	TkEscape =		0x29
	TkBackspace =		0x2A
	TkTab =			0x2B
	TkSpace =		0x2C
	TkMinus =		0x2D
	TkEquals =		0x2E
	TkLbracket =		0x2F
	TkRbracket =		0x30
	TkBackslash =		0x31
	TkSemicolon =		0x33
	TkApostrophe =		0x34
	TkGrave =		0x35
	TkComma =		0x36
	TkPeriod =		0x37
	TkSlash =		0x38
	TkF1 =			0x3A
	TkF2 =			0x3B
	TkF3 =			0x3C
	TkF4 =			0x3D
	TkF5 =			0x3E
	TkF6 =			0x3F
	TkF7 =			0x40
	TkF8 =			0x41
	TkF9 =			0x42
	TkF10 =			0x43
	TkF11 =			0x44
	TkF12 =			0x45
	TkPause =		0x48
	TkInsert =		0x49
	TkHome =		0x4A
	TkPageup =		0x4B
	TkDelete =		0x4C
	TkEnd =			0x4D
	TkPagedown =		0x4E
	TkRight =		0x4F
	TkLeft =		0x50
	TkDown =		0x51
	TkUp =			0x52
	TkKpDivide =		0x54
	TkKpMultiply =		0x55
	TkKpMinus =		0x56
	TkKpPlus =		0x57
	TkKpEnter =		0x58
	TkKp_1 =		0x59
	TkKp_2 =		0x5A
	TkKp_3 =		0x5B
	TkKp_4 =		0x5C
	TkKp_5 =		0x5D
	TkKp_6 =		0x5E
	TkKp_7 =		0x5F
	TkKp_8 =		0x60
	TkKp_9 =		0x61
	TkKp_0 =		0x62
	TkKpPeriod =		0x63
	TkShift =		0x70
	TkControl =		0x71

	TkMouseLeft =		0x80
	TkMouseRight =		0x81
	TkMouseMiddle =		0x82
	TkMouseX1 =		0x83
	TkMouseX2 =		0x84
	TkMouseMove =		0x85
	TkMouseScroll =		0x86
	TkMouseX =		0x87
	TkMouseY =		0x88
	TkMousePixelX =		0x89
	TkMousePixelY =		0x8A
	TkMouseWheel =		0x8B
	TkMouseClicks =		0x8C

	TkKeyReleased =		0x100

	TkWidth =		0xC0
	TkHeight =		0xC1
	TkCellWidth =		0xC2
	TkCellHeight =		0xC3
	TkColor =		0xC4
	TkBkcolor =		0xC5
	TkLayer =		0xC6
	TkComposition =		0xC7
	TkChar =		0xC8
	TkWchar =		0xC9
	TkEvent =		0xCA
	TkFullscreen =		0xCB

	TkClose =		0xE0
	TkResized =		0xE1

	TkInputNone =		0
	TkInputCancelled =	-1
)

func Open() error {
	rv := C.terminal_open()
	if rv == 0 {
		return errors.New("failed to start bearlibterminal")
	}
	return nil
}

func Close() {
	C.terminal_close()
}

func Set(opts string) error {
	str := C.CString(opts)
	rv := C.terminal_set(str)
	C.free(unsafe.Pointer(str))
	if rv == 0 {
		return errors.New("failed to set options")
	}
	return nil
}

func Color(c uint32) {
	C.terminal_color(C.color_t(c))
}

func BkColor(c uint32) {
	C.terminal_bkcolor(C.color_t(c))
}

func Composition(q bool) {
	if q {
		C.terminal_composition(C.TK_ON)
	} else {
		C.terminal_composition(C.TK_OFF)
	}
}

func Layer(l uint8) {
	C.terminal_layer(C.int(l))
}

func Clear() {
	C.terminal_clear()
}

func ClearArea(x, y, w, h int) {
	C.terminal_clear_area(C.int(x), C.int(y), C.int(w), C.int(h))
}

func Crop(x, y, w, h int) {
	C.terminal_crop(C.int(x), C.int(y), C.int(w), C.int(h))
}

func Refresh() {
	C.terminal_refresh()
}

func Put(x, y int, c rune) {
	C.terminal_put(C.int(x), C.int(y), C.int(c))
}

func Pick(x, y int, l uint8) rune {
	return rune(C.terminal_pick(C.int(x), C.int(y), C.int(l)))
}

func PickColor(x, y int, l uint8) uint32 {
	return uint32(C.terminal_pick_color(C.int(x), C.int(y), C.int(l)))
}

func PickBkColor(x, y int) uint32 {
	return uint32(C.terminal_pick_bkcolor(C.int(x), C.int(y)))
}

func PutExt() {
	// TODO
}

func Print(x, y int, s string) {
	str := C.CString(s)
	C.terminal_print(C.int(x), C.int(y), str)
	C.free(unsafe.Pointer(str))
}

func Measure(s string) {
	str := C.CString(s)
	C.terminal_measure(str)
	C.free(unsafe.Pointer(str))
}

func State(inp Input) int {
	return int(C.terminal_state(C.int(inp)))
}

func Check(inp Input) bool {
	if C.terminal_state(C.int(inp)) == 0 {
		return false
	} else {
		return true
	}
}

func HasInput() bool {
	if C.terminal_has_input() == 0 {
		return false
	} else {
		return true
	}
}

func Read() Input {
	return Input(C.terminal_read())
}

func Peek() Input {
	return Input(C.terminal_peek())
}

func ReadStr() {
	// TODO
}

func Delay(ms int) {
	C.terminal_delay(C.int(ms))
}

func ColorFromName(name string) uint32 {
	str := C.CString(name)
	col := C.color_from_name(str)
	C.free(unsafe.Pointer(str))
	return uint32(col)
}

func ColorFromArgb(a, r, g, b uint8) uint32 {
	ca := C.uint8_t(a)
	cr := C.uint8_t(r)
	cg := C.uint8_t(g)
	cb := C.uint8_t(b)
	return uint32(C.color_from_argb(ca, cr, cg, cb))
}