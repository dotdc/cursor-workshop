# Gopher Bookshop - Go Workshop

Welcome to the "Gopher Bookshop", the Go variant of this [Cursor](https://cursor.com) workshop!

## Quick start

### Prerequisites

Make sure you have **Go**, **Make**, and **Cursor** installed on your computer.

**Install Cursor:**

- <https://cursor.com/download>

**Install Go and Make:**

| OS | Installation command |
| --- | --- |
| Ubuntu | `sudo apt install -y golang-go make` |
| Fedora | `sudo dnf install -y golang make` |
| Arch Linux | `sudo pacman -Sy go make` |
| MacOS | `brew install go make` |

### Run application

```bash
make
```

> [!IMPORTANT]
> Before starting the workshop, delete the other `*-workshop/` folders (or add a Cursor rule that ignores them) so Cursor's LLM context stays focused on the Go version.

## Workshop Quests

### The Frenchy's typo

Users of the Books API don't understand what `rarity` refers to. In fact, our French developer accidentally misused `rarity` instead of `rating`. Please fix that in `app/models.go` using Cursor's [tab](https://cursor.com/docs/tab/overview) autocompletion.

<details>
<summary>Need a hint?</summary><br>
Replace it first in the Book struct (line 7), then see if <b>tab</b> completion does the rest!
</details>
<br>

### From Shakespeare to Molière

The `.cursor/rules/pirate.mdc` file forces Cursor to answer like a pirate. While you enjoy the Monkey Island vibe, you would prefer that Cursor respond in French to sharpen your French skills. Ask Cursor to update the rule so replies default to French unless `SHAKESPEARE` is present in the prompt, allowing your British colleague to bypass it.

<details>
<summary>Need a hint?</summary><br>
Use Cursor <b>Agent</b> mode to rewrite the rule so it responds in French, but do not forget the SHAKESPEARE bypass for your British colleague!
</details>
<br>

### The Eternal Year Problem

The year in the footer is hardcoded to `2025` in `app/templates/base.html`. As 2026 is around the corner, and since we don't want to change it manually every year, see if Cursor can help display the current year automatically.

<details>
<summary>Need a hint?</summary><br>
You could either <b>Ask</b> it directly, or call an <b>Agent</b>!
</details>
<br>

### The Missing Search Feature

The search feature on the Books page was never developed by our lazy French developer.<br>
Can you add it?

<details>
<summary>Need a hint?</summary><br>
Copy paste the instruction to your Cursor <b>Agent</b>!
</details>
<br>

### The Caffeinated Artist

Your coffee addiction and artistic soul are telling you that one of the buttons deserves a better color.
Take a screenshot of one of the buttons, then ask Cursor to change its color to the legendary [Starbucks Green](https://brandpalettes.com/starbucks-coffee-color-codes/) (`#00704A`). The other buttons should keep their original color.

<details>
<summary>Need a hint?</summary><br>
You need an artistic <b>Agent</b> for this one!
</details>
<br>

### The Hidden Easter Egg

We suspect the original developer hid an easter egg in the Gopher Bookshop instead of working on the search feature. See if Cursor can help you uncover it!

<details>
<summary>Need a hint?</summary><br>
Just <b>Ask</b>...
</details>
<br>

### Trekkie Health Check

The function `GetHealthStatus()` in `app/services.go` currently returns a hardcoded `OK`. Since our customers are Star Trek fans, we need to follow their Trekkie [odd-even rule](https://screenrant.com/star-trek-movies-odd-number-curse-explained/), they've requested that the status should be `GOOD` when the request happens on an "even" second, and `BAD` when on an "odd" second.

<details>
<summary>Need a hint?</summary><br>
Good luck with these weirdos!
</details>
<br>

### The Rust Rewrite

You've finally realized that Python is slow and decided to rewrite the whole thing in Rust! See if Cursor can help you!

<details>
<summary>Need a hint?</summary><br>
Maybe it's the good time to enable <b>Max mode</b>!
</details>
<br>

### The Forgotten TDD Course

It seems the original developer completely missed the Test-Driven Development course, there are no tests at all! Use Cursor to add proper tests for this application so we can finally trust it.

<details>
<summary>Need a hint?</summary><br>
Maybe <b>Ask</b> Cursor to generate a few, then refine them with an Agent!
</details>
<br>

### The Documentalist

You've successfully made the Gopher Bookshop production-ready with your godlike developer skills (and maybe with Cursor's help!). But since you're a nice person who doesn't want the next developer to eat their hat in despair, you've decided to use Cursor to create a `MAINTAINERS_GUIDE.md` explaining the ins and outs of this software masterpiece!

<details>
<summary>Need a hint?</summary><br>
Gods don't need any hints!
</details>

---

Hope you will learn something new and have fun during this [Cursor](https://cursor.com) workshop! 🎉

## Credits

- Gopher logo (static/img/gopher-logo.png) sourced from Pixilart: <https://www.pixilart.com/photo/pixel-gopher-e12f7b080c61303>
