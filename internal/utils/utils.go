package utils

import (
	_ "embed"
	"github.com/inkyblackness/imgui-go/v4"
	"log"
	"os"
)

var (
	//go:embed far.otf
	FAReg []byte
	//go:embed far.otf
	FALig []byte
	//go:embed RobotoReg.ttf
	Roboto []byte
	//go:embed RobotoBold.ttf
	RobotoBold []byte

	FontAwesome      imgui.Font
	FontAwesomeLight imgui.Font
	FontRobotoBold   imgui.Font
)

// GetHomeDirectory Define users home directory
func GetHomeDirectory() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return home
}

func LoadFonts() {
	fonts := imgui.CurrentIO().Fonts()

	rangeBuilder := imgui.GlyphRangesBuilder{}

	rangeBuilder.Add(0x0020, 0xF12D)

	newRanges := rangeBuilder.Build()

	fonts.AddFontFromMemoryTTF(Roboto, 18)
	FontAwesome = fonts.AddFontFromMemoryTTFV(FAReg, 16, imgui.DefaultFontConfig, newRanges.GlyphRanges)
	FontAwesomeLight = fonts.AddFontFromMemoryTTFV(FALig, 16, imgui.DefaultFontConfig, newRanges.GlyphRanges)
	FontRobotoBold = fonts.AddFontFromMemoryTTFV(RobotoBold, 18, imgui.DefaultFontConfig, newRanges.GlyphRanges)

	fonts.Build()
}
