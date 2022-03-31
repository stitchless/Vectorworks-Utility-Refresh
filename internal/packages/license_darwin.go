package packages

import (
	"bufio"
	"bytes"
	"fmt"
	"howett.net/plist"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

type LicenseOpts struct {
	serial map[string]string `plist:"NNA User License"`
}

func GetSerialLocation(installation Installation) string {
	switch installation.ModuleName {
	case ModuleVectorworks:
		return app.HomeDirectory + "/Library/Preferences/net.nemetschek.vectorworks.license." + installation.Year + ".plist"
	case ModulesVision:
		return app.HomeDirectory + "/Library/Preferences/net.vectorworks.vision.license." + installation.Year + ".plist"
	}

	return ""
}

// getSerial will read in a plist, decode it and return a keyed value as a string value
func getSerial(installation Installation) (string, error) {
	serialLocation := GetSerialLocation(installation)

	// Read in plist
	plistFile, err := ioutil.ReadFile(serialLocation)
	if err != nil {
		return "", err
	}

	buffer := bytes.NewReader(plistFile)

	// parse and return plist serial
	var plistData LicenseOpts
	decoder := plist.NewDecoder(buffer)
	err = decoder.Decode(&plistData.serial)
	if err != nil {
		return "", err
	}

	return plistData.serial[`NNA User License`], nil
}

func ReplaceOldSerial(installation Installation, newSerial string) error {
	licenseLocation := GetSerialLocation(installation)
	plistFile, err := os.Open(licenseLocation)
	if err != nil {
		return err
	}

	err = plistFile.Truncate(0)
	if err != nil {
		return err
	}

	newSerial = cleanSerial(newSerial)

	plistData := &LicenseOpts{
		serial: map[string]string{"NNA User License": newSerial},
	}

	fmt.Println(plistData.serial)
	buffer := &bytes.Buffer{}
	encoder := plist.NewEncoder(buffer)

	err = encoder.Encode(plistData.serial)
	if err != nil {
		return err
	}

	err = os.WriteFile(licenseLocation, buffer.Bytes(), 0644)
	if err != nil {
		return err
	}

	w := bufio.NewWriter(buffer)
	n4, err := w.WriteString("buffered\n")
	if err != nil {
		return err
	}

	fmt.Printf("wrote %d bytes\n", n4)

	err = w.Flush()
	if err != nil {
		return err
	}

	err = refreshPList()
	if err != nil {
		return err
	}

	return nil
}

func refreshPList() error {
	fmt.Println("Refreshing plist files...")
	// osascript -e 'do shell script "sudo killall -u $USER cfprefsd" with administrator privileges'
	cmd := exec.Command(`osascript`, "-s", "h", "-e", `do shell script "sudo killall -u $USER cfprefsd" with administrator privileges`)
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	log.SetOutput(os.Stderr)

	if err = cmd.Start(); err != nil {
		log.Fatal(err)
	}

	cmdErrOutput, err := ioutil.ReadAll(stderr)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", cmdErrOutput)

	if err = cmd.Wait(); err != nil {
		return err
	}

	return nil
}
