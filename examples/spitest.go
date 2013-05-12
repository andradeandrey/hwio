// SPI example
//
// Connect SPI MISO to MOSI as a loopback, and send data to yourself ;-)
//
// This should work on a BeagleBone; it uses SPI0 which is on GPIO0_5.

package main

import (
	"fmt"
	"github.com/mrmorphic/hwio"
	"hwio"
	"math/rand"
	"os"
	"time"
)

func main() {
	// get an SPI clock pin by name. Convention for SPI is that you assign the clock pin, and the driver will assign all pins for the
	// module (i.e. SCLK, MISO, MOSI, CE) so that you can't then allocate them for GPIO and get unpredictable result.
	spi0, err := hwio.GetPin("GPIO0_5")

	// Generally we wouldn't expect an error, but perhaps someone is running this a
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Set the mode of the pin to output. This will return an error if, for example,
	// we were trying to set an analog input to an output.
	err = hwio.PinMode(spi0, hwio.SPI_CLOCK)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	nErrors := 0
	rand.Seed(time.Now().UTC().UnixNano())
	for {
		// get a random byte
		randomByte := rand.Intn(255)

		// send the byte
		e := hwio.SPIWrite(spi0, randomByte, 8, 0)
		if e != nil {
			fmt.Printf("Error writing SPI value (%d): %s", randomByte, e)
			break
		}

		// there should be a byte available
		// read that byte and compare with what was sent
		valueRead, e := SPIRead(spi)

		if valueRead != uint(randomByte) {
			fmt.Printf("Data didn't match; written=%d, read=%d\n", randomByte, valueRead)
			nErrors++
			if nErrors > 100 {
				fmt.Printf("Too many errors")
				break
			}
		}
	}
}
