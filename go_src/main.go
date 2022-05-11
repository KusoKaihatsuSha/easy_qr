package main

import "C"
import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/skip2/go-qrcode"
	"golang.design/x/clipboard"
)

var p = fmt.Println

//export PrintQr
func PrintQr(b *C.char, level C.int) *C.char {
	return C.CString(PrintFromBase64toBase64(C.GoString(b), int(level), true))
}

//export Rundll32ToClipboard
func Rundll32ToClipboard() {
	text := os.Args[3]
	level, err := strconv.Atoi(os.Args[4])
	if err != nil {
		p(err)
	}
	clipboard.Write(clipboard.FmtImage, Print(text, level))
}

//export Rundll32save
func Rundll32save() {
	text := os.Args[3]
	level, err := strconv.Atoi(os.Args[4])
	if err != nil {
		p(err)
	}
	ioutil.WriteFile(os.Args[5], Print(text, level), 644)
}

//export Print
func Print(text string, level int) []byte {
	err := clipboard.Init()
	if err != nil {
		p(err)
	}
	var lvl qrcode.RecoveryLevel
	switch level {
	case 0:
		lvl = qrcode.Low
	case 1:
		lvl = qrcode.Medium
	case 2:
		lvl = qrcode.High
	case 3:
		lvl = qrcode.Highest
	default:
		lvl = qrcode.Low
	}
	qr, err := qrcode.New(string(text), lvl)
	if err != nil {
		p(err)
	}
	rand.Seed(time.Now().UnixNano())
	png, err := qr.PNG(400)
	if err != nil {
		p(err)
	}
	return png
}

//export PrintFromBase64toBase64
func PrintFromBase64toBase64(textBase64 string, level int, flagClipboard bool) string {
	text, err := base64.StdEncoding.DecodeString(textBase64)
	if err != nil {
		p(err)
	}
	png := Print(string(text), level)
	if flagClipboard {
		clipboard.Write(clipboard.FmtImage, png)
	}
	return base64.StdEncoding.EncodeToString(png)
}

func main() {
	_ = ""
}
