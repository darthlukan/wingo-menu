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
	bgColor, fgColor, height, width, display int
	extType, states                          []string
}

func (m *Menu) SetBgColor(colorCode int) {
	m.bgColor = colorCode
	return
}

func (m *Menu) GetBgColor() int {
	bgColor := m.bgColor
	return bgColor
}

func (m *Menu) SetFgColor(colorCode int) {
	m.fgColor = colorCode
	return
}

func (m *Menu) GetFgColor() int {
	fgColor := m.fgColor
	return fgColor
}

func (m *Menu) SetHeight(height int) {
	m.height = height
	return
}

func (m *Menu) GetHeight() int {
	height := m.height
	return height
}

func (m *Menu) SetWidth(width int) {
	m.width = width
	return
}

func (m *Menu) GetWidth() int {
	width := m.width
	return width
}

func (m *Menu) SetDisplay(display int) {
	m.display = display
	return
}

func (m *Menu) GetDisplay() int {
	display := m.display
	return display
}

func (m *Menu) SetExtType(extType []string) {
	m.extType = extType
	return
}

func (m *Menu) GetExtType() []string {
	extType := m.extType
	return extType
}

func (m *Menu) SetStates(states []string) {
	m.states = states
	return
}

func (m *Menu) GetStates() []string {
	states := m.states
	return states
}

func NewMenu(X *xgbutil.XUtil) {
	menu := new(Menu)

	window, err := xwindow.Generate(X)
	if err != nil {
		log.Fatal(err)
	}

	menu.SetBgColor(0x191919)
	menu.SetFgColor(0xd3d3d3)
	menu.SetHeight(100)
	menu.SetWidth(100)
	menu.SetDisplay(0)
	menu.SetExtType([]string{"_NET_WM_WINDOW_TYPE_DOCK"})
	states := []string{"_NET_WM_STATE_SKIP_PAGER", "_NET_WM_STATE_SKIP_TASKBAR", "_NET_WM_STATE_MODAL"}
	menu.SetStates(states)

	fmt.Printf("menu: %v\n", menu)
	window.Create(X.RootWin(), 100, 100, menu.GetWidth(), menu.GetHeight(),
		xproto.CwBackPixel|xproto.CwEventMask, uint32(menu.GetBgColor()),
		xproto.EventMaskButtonRelease)

	ewmh.WmWindowTypeSet(X, window.Id, menu.GetExtType())
	ewmh.WmStateSet(X, window.Id, menu.GetStates())
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
