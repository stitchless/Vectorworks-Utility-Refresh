package packages

import (
	"fmt"
	"github.com/jpeizer/Vectorworks-Utility-Refresh/internal/utils"
	"strings"
)

type SoftwareModule = string

type Software struct {
	Name          SoftwareModule
	Installations []Installation
}

type Application struct {
	SoftwarePackages map[string]Software
	HomeDirectory    string
}

const (
	ModuleVectorworks    SoftwareModule = "Vectorworks"
	ModulesVision        SoftwareModule = "Vision"
	ModulesCloudServices SoftwareModule = "VCS"
)

var (
	Vectorworks   = Software{ModuleVectorworks, []Installation{}}
	Vision        = Software{ModulesVision, []Installation{}}
	CloudServices = Software{ModulesCloudServices, []Installation{}}
	AllSoftware   = []Software{Vectorworks, Vision, CloudServices}
)

type application *Application

var app application

func (s *Software) GetInstallation(year string) *Installation {
	for _, installation := range s.Installations {
		if installation.Year == year {
			return &installation
		}
	}
	return nil
}

func (s *Software) Refresh() {

	years, err := FindInstallationYears(s.Name)
	if err != nil {
		return
	}

	for i, year := range years {
		installation := Installation{
			ID:             i,
			Year:           year,
			SoftwareModule: s.Name,
		}

		serial, err := getSerial(installation)
		if err != nil {
			continue
		}

		installation.setProperties()
		installation.setUserData()
		installation.setRMCache()
		installation.setLogFileSent()
		installation.setLogFile()

		serialStart := strings.Split(serial[0:6], "")
		installation.License = License{
			Serial:     serial,
			Local:      getLocal(serialStart),
			Platform:   getPlatform(serialStart),
			Activation: getActivation(serialStart),
			Type:       getType(serialStart),
		}

		s.Installations = append(s.Installations, installation)
	}
}

func GetApplication() *Application {
	app = &Application{
		SoftwarePackages: map[string]Software{},
		HomeDirectory:    utils.GetHomeDirectory(),
	}

	for _, software := range AllSoftware {
		software.Refresh()
		if len(software.Installations) > 0 {
			app.SoftwarePackages[software.Name] = software
		}
	}
	fmt.Printf("%+v\n", app)
	return app
}
