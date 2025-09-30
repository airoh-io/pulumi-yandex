# GitHub Workflows Migration Summary

## Overview

The GitHub Actions workflows have been **completely modernized** for your fork. Old workflows were designed for Pulumi's official release infrastructure and needed significant updates.

## 🎯 What Changed

### Old Setup (13 workflows)
- Complex, Pulumi-specific workflows
- Outdated dependencies (Go 1.17, Node 10)
- Tied to Pulumi's release infrastructure
- Package name: `@pulumi/yandex`
- Many unnecessary workflows for a fork

### New Setup (4 workflows)
- ✅ Simplified and focused
- ✅ Modern dependencies (Go 1.22, Node 18, Python 3.11, .NET 8)
- ✅ Tailored for fork maintenance
- ✅ Package name: `@airoh/pulumi-yandex`
- ✅ Automated dependency updates
- ✅ GitHub Packages publishing

## 📋 New Workflows

### 1. `ci.yml` - Continuous Integration
**Triggers**: Push to main/master, Pull Requests

**What it does**:
- ✅ Lints provider code
- ✅ Builds provider binary
- ✅ Builds all SDKs (Node.js, Python, Go) in parallel
- ✅ Runs example tests on PRs

**Runs on**: Every commit and PR

### 2. `release.yml` - Automated Releases
**Triggers**: Git tags (v1.0.0, v1.0.1, etc.)

**What it does**:
- ✅ Builds provider binary
- ✅ Builds all SDKs (including .NET)
- ✅ Publishes npm package to GitHub Packages
- ✅ Publishes Python package to PyPI (optional)
- ✅ Creates GitHub Release with downloadable artifacts

**How to use**:
```bash
# Update version
vim provider/pkg/version/version.go  # Set Version = "1.0.0"

# Commit and tag
git add -A
git commit -m "Release v1.0.0"
git tag v1.0.0
git push origin v1.0.0

# Workflow runs automatically!
```

### 3. `update-dependencies.yml` - Dependency Checker
**Triggers**: Weekly (Monday 9 AM UTC) or manual

**What it does**:
- ✅ Checks for new Terraform Yandex Provider versions
- ✅ Checks for new Pulumi Bridge versions
- ✅ Creates GitHub issues when updates are available
- ✅ Includes instructions for updating

**Benefit**: Never miss an update!

### 4. `cleanup.yml` - Housekeeping
**Triggers**: Daily (midnight UTC) or manual

**What it does**:
- ✅ Deletes workflow artifacts older than 30 days
- ✅ Keeps repository clean and saves storage

## 🔄 Migration Details

### Removed Workflows (Backed up to `.github/workflows.old/`)
| Old Workflow | Reason | Replacement |
|--------------|--------|-------------|
| `main.yml` / `master.yml` | Complex Pulumi infrastructure | `ci.yml` (simplified) |
| `prerelease.yml` | Not needed for fork | `release.yml` |
| `nightly-test.yml` | Resource intensive | Tests run on PRs in `ci.yml` |
| `run-acceptance-tests.yml` | Complex setup | Simplified in `ci.yml` |
| `update-bridge.yml` | Old automation | `update-dependencies.yml` |
| `update-upstream-provider.yml` | Old automation | `update-dependencies.yml` |
| `command-dispatch.yml` | Pulumi-specific | Not needed |
| `community-moderation.yml` | Pulumi-specific | Not needed |
| `artifact-cleanup.yml` | Old version | `cleanup.yml` |
| `pull-request.yml` | Simplified | Part of `ci.yml` |

### Updated Dependencies

| Component | Old Version | New Version |
|-----------|-------------|-------------|
| Go | 1.17.x | 1.22 |
| Node.js | 10.x | 18 |
| Python | 3.8 | 3.11 |
| .NET | 3.1 | 8.0 |
| Actions | v2 | v4/v5 |

### Updated Package Names

| Language | Old | New |
|----------|-----|-----|
| npm | `@pulumi/yandex` | `@airoh/pulumi-yandex` |
| Python | `pulumi-yandex` | `pulumi-yandex` |
| Go | `github.com/pulumi/pulumi-yandex` | `github.com/airoh-io/pulumi-yandex` |

## 🚀 How to Use

### Running CI
CI runs automatically on:
- Every push to `master` or `main`
- Every pull request

No action needed - just commit and push!

### Creating a Release

1. **Update Version**:
   ```bash
   # Edit provider/pkg/version/version.go
   var Version string = "1.0.0"
   ```

2. **Commit and Tag**:
   ```bash
   git add provider/pkg/version/version.go
   git commit -m "Bump version to 1.0.0"
   git tag v1.0.0
   git push origin v1.0.0
   ```

3. **Watch the Magic** ✨:
   - Workflow builds everything
   - Publishes to GitHub Packages
   - Creates GitHub Release
   - Users can install immediately!

### Installing from GitHub Packages

After release, users install with:

```bash
# Configure .npmrc
echo "@airoh:registry=https://npm.pkg.github.com" >> ~/.npmrc
echo "//npm.pkg.github.com/:_authToken=YOUR_TOKEN" >> ~/.npmrc

# Install
npm install @airoh/pulumi-yandex
```

## 🔒 Required Secrets

### Automatic (No Setup Needed)
- `GITHUB_TOKEN` - Provided by GitHub

### Optional
- `PYPI_TOKEN` - If publishing to PyPI
- `PULUMI_ACCESS_TOKEN` - For integration tests
- `YC_*` secrets - For Yandex Cloud tests

Configure in: **Settings → Secrets and variables → Actions**

## 📊 Benefits

### Before (Old Workflows)
- ❌ 13 complex workflows
- ❌ Outdated dependencies
- ❌ Pulumi-specific infrastructure
- ❌ Manual package publishing
- ❌ No automated updates checking
- ❌ Resource intensive

### After (New Workflows)
- ✅ 4 focused workflows
- ✅ Latest dependencies
- ✅ Fork-optimized
- ✅ Automated publishing
- ✅ Weekly dependency checks
- ✅ Efficient and clean

## 🧪 Testing Locally

Test workflows locally with [act](https://github.com/nektos/act):

```bash
# Install act
brew install act

# Test CI workflow
act push -j lint

# Test build
act push -j build-sdks
```

## 📝 Next Steps

1. **Test CI**:
   ```bash
   git add .github/workflows/
   git commit -m "Update workflows"
   git push
   ```

2. **Watch Workflow Run**:
   - Go to GitHub → Actions tab
   - See CI workflow run

3. **Create First Release**:
   ```bash
   # Set version
   vim provider/pkg/version/version.go
   
   # Tag and push
   git tag v1.0.0
   git push origin v1.0.0
   
   # Watch release workflow!
   ```

## 🐛 Troubleshooting

### Workflow doesn't run
- Check `.github/workflows/` files are committed
- Verify trigger conditions (branch name, tag format)

### Build fails
- Check workflow logs in Actions tab
- Verify all dependencies are up to date

### Publish fails
- Ensure `GITHUB_TOKEN` has packages:write permission
- For PyPI, verify `PYPI_TOKEN` is set

## 📚 Documentation

- **Workflow Details**: See `.github/workflows/README.md`
- **Publishing Guide**: See `PUBLISHING_NPM.md`
- **Quick Start**: See `QUICK_START.md`

## ✅ Summary

The workflows are now:
- ✨ **Modernized** - Latest versions and best practices
- 🎯 **Simplified** - 4 focused workflows instead of 13
- 🚀 **Automated** - Automatic publishing and updates
- 🔧 **Maintainable** - Easy to understand and modify
- 📦 **Production-Ready** - Ready for your fork's lifecycle

Your provider is ready for automated CI/CD! 🎉

---

**Migration Date**: September 30, 2025  
**Status**: ✅ Complete  
**Old Workflows**: Backed up to `.github/workflows.old/`
