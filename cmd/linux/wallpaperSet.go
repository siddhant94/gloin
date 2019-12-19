package linux

import (
	"fmt"
	"os/exec"
	"strings"
)

type Gsettings struct {
	Base     string
	BgSchema string
	Key      string
}

const backgroundSchema = "org.gnome.desktop.background"
const baseCommand = "gsettings"
const key = "picture-uri"
const prepend = "file://"

var gSettings Gsettings

func (gsettings Gsettings) GetCmd() (string, []string) {
	// return gsettings.Base + " get " + gsettings.BgSchema + " " + gsettings.Key
	return gsettings.Base, []string{"get", gsettings.BgSchema, gsettings.Key}
}

func (gsettings Gsettings) SetCmd(value string) (string, []string) {
	// return gsettings.Base + " set " + gsettings.BgSchema + " " + gsettings.Key + " " + prepend + value
	return gsettings.Base, []string{"set", gsettings.BgSchema, gsettings.Key, (prepend + value)}
}

func init() {
	gSettings = Gsettings{
		Base:     baseCommand,
		BgSchema: backgroundSchema,
		Key:      key,
	}
}

func SetWallpaper(imageDir string, latestImage string) {
	name, osCmd := gSettings.GetCmd()
	op, err := exec.Command(name, osCmd...).CombinedOutput()
	if err != nil {
		fmt.Printf("Error getting current wallpaper : %v\n %v\n", err, string(op))
		return
	}
	currentWallpaper := string(op)
	// Remove `prepend` i.e. `file://`
	currentWallpaper = strings.Replace(currentWallpaper, prepend, "", 1)
	fileUri := imageDir + "/" + latestImage
	name, osCmd = gSettings.SetCmd(fileUri)
	op, err = exec.Command(name, osCmd...).CombinedOutput()
	if err != nil {
		fmt.Printf("Error setting wallpaper (%v) : %v\n %v\n", fileUri, err, string(op))
		return
	}
	fmt.Println("Successfully changed desktop background")
}
