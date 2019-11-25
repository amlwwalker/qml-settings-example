import QtQuick 2.7          //Rectangle
import QtQuick.Controls 2.1 //StackView
import QtQuick.Layouts 1.3  //GridLayout

ApplicationWindow {
  id: app
  visible: true
  title: "Pickle It"
  minimumWidth: 1200; minimumHeight: 768
  // visibility: "FullScreen"
  width: minimumWidth; height: minimumHeight
  ColumnLayout {
      id: mainLayout
      // anchors.fill: parent
      anchors.right: parent.right
      anchors.left: parent.left
      GroupBox {
          id: gridBox
          Layout.fillWidth: true
          GridLayout {
              id: gridLayout
              rows: 5
              flow: GridLayout.TopToBottom
              anchors.fill: parent

              Label { text: "Happy to send back statistics about how the app is performing"; Layout.fillWidth: true  }
              Label { text: "Save a screen shot of the file when creating versions"; Layout.fillWidth: true  }
              Label { text: "Notify when a new file is ready"; Layout.fillWidth: true  }
              Label { text: "Notify when a patch is created"; Layout.fillWidth: true  }
              Label { text: "Override existing files (or save patches next to them"; Layout.fillWidth: true  }

              CheckBox { id: statisticsCheck; checked: systemSettings.sendStatistics}
              CheckBox { id: screenshotCheck; checked: systemSettings.saveScreenshots}
              CheckBox { id: newFileReadySystemNotifyCheck; checked: systemSettings.systemNewFileNotification}
              CheckBox { id: patchSystemNotifyCheck; checked: systemSettings.systemNewPatchNotification}
              CheckBox { id: overrideExistingCheck; checked: systemSettings.overrideMasterFile}
          }
      }
  }
}
