"""Entry point of the Snake Bookshop application."""
from __future__ import annotations

import os

from app import create_app


def main() -> None:
    """Run the server."""
    app = create_app()
    port = int(os.environ.get("PORT", "8081"))
    app.run(debug=True, port=port)


if __name__ == "__main__":
    main()
