#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
DEFAULT_REPO="git@github.com:1378496782/goframe-micro-shop.git"

GITHUB_REPO="${GITHUB_REPO:-$DEFAULT_REPO}"
PUBLISH_DIR="${PUBLISH_DIR:-$ROOT_DIR/.github-publish/goframe-micro-shop}"
COMMIT_MSG="chore: publish sanitized project snapshot"
DO_PUSH=0

usage() {
  cat <<'EOF'
Usage:
  scripts/publish-github.sh [options]

Options:
  -m, --message <msg>       Commit message.
      --push                Push to GitHub after commit.
      --repo <git-url>      GitHub repository URL.
      --publish-dir <dir>   Local publish worktree directory.
  -h, --help                Show help.

Examples:
  scripts/publish-github.sh -m "docs: update project readme"
  scripts/publish-github.sh -m "feat: update order flow" --push
EOF
}

while [[ $# -gt 0 ]]; do
  case "$1" in
    -m|--message)
      COMMIT_MSG="${2:?missing commit message}"
      shift 2
      ;;
    --push)
      DO_PUSH=1
      shift
      ;;
    --repo)
      GITHUB_REPO="${2:?missing repository URL}"
      shift 2
      ;;
    --publish-dir)
      PUBLISH_DIR="${2:?missing publish directory}"
      shift 2
      ;;
    -h|--help)
      usage
      exit 0
      ;;
    *)
      echo "Unknown option: $1" >&2
      usage >&2
      exit 2
      ;;
  esac
done

need_cmd() {
  if ! command -v "$1" >/dev/null 2>&1; then
    echo "Missing required command: $1" >&2
    exit 1
  fi
}

need_cmd git
need_cmd rsync
need_cmd python3

echo "Project root: $ROOT_DIR"
echo "Publish dir : $PUBLISH_DIR"
echo "GitHub repo : $GITHUB_REPO"

if [[ ! -d "$PUBLISH_DIR/.git" ]]; then
  mkdir -p "$(dirname "$PUBLISH_DIR")"
  git clone "$GITHUB_REPO" "$PUBLISH_DIR"
fi

git -C "$PUBLISH_DIR" checkout main >/dev/null
git -C "$PUBLISH_DIR" pull --ff-only origin main

echo "Syncing source files..."
rsync -a --delete \
  --exclude '.git/' \
  --exclude '.github-publish/' \
  --exclude 'node_modules/' \
  --exclude 'dist/' \
  --exclude 'build/' \
  --exclude 'unpackage/' \
  --exclude '.qoder/' \
  --exclude '.idea/' \
  --exclude '.vscode/' \
  --exclude '.DS_Store' \
  --exclude '*.log' \
  --exclude '*.tmp' \
  --exclude '*.bak' \
  --exclude '*~' \
  --exclude '*.exe' \
  --exclude '*.exe~' \
  --exclude 'project.private.config.json' \
  --exclude 'gfast-app' \
  --exclude '.env' \
  --exclude '.env.*' \
  "$ROOT_DIR/" "$PUBLISH_DIR/"

echo "Removing files that should not be published..."
python3 - "$PUBLISH_DIR" <<'PY'
from pathlib import Path
import shutil
import sys

root = Path(sys.argv[1]).resolve()
skip_dirs = {".git"}
remove_dirs = {
    ".github-publish",
    ".idea",
    ".qoder",
    ".vscode",
    "build",
    "dist",
    "node_modules",
    "unpackage",
}
remove_exact_files = {
    ".DS_Store",
    ".env",
    "gfast-app",
    "project.private.config.json",
}
remove_suffixes = {
    ".bak",
    ".exe",
    ".log",
    ".tmp",
}

removed = []

def is_under_skip(path: Path) -> bool:
    return any(part in skip_dirs for part in path.relative_to(root).parts)

for path in sorted(root.rglob("*"), key=lambda item: len(item.parts), reverse=True):
    if is_under_skip(path):
        continue
    if path.is_dir() and path.name in remove_dirs:
        shutil.rmtree(path)
        removed.append(str(path.relative_to(root)) + "/")
        continue
    if not path.is_file():
        continue
    name = path.name
    if (
        name in remove_exact_files
        or name.endswith("~")
        or path.suffix.lower() in remove_suffixes
        or (name.startswith(".env.") and name != ".env.example")
    ):
        path.unlink()
        removed.append(str(path.relative_to(root)))

if removed:
    print("Removed publish-only excluded paths:")
    for item in removed[:80]:
        print(f"  - {item}")
    if len(removed) > 80:
        print(f"  ... {len(removed) - 80} more")
else:
    print("No excluded publish paths needed removal.")
PY

echo "Sanitizing public copy..."
python3 - "$PUBLISH_DIR" <<'PY'
from pathlib import Path
import re
import sys

root = Path(sys.argv[1]).resolve()

skip_dirs = {
    ".git",
    ".github-publish",
    "node_modules",
    "dist",
    "build",
    "unpackage",
    ".qoder",
}

skip_files = {
    "scripts/publish-github.sh",
}

global_regex_replacements = [
    (
        re.compile(rb"(?m)^([ \t]*privateKey:[ \t]*\|[ \t]*\n)(?:[ \t]+.*\n)+"),
        rb"\1    CHANGE_ME_WECHAT_PRIVATE_KEY\n",
    ),
    (
        re.compile(rb"(mysql:[^:\s\"']+:)[^@\s\"']+(@tcp)"),
        rb"\1CHANGE_ME_MYSQL_PASSWORD\2",
    ),
    (
        re.compile(rb"(-u[a-zA-Z0-9_]+[ \t]+-p)[^ \t\r\n\\]+"),
        rb"\1CHANGE_ME_MYSQL_PASSWORD",
    ),
    (
        re.compile(rb"(JWTSecretKey[ \t]*=[ \t]*[\"'])(?!CHANGE_ME_)[^\"'\r\n]+([\"'])"),
        rb"\1CHANGE_ME_JWT_SECRET\2",
    ),
    (
        re.compile(rb"(token=)[A-Za-z0-9%+/=_-]{24,}"),
        rb"\1DEMO_TOKEN_REDACTED",
    ),
    (
        re.compile(rb"([\"']token[\"'][ \t]*:[ \t]*[\"'])(?!CHANGE_ME_|DEMO_)[A-Za-z0-9+/=_-]{24,}([\"'])"),
        rb"\1DEMO_TOKEN_REDACTED\2",
    ),
    (
        re.compile(rb"7ZUSfV[A-Za-z0-9%+/=_-]{20,}"),
        b"DEMO_TOKEN_REDACTED",
    ),
    (
        re.compile(rb"cps_key=[a-f0-9]{32}"),
        b"cps_key=DEMO_CPS_KEY",
    ),
    (
        re.compile(rb"GF_SECURITY_ADMIN_PASSWORD=[^\r\n]+"),
        b"GF_SECURITY_ADMIN_PASSWORD=CHANGE_ME_GRAFANA_ADMIN_PASSWORD",
    ),
]

config_regex_replacements = [
    (
        re.compile(rb"(?im)^([ \t]*[A-Za-z0-9_.-]*(?:password|passwd|secret|token|accessKey|secretKey|apiV3Key|mchId|serialNo|appId|aesKey|encryptKey|encodingAESKey)[A-Za-z0-9_.-]*[ \t]*:[ \t]*)([\"']?)(?!CHANGE_ME_|DEMO_)[^\"'\r\n#]+([\"']?)"),
        rb"\1\2CHANGE_ME_SECRET\3",
    ),
    (
        re.compile(rb"([\"'](?:password|passwd|secret|token|accessKey|secretKey|apiV3Key|mchId|serialNo|appId|aesKey|encryptKey|encodingAESKey)[\"'][ \t]*:[ \t]*[\"'])(?!CHANGE_ME_|DEMO_)[^\"'\r\n]+([\"'])", re.I),
        rb"\1CHANGE_ME_SECRET\2",
    ),
    (
        re.compile(rb"(?im)^([A-Z0-9_]*(?:PASSWORD|PASSWD|SECRET|TOKEN|KEY)[A-Z0-9_]*=)(?!CHANGE_ME_|DEMO_)[^\r\n]+"),
        rb"\1CHANGE_ME_SECRET",
    ),
]

config_suffixes = {
    ".conf",
    ".env",
    ".ini",
    ".json",
    ".md",
    ".properties",
    ".sql",
    ".toml",
    ".txt",
    ".yaml",
    ".yml",
}

def should_skip(path: Path) -> bool:
    rel_parts = path.relative_to(root).parts
    rel_path = "/".join(rel_parts)
    return rel_path in skip_files or any(part in skip_dirs for part in rel_parts)

def is_config_like(path: Path) -> bool:
    name = path.name.lower()
    if path.suffix.lower() in config_suffixes:
        return True
    return name in {"dockerfile", "makefile"}

changed = []
for path in root.rglob("*"):
    if not path.is_file() or should_skip(path):
        continue
    try:
        data = path.read_bytes()
    except OSError:
        continue

    # Avoid rewriting large generated/binary assets. Config/code/sql docs are small text files.
    if len(data) > 5 * 1024 * 1024 or b"\0" in data[:4096]:
        continue

    updated = data
    for pattern, replacement in global_regex_replacements:
        updated = pattern.sub(replacement, updated)
    if is_config_like(path):
        for pattern, replacement in config_regex_replacements:
            updated = pattern.sub(replacement, updated)

    if updated != data:
        path.write_bytes(updated)
        changed.append(str(path.relative_to(root)))

if changed:
    print("Sanitized files:")
    for item in changed[:80]:
        print(f"  - {item}")
    if len(changed) > 80:
        print(f"  ... {len(changed) - 80} more")
else:
    print("No known secret placeholders needed replacement.")

blocked_regexes = [
    re.compile(rb"AKIA[0-9A-Z]{16}"),
    re.compile(rb"sk-[A-Za-z0-9]{20,}"),
    re.compile(rb"xox[baprs]-[A-Za-z0-9-]{20,}"),
    re.compile(rb"-----BEGIN [A-Z ]*PRIVATE KEY-----"),
]

hits = []
for path in root.rglob("*"):
    if not path.is_file() or should_skip(path):
        continue
    try:
        data = path.read_bytes()
    except OSError:
        continue
    if len(data) > 5 * 1024 * 1024 or b"\0" in data[:4096]:
        continue

    rel = str(path.relative_to(root))
    for pattern in blocked_regexes:
        if pattern.search(data):
            hits.append(f"{rel}: matched blocked regex {pattern.pattern.decode('ascii', 'ignore')}")

if hits:
    print("\nPotential secrets still exist in the publish copy:", file=sys.stderr)
    for hit in hits[:50]:
        print(f"  - {hit}", file=sys.stderr)
    if len(hits) > 50:
        print(f"  ... {len(hits) - 50} more", file=sys.stderr)
    sys.exit(1)

print("Secret scan passed.")
PY

echo "Checking publish diff..."
git -C "$PUBLISH_DIR" status --short

git -C "$PUBLISH_DIR" add -A
if git -C "$PUBLISH_DIR" diff --cached --quiet; then
  echo "No public changes to commit."
else
  git -C "$PUBLISH_DIR" commit --quiet -m "$COMMIT_MSG"
  echo "Committed public snapshot: $COMMIT_MSG"
fi

if [[ "$DO_PUSH" -eq 1 ]]; then
  git -C "$PUBLISH_DIR" push origin main
  echo "Published to GitHub."
else
  echo
  echo "Publish copy is ready but not pushed."
  echo "Review it with:"
  echo "  git -C \"$PUBLISH_DIR\" status"
  echo "  git -C \"$PUBLISH_DIR\" diff --stat origin/main...HEAD"
  echo
  echo "Push when ready:"
  echo "  git -C \"$PUBLISH_DIR\" push origin main"
fi
