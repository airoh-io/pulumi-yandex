# Workflows Summary - @airoh-io/pulumi-yandex

## ✅ Active Workflows (4)

### 1. 🔄 `ci.yml` - Continuous Integration
- **Triggers**: Every push, every PR
- **Duration**: ~3-5 minutes
- **Actions**:
  - Lint provider code
  - Build provider binary
  - Build Node.js SDK (primary language)
  - Run tests on PRs

### 2. 🚀 `release.yml` - Automated Release
- **Triggers**: Git tags (v1.0.0, v1.0.1, etc.)
- **Duration**: ~5-8 minutes
- **Actions**:
  - Build provider binary
  - Build Node.js SDK
  - Publish `@airoh-io/pulumi-yandex` to GitHub Packages
  - Create GitHub Release with downloadable artifacts

**Note**: Python, Go, and .NET SDKs available in release assets but not auto-published.

### 3. 🔍 `update-dependencies.yml` - Dependency Monitor
- **Triggers**: Weekly (Monday 9 AM) or manual
- **Duration**: <1 minute
- **Actions**:
  - Check Terraform Provider updates
  - Check Pulumi Bridge updates
  - Create issues when updates found

### 4. 🧹 `cleanup.yml` - Artifact Cleanup
- **Triggers**: Daily (midnight) or manual
- **Duration**: <1 minute
- **Actions**:
  - Delete artifacts >30 days old

## Quick Commands

### Trigger CI
```bash
git add .
git commit -m "Update code"
git push
```

### Create Release (Publishes npm package automatically)
```bash
# 1. Update version
echo 'var Version string = "1.0.0"' > provider/pkg/version/version.go

# 2. Tag and push
git add provider/pkg/version/version.go
git commit -m "Release v1.0.0"
git tag v1.0.0
git push origin v1.0.0
```

### Build Other SDKs Locally (When Needed)
```bash
make build_python  # Build Python SDK
make build_go      # Build Go SDK
make build_dotnet  # Build .NET SDK
```

## What Happens on Release

When you push a tag like `v1.0.0`:

1. ⚙️ **Build Provider** (2 min)
   - Provider binary compiled

2. 📦 **Build Node.js SDK** (3 min)
   - TypeScript compiled
   - Ready for publishing

3. 🚀 **Publish to GitHub Packages** (1 min)
   - `@airoh-io/pulumi-yandex` published
   - Installable immediately

4. 📝 **GitHub Release** (1 min)
   - Release created
   - Provider binary attached
   - Node.js SDK attached

**Total**: ~5-8 minutes from tag push to published package!

## Primary Language: Node.js/TypeScript

This workflow is optimized for Node.js as the primary SDK:
- ✅ Built automatically on every commit
- ✅ Published automatically on release
- ✅ Available via GitHub Packages

**Other SDKs**:
- Available but not auto-published
- Can be built locally: `make build_python`, `make build_go`, `make build_dotnet`
- Included as artifacts in GitHub Releases

## Monitoring

View workflow runs:
- **GitHub → Actions tab**
- Filter by workflow name
- Check logs for any failures

## Configuration

### Required Secrets (Auto-provided)
- `GITHUB_TOKEN` - Automatically provided by GitHub

### Optional Secrets (If Needed)
- `PULUMI_ACCESS_TOKEN` - For integration tests
- `YC_*` - For Yandex Cloud tests

## Migration Complete ✅

- ✅ Old workflows backed up to `.github/workflows.old/`
- ✅ 4 new modernized workflows created
- ✅ Focused on Node.js as primary SDK
- ✅ Other SDKs buildable on demand
- ✅ Automated GitHub Packages publishing
- ✅ Dependency update monitoring

**Status**: Production Ready 🚀
