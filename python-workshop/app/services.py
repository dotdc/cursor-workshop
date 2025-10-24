"""Service layer with helpers that support the Flask routes."""
from __future__ import annotations

import re
from functools import lru_cache
from pathlib import Path

from markupsafe import Markup
import markdown

from .models import Book, load_books


def get_books() -> list[Book]:
    """Return the seeded books."""
    return load_books()


def get_health_status() -> str:
    """Return a static health status string."""
    return "OK"


def format_rarity(rating: float) -> str:
    return f"{rating:.1f}★"


@lru_cache(maxsize=1)
def get_workshop_goals_html() -> Markup:
    """Render the Workshop Goals section from the README as HTML."""
    readme_path = Path(__file__).resolve().parent.parent / "README.md"
    if not readme_path.exists():
        return Markup("<p>Workshop goals are unavailable.</p>")

    readme_text = readme_path.read_text(encoding="utf-8")
    section = _extract_workshop_goals(readme_text)
    if not section:
        return Markup("<p>Workshop goals are unavailable.</p>")

    html = markdown.markdown(section, extensions=["extra"])
    html = re.sub(
        r"(<h3[^>]*>)(.*?)(</h3>)",
        r"\1✶ \2 ✶\3\n<hr style=\"color:red\"><br>",
        html,
        flags=re.DOTALL,
    )
    return Markup(html)


def _extract_workshop_goals(readme_text: str) -> str:
    """Return the README snippet that corresponds to the Workshop goals/quests section."""
    match = re.search(r"^##\s+Workshop[^\n]*", readme_text, re.MULTILINE)
    if not match:
        return ""

    section = readme_text[match.end() :]
    # Stop at the next top-level heading if present.
    next_heading_index = section.find("\n## ")
    if next_heading_index != -1:
        section = section[:next_heading_index]

    # Cut off any trailing horizontal rule and closing note.
    hr_index = section.find("\n---")
    if hr_index != -1:
        section = section[:hr_index]

    lines = section.splitlines()
    # Remove potential leading blank line from slicing and trim.
    while lines and not lines[0].strip():
        lines.pop(0)
    return "\n".join(lines).strip()
