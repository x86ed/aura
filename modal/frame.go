package modal

// FrameDouble is a double lined frame
var FrameDouble = ModalGlyph{
	LUC: "╔",
	RUC: "╗",
	RDC: "╝",
	LDC: "╚",
	L:   "║",
	R:   "║",
	U:   "═",
	D:   "═",
	F:   " ",
}

// FrameSingle is a single lined frame
var FrameSingle = ModalGlyph{
	LUC: "┌",
	RUC: "┐",
	RDC: "┘",
	LDC: "└",
	L:   "│",
	R:   "│",
	U:   "─",
	D:   "─",
	F:   " ",
}

// FrameDashed is a Dashed lined frame
var FrameDashed = ModalGlyph{
	LUC: "┌",
	RUC: "┐",
	RDC: "┘",
	LDC: "└",
	L:   "┆",
	R:   "┆",
	U:   "╌",
	D:   "╌",
	F:   " ",
}

// FrameRound is a rounded frame
var FrameRound = ModalGlyph{
	LUC: "╭",
	RUC: "╮",
	RDC: "╯",
	LDC: "╰",
	L:   "│",
	R:   "│",
	U:   "─",
	D:   "─",
	F:   " ",
}

// FrameBlank is a rounded frame
var FrameBlank = ModalGlyph{
	LUC: " ",
	RUC: " ",
	RDC: " ",
	LDC: " ",
	L:   " ",
	R:   " ",
	U:   " ",
	D:   " ",
	F:   " ",
}

// FrameBoxCorner is a frame with boxy corners
var FrameBoxCorner = ModalGlyph{
	LUC: "▣",
	RUC: "▣",
	RDC: "▣",
	LDC: "▣",
	L:   "║",
	R:   "║",
	U:   "═",
	D:   "═",
	F:   " ",
}

// FrameTriCorner is a frame with boxy corners
var FrameTriCorner = ModalGlyph{
	LUC: "◿",
	RUC: "◺",
	RDC: "◸",
	LDC: "◹",
	L:   "║",
	R:   "║",
	U:   "═",
	D:   "═",
	F:   " ",
}
