package kde

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"os/user"
	"path"
	"path/filepath"
	"time"
)

func ChangeWallpaper(imageData *[]byte) {
	localWallpaperPath := persistWallpaper(imageData)
	setWallpaper(localWallpaperPath)
}

func setWallpaper(path string) {
	changeWallpaperCommand :=
		"dbus-send --session --dest=org.kde.plasmashell --type=method_call /PlasmaShell org.kde.PlasmaShell.evaluateScript 'string:" +
			"var Desktops = desktops();" +
			"for (i=0;i<Desktops.length;i++) {" +
			"        d = Desktops[i];" +
			"        d.wallpaperPlugin = \"org.kde.image\";" +
			"        d.currentConfigGroup = Array(\"Wallpaper\"," +
			"                                    \"org.kde.image\"," +
			"                                    \"General\");" +
			"        d.writeConfig(\"Image\", \"file://" + path + "\");" +
			"        d.writeConfig(\"FillMode\", 1);" +
			"}'"
	executeCommand(changeWallpaperCommand)
}

func persistWallpaper(data *[]byte) string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	wallpaperDirectory := path.Join(usr.HomeDir, ".wallpaper")
	localWallpaperPath := path.Join(wallpaperDirectory, randomName()+".png")
	removeContents(wallpaperDirectory)

	err = ioutil.WriteFile(localWallpaperPath, *data, 0644)

	if err != nil {
		log.Fatal(err)
	}

	return localWallpaperPath
}

func randomName() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 15)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func executeCommand(cmd string) string {
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

func removeContents(dir string) {
	d, err := os.Open(dir)
	if err != nil {
		log.Fatal(err)
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		log.Fatal(err)
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			log.Fatal(err)
		}
	}
}
