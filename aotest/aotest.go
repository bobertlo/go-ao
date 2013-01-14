package main

/* aotest.go
 * 
 * This file is based on the example ao_example.c which was
 * written by Stand Seibert - July 2001 and realeased into 
 * the public domain. This file is also in the public domain.
 *
 */

import (
	"math"
	"github.com/bobertlo/go-ao/ao"
)

func main() {
	ao.Initialize()

	driver, err := ao.DefaultDriverId()
	if err != nil {
		panic("could not find audio driver!")
	}

	f := ao.Format{}
	f.Bits = 16
	f.Rate = 44100
	f.Channels = 2
	f.Byte_format = ao.FMT_LITTLE
	player, err := ao.NewPlayer(driver, f, nil)
	if err != nil {
		panic("could not open audio device")
	}

	buffer := make([]byte, f.Bits/8 * f.Channels * f.Rate)
	for i := 0; i < f.Rate; i++ {
		sample := int(0.75 * 32768.0 * math.Sin(2.0 * 3.14159 * 440.0 * float64(i) / float64(f.Rate)))
		buffer[4*i] = byte(sample & 0xff)
		buffer[4*i+2] = buffer[4*i]
		buffer[4*i+1] = byte((sample >> 8) & 0xff)
		buffer[4*i+3] = buffer[4*i+1]
	}
	player.Play(buffer)

	player.Close()
}
