package main

import (
	"fmt"
	"os"

	"github.com/inkyblackness/imgui-go/v4"

	"github.com/jpeizer/Vectorworks-Utility-Refresh/internal/app"
	"github.com/jpeizer/Vectorworks-Utility-Refresh/internal/platforms"
	"github.com/jpeizer/Vectorworks-Utility-Refresh/internal/renderers"
)

// TODO: set theme from https://github.com/AllenDang/giu/blob/b3e4ae718b3a78be20ccc69748f7fc96ea7b65f1/MasterWindow.go

func main() {
	context := imgui.CreateContext(nil)
	defer context.Destroy()
	io := imgui.CurrentIO()

	platform, err := platforms.NewGLFW(io, platforms.GLFWClientAPIOpenGL3)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}
	defer platform.Dispose()

	renderer, err := renderers.NewOpenGL3(io)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}
	defer renderer.Dispose()

	app.Run(platform, renderer)
}
