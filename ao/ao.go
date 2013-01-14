package ao

/*
#include <ao/ao.h>
#cgo LDFLAGS: -lao 
*/
import "C"

import (
	"fmt"
	"unsafe"
)

const (
	FMT_LITTLE = int(C.AO_FMT_LITTLE)
	FMT_BIG    = int(C.AO_FMT_BIG)
	FMT_NATIVE = int(C.AO_FMT_NATIVE)
)

type Format struct {
	Bits int
	Rate int
	Channels int
	Byte_format int
	Matrix string
}

type Player struct {
	handle *C.ao_device
	format Format
}

func Initialize() {
	C.ao_initialize();
}

func DefaultDriverId() (int, error) {
	id := int(C.ao_default_driver_id())
	if id < 0 {
		return 0, fmt.Errorf("no sound device found")
	}
	return id, nil
}


func NewPlayer(driver int, f Format, options map[string]string) (*Player, error) {
	var cf C.ao_sample_format
	cf.bits = C.int(f.Bits)
	cf.rate = C.int(f.Rate)
	cf.channels = C.int(f.Channels)
	cf.byte_format = C.int(f.Byte_format)
	if f.Matrix == "" {
		cf.matrix = nil
	} else {
		// this string is 
		cf.matrix = C.CString(f.Matrix)
	}
	// FIXME: handle options (currently ignored)
	devp := C.ao_open_live(C.int(driver), &cf, nil)
	if cf.matrix != nil {
		C.free(unsafe.Pointer(cf.matrix))
	}
	if devp == nil {
		return nil, fmt.Errorf("could not initialize ao device")
	}
	p := new(Player)
	p.handle = devp
	p.format = f
	return p, nil
}

// FIXME: handle errors
func (p *Player) Play(buf []byte) {
	C.ao_play(p.handle, (*C.char)(unsafe.Pointer(&buf[0])), C.uint_32(len(buf)))
}

// FIXME: handle errors
func (p *Player) Close() {
	C.ao_close(p.handle)
}
