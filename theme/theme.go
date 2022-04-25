package theme

import (
	"fyne.io/fyne/v2"
	fyne_theme "fyne.io/fyne/v2/theme"
	"image/color"
)

type FzHtJtTtfDarkTheme struct {
}

type FzHtJtTtfLightTheme struct {
}

func (t *FzHtJtTtfDarkTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	return fyne_theme.DarkTheme().Color(n, v)
}

func (t *FzHtJtTtfDarkTheme) Font(fyne.TextStyle) fyne.Resource {
	return resourceFangZhengHeiTiJianTiTtf
}

func (t *FzHtJtTtfDarkTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return fyne_theme.DarkTheme().Icon(n)
}

func (t *FzHtJtTtfDarkTheme) Size(n fyne.ThemeSizeName) float32 {
	return fyne_theme.DarkTheme().Size(n)
}

func (t *FzHtJtTtfLightTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	return fyne_theme.LightTheme().Color(n, v)
}

func (t *FzHtJtTtfLightTheme) Font(fyne.TextStyle) fyne.Resource {
	return resourceFangZhengHeiTiJianTiTtf
}

func (t *FzHtJtTtfLightTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return fyne_theme.LightTheme().Icon(n)
}

func (t *FzHtJtTtfLightTheme) Size(n fyne.ThemeSizeName) float32 {
	return fyne_theme.LightTheme().Size(n)
}

func DarkTheme() fyne.Theme {
	return &FzHtJtTtfDarkTheme{}
}

func LightTheme() fyne.Theme {
	return &FzHtJtTtfLightTheme{}
}
