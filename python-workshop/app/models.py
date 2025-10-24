"""Data models and seed data for the Snake Bookshop workshop."""
from __future__ import annotations

from dataclasses import asdict, dataclass
from functools import lru_cache
from typing import Any


@dataclass(slots=True)
class Book:
    """Representation of a pixel-art Snake Bookshop volume."""

    title: str
    collection: str
    rarity: float
    blurb: str
    accent: str
    accent_dark: str
    icon: str

    def model_dump(self) -> dict[str, Any]:
        """Return a JSON-safe representation of the dataclass."""
        return asdict(self)


_BOOKS: tuple[Book, ...] = (
    Book(
        title="Cursors History",
        collection="Lore Compendium",
        rarity=4.8,
        blurb="Chronicles the serpentine saga of pointing devices from cave glyphs to holo UI.",
        accent="#f97316",
        accent_dark="#9a3412",
        icon="img/cursors-history.svg",
    ),
    Book(
        title="Cursor Design",
        collection="Atelier Series",
        rarity=4.6,
        blurb="Pixel-perfect guidance for crafting charismatic cursors in 8-bit workshops.",
        accent="#22d3ee",
        accent_dark="#0e7490",
        icon="img/cursor-design.svg",
    ),
    Book(
        title="Magic Cursors",
        collection="Arcana Shelf",
        rarity=4.9,
        blurb="Spellbound glyphs infused with mana trails and shimmer particles.",
        accent="#a855f7",
        accent_dark="#6b21a8",
        icon="img/magic-cursors.svg",
    ),
    Book(
        title="Gaming Cursors",
        collection="Arcade Manuals",
        rarity=4.4,
        blurb="High-DPI sprites with crit chance stats for competitive pixel arenas.",
        accent="#ef4444",
        accent_dark="#7f1d1d",
        icon="img/gaming-cursors.svg",
    ),
    Book(
        title="Open Source Cursors",
        collection="Community Guild",
        rarity=4.5,
        blurb="Collaborative cursors with permissive licenses and remix guides.",
        accent="#22c55e",
        accent_dark="#166534",
        icon="img/open-source-cursors.svg",
    ),
    Book(
        title="Shitty Cursors",
        collection="Glitch Annex",
        rarity=3.2,
        blurb="A comedic anthology of cursed pointers that wobble, blink, and misbehave.",
        accent="#facc15",
        accent_dark="#854d0e",
        icon="img/shitty-cursors.svg",
    ),
    Book(
        title="Weird Cursors",
        collection="Oddities Cabinet",
        rarity=4.1,
        blurb="Experimental forms featuring tentacles, portals, and non-Euclidean trails.",
        accent="#38bdf8",
        accent_dark="#075985",
        icon="img/weird-cursors.svg",
    ),
    Book(
        title="Funny Cursors",
        collection="Comedy Stacks",
        rarity=4.3,
        blurb="Prankster cursors with slapstick animations and honk-triggered clicks.",
        accent="#fb7185",
        accent_dark="#9f1239",
        icon="img/funny-cursors.svg",
    ),
    Book(
        title="Cursor IDE",
        collection="Workshop Prime",
        rarity=5.0,
        blurb="Official guide to Cursor IDE shortcuts, pair-programming tricks, and AI rituals.",
        accent="#2563eb",
        accent_dark="#1e3a8a",
        icon="img/cursor-ide.svg",
    ),
)


@lru_cache(maxsize=1)
def load_books() -> list[Book]:
    """Return a copy of the seeded books."""
    return list(_BOOKS)
