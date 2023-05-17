package html2image

var DPI = float64(72)

func calcCharacterPx(text string, fontSize float64) float64 {
	return (calCharacterLen(text) * fontSize * DPI / 72) / 3
}

func calCharacterLen(str string) float64 {
	sl := 0.0
	rs := []rune(str)
	for _, r := range rs {
		rInt := int(r)
		if rInt < 128 {
			sl += 1.7
		} else {
			sl += 3
		}
	}
	return sl
}

func splitMultiLineText(text string, size float64, domX1, parentX2, parentX1 int) []string {
	firstLineWidth := float64(parentX2-domX1) - size
	maxLineWidth := float64(parentX2-parentX1) - size

	var multiTextGroupByLine []string
	var tmpStr string
	var tmpWidth float64
	for idx, value := range text {
		subStr := string(value)
		strWidth := (calCharacterLen(subStr) * size * DPI / 72) / 3
		tmpWidth += strWidth
		tmpStr += subStr
		if (idx == 0 && tmpWidth > firstLineWidth) || (idx > 0 && tmpWidth > maxLineWidth) {
			multiTextGroupByLine = append(multiTextGroupByLine, tmpStr)
			tmpWidth = 0
			tmpStr = ""
		}
	}
	if tmpStr != "" {
		multiTextGroupByLine = append(multiTextGroupByLine, tmpStr)
	}

	return multiTextGroupByLine
}
