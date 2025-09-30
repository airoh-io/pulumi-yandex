# GitHub Actions Workflows

This directory contains modernized CI/CD workflows for the Pulumi Yandex provider.

## Workflows

### 🔄 CI (`ci.yml`)
**Trigger**: Push to `master`/`main`, Pull Requests

**Purpose**: Continuous Integration
- Lints the provider code
- Builds provider binaries
- Builds Node.js SDK (primary language)
- Runs example tests (on PRs)

**Jobs**:
1. **lint** - Build and validate provider
2. **build-sdks** - Build Node.js SDK
3. **test** - Run example tests

### 🚀 Release (`release.yml`)
**Trigger**: Git tags matching `v*` (e.g., `v1.0.0`)

**Purpose**: Automated release and package publishing
- Builds provider binary
- Builds Node.js SDK
- Publishes npm package to GitHub Packages
- Creates GitHub Release with all artifacts

**Jobs**:
1. **build-provider** - Build provider binary
2. **build-sdks** - Build Node.js SDK
3. **publish-npm** - Publish to GitHub Packages
4. **create-release** - Create GitHub Release

**Note**: Python, Go, and .NET SDKs can be built manually with `make build_python`, `make build_go`, or `make build_dotnet` when needed.

### 🔍 Update Dependencies (`update-dependencies.yml`)
**Trigger**: Weekly (Monday 9 AM UTC) or manual

**Purpose**: Automated dependency update checks
- Checks for new Terraform Yandex Provider versions
- Checks for new Pulumi Bridge versions
- Creates GitHub issues when updates are available

### 🧹 Cleanup (`cleanup.yml`)
**Trigger**: Daily (midnight UTC) or manual

**Purpose**: Housekeeping
- Deletes workflow artifacts older than 30 days
- Keeps the repository clean

## Required Secrets

Configure these in GitHub Settings → Secrets and variables → Actions:

### For npm Publishing (GitHub Packages)
- `GITHUB_TOKEN` - Automatically provided by GitHub

### For PyPI Publishing (Optional)
- `PYPI_TOKEN` - Your PyPI API token
- Set repository variable `PUBLISH_PYPI=true` to enable

### For Testing (Optional)
- `PULUMI_ACCESS_TOKEN` - For running Pulumi tests
- `YC_CLOUD_ID` - Yandex Cloud ID
- `YC_FOLDER_ID` - Yandex Folder ID
- `YC_SERVICE_ACCOUNT_KEY_FILE` - Service account credentials

## Version Requirements

The workflows use these versions (defined in env vars):
- **Go**: 1.22
- **Node.js**: 18
- **Python**: 3.11
- **.NET**: 8.0

Update in each workflow file if needed.

## How to Trigger a Release

1. **Update Version**:
   ```bash
   # In provider/pkg/version/version.go
   var Version string = "1.0.0"
   ```

2. **Commit and Tag**:
   ```bash
   git add provider/pkg/version/version.go
   git commit -m "Release v1.0.0"
   git tag v1.0.0
   git push origin v1.0.0
   ```

3. **Workflow Automatically**:
   - Builds provider and all SDKs
   - Publishes npm package to GitHub Packages
   - Creates GitHub Release with downloadable artifacts

## Publishing to GitHub Packages

The npm package is automatically published to GitHub Packages on release.

### Installing from GitHub Packages

Users need to configure `.npmrc`:
```bash
@airoh-io:registry=https://npm.pkg.github.com
//npm.pkg.github.com/:_authToken=YOUR_GITHUB_TOKEN
```

Then install:
```bash
npm install @airoh-io/pulumi-yandex
```

## Publishing to PyPI (Optional)

To enable PyPI publishing:

1. **Get PyPI Token**:
   - Create account on https://pypi.org
   - Generate API token

2. **Add Secret**:
   - Go to repository Settings → Secrets
   - Add `PYPI_TOKEN` with your token

3. **Enable Publishing**:
   - Go to repository Settings → Variables
   - Add variable `PUBLISH_PYPI` with value `true`

## Local Testing

Test workflows locally with [act](https://github.com/nektos/act):

```bash
# Install act
brew install act

# Test CI workflow
act -j lint

# Test release (dry-run)
act -j build-provider -e test-event.json
```

## Migration from Old Workflows

The old workflows have been backed up to `.github/workflows.old/`.

### Key Changes:
1. **Simplified** - Removed complex Pulumi-specific infrastructure
2. **Modernized** - Updated to latest GitHub Actions versions
3. **Focused** - Tailored for fork maintenance
4. **Updated** - Go 1.22, Node 18, Python 3.11, .NET 8
5. **Package Name** - Uses `@airoh-io/pulumi-yandex`

### Removed Workflows:
- `artifact-cleanup.yml` - Replaced with simplified `cleanup.yml`
- `command-dispatch.yml` - Not needed for fork
- `community-moderation.yml` - Not needed for fork
- `main.yml` / `master.yml` - Replaced with `ci.yml`
- `nightly-test.yml` - Not needed for fork
- `prerelease.yml` - Simplified into `release.yml`
- `run-acceptance-tests.yml` - Simplified into `ci.yml`
- `update-bridge.yml` - Replaced with `update-dependencies.yml`
- `update-upstream-provider.yml` - Replaced with `update-dependencies.yml`

## Troubleshooting

### Build Fails on Python SDK
**Issue**: setuptools not found  
**Fix**: The workflow now installs setuptools automatically

### npm Publish Fails
**Issue**: Authentication error  
**Fix**: Ensure `GITHUB_TOKEN` has `write:packages` permission (automatically granted in workflows)

### Release Not Created
**Issue**: Tag doesn't trigger workflow  
**Fix**: Ensure tag format is `v*` (e.g., `v1.0.0`)

## Custom Workflows

You can add custom workflows for:
- **Automated testing with real infrastructure**
- **Documentation generation**
- **Security scanning**
- **Code coverage reports**

Place additional `.yml` files in this directory.

## Support

For issues with workflows, check:
1. Workflow run logs in GitHub Actions tab
2. [GitHub Actions Documentation](https://docs.github.com/en/actions)
3. Create an issue in the repository

---

**Last Updated**: September 30, 2025  
**Workflow Version**: 2.0 (Modernized for Fork)
