package ui

import (
	"fmt"
	"github.com/inkyblackness/imgui-go/v4"
	"github.com/jpeizer/Vectorworks-Utility-Refresh/internal/software"
)

var (
	toggleSerialDetails bool
)

// RenderSoftware shows serials of found supported software
func RenderSoftware() {

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
					//imgui.PushFont(FontAwesome)
					if imgui.Button("\uF013" + "##" + installation.Year + "licenseButton") {
						toggleSerialDetails = !toggleSerialDetails
					}
					//imgui.PopFont()
					// Show License Tags
					if toggleSerialDetails {
						imgui.BeginTableV("##softwareTagsTable", 4, imgui.TableFlagsSizingFixedFit, imgui.Vec2{X: -1, Y: 30}, 0)
						imgui.TableNextColumn()
						imgui.Text(installation.License.Platform)
						imgui.SameLine()
						imgui.Dummy(imgui.Vec2{X: 20, Y: -1})
						imgui.TableNextColumn()
						imgui.Text(installation.License.Local)
						imgui.SameLine()
						imgui.Dummy(imgui.Vec2{X: 20, Y: -1})
						imgui.TableNextColumn()
						imgui.Text(installation.License.Activation)
						imgui.SameLine()
						imgui.Dummy(imgui.Vec2{X: 20, Y: -1})
						imgui.TableNextColumn()
						imgui.Text(installation.License.Type)
						imgui.EndTable()
					}
					imgui.Dummy(imgui.Vec2{X: -1, Y: 5})
					imgui.BeginChildV("##softwareContentChild", imgui.ContentRegionAvail(), true, 0)

					//////////
					// Edit Serial
					//////////
					imgui.BeginTableV("##"+installation.Year+"SoftwareActions", 5, imgui.TableFlagsSizingFixedFit, imgui.Vec2{X: -1, Y: 30}, -1)
					//imgui.TableNextRow(0, 30)
					imgui.NextColumn()
					imgui.Text("Resource Manager Cache")
					imgui.NextColumn()
					imgui.Button("Clean")
					imgui.NextColumn()
					imgui.Button("Rename")
					imgui.NextColumn()
					imgui.Text("Test")
					imgui.NextColumn()
					imgui.Text("Last")
					imgui.EndTable()
					//////////
					// Clear User Data
					//////////

					imgui.EndChild()
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
