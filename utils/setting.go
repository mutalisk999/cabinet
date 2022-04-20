package utils

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

func WinSizeX1() fyne.Size {
	return fyne.NewSize(800, 400)
}

func WinSizeX1_5() fyne.Size {
	return fyne.NewSize(1200, 600)
}

func WinSizeX2() fyne.Size {
	return fyne.NewSize(1600, 800)
}

func ThemeDefault() fyne.Theme {
	return theme.DefaultTheme()
}

func ThemeDark() fyne.Theme {
	return theme.DarkTheme()
}

func ThemeLight() fyne.Theme {
	return theme.LightTheme()
}
