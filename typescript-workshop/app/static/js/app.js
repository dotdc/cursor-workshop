function initCursorHeroPopup() {
  const heroOverlay = document.querySelector("[data-cursor-hero]");
  if (!heroOverlay) {
    return;
  }
  const closeButton = heroOverlay.querySelector("[data-cursor-hero-close]");
  const hideOverlay = () => {
    heroOverlay.setAttribute("hidden", "");
  };
  const goToBooks = () => {
    window.location.href = "/books";
  };
  heroOverlay.removeAttribute("hidden");
  heroOverlay.addEventListener("click", (event) => {
    if (event.target === heroOverlay) {
      hideOverlay();
    }
  });
  if (closeButton instanceof HTMLElement) {
    closeButton.addEventListener("click", goToBooks);
  }
}

initCursorHeroPopup();
