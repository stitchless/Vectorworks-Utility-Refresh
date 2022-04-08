package app

import (
	"fmt"
	"github.com/jpeizer/Vectorworks-Utility-Refresh/internal/ui"
	"time"

	"github.com/inkyblackness/imgui-go/v4"
)

// Platform covers mouse/keyboard/gamepad inputs, cursor shape, timing, windowing.
type Platform interface {
	// ShouldStop is regularly called as the abort condition for the program loop.
	ShouldStop() bool
	// ProcessEvents is called once per render loop to dispatch any pending events.
	ProcessEvents()
	// DisplaySize returns the dimension of the display.
	DisplaySize() [2]float32
	// FramebufferSize returns the dimension of the framebuffer.
	FramebufferSize() [2]float32
	// NewFrame marks the beginning of a render pass. It must update the imgui IO state according to user input (mouse, keyboard, ...)
	NewFrame()
	// PostRender marks the completion of one render pass. Typically, this causes the display buffer to be swapped.
	PostRender()
	// ClipboardText returns the current text of the clipboard, if available.
	ClipboardText() (string, error)
	// SetClipboardText sets the text as the current text of the clipboard.
	SetClipboardText(text string)
}

type clipboard struct {
	platform Platform
}

func (board clipboard) Text() (string, error) {
	return board.platform.ClipboardText()
}

func (board clipboard) SetText(text string) {
	board.platform.SetClipboardText(text)
}

// Renderer covers rendering imgui draw data.
type Renderer interface {
	// PreRender causes the display buffer to be prepared for new output.
	PreRender(clearColor [3]float32)
	// Render draws the provided imgui draw data.
	Render(displaySize [2]float32, framebufferSize [2]float32, drawData imgui.DrawData)
}

const (
	//millisPerSecond = 1000
	sleepDuration = time.Millisecond * 10
)

// Run implements the main program loop of the demo. It returns when the platform signals to stop.
// This demo application shows some basic features of ImGui, as well as exposing the standard demo window.
func Run(p Platform, r Renderer) {
	imgui.CurrentIO().SetClipboard(clipboard{platform: p})

	showDemoWindow := false
	clearColor := [3]float32{0.0, 0.0, 0.0}
	showDebugWindow := true

	for !p.ShouldStop() {
		p.ProcessEvents()

		// Signal start of a new frame
		p.NewFrame()
		imgui.NewFrame()

		//1. Show a simple window.
		//Tip: if we don't call imgui.Begin()/imgui.End() the widgets automatically appears in a window called "Debug".
		if showDebugWindow {
			// Sets window size to be the same as the framebuffer size
			size := p.DisplaySize()
			imgui.SetNextWindowPos(imgui.Vec2{X: 0, Y: 0})
			imgui.SetNextWindowSize(imgui.Vec2{X: size[0], Y: size[1]})

			// Set the window to be a fixed size
			imgui.BeginV("Main Window", &showDebugWindow,
				imgui.WindowFlagsNoTitleBar|
					imgui.WindowFlagsNoCollapse|
					imgui.WindowFlagsNoMove|
					imgui.WindowFlagsNoResize)

			numberOfButtons := len(ui.AllActiveFeatures)

			imgui.BeginTableV("Features", numberOfButtons, 0, imgui.Vec2{X: -1, Y: 30}, -1)
			imgui.TableNextRowV(0, 30)

			for _, activeFeature := range ui.AllActiveFeatures {
				if ui.CurrentFeature == "" {
					ui.CurrentFeature = activeFeature
				}

				imgui.TableNextColumn()
				imgui.PushID(activeFeature.String())

				var pressed bool

				if ui.CurrentFeature == activeFeature {
					imgui.PushStyleColor(imgui.StyleColorButton, imgui.Vec4{X: 0.06, Y: 0.53, Z: 0.98, W: 1.00})
				} else {
					imgui.PushStyleColor(imgui.StyleColorButtonActive, imgui.Vec4{X: 0.06, Y: 0.53, Z: 0.98, W: 1.00})
				}

				if imgui.ButtonV(activeFeature.String(), imgui.Vec2{X: -1, Y: 30}) {
					activeFeature.SetActive()
					pressed = !pressed
					fmt.Println("Active Feature:", activeFeature.String())
				}
				imgui.PopStyleColor()

				imgui.PopID()
			}

			imgui.EndTable()

			ui.CurrentFeature.Render()
			imgui.End()
		}

		// 3. Show the ImGui demo window. Most of the sample code is in imgui.ShowDemoWindow().
		// Read its code to learn more about Dear ImGui!
		if showDemoWindow {
			// Normally user code doesn't need/want to call this because positions are saved in .ini file anyway.
			// Here we just want to make the demo initial state a bit more friendly!
			const demoX = 650
			const demoY = 20
			imgui.SetNextWindowPosV(imgui.Vec2{X: demoX, Y: demoY}, imgui.ConditionFirstUseEver, imgui.Vec2{})

			imgui.ShowDemoWindow(&showDemoWindow)
		}
		//if showGoDemoWindow {
		//	demo.Show()
		//}

		// Rendering
		imgui.Render() // This call only creates the draw data list. Actual rendering to framebuffer is done below.

		r.PreRender(clearColor)
		// At this point, the application could perform its own rendering...
		//app.RenderScene()

		r.Render(p.DisplaySize(), p.FramebufferSize(), imgui.RenderedDrawData())
		p.PostRender()

		// sleep to avoid 100% CPU usage for this demo
		<-time.After(sleepDuration)
	}
}
