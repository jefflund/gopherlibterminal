package gopherlibterminal

// #cgo CFLAGS:
// #cgo amd64 386 CFLAGS: -DX86=1
// #cgo LDFLAGS: -lBearLibTerminal -Wl,-rpath .
// #include <stdlib.h>
// #include <BearLibTerminal.h>
import "C"
import "errors"
import "unsafe"

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

}

func Print(x, y int, s string) {
	str := C.CString(s)
	C.terminal_print(C.int(x), C.int(y), str)
	C.free(unsafe.Pointer(str))
}

func Measure() {

}

func State() {

}

func Check() {

}

func HasInput() bool {
	if C.terminal_has_input() == 0 {
		return false
	} else {
		return true
	}
}

func Read() {

}

func Peek() {

}

func ReadStr() {

}

func Delay(ms int) {
	C.terminal_delay(C.int(ms))
}

func ColorFromName() {

}

func ColorFromArgb() {

}