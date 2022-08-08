package qrcode

import (
	// import built-in packages
	"fmt"
	// import third-party packages
	"github.com/mattn/go-colorable"
	"github.com/skip2/go-qrcode"
)

func QrcodeTerminal(text string) {
	new().get(text).Print()
}

type consoleColors struct {
	BrightBlack string
	BrightWhite string
}
type qrcodeRecoveryLevel qrcode.RecoveryLevel
type qrcodeRecoveryLevels struct {
	Medium qrcodeRecoveryLevel
}

var (
	ConsoleColors consoleColors = consoleColors{
		BrightBlack: "\033[48;5;0m  \033[0m",
		BrightWhite: "\033[48;5;7m  \033[0m"}
	QRCodeRecoveryLevels = qrcodeRecoveryLevels{
		Medium: qrcodeRecoveryLevel(qrcode.Medium)}
)

type QRCodeString string

func (v *QRCodeString) Print() {
	fmt.Fprint(outer, *v)
}

type qrcodeTerminal struct {
	front string
	back  string
	level qrcodeRecoveryLevel
}

func (v *qrcodeTerminal) get(content interface{}) (result *QRCodeString) {
	var qr *qrcode.QRCode
	var err error
	if t, ok := content.(string); ok {
		qr, err = qrcode.New(t, qrcode.RecoveryLevel(v.level))
	} else if t, ok := content.([]byte); ok {
		qr, err = qrcode.New(string(t), qrcode.RecoveryLevel(v.level))
	}
	if qr != nil && err == nil {
		data := qr.Bitmap()
		result = v.getQRCodeString(data)
	}
	return
}

func new() *qrcodeTerminal {
	front, back, level := ConsoleColors.BrightBlack, ConsoleColors.BrightWhite, QRCodeRecoveryLevels.Medium
	obj := qrcodeTerminal{
		front: front,
		back:  back,
		level: level}
	return &obj
}

func (v *qrcodeTerminal) getQRCodeString(data [][]bool) (result *QRCodeString) {
	str := ""
	for ir, row := range data {
		lr := len(row)
		if ir == 0 || ir == 1 || ir == 2 ||
			ir == lr-1 || ir == lr-2 || ir == lr-3 {
			continue
		}
		for ic, col := range row {
			lc := len(data)
			if ic == 0 || ic == 1 || ic == 2 ||
				ic == lc-1 || ic == lc-2 || ic == lc-3 {
				continue
			}
			if col {
				str += fmt.Sprint(v.front)
			} else {
				str += fmt.Sprint(v.back)
			}
		}
		str += fmt.Sprintln()
	}
	obj := QRCodeString(str)
	result = &obj
	return
}

var outer = colorable.NewColorableStdout()
