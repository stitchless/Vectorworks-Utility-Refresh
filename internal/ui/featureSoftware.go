package ui

import (
	"fmt"
	"github.com/inkyblackness/imgui-go/v4"
	"github.com/jpeizer/Vectorworks-Utility-Refresh/internal/packages"
	"github.com/jpeizer/Vectorworks-Utility-Refresh/internal/utils"
	"strings"
)

var (
	toggleSerialDetails bool
)

// RenderSoftware shows serials of found supported packages
func RenderSoftware() {
	// Start of packages tab bar
	imgui.BeginTabBar("##SoftwareTabBar")
	// Run for all active supported packages
	for _, swPkg := range application.SoftwarePackages {
		// Test for installations of active packages prior to making a table
		if len(swPkg.Installations) == 0 {
			continue
		}
		// Insert new tab for each installed supported packages
		imgui.PushID(swPkg.Name)
		if imgui.BeginTabItem(swPkg.Name) {
			// Begin of packages year tab bar
			imgui.BeginTabBar("TabBar")
			// Find all installed packages versions
			for _, installation := range swPkg.Installations {
				// Insert a new tab for all packages versions found
				imgui.PushID(installation.Year)
				if imgui.BeginTabItem(installation.Year) {
					// ----------------------------
					// LAYOUT FOR SOFTWARE FEATURES
					// ----------------------------
					// Software serial label
					imgui.Dummy(imgui.Vec2{X: imgui.ContentRegionAvail().X, Y: 5})
					//imgui.PushFont()
					imgui.PushItemWidth(350)
					// Flags 2 InputTextFlagsCharsUppercase | 4 InputTextFlagsAutoSelectAll | InputTextFlagsEnterReturnsTrue
					if imgui.InputTextV("##EditedSerial", &installation.License.Serial, 1<<2|1<<4|1<<5, nil) {
						err := packages.ReplaceOldSerial(*installation, installation.License.Serial)
						if err != nil {
							_ = fmt.Errorf("error replacing old serial: %s", err)
						}
						//err := packages.GenerateInstalledSoftwareMap()
						//if err != nil {
						//	fmt.Errorf("error updating internal installation data after serial update %v", err)
						//}
					}
					imgui.PopItemWidth()
					//imgui.PopFont()
					if imgui.IsItemHovered() {
						// Introduce timer for the tooltips
						// https://gist.github.com/toutougabi/f56309cb9f802f34eeddda65eb27cad2
						imgui.SetTooltip("Insert new serial and press enter to update")
					}

					// Installation Info Button
					imgui.SameLine()
					imgui.PushFont(utils.FontAwesomeLight)
					if imgui.ButtonV("\uF05A Info", imgui.Vec2{X: 70, Y: 25}) {
						toggleSerialDetails = !toggleSerialDetails
					}

					// License Cleanup Button
					if imgui.BeginPopupModalV("Clean Software", nil, imgui.WindowFlagsAlwaysAutoResize) {
						imgui.Text("Select the options below to clean the packages.")

						imgui.Separator()
						imgui.Dummy(imgui.Vec2{X: 0, Y: 5})

						imgui.Checkbox("Remove resource manager cache", &installation.CleanOptions.RemoveRMC)
						imgui.Checkbox("Remove user folder data", &installation.CleanOptions.RemoveUserData)
						imgui.Checkbox("Remove user settings", &installation.CleanOptions.RemoveUserSettings)
						imgui.Checkbox("Remove installer files", &installation.CleanOptions.RemoveInstallerSettings)
						imgui.Checkbox("Remove all user data", &installation.CleanOptions.RemoveAllData)

						imgui.Dummy(imgui.Vec2{X: 0, Y: 5})
						imgui.Separator()
						imgui.Dummy(imgui.Vec2{X: 0, Y: 5})

						if imgui.Button("Remove") {
							err := installation.Clean()
							if err != nil {
								_ = fmt.Errorf("error cleaning installation: %s", err)
							}
							imgui.CloseCurrentPopup()
							application.Refresh()
						}
						imgui.SameLine()
						if imgui.Button("Cancel") {
							imgui.CloseCurrentPopup()
						}
						imgui.EndPopup()
					}

					imgui.SameLine()
					if imgui.ButtonV("\uF12D Clean"+"##CleanDialog", imgui.Vec2{X: 80, Y: 25}) {
						imgui.OpenPopup("Clean Software")
					}
					if imgui.IsItemHovered() {
						imgui.SetTooltip("Clean up this installation of packages")
					}

					imgui.PopFont()

					// Show License Tags
					if toggleSerialDetails {
						imgui.Text("Platform:")
						imgui.SameLine()
						imgui.PushFont(utils.FontRobotoBold)
						imgui.Text(installation.License.Platform)
						imgui.PopFont()

						imgui.SameLine()

						imgui.Text("     Region:")
						imgui.SameLine()
						imgui.PushFont(utils.FontRobotoBold)
						imgui.Text(installation.License.Local)
						imgui.PopFont()

						imgui.SameLine()

						imgui.Text("     Activation Type:")
						imgui.SameLine()
						imgui.PushFont(utils.FontRobotoBold)
						imgui.Text(installation.License.Activation)
						imgui.PopFont()

						imgui.SameLine()

						imgui.Text("     License Type:")
						imgui.SameLine()
						imgui.PushFont(utils.FontRobotoBold)
						imgui.Text(installation.License.Type)
						imgui.PopFont()
					}

					// Add spacer between serial and content body
					imgui.Dummy(imgui.Vec2{X: imgui.ContentRegionAvail().X, Y: 2})

					// ----------------------------
					// LAYOUT FOR SOFTWARE FEATURES
					// ----------------------------
					imgui.BeginGroup()
					imgui.PushFont(utils.FontRobotoBold)
					imgui.Text("Output")
					imgui.PopFont()
					imgui.BeginChildV("##optionsChild", imgui.Vec2{X: 0, Y: 0}, true, 0)
					// Text field for any output.
					if len(application.SoftwareOutputString) == 0 {
						imgui.Text("No output data yet...")
					}
					imgui.Text(strings.Join(application.SoftwareOutputString, "\n"))
					imgui.EndChild()
					imgui.EndGroup()

					// End TABS
					imgui.EndTabItem()
				}
				imgui.PopID()
			}
			// Ending the packages version tab bar
			imgui.EndTabBar()
			// Ending the packages name tab content
			imgui.EndTabItem()
		}
		imgui.PopID()
	}
	// Ending the packages name tab bar
	imgui.EndTabBar()
}
