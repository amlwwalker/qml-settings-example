package main

import (
	"fmt"
	"sync"

	"github.com/therecipe/qt/core"
)

var once sync.Once
var instance *Controller

type SystemSettings struct {
	core.QObject

	// to save the object to bolt/storm, because of the above
	// core.QObject, should i give the below each a name and then add a tag such as
	// `storm:"id,increment"` and see if I can ignore core.QObject?
	_ bool `property:"sendStatistics"`
	_ bool `property:"saveScreenshots"`
	_ bool `property:"systemNewFileNotification"`
	_ bool `property:"systemNewPatchNotification"`
	_ bool `property:"overrideMasterFile"`
}
type Controller struct {
	core.QObject
	UserSettings *SystemSettings
}

func ControllerInstance() *Controller {
	once.Do(func() {
		instance = NewController(nil)
	})
	return instance
}

func (c *Controller) InitWith() {
	c.UserSettings = NewSystemSettings(nil)

	c.UserSettings.SetSendStatistics(true)
	c.UserSettings.ConnectSendStatisticsChanged(func(setting bool) {
		fmt.Printf("settings 1%+v\r\n", c.UserSettings)
		//now i want to send the properties/settings and store them in a database (boltdb.)
	})

	c.UserSettings.SetSaveScreenshots(true)
	c.UserSettings.SetSystemNewFileNotification(false)
	c.UserSettings.SetSystemNewPatchNotification(false)
	c.UserSettings.SetOverrideMasterFile(true)
}
