package main

import (
	"github.com/adrg/xdg"
)

const (
	SettingsPath                = "MasterPlan/settings.json"
	SettingsLegacyPath          = "masterplan-settings.json"
	SettingsTheme               = "Theme"
	SettingsDownloadDirectory   = "DownloadDirectory"
	SettingsWindowPosition      = "WindowPosition"
	SettingsSaveWindowPosition  = "SaveWindowPosition"
	SettingsCustomFontPath      = "CustomFontPath"
	SettingsFontSize            = "FontSize"
	SettingsKeybindings         = "Keybindings"
	SettingsTargetFPS           = "TargetFPS"
	SettingsUnfocusedFPS        = "UnfocusedFPS"
	SettingsDisableSplashscreen = "DisableSplashscreen"
	SettingsBorderlessWindow    = "BorderlessWindow"
	SettingsRecentPlanList      = "RecentPlanList"
	SettingsAlwaysShowNumbering = "AlwaysShowNumbering"
)

func NewProgramSettings() *Properties {

	props := NewProperties()
	props.Get(SettingsTheme).Set("Moonlight")
	props.Get(SettingsDownloadDirectory).Set("")
	props.Get(SettingsTargetFPS).Set(60)
	props.Get(SettingsUnfocusedFPS).Set(60)
	props.Get(SettingsFontSize).Set(30)
	props.Get(SettingsDownloadDirectory).Set("")

	path, _ := xdg.ConfigFile(SettingsPath)

	// Attempt to load the property here
	if FileExists(path) {
		props.Load(path)
	}

	props.OnChange = func(property *Property) {
		props.Save(path)
	}

	props.Save(path)

	// TargetFPS:         60,
	// 	UnfocusedFPS:      60,
	// 	WindowPosition:    sdl.Rect{-1, -1, 0, 0},
	// 	Theme:             "Moonlight", // Default theme
	// 	Keybindings:       NewKeybindings(),
	// 	FontSize:          30,
	// 	DownloadDirectory: "",

	return props
}

// func (ps *OldProgramSettings) CleanUpRecentPlanList() {

// 	newList := []string{}
// 	for i, s := range ps.RecentPlanList {
// 		_, err := os.Stat(s)
// 		if err == nil {
// 			newList = append(newList, ps.RecentPlanList[i]) // We could alter the slice to cut out the strings that are invalid, but this is visually cleaner and easier to understand
// 		}
// 	}
// 	ps.RecentPlanList = newList
// }