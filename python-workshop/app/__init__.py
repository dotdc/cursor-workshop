"""Snake Bookshop application."""
from __future__ import annotations

from flask import Flask

from .routes import register_routes
from .services import format_rarity


def create_app() -> Flask:
    """Create and configure the Flask application."""
    app = Flask(__name__)
    app.config["JSON_SORT_KEYS"] = False

    register_routes(app)

    @app.context_processor
    def inject_template_helpers() -> dict[str, object]:
        """Expose formatting helpers to Jinja templates."""
        return {"format_rarity": format_rarity}

    return app
