<h1 align="center">🏹 Hookshot</h1>

<p align="center">
  A TUI extension for the GitHub CLI that lets you browse repository <b>webhooks</b>,
see <b>deliveries</b>, and inspect <b>payloads/responses</b>.
  <br />
  <br />
  <img src="https://img.shields.io/github/release/marcelblijleven/gh-hookshot.svg" alt="Latest Release">
</p>

## Motivation

I work on applications that rely heavily on GitHub webhooks across hundreds of
repositories, so I wanted a fast, visual, and developer-friendly way to interact
with webhooks right inside my terminal.

My main use cases are:

* Quickly inspecting webhook deliveries without leaving the terminal
* Debugging failing hooks faster, copying payloads easily without opening the browser (upcoming feature)
* Redelivering hooks when necessary (upcoming feature)

## Quickstart

Available flags:

* `--owner` — repository owner (required)
* `--repo` — repository name (required)

> [!IMPORTANT]
> You’ll need admin permission on the target repo to view webhooks and deliveries; otherwise the UI will show a clear error.

### Keybindings

* Move: `↑/k` `↓/j` `←/h` `→/l`
* Select: `Enter`
* Toggle Help: `?`
* Quit: `q` `esc` `ctrl+c`

## Features

* **List webhooks** configured on a repository
* **Browse deliveries** per webhook (including redeliveries)
* **View delivery details** (JSON request/response rendered nicely in the terminal)

### Upcoming features

* [ ] Redeliver deliveries
* [ ] Copy delivery content to clipboard
* [ ] Layout improvements
* [ ] Determine owner/repo based on active git directory

## Installation

> [!IMPORTANT]
> Hookshot requires the official [GitHub CLI](https://github.com/cli/cli) to be installed and authenticated.

Install the extension

```bash
gh extension install marcelblijleven/gh-hookshot
```

## Developing

During development, you can install from a local checkout:

```bash
cd gh-hookshot
gh extension install .
go build && gh hookshot [FLAGS]
```

## Tech Stack

* [**Bubble Tea**](https://github.com/charmbracelet/bubbletea)
* [**Lip Gloss**](https://github.com/charmbracelet/lipgloss)
* [**Glamour**](https://github.com/charmbracelet/glamour)
* [**go-gh**](https://github.com/cli/go-gh)
* [**Cobra**](https://github.com/spf13/cobra)
