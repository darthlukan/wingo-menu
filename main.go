package main

import (
	"fmt"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/ewmh"
	"github.com/BurntSushi/xgbutil/mousebind"
	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/BurntSushi/xgbutil/xwindow"
	"log"
)

type Menu struct {
	BgColor, FgColor, Height, Width, Display int
	ExtType, States                          []string
}

func NewMenu(X *xgbutil.XUtil) {
	menu := new(Menu)

	window, err := xwindow.Generate(X)
	if err != nil {
		log.Fatal(err)
	}

	menu.BgColor = 0x191919
	menu.FgColor = 0xd3d3d3
	menu.Height = 100
	menu.Width = 100
	menu.Display = 0
	menu.ExtType = []string{"_NET_WM_WINDOW_TYPE_MENU"}
	states := []string{"_NET_WM_STATE_SKIP_PAGER", "_NET_WM_STATE_SKIP_TASKBAR"}
	menu.States = states

	fmt.Printf("menu: %v\n", menu)
	window.Create(X.RootWin(), 100, 100, menu.Width, menu.Height,
		xproto.CwBackPixel|xproto.CwEventMask, uint32(menu.BgColor),
		xproto.EventMaskButtonRelease)

	ewmh.WmWindowTypeSet(X, window.Id, menu.ExtType)
	ewmh.WmStateSet(X, window.Id, menu.States)
	window.Map()
}

func main() {
	fmt.Printf("Hello World!\n")

	X, err := xgbutil.NewConn()
	if err != nil {
		log.Fatal(err)
	}

	mousebind.Initialize(X)
	NewMenu(X)
	xevent.Main(X)
}
