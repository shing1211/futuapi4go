# Translating

Thanks for helping translate futuapi4go!

## Canonical language

**English is authoritative.** Every README translation links back to
[`README.md`](./README.md). If a translation and the English version disagree,
the English version is correct.

## Supported languages

| Locale | Language | File |
|--------|----------|------|
| `en` | English (canonical) | [`README.md`](./README.md) |
| `zh-Hans` | 简体中文 | [`README.zh-Hans.md`](./README.zh-Hans.md) |
| `zh-Hant` | 繁體中文 | [`README.zh-Hant.md`](./README.zh-Hant.md) |
| `ja` | 日本語 | [`README.ja.md`](./README.ja.md) |
| `ko` | 한국어 | [`README.ko.md`](./README.ko.md) |
| `es` | Español | [`README.es.md`](./README.es.md) |

`README.zh-CN.md` is a redirect stub kept only for link compatibility; do not
edit it.

## Scope

- **Translated:** `README.md` only.
- **Not translated:** `LICENSE`, `NOTICE`, `DISCLAIMER.md`, `SECURITY.md`, and the
  `docs/` set. English remains canonical for these. Do not translate legal text.

## Adding a language

1. Copy `README.md` to `README.<locale>.md` (BCP-47; prefer script tags for
   Chinese: `zh-Hans` / `zh-Hant`).
2. Add the language to the switcher line in **every** `README*.md` file.
3. Add a translation banner at the top (below the title), including the commit
   you translated from:

   ```markdown
   > Translation of the canonical English [README](./README.md). English is authoritative.
   > Last synced: <short-sha>
   ```

4. Run `python3 scripts/check_i18n.py` and fix anything it reports.
5. Open a pull request.

## Keeping in sync

When the English README changes materially:

- Update the affected translation(s).
- Bump the `Last synced:` commit in each updated translation.

Translations may lag; the banner makes the lag explicit. Where a translation is
markedly out of date, the maintainers may add a note pointing readers to the
English version.

## Rules

- Do **not** re-derive protocol counts or versions. Copy the current values from
  [`docs/VERSION_MAP.md`](./docs/VERSION_MAP.md) and the README badges.
- Keep code blocks identical to English except for comments.
- Keep the switcher line and banner format exact — `scripts/check_i18n.py`
  enforces both.
- Translation quality over quantity: a wrong translation of a trading SDK can
  cause real harm.
