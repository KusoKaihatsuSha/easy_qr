package main

// #include <stdio.h>
// #include <stdlib.h>
//
// static char* myprint(char* s) {
//   return s;
// }
import "C"
import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"

	"github.com/skip2/go-qrcode"
	. "github.com/ulvham/helper"
)

var p = fmt.Println

//export PrintQr
func PrintQr(b *C.char, l C.int) *C.char {
	// build args
	// -o qrcode.dll -buildmode=c-shared
	p(true)
	text := C.GoString(b)
	text_, _ := base64.StdEncoding.DecodeString(text)
	Dbg(text_)
	level := l

	p("return in GO dll", string(text_), l)

	var level_ qrcode.RecoveryLevel
	switch level {
	case 0:
		level_ = qrcode.Low
	case 1:
		level_ = qrcode.Medium
	case 2:
		level_ = qrcode.High
	case 3:
		level_ = qrcode.Highest
	default:
		level_ = qrcode.Low
	}

	qr_element, err := qrcode.New(string(text_), level_)
	if err != nil {
		p(err)
	}

	rand.Seed(time.Now().UnixNano())

	png_, _ := qr_element.PNG(400)
	se := base64.StdEncoding.EncodeToString(png_)

	return C.CString(se)
}

func main() {
	//p(nil);
}
