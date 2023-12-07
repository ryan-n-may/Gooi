package main

import (
	event 		"gooi/base/event"
	comp 		"gooi/base/components"
	listeners 	"gooi/base/listeners"
	windows 	"gooi/base/windows"
	ompo   		"gooi/base/compositions"
	cons        "gooi/base/constants"
	colours     "gooi/base/colours"
	intf 		"gooi/interfaces"
)


func main() {
/** Creating Window **/
	var A = windows.NewWindow("Tabs Alignment Demo", 800, 500)
	A.OpenWindow()
	A.SetBackgroundColour(colours.BLUE)
/** Creating mouse listener and assinging it to the window **/
	var ML = listeners.CreateMouseHandler("Mouse Handler")
	A.SetMouseHandler(ML)
/** Creating event listener/handler and assigning it to the window **/
	var E = event.NewEventHandler()
	A.GetWindowCanvas().SetEventHandler(E)
/** Creating event listener **/
	var KL = listeners.CreateKeyListener("KeyListener", A.GetWindowCanvas())
	
	var columnComposition = ompo.NewColumnComposition(
		"MasterColumnComposition", A.WindowCanvas, A.WindowCanvas, 0, 0, 0, 1.0, 1.0, cons.ALIGN_LEFT)

	var titleCardComposition = createTitleCardArea(A.WindowCanvas, columnComposition, KL)
	//var tabsComposition = createTabArea(A.WindowCanvas, columnComposition, E, ML)


	columnComposition.AddDisplayable(titleCardComposition)
	//columnComposition.AddDisplayable(tabsComposition)

	A.GetWindowCanvas().AddDisplayable(columnComposition)

	A.RunWindow()
}

func createTitleCardArea(
	C intf.Canvas_Interface, 
	M intf.Displayable,
	KL *listeners.KeyHandler_Struct,
) intf.Displayable {
	var NewBoxComposition = ompo.NewBoxComposition(
		"generalBox", C, M,
		0, 0, 0, 1, 1,
		cons.ALIGN_TOP_LEFT,
		colours.WHITE)

	var rowComposition = ompo.NewRowComposition("titleCardRowComposition", C, NewBoxComposition, 0, 0, 0, 1.0, 1.0, cons.ALIGN_CENTRE_ROW)

	var label = comp.NewLabel(C, M, "Connection Name:", 0, 0, 0, "luxi", "base/components/fonts/luxi.ttf", 16)
	var textInput = comp.NewTextInput(C, M, "Connection Name", "Connection", KL, 200, 20, 0, 0, 0, 20, colours.LIGHT_GRAY, "luxi", "base/components/fonts/luxi.ttf", 16)

	NewBoxComposition.AddDisplayable(rowComposition)

	rowComposition.AddDisplayable(label)
	rowComposition.AddDisplayable(textInput)

	return NewBoxComposition
}

func createBottomBarArea(){

}

func createTabArea(
	C intf.Canvas_Interface, 
	M intf.Displayable,
	E intf.EventHandler_Interface,
	ML intf.MouseHandler_Interface,
) intf.Displayable {
	var tabAreaComposition = ompo.NewTabComposition(
		"Tabs",
		C,
		M,
		E, 
		ML,
		[]string{"General", "Wi-Fi" /*, "Security", "IPv4", "IPv6"*/},
		0, 0, 0, 1.0, 0.7,
		"luxi", "base/components/fonts/luxi.ttf", 16,
	)
	var generalBox = ompo.NewBoxComposition(
		"generalBox", C, tabAreaComposition,
		0, 0, 0, 1, 0.5,
		cons.ALIGN_TOP_LEFT,
		colours.LIGHT_GRAY)
	
	var generalColumnLayout = ompo.NewColumnComposition(
		"General", C, generalBox, 0, 0, 0, 1.0, 1.0, cons.ALIGN_LEFT)
	
	var connectionPriorityCheckBox = comp.CreateCheckbox(
		C, generalColumnLayout, "Connect Automatically with Priority",
		10, 0, 0, 0, &event.NULL_EVENT, 
		"luxi", "base/components/fonts/luxi.ttf", 16)
	var automaticallyConnectToVPN = comp.CreateCheckbox(
		C, generalColumnLayout, "Connect Automatically to VPN",
		10, 0, 0, 0, &event.NULL_EVENT, 
		"luxi", "base/components/fonts/luxi.ttf", 16)
	var allUsersMayConnectToThisNetwork = comp.CreateCheckbox(
		C, generalColumnLayout, "All Users May Connect",
		10, 0, 0, 0, &event.NULL_EVENT, 
		"luxi", "base/components/fonts/luxi.ttf", 16)
	var reconnectButton = comp.CreateButton(
		C, generalColumnLayout, "Reconnect", 
		200, 30, 10, 0, 0, 0, 
		"luxi", "base/components/fonts/luxi.ttf", 16,
		&event.NULL_EVENT, 200)

	generalColumnLayout.AddDisplayable(reconnectButton)
	generalColumnLayout.AddDisplayable(connectionPriorityCheckBox)
	generalColumnLayout.AddDisplayable(automaticallyConnectToVPN)
	generalColumnLayout.AddDisplayable(allUsersMayConnectToThisNetwork)
	generalBox.AddDisplayable(generalColumnLayout)

	var wifiBox = ompo.NewBoxComposition(
		"generalBox", C, tabAreaComposition,
		0, 0, 0, 1, 1,
		cons.ALIGN_TOP_LEFT,
		colours.RED)
	var wifiButton = comp.CreateButton(
		C, wifiBox, "wifiButton", 
		200, 30, 10, 0, 0, 0, 
		"luxi", "base/components/fonts/luxi.ttf", 16,
		&event.NULL_EVENT, 200)
	wifiBox.AddDisplayable(wifiButton)

	var securityBox = ompo.NewBoxComposition(
		"securityBox", C, tabAreaComposition,
		0, 0, 0, 1, 1,
		cons.ALIGN_CENTRE,
		colours.RED)
	var securityButton = comp.CreateButton(
		C, securityBox, "securityButton", 
		200, 30, 10, 0, 0, 0, 
		"luxi", "base/components/fonts/luxi.ttf", 16,
		&event.NULL_EVENT, 200)
	securityBox.AddDisplayable(securityButton)

	var ipv4Box = ompo.NewBoxComposition(
		"ipv4Box", C, tabAreaComposition,
		0, 0, 0, 1, 1,
		cons.ALIGN_CENTRE,
		colours.RED)
	var ipv4Button = comp.CreateButton(
		C, ipv4Box, "ipv4Button", 
		200, 30, 10, 0, 0, 0, 
		"luxi", "base/components/fonts/luxi.ttf", 16,
		&event.NULL_EVENT, 200)
	ipv4Box.AddDisplayable(ipv4Button)

	var ipv6Box = ompo.NewBoxComposition(
		"ipv6Box", C, tabAreaComposition,
		0, 0, 0, 1, 1,
		cons.ALIGN_CENTRE,
		colours.RED)
	var ipv6Button = comp.CreateButton(
		C, ipv6Box, "ipv6Button", 
		200, 30, 10, 0, 0, 0, 
		"luxi", "base/components/fonts/luxi.ttf", 16,
		&event.NULL_EVENT, 200)
	ipv6Box.AddDisplayable(ipv6Button)



	tabAreaComposition.AddDisplayable(wifiBox, 0)
	tabAreaComposition.AddDisplayable(generalBox, 1)

	return tabAreaComposition
}