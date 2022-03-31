package ui

import (
	"github.com/inkyblackness/imgui-go/v4"
	"github.com/jpeizer/Vectorworks-Utility-Refresh/internal/packages"
)

var application = packages.GetApplication()

// featureName provides the user readable string for a supported packages package
type featureName string

// hard coded feature names that are possible for all implemented packages packages
const (
	featureSoftware         featureName = "Software"
	//featureTraceApplication featureName = "Trace Application"
	featureDemoWindow       featureName = "Demo Window"
	featureSettings         featureName = "Settings"
)

// AllActiveFeatures is a list of all the currently supported features the application supports
var AllActiveFeatures = []featureName{
	featureSoftware,
	//featureTraceApplication,
	featureDemoWindow,
	//featureSettings,
}

// CurrentFeature is used to control the flow of actively rendered features
var CurrentFeature featureName

// String returns the string representation of the feature name
func (f featureName) String() string {
	return string(f)
}

// IsActive returns true if the feature is currently active
func (f featureName) IsActive() bool {
	return CurrentFeature == f
}

func (f featureName) SetActive() {
	CurrentFeature = f
}

func (f featureName) Render() {
	switch f {
	case featureSoftware:
		RenderSoftware()
	//case featureTraceApplication:
	//	RenderTraceApplication()
	case featureDemoWindow:
		open := true
		imgui.ShowDemoWindow(&open)
	//case featureSettings:
	//	RenderSettings()
	default:
		RenderSoftware()
	}
}
