#!/usr/bin/env python3
"""
check_i18n.py — keep README translations consistent with the English canonical.

Checks:
  1. Every expected README.<locale>.md exists.
  2. Every README contains the canonical language switcher, verbatim.
  3. Every non-English README carries the translation banner ("Last synced:").
  4. Every translation has the same number of Markdown headings as English
     (catches dropped/merged sections), ignoring fenced code blocks.
  5. README.zh-CN.md exists as a redirect stub to README.zh-Hans.md.

Exits non-zero on any failure. See TRANSLATING.md.
"""

import os
import re
import sys

ROOT = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))

# locale -> (native label, filename)
LANGUAGES = [
    ("en", "English", "README.md"),
    ("zh-Hans", "简体中文", "README.zh-Hans.md"),
    ("zh-Hant", "繁體中文", "README.zh-Hant.md"),
    ("ja", "日本語", "README.ja.md"),
    ("ko", "한국어", "README.ko.md"),
    ("es", "Español", "README.es.md"),
]
CANONICAL = "README.md"
STUB = "README.zh-CN.md"
STUB_TARGET = "README.zh-Hans.md"

SWITCHER_RE = re.compile(r"^\[[^\]]+\]\(\./README[^)]*\)( · \[[^\]]+\]\(\./README[^)]*\))+$", re.M)
HEADING_RE = re.compile(r"^(#{1,6})\s")
BANNER_RE = re.compile(r"Last synced:\s*\S+")


def switcher_line() -> str:
    return " · ".join(f"[{label}](./{file})" for _, label, file in LANGUAGES)


def read(path: str) -> str:
    with open(os.path.join(ROOT, path), encoding="utf-8") as fh:
        return fh.read()


def heading_count(text: str) -> int:
    count = 0
    fence = False
    for line in text.splitlines():
        if line.lstrip().startswith("```"):
            fence = not fence
            continue
        if not fence and HEADING_RE.match(line):
            count += 1
    return count


def main() -> int:
    failures = []
    expected_switcher = switcher_line()

    # 1 + 2 + 3
    base_headings = None
    for locale, _label, filename in LANGUAGES:
        path = os.path.join(ROOT, filename)
        if not os.path.exists(path):
            failures.append(f"missing file: {filename}")
            continue
        text = read(filename)

        if expected_switcher not in text:
            failures.append(f"{filename}: canonical language switcher not found")

        if locale != "en" and not BANNER_RE.search(text):
            failures.append(f"{filename}: missing translation banner (Last synced:)")

        if locale == "en":
            base_headings = heading_count(text)
        elif base_headings is not None and heading_count(text) != base_headings:
            failures.append(
                f"{filename}: heading count {heading_count(text)} != English {base_headings}"
            )

    # 4
    if os.path.exists(os.path.join(ROOT, STUB)):
        stub = read(STUB)
        if STUB_TARGET not in stub:
            failures.append(f"{STUB}: does not redirect to {STUB_TARGET}")
    else:
        failures.append(f"missing redirect stub: {STUB}")

    if failures:
        print("i18n check failed:", file=sys.stderr)
        for item in failures:
            print(f"  {item}", file=sys.stderr)
        return 1

    print(f"i18n OK: {len(LANGUAGES)} languages consistent")
    return 0


if __name__ == "__main__":
    sys.exit(main())
