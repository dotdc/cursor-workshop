import fs from "node:fs";
import path from "node:path";
import { marked } from "marked";

import { cloneBooks, Book } from "./models.js";

let workshopHTML: string | null = null;

export function getBooks(): Book[] {
  return cloneBooks();
}

export function getHealthStatus(): string {
  return "OK";
}

export function formatRarity(rating: number): string {
  return `${rating.toFixed(1)}★`;
}

export function loadWorkshopGoalsHTML(basePath: string): string {
  if (workshopHTML) {
    return workshopHTML;
  }

  const readmePath = path.join(basePath, "README.md");

  let raw: string;
  try {
    raw = fs.readFileSync(readmePath, "utf-8");
  } catch (error) {
    workshopHTML = `<p>Workshop goals are unavailable.</p>`;
    return workshopHTML;
  }

  const section = extractWorkshopSection(raw);
  if (!section) {
    workshopHTML = `<p>Workshop goals are unavailable.</p>`;
    return workshopHTML;
  }

  const rendered = marked(section);
  const html = typeof rendered === "string" ? rendered : String(rendered);
  workshopHTML = ensureHintSpacing(decorateHeadings(html));
  return workshopHTML;
}

function extractWorkshopSection(content: string): string {
  const heading = /^(##\s+Workshop[^\n]*)/m;
  const match = content.match(heading);

  if (!match || match.index === undefined) {
    return "";
  }

  let section = content.slice(match.index + match[0].length);

  const nextHeadingIndex = section.search(/\n##\s/);
  if (nextHeadingIndex !== -1) {
    section = section.slice(0, nextHeadingIndex);
  }

  const separatorIndex = section.indexOf("\n---");
  if (separatorIndex !== -1) {
    section = section.slice(0, separatorIndex);
  }

  const lines = section.split("\n").filter((line, index, arr) => {
    if (index === 0 || index === arr.length - 1) {
      return line.trim() !== "";
    }
    return true;
  });

  return lines.join("\n").trim();
}

const headingDecorator = /(<h3[^>]*>)(.*?)(<\/h3>)/gis;

function decorateHeadings(html: string): string {
  return html.replace(headingDecorator, (_match, open, inner, close) => {
    return `${open}✶ ${inner} ✶${close}\n<hr style="color:red"><br>`;
  });
}

const hintBlockSpacing = /<\/details>\s*<br\s*\/?>/gi;

function ensureHintSpacing(html: string): string {
  return html.replace(hintBlockSpacing, (match) => `${match}<br>`);
}
