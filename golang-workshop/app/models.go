package app

// Book represents a pixel-art volume in the Gopher Bookshop.
type Book struct {
	Title      string  `json:"title"`
	Collection string  `json:"collection"`
	Rarity     float64 `json:"rarity"`
	Blurb      string  `json:"blurb"`
	Accent     string  `json:"accent"`
	AccentDark string  `json:"accent_dark"`
	Icon       string  `json:"icon"`
}

// seededBooks contains the static catalogue used across the workshop.
var seededBooks = []Book{
	{
		Title:      "Cursors History",
		Collection: "Lore Compendium",
		Rarity:     4.8,
		Blurb:      "Chronicles the serpentine saga of pointing devices from cave glyphs to holo UI.",
		Accent:     "#f97316",
		AccentDark: "#9a3412",
		Icon:       "img/cursors-history.svg",
	},
	{
		Title:      "Cursor Design",
		Collection: "Atelier Series",
		Rarity:     4.6,
		Blurb:      "Pixel-perfect guidance for crafting charismatic cursors in 8-bit workshops.",
		Accent:     "#22d3ee",
		AccentDark: "#0e7490",
		Icon:       "img/cursor-design.svg",
	},
	{
		Title:      "Magic Cursors",
		Collection: "Arcana Shelf",
		Rarity:     4.9,
		Blurb:      "Spellbound glyphs infused with mana trails and shimmer particles.",
		Accent:     "#a855f7",
		AccentDark: "#6b21a8",
		Icon:       "img/magic-cursors.svg",
	},
	{
		Title:      "Gaming Cursors",
		Collection: "Arcade Manuals",
		Rarity:     4.4,
		Blurb:      "High-DPI sprites with crit chance stats for competitive pixel arenas.",
		Accent:     "#ef4444",
		AccentDark: "#7f1d1d",
		Icon:       "img/gaming-cursors.svg",
	},
	{
		Title:      "Open Source Cursors",
		Collection: "Community Guild",
		Rarity:     4.5,
		Blurb:      "Collaborative cursors with permissive licenses and remix guides.",
		Accent:     "#22c55e",
		AccentDark: "#166534",
		Icon:       "img/open-source-cursors.svg",
	},
	{
		Title:      "Shitty Cursors",
		Collection: "Glitch Annex",
		Rarity:     3.2,
		Blurb:      "A comedic anthology of cursed pointers that wobble, blink, and misbehave.",
		Accent:     "#facc15",
		AccentDark: "#854d0e",
		Icon:       "img/shitty-cursors.svg",
	},
	{
		Title:      "Weird Cursors",
		Collection: "Oddities Cabinet",
		Rarity:     4.1,
		Blurb:      "Experimental forms featuring tentacles, portals, and non-Euclidean trails.",
		Accent:     "#38bdf8",
		AccentDark: "#075985",
		Icon:       "img/weird-cursors.svg",
	},
	{
		Title:      "Funny Cursors",
		Collection: "Comedy Stacks",
		Rarity:     4.3,
		Blurb:      "Prankster cursors with slapstick animations and honk-triggered clicks.",
		Accent:     "#fb7185",
		AccentDark: "#9f1239",
		Icon:       "img/funny-cursors.svg",
	},
	{
		Title:      "Cursor IDE",
		Collection: "Workshop Prime",
		Rarity:     5.0,
		Blurb:      "Official guide to Cursor IDE shortcuts, pair-programming tricks, and AI rituals.",
		Accent:     "#2563eb",
		AccentDark: "#1e3a8a",
		Icon:       "img/cursor-ide.svg",
	},
}

// CloneBooks returns a defensive copy of the seeded data to avoid mutations.
func CloneBooks() []Book {
	books := make([]Book, len(seededBooks))
	copy(books, seededBooks)
	return books
}
