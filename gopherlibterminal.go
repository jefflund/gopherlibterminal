// Package glt or GopherLibTerminal is a wrapper over BearLibTerminal, a
// terminal emulator over SDL.
package glt

// #cgo CFLAGS:
// #cgo amd64 386 CFLAGS: -DX86=1
// #cgo LDFLAGS: -lBearLibTerminal -Wl,-rpath .
// #include <stdlib.h>
// #include <BearLibTerminal.h>
import "C"
import "errors"
import "unsafe"

// Input represents a scancode for events/states.
type Input int

// Input constants.
const (
	TkA          Input = 0x04
	TkB          Input = 0x05
	TkC          Input = 0x06
	TkD          Input = 0x07
	TkE          Input = 0x08
	TkF          Input = 0x09
	TkG          Input = 0x0A
	TkH          Input = 0x0B
	TkI          Input = 0x0C
	TkJ          Input = 0x0D
	TkK          Input = 0x0E
	TkL          Input = 0x0F
	TkM          Input = 0x10
	TkN          Input = 0x11
	TkO          Input = 0x12
	TkP          Input = 0x13
	TkQ          Input = 0x14
	TkR          Input = 0x15
	TkS          Input = 0x16
	TkT          Input = 0x17
	TkU          Input = 0x18
	TkV          Input = 0x19
	TkW          Input = 0x1A
	TkX          Input = 0x1B
	TkY          Input = 0x1C
	TkZ          Input = 0x1D
	Tk1          Input = 0x1E
	Tk2          Input = 0x1F
	Tk3          Input = 0x20
	Tk4          Input = 0x21
	Tk5          Input = 0x22
	Tk6          Input = 0x23
	Tk7          Input = 0x24
	Tk8          Input = 0x25
	Tk9          Input = 0x26
	Tk0          Input = 0x27
	TkReturn     Input = 0x28
	TkEnter      Input = 0x28
	TkEscape     Input = 0x29
	TkBackspace  Input = 0x2A
	TkTab        Input = 0x2B
	TkSpace      Input = 0x2C
	TkMinus      Input = 0x2D
	TkEquals     Input = 0x2E
	TkLbracket   Input = 0x2F
	TkRbracket   Input = 0x30
	TkBackslash  Input = 0x31
	TkSemicolon  Input = 0x33
	TkApostrophe Input = 0x34
	TkGrave      Input = 0x35
	TkComma      Input = 0x36
	TkPeriod     Input = 0x37
	TkSlash      Input = 0x38
	TkF1         Input = 0x3A
	TkF2         Input = 0x3B
	TkF3         Input = 0x3C
	TkF4         Input = 0x3D
	TkF5         Input = 0x3E
	TkF6         Input = 0x3F
	TkF7         Input = 0x40
	TkF8         Input = 0x41
	TkF9         Input = 0x42
	TkF10        Input = 0x43
	TkF11        Input = 0x44
	TkF12        Input = 0x45
	TkPause      Input = 0x48
	TkInsert     Input = 0x49
	TkHome       Input = 0x4A
	TkPageup     Input = 0x4B
	TkDelete     Input = 0x4C
	TkEnd        Input = 0x4D
	TkPagedown   Input = 0x4E
	TkRight      Input = 0x4F
	TkLeft       Input = 0x50
	TkDown       Input = 0x51
	TkUp         Input = 0x52
	TkKpDivide   Input = 0x54
	TkKpMultiply Input = 0x55
	TkKpMinus    Input = 0x56
	TkKpPlus     Input = 0x57
	TkKpEnter    Input = 0x58
	TkKp1        Input = 0x59
	TkKp2        Input = 0x5A
	TkKp3        Input = 0x5B
	TkKp4        Input = 0x5C
	TkKp5        Input = 0x5D
	TkKp6        Input = 0x5E
	TkKp7        Input = 0x5F
	TkKp8        Input = 0x60
	TkKp9        Input = 0x61
	TkKp0        Input = 0x62
	TkKpPeriod   Input = 0x63
	TkShift      Input = 0x70
	TkControl    Input = 0x71

	TkMouseLeft   Input = 0x80
	TkMouseRight  Input = 0x81
	TkMouseMiddle Input = 0x82
	TkMouseX1     Input = 0x83
	TkMouseX2     Input = 0x84
	TkMouseMove   Input = 0x85
	TkMouseScroll Input = 0x86
	TkMouseX      Input = 0x87
	TkMouseY      Input = 0x88
	TkMousePixelX Input = 0x89
	TkMousePixelY Input = 0x8A
	TkMouseWheel  Input = 0x8B
	TkMouseClicks Input = 0x8C

	TkKeyReleased Input = 0x100

	TkWidth       Input = 0xC0
	TkHeight      Input = 0xC1
	TkCellWidth   Input = 0xC2
	TkCellHeight  Input = 0xC3
	TkColor       Input = 0xC4
	TkBkcolor     Input = 0xC5
	TkLayer       Input = 0xC6
	TkComposition Input = 0xC7
	TkChar        Input = 0xC8
	TkWchar       Input = 0xC9
	TkEvent       Input = 0xCA
	TkFullscreen  Input = 0xCB

	TkClose   Input = 0xE0
	TkResized Input = 0xE1

	TkInputNone      Input = 0
	TkInputCancelled Input = -1
)

// This should be called first. This initializes the library and
// readies the terminal with a size of 80x25 and white on black. This
// doesn't create a window, the first call to Refresh() will. Returns
// error on failure.
func Open() error {
	rv := C.terminal_open()
	if rv == 0 {
		return errors.New("failed to start bearlibterminal")
	}
	return nil
}

// Should be called when finished using the library.
func Close() {
	C.terminal_close()
}

// This function accepts a configuration string to configure various
// library options, like managing fonts, tilesets, etc. The format is
// semicolon separated. An example taken from the official docs:
//
// window.title='game';
// font: UbuntuMono-R.ttf, size=12;
// ini.settings.tile-size=16;
//
// See more at
// http://foo.wyrd.name/en:bearlibterminal:reference:configuration
func Set(opts string) error {
	str := C.CString(opts)
	rv := C.terminal_set(str)
	C.free(unsafe.Pointer(str))
	if rv == 0 {
		return errors.New("failed to set options")
	}
	return nil
}

// Sets the foreground to a 32-bit color of the form 0xaarrggbb. Can
// be queried using State(TkColor).
func Color(c uint32) {
	C.terminal_color(C.color_t(c))
}

// Sets the background to a 32-bit color of the form 0xaarrggbb. Can
// be querid using State(TkBkcolor).
func BkColor(c uint32) {
	C.terminal_bkcolor(C.color_t(c))
}

// Setting composition to true makes it so writing to a cell does not
// overwrite the cell's contents. Can be queried using
// State(TkComposition).
func Composition(q bool) {
	if q {
		C.terminal_composition(C.TK_ON)
	} else {
		C.terminal_composition(C.TK_OFF)
	}
}

//
func Layer(l uint8) {
	C.terminal_layer(C.int(l))
}

//
func Clear() {
	C.terminal_clear()
}

//
func ClearArea(x, y, w, h int) {
	C.terminal_clear_area(C.int(x), C.int(y), C.int(w), C.int(h))
}

//
func Crop(x, y, w, h int) {
	C.terminal_crop(C.int(x), C.int(y), C.int(w), C.int(h))
}

//
func Refresh() {
	C.terminal_refresh()
}

//
func Put(x, y int, c rune) {
	C.terminal_put(C.int(x), C.int(y), C.int(c))
}

//
func Pick(x, y int, l uint8) rune {
	return rune(C.terminal_pick(C.int(x), C.int(y), C.int(l)))
}

//
func PickColor(x, y int, l uint8) uint32 {
	return uint32(C.terminal_pick_color(C.int(x), C.int(y), C.int(l)))
}

//
func PickBkColor(x, y int) uint32 {
	return uint32(C.terminal_pick_bkcolor(C.int(x), C.int(y)))
}

//
func PutExt(x, y, dx, dy int, c rune, corners []uint32) {
	code := C.int(c)
	corn := (*C.color_t)(unsafe.Pointer(&corners[0]))
	C.terminal_put_ext(C.int(x), C.int(y), C.int(dx), C.int(dy), code, corn)
}

//
func Print(x, y int, s string) {
	str := C.CString(s)
	C.terminal_print(C.int(x), C.int(y), str)
	C.free(unsafe.Pointer(str))
}

//
func Measure(s string) {
	str := C.CString(s)
	C.terminal_measure(str)
	C.free(unsafe.Pointer(str))
}

//
func State(inp Input) int {
	return int(C.terminal_state(C.int(inp)))
}

//
func Check(inp Input) bool {
	if C.terminal_state(C.int(inp)) == 0 {
		return false
	} else {
		return true
	}
}

//
func HasInput() bool {
	if C.terminal_has_input() == 0 {
		return false
	} else {
		return true
	}
}

//
func Read() Input {
	return Input(C.terminal_read())
}

//
func Peek() Input {
	return Input(C.terminal_peek())
}

//
func ReadStr(x, y int, buffer []byte) int {
	buf := (*C.char)(unsafe.Pointer(&buffer[0]))
	l := C.int(len(buffer))
	rv := C.terminal_read_str(C.int(x), C.int(y), buf, l)
	return int(rv)
}

//
func Delay(ms int) {
	C.terminal_delay(C.int(ms))
}

//
func ColorFromName(name string) uint32 {
	str := C.CString(name)
	col := C.color_from_name(str)
	C.free(unsafe.Pointer(str))
	return uint32(col)
}

//
func ColorFromArgb(a, r, g, b uint8) uint32 {
	ca := C.uint8_t(a)
	cr := C.uint8_t(r)
	cg := C.uint8_t(g)
	cb := C.uint8_t(b)
	return uint32(C.color_from_argb(ca, cr, cg, cb))

}
