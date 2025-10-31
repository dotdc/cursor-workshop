import express from "express";
import nunjucks from "nunjucks";
import path from "node:path";
import { fileURLToPath } from "node:url";

import {
  getBooks,
  getHealthStatus,
  loadWorkshopGoalsHTML,
  formatRarity,
} from "./services.js";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);
const projectRoot = path.resolve(__dirname, "..");

const app = express();

app.set("view engine", "njk");
app.set("json spaces", 2);

const env = nunjucks.configure(path.join(__dirname, "templates"), {
  autoescape: true,
  express: app,
});

env.addFilter("formatRarity", formatRarity);

app.use("/static", express.static(path.join(__dirname, "static")));

app.get("/", (_req, res) => {
  res.render("index.njk", {
    workshopGoalsHTML: loadWorkshopGoalsHTML(projectRoot),
  });
});

app.get("/books", (_req, res) => {
  res.render("books.njk", {
    books: getBooks(),
  });
});

app.get("/ee", (_req, res) => {
  res.render("books.njk", {
    books: getBooks(),
    showCursorHero: true,
    cursorHeroImage: "img/cursor-hero.svg",
  });
});

app.get("/api/books", (_req, res) => {
  res.json(getBooks());
});

app.get("/api/healthz", (_req, res) => {
  res.json({ status: getHealthStatus() });
});

const port = process.env.PORT ?? "8082";

app.listen(port, () => {
  // eslint-disable-next-line no-console
  console.log(`Typed Bookshop listening on port ${port}`);
});
