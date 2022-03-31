package packages

type Installation struct {
	ID int
	ModuleName
	Year         string
	License      License
	Properties   []string
	RMCache      string
	LogFile      string
	LogFileSent  string
	Directories  []string
	CleanOptions InstallationOptions
}

type InstallationOptions struct {
	RemoveRMC               bool
	RemoveUserData          bool
	RemoveUserSettings      bool
	RemoveInstallerSettings bool
	RemoveAllData           bool
}

//
//var AllInstalledSoftwareMap = make(map[ModuleName][]Installation)
//
//// Generate installation map on launch.
//func Refresh() {
//	err := GenerateInstalledSoftwareMap()
//	if err != nil {
//		fmt.Errorf("could not generate installation map for %s: %v", AllActiveSoftwareNames, err)
//	}
//}
//
//func (i *Installation) Refresh() string {
//	return ""
//}
//
//// GenerateInstalledSoftwareMap creates a map
//// key: ModuleName
//// Value: [] Installation
//func GenerateInstalledSoftwareMap() error {
//	for _, softwareName := range AllActiveSoftwareNames {
//		installations, err := FindInstallationsBySoftware(softwareName)
//		if err != nil {
//			return fmt.Errorf("error: packages search failed - %v", err)
//		}
//		if len(installations) != 0 {
//			AllInstalledSoftwareMap[softwareName] = installations
//		}
//	}
//	return nil
//}
//
//func getActivation(serial []string) string {
//	out, OK := licenseActivationMap[serial[0]]
//	if OK {
//		return out
//	}
//	return "Activation not found"
//}
//
//func getPlatform(serial []string) string {
//	out, OK := licensePlatformMap[serial[2]]
//	if OK {
//		return out
//	}
//	return "Platform not found"
//}
//
//func getLocal(serial []string) string {
//	local := strings.Join(serial[3:5], "")
//	out, OK := licenseLocalMap[local]
//	if OK {
//		return out
//	}
//	return "Local not found"
//}
//
//func getType(serial []string) string {
//	out, OK := licenseTypeMap[serial[5]]
//	if OK {
//		return out
//	}
//	return "License type not found"
//}
