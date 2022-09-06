package erase

const (
	prefix  = string("\u001b[")
	sSuffix = "J"
	lSuffix = "K"
)

// ToScreenEnd erases from the cursor to the end of the screen
func ToScreenEnd() string {
	return prefix + "0" + sSuffix
}

// ToScreenBegin erases from the cursor to the screen beginning
func ToScreenBegin() string {
	return prefix + "1" + sSuffix
}

// Screen erases the screen
func Screen() string {
	return prefix + "2" + sSuffix
}

// SavedLines deletes the saved lines
func SavedLines() string {
	return prefix + "3" + sSuffix
}

// ToLineEnd deletes until the line end
func ToLineEnd() string {
	return prefix + "0" + lSuffix
}

// ToLineStart deletes untill the line start
func ToLineStart() string {
	return prefix + "1" + lSuffix
}

// Line deletes the line
func Line() string {
	return prefix + "2" + lSuffix
}
