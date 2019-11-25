package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/widgets"
)

func main() {

	// configure UI
	var path string
	if false {
		path = "qrc:/qml/view.qml"
	} else {
		path = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "amlwwalker", "example-qml", "settings-example", "qml", "main.qml")
	}
	core.QCoreApplication_SetOrganizationName("Settings Example") //needed to fix an QML Settings issue on windows
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	widgets.NewQApplication(len(os.Args), os.Args)
	engine := qml.NewQQmlApplicationEngine(nil)
	ControllerInstance().InitWith()

	engine.AddImportPath("qrc:/qml/")
	if true {
		engine.Load(core.NewQUrl3(path, 0))
	} else {
		engine.Load(core.NewQUrl3("qrc:/qml/main.qml", 0))
	}

	view := gui.NewQWindowFromPointer(engine.RootObjects()[0].Pointer()) //get a handle on the view as a OS object
	fmt.Printf("userSettings %+v\r\n", ControllerInstance().UserSettings)
	engine.RootContext().SetContextProperty("systemSettings", ControllerInstance().UserSettings)

	view.Show()
	widgets.QApplication_Exec()
}
