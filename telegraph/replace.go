package telegraph

import "strings"

func replace(title string) string {
	title = strings.Replace(title, "\n", "", -1)
	title = strings.Replace(title, "［", "", -1)
	title = strings.Replace(title, "］", "", -1)
	title = strings.Replace(title, "\u00A0", "", -1)
	title = strings.Replace(title, "，", "", -1)
	title = strings.Replace(title, "《", "", -1)
	title = strings.Replace(title, "》", "", -1)
	title = strings.Replace(title, " ", "", -1)
	title = strings.Replace(title, "[", "", -1)
	title = strings.Replace(title, "]", "", -1)
	title = strings.Replace(title, "（", "", -1)
	title = strings.Replace(title, "）", "", -1)
	//？
	title = strings.Replace(title, "？", "", -1)

	for strings.Contains(title, " ") {
		title = strings.Replace(title, " ", "", -1)
	}
	return title
}
