#!/usr/bin/env python3
"""Check every Markdown link in the repository.

Verifies that each relative link target exists and that each `#anchor`
matches a heading in the target file. External (http/https) links and
`[[wiki-links]]` (which resolve inside a GM's world, not here) are skipped.

Run from the repository root:  python3 scripts/check-links.py
Exit status 1 if anything is broken, so CI fails on a bad link.
"""
from __future__ import annotations

import os
import re
import sys

SKIP_DIRS = {".git", "dist", "node_modules", ".github/ISSUE_TEMPLATE"}
LINK = re.compile(r"\[[^\]]*\]\(([^)\s]+)\)")
HEADING = re.compile(r"^#+\s+(.*)$", re.M)
# Fenced code blocks: links inside them are examples, not real links.
FENCE = re.compile(r"^```.*?^```", re.M | re.S)


def slug(heading: str) -> str:
    """GitHub's heading anchor: lowercase, punctuation dropped, spaces to dashes."""
    text = heading.strip().lower()
    text = re.sub(r"[^\w\- ]", "", text)
    return text.replace(" ", "-")


def anchors(path: str, cache: dict[str, set[str]]) -> set[str]:
    if path not in cache:
        with open(path, encoding="utf-8") as fh:
            cache[path] = {slug(h) for h in HEADING.findall(fh.read())}
    return cache[path]


def main() -> int:
    root = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
    os.chdir(root)
    cache: dict[str, set[str]] = {}
    problems: list[str] = []
    checked = 0

    for folder, dirs, files in os.walk("."):
        dirs[:] = [d for d in dirs if d not in SKIP_DIRS and not d.startswith(".git")]
        for name in sorted(files):
            if not name.endswith(".md"):
                continue
            src = os.path.normpath(os.path.join(folder, name))
            with open(src, encoding="utf-8") as fh:
                body = FENCE.sub("", fh.read())
            for target in LINK.findall(body):
                if target.startswith(("http://", "https://", "mailto:", "#")):
                    continue
                path, _, anchor = target.partition("#")
                dest = os.path.normpath(os.path.join(folder, path)) if path else src
                checked += 1
                if not os.path.exists(dest):
                    problems.append(f"{src}: missing target {target}")
                elif anchor and dest.endswith(".md") and slug(anchor) not in anchors(dest, cache):
                    problems.append(f"{src}: no heading for {target}")

    for line in problems:
        print(line)
    print(f"checked {checked} links, {len(problems)} broken")
    return 1 if problems else 0


if __name__ == "__main__":
    sys.exit(main())
