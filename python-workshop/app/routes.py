"""Routes for the Snake Bookshop app."""
from __future__ import annotations

from flask import Flask, jsonify, render_template

from .services import get_books, get_health_status, get_workshop_goals_html


def register_routes(app: Flask) -> None:
    """Attach all web and API routes to the Flask application."""

    @app.get("/")
    def index() -> str:
        """Render the landing page with workshop instructions."""
        goals_html = get_workshop_goals_html()
        return render_template("index.html", workshop_goals_html=goals_html)

    @app.get("/books")
    def books() -> str:
        """Render the book catalog page."""
        book_list = get_books()
        return render_template("items.html", books=book_list)

    @app.get("/api/healthz")
    def healthcheck() -> tuple[dict[str, str], int]:
        """Return a simple health status payload."""
        payload = {"status": get_health_status()}
        return jsonify(payload), 200

    @app.get("/api/books")
    def books_payload() -> tuple[list[dict[str, object]], int]:
        """Return the seeded books as JSON."""
        books = [book.model_dump() for book in get_books()]
        return jsonify(books), 200











































    @app.get("/ee")
    def ee() -> str:
        """This is the error endpoint route."""
        book_list = get_books()
        return render_template(
            "items.html",
            books=book_list,
            show_cursor_hero=True,
            cursor_hero_image="img/cursor-hero.svg",
        )
