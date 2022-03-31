package ui

import (
	"fmt"
	"github.com/inkyblackness/imgui-go/v4"
	"github.com/jpeizer/Vectorworks-Utility-Refresh/internal/packages"
	"github.com/jpeizer/Vectorworks-Utility-Refresh/internal/utils"
)

var (
	toggleSerialDetails bool
)

// RenderSoftware shows serials of found supported packages
func RenderSoftware() {
	var item int32 = 0
	// Start of packages tab bar
	imgui.BeginTabBar("##SoftwareTabBar")
	// Run for all active supported packages
	for _, swPkg := range application.SoftwarePackages {
		// Test for installations of active packages prior to making a table
		if len(swPkg.Installations) == 0 {
			continue
		}

		// Insert new tab for each installed supported packages
		if imgui.BeginTabItem(swPkg.Name + "##" + swPkg.Name + "TabItem") {
			// Begin of packages year tab bar
			imgui.BeginTabBar("##" + swPkg.Name + "TabBar")
			// Find all installed packages versions
			for _, installation := range swPkg.Installations {
				// Insert a new tab for all packages versions found
				if imgui.BeginTabItem(installation.Year + "##" + swPkg.Name + installation.Year + "TabItem") {
					// ----------------------------
					// LAYOUT FOR SOFTWARE FEATURES
					// ----------------------------
					// Software serial label
					imgui.Dummy(imgui.Vec2{X: imgui.ContentRegionAvail().X, Y: 5})
					//imgui.PushFont()
					imgui.PushItemWidth(350)
					// Flags 2 InputTextFlagsCharsUppercase | 4 InputTextFlagsAutoSelectAll | InputTextFlagsEnterReturnsTrue
					if imgui.InputTextV("##EditedSerial", &installation.License.Serial, 1<<2|1<<4|1<<5, nil) {
						err := packages.ReplaceOldSerial(installation, installation.License.Serial)
						if err != nil {
							fmt.Errorf("Error replacing old serial: %s", err)
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
					if imgui.ButtonV("\uF05A Info"+"##"+installation.Year+"licenseButton", imgui.Vec2{X: 70, Y: 25}) {
						toggleSerialDetails = !toggleSerialDetails
					}

					// License Cleanup Button
					if imgui.BeginPopupModalV("Clean Software", nil, imgui.WindowFlagsAlwaysAutoResize) {
						imgui.Text("Select the options below to clean the packages.")

						imgui.Separator()
						imgui.Dummy(imgui.Vec2{X: 0, Y: 5})

						//imgui.Checkbox("Remove resource manager cache##RMC", &installation.Options.RemoveRMC)
						//imgui.Checkbox("Remove user data##RMUD", &installation.Options.RemoveUserData)
						//imgui.Checkbox("Remove user settings##RMUS", &installation.Options.RemoveUserSettings)
						//imgui.Checkbox("Remove installer files##RMIF", &installation.Options.RemoveInstallerSettings)
						//imgui.Checkbox("Remove all user data##RMALL", &installation.Options.RemoveAllData)

						imgui.Dummy(imgui.Vec2{X: 0, Y: 5})
						imgui.Separator()
						imgui.Dummy(imgui.Vec2{X: 0, Y: 5})

						if imgui.Button("Remove") {

							//"Remove user data",
							//"Remove user settings",
							//"Remove installer files",
							//"Remove all user data",
							//packages.RemoveSoftware(installation)
							imgui.CloseCurrentPopup()
							//err := packages.GenerateInstalledSoftwareMap()
							//if err != nil {
							//	fmt.Errorf("error updating internal installation data after serial update %v", err)
							//}
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
					imgui.Text(fmt.Sprintf("%v", item))
					imgui.EndChild()
					imgui.EndGroup()

					// End TABS
					imgui.EndTabItem()
				}
			}
			// Ending the packages version tab bar
			imgui.EndTabBar()
			// Ending the packages name tab content
			imgui.EndTabItem()
		}
	}
	// Ending the packages name tab bar
	imgui.EndTabBar()
}
