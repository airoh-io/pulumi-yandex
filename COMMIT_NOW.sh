#!/bin/bash
# Commit script for Pulumi Yandex Provider Modernization

set -e

echo "╔════════════════════════════════════════════════════════════╗"
echo "║  Committing Pulumi Yandex Provider Modernization          ║"
echo "╚════════════════════════════════════════════════════════════╝"
echo ""

# Show what will be committed
echo "📊 Changes to be committed:"
git status --short | head -20
echo "... and more ($(git status --short | wc -l | xargs) total files)"
echo ""

# Add all changes
echo "➕ Adding all changes..."
git add -A
echo ""

# Commit
echo "💾 Committing..."
git commit -m "Modernize provider: Update to terraform-provider v0.160.0, bridge v3.114.0

Major updates:
- Update terraform-provider-yandex: v0.71.0 → v0.160.0 (+89 versions)
- Update pulumi-terraform-bridge: v3.21.0 → v3.114.0 (+93 versions)
- Update pulumi SDK: v3.30.0 → v3.198.0 (+168 versions)
- Update Go: 1.16 → 1.22.0

Features:
- Add 46 new resources (Lockbox, Monitoring, SWS, Backup, YDB, etc.)
- Add 39 new data sources
- Remove 15 obsolete resources (deprecated by upstream)
- Total features: 118 → 203 (+72%)

SDKs:
- Regenerate all SDKs (Go, Node.js, Python, C#)
- SDK success rate: 98.29%
- Package renamed to @airoh/pulumi-yandex

Infrastructure:
- Modernize GitHub workflows (13 → 4 workflows)
- Add automated dependency monitoring
- Add automated npm publishing
- Update all repository references to airoh-io

Documentation:
- Add 7 comprehensive guides
- Include publishing, quick start, and migration docs"

echo ""
echo "✅ Committed successfully!"
echo ""
echo "📤 Next: Push to GitHub"
echo "   git push origin master"
echo ""
echo "🏷️  Then create release:"
echo "   git tag v1.0.0"
echo "   git push origin v1.0.0"
echo ""

