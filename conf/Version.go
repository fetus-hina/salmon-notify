package conf

import "fmt"

const AppName string = "salmon-notify"
const AppVersion string = "0.2.1"
const AppURL string = "https://github.com/fetus-hina/salmon-notify"

func GetUserAgentString() string {
	return fmt.Sprintf(
		"%s/%s (+%s)",
		AppName,
		AppVersion,
		AppURL,
	)
}

func CliVersion() string {
	return fmt.Sprintf(
		"%s version %s\n"+
			"Copyright (C) 2019-2022 AIZAWA Hina\n"+
			"Released under the MIT License\n"+
			"%s",
		AppName,
		AppVersion,
		AppURL,
	)
}
