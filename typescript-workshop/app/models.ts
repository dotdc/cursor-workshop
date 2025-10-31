export interface Book {
  title: string;
  collection: string;
  rarity: number;
  blurb: string;
  accent: string;
  accentDark: string;
  icon: string;
}

const seededBooks: Book[] = [
  {
    title: "Cursors History",
    collection: "Lore Compendium",
    rarity: 4.8,
    blurb:
      "Chronicles the serpentine saga of pointing devices from cave glyphs to holo UI.",
    accent: "#f97316",
    accentDark: "#9a3412",
    icon: "img/cursors-history.svg",
  },
  {
    title: "Cursor Design",
    collection: "Atelier Series",
    rarity: 4.6,
    blurb: "Pixel-perfect guidance for crafting charismatic cursors in 8-bit workshops.",
    accent: "#22d3ee",
    accentDark: "#0e7490",
    icon: "img/cursor-design.svg",
  },
  {
    title: "Magic Cursors",
    collection: "Arcana Shelf",
    rarity: 4.9,
    blurb: "Spellbound glyphs infused with mana trails and shimmer particles.",
    accent: "#a855f7",
    accentDark: "#6b21a8",
    icon: "img/magic-cursors.svg",
  },
  {
    title: "Gaming Cursors",
    collection: "Arcade Manuals",
    rarity: 4.4,
    blurb: "High-DPI sprites with crit chance stats for competitive pixel arenas.",
    accent: "#ef4444",
    accentDark: "#7f1d1d",
    icon: "img/gaming-cursors.svg",
  },
  {
    title: "Open Source Cursors",
    collection: "Community Guild",
    rarity: 4.5,
    blurb: "Collaborative cursors with permissive licenses and remix guides.",
    accent: "#22c55e",
    accentDark: "#166534",
    icon: "img/open-source-cursors.svg",
  },
  {
    title: "Shitty Cursors",
    collection: "Glitch Annex",
    rarity: 3.2,
    blurb: "A comedic anthology of cursed pointers that wobble, blink, and misbehave.",
    accent: "#facc15",
    accentDark: "#854d0e",
    icon: "img/shitty-cursors.svg",
  },
  {
    title: "Weird Cursors",
    collection: "Oddities Cabinet",
    rarity: 4.1,
    blurb: "Experimental forms featuring tentacles, portals, and non-Euclidean trails.",
    accent: "#38bdf8",
    accentDark: "#075985",
    icon: "img/weird-cursors.svg",
  },
  {
    title: "Funny Cursors",
    collection: "Comedy Stacks",
    rarity: 4.3,
    blurb: "Prankster cursors with slapstick animations and honk-triggered clicks.",
    accent: "#fb7185",
    accentDark: "#9f1239",
    icon: "img/funny-cursors.svg",
  },
  {
    title: "Cursor IDE",
    collection: "Workshop Prime",
    rarity: 5.0,
    blurb: "Official guide to Cursor IDE shortcuts, pair-programming tricks, and AI rituals.",
    accent: "#2563eb",
    accentDark: "#1e3a8a",
    icon: "img/cursor-ide.svg",
  },
];

export function cloneBooks(): Book[] {
  return seededBooks.map((book) => ({ ...book }));
}
