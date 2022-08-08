package server

import (
	// import built-in packages
	"fmt"
	// import local packages
	"github.com/impossible98/gohttpd/global"
	"github.com/impossible98/gohttpd/pkg/color"
	"github.com/impossible98/gohttpd/pkg/ip"
	"github.com/impossible98/gohttpd/pkg/qrcode"
)

func RunTip(port string) {
	fmt.Println()
	fmt.Println("  " + color.Blue(global.APP_NAME+" v"+global.VERSION) + " " + color.Green("server running at:"))
	fmt.Println()
	fmt.Println("  > Local:   " + color.Blue("http://localhost:"+port+"/"))
	fmt.Println("  > Network: " + color.Blue("http://"+ip.LocalIP()+":"+port+"/"))
	fmt.Println()
	fmt.Println(color.Green("  You can login by scanning the qrcode:"))
	fmt.Println()
	qrcode.QrcodeTerminal("http://" + ip.LocalIP() + ":" + port + "/")
	fmt.Println()
}
