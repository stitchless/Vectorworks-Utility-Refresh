package ui

import (
	"fmt"
	"github.com/inkyblackness/imgui-go/v4"
	"github.com/jpeizer/Vectorworks-Utility-Refresh/internal/software"
	"github.com/jpeizer/Vectorworks-Utility-Refresh/internal/utils"
)

var (
	toggleSerialDetails bool
)

// RenderSoftware shows serials of found supported software
func RenderSoftware() {
	var item int32 = 0

	// Start of software tab bar
	imgui.BeginTabBar("##SoftwareTabBar")
	// Run for all active supported software
	for _, softwareName := range software.AllActiveSoftwareNames {
		// Test for installations of active software prior to making a table
		if len(software.AllInstalledSoftwareMap[softwareName]) == 0 {
			continue
		}
		// Insert new tab for each installed supported software
		if imgui.BeginTabItem(softwareName + "##" + softwareName + "TabItem") {
			// Begin of software year tab bar
			imgui.BeginTabBar("##" + softwareName + "TabBar")
			// Find all installed software versions
			for _, installation := range software.AllInstalledSoftwareMap[softwareName] {
				// Insert a new tab for all software versions found
				if imgui.BeginTabItem(installation.Year + "##" + softwareName + installation.Year + "TabItem") {
					// ----------------------------
					// LAYOUT FOR SOFTWARE FEATURES
					// ----------------------------
					// Software serial label
					imgui.Dummy(imgui.Vec2{X: imgui.ContentRegionAvail().X, Y: 5})
					//imgui.PushFont()
					imgui.PushItemWidth(350)
					// Flags 2 InputTextFlagsCharsUppercase | 4 InputTextFlagsAutoSelectAll | InputTextFlagsEnterReturnsTrue
					if imgui.InputTextV("##EditedSerial", &installation.License.Serial, 1<<2|1<<4|1<<5, nil) {
						software.ReplaceOldSerial(installation, installation.License.Serial)
						err := software.GenerateInstalledSoftwareMap()
						if err != nil {
							fmt.Errorf("error updating internal installation data after serial update %v", err)
						}
					}
					imgui.PopItemWidth()
					//imgui.PopFont()
					if imgui.IsItemHovered() {
						// Introduce timer for the tooltips
						// https://gist.github.com/toutougabi/f56309cb9f802f34eeddda65eb27cad2
						imgui.SetTooltip("Insert new serial and press enter to update")
					}

					// Cog Icon button
					imgui.SameLine()
					imgui.PushFont(utils.FontAwesomeLight)
					if imgui.Button("\uF05A" + "##" + installation.Year + "licenseButton") {
						toggleSerialDetails = !toggleSerialDetails
					}

					//installation.CreateModal()
					var open bool = true
					imgui.BeginPopupModalV("##RemoveSoftware", &open, 0)
					if imgui.Button("Cancel") {
						imgui.CloseCurrentPopup()
					}

					imgui.SameLine()
					if imgui.Button("\uF12D" + "##CleanDialog") {
						imgui.OpenPopup("RemoveSoftware")
					}
					if imgui.IsItemHovered() {
						imgui.SetTooltip("Clean up this installation of software")
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
					//imgui.BeginChildV("Options ##"+softwareName+installation.Year+"licenseChild", imgui.Vec2{X: 0, Y: 0}, true, 0)

					groupWidth := imgui.ContentRegionAvail().X * .30

					imgui.PushItemWidth(groupWidth)

					imgui.BeginGroup()
					imgui.PushFont(utils.FontRobotoBold)
					imgui.Text("Options")
					imgui.PopFont()

					//imgui.Text("Remove Resource Manage Cache")
					//imgui.SameLine()
					if imgui.ButtonV("Remove Resource Manager Cache##RMRMCacheButton", imgui.Vec2{X: groupWidth, Y: 0}) {
						err := software.RemoveResourceManagerCache(installation)
						if err != nil {
							fmt.Errorf("error removing Resource Manager Cache %v", err)
						}

						fmt.Println("Removed Resource Manager Cache")
					}

					if imgui.ButtonV("Remove User Data##RMUserDataButton", imgui.Vec2{X: groupWidth, Y: 0}) {
						err := software.RemoveUserData(installation)
						if err != nil {
							fmt.Errorf("error removing User Data %v", err)
						}

						fmt.Println("Removed User Data")
					}

					imgui.ListBox("##optionsBody", &item, []string{
						"Remove resource manager cache",
						"Remove user data",
						"Remove user settings",
						"Remove installer files",
						"Remove all user data",
					})

					imgui.Text("Select an option above to modify your install")
					imgui.EndGroup()
					imgui.PopItemWidth()

					imgui.SameLine()
					imgui.BeginGroup()
					imgui.PushFont(utils.FontRobotoBold)
					imgui.Text("Output")
					imgui.PopFont()
					imgui.BeginChildV("##optionsChild", imgui.Vec2{X: 0, Y: 0}, true, 0)
					imgui.Text(fmt.Sprintf("%v", item))
					imgui.EndChild()
					imgui.EndGroup()

					//imgui.EndChild()
					//imgui.LabelText("##"+softwareName+installation.Year+"licenseLabel", "Clear Data Options")
					//imgui.BeginTableV("##"+installation.Year+"SoftwareActions", 5, imgui.TableFlagsSizingFixedFit, imgui.Vec2{X: -1, Y: 30}, -1)
					//imgui.TableNextRow()
					//imgui.TableNextColumn()
					//if imgui.Button(" Resource Manager Cache " + "##" + installation.Year + "RMC") {
					//	//software.ClearInstalledSoftware(installation)
					//	//err := software.GenerateInstalledSoftwareMap()
					//	//if err != nil {
					//	//	fmt.Errorf("error updating internal installation data after serial update %v", err)
					//	//}
					//}
					//imgui.EndTable()

					////imgui.Dummy(imgui.Vec2{X: -1, Y: 5})
					//imgui.BeginChildV("##softwareContentChild", imgui.ContentRegionAvail(), true, 0)
					//
					////////////
					//// Edit Serial
					////////////
					//imgui.Text("Edit Serial")
					//
					////////////
					//// Clear User Data
					////////////
					//
					//imgui.EndChild()
					// ----------------------------
					// Ending the active software version tab content
					imgui.EndTabItem()
				}
			}
			// Ending the software version tab bar
			imgui.EndTabBar()
			// Ending the software name tab content
			imgui.EndTabItem()
		}
	}
	// Ending the software name tab bar
	imgui.EndTabBar()
}
