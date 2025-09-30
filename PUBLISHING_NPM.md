# Publishing @airoh-io/pulumi-yandex as Private npm Package

## Package Information
- **Name**: `@airoh-io/pulumi-yandex`
- **Current Version**: `v0.0.1-alpha.1650911783+5f483d9c.dirty`
- **Built Package Location**: `sdk/nodejs/bin/`

## Prerequisites

Before publishing, you should:

### 1. Set a Proper Version
Update the version in `provider/pkg/version/version.go`:

```go
var Version string = "1.0.0"  // or your desired version like "0.1.0"
```

Then rebuild:
```bash
make build_nodejs
```

### 2. Create npm Account (if using npm registry)
```bash
npm login
# Enter your npm credentials
```

## Option 1: npm Private Packages (Recommended)

### Requirements
- npm Teams subscription ($7/month)
- Organization created on npm (@airoh)

### Steps

1. **Create Organization on npm**
   - Go to https://www.npmjs.com/
   - Click "Add Organization"
   - Name it `airoh-io`
   - Subscribe to npm Teams

2. **Publish the Package**
   ```bash
   cd sdk/nodejs/bin
   
   # Login if not already
   npm login
   
   # Publish as private
   npm publish --access restricted
   ```

3. **Grant Access to Team Members**
   ```bash
   # Add team members in npm web interface
   # Or via CLI:
   npm team add airoh-io:developers username
   ```

4. **Install in Projects**
   ```bash
   # Team members can install with:
   npm install @airoh-io/pulumi-yandex
   ```

## Option 2: GitHub Packages (Free)

### Requirements
- GitHub account with write access to airoh-io/pulumi-yandex
- Personal Access Token (PAT)

### Steps

1. **Create GitHub Personal Access Token**
   - Go to GitHub Settings → Developer settings → Personal access tokens
   - Create token with `write:packages` and `read:packages` scopes
   - Save the token securely

2. **Configure npm to Use GitHub Packages**
   
   Create or edit `~/.npmrc`:
   ```bash
   @airoh-io:registry=https://npm.pkg.github.com
   //npm.pkg.github.com/:_authToken=YOUR_GITHUB_TOKEN
   ```

3. **Update package.json for GitHub**
   
   In `sdk/nodejs/package.json`, add:
   ```json
   {
     "publishConfig": {
       "registry": "https://npm.pkg.github.com"
     }
   }
   ```

4. **Rebuild and Publish**
   ```bash
   # Rebuild with updated config
   make build_nodejs
   
   # Publish
   cd sdk/nodejs/bin
   npm publish
   ```

5. **Install in Projects**
   
   Users need to configure their `.npmrc`:
   ```bash
   @airoh-io:registry=https://npm.pkg.github.com
   //npm.pkg.github.com/:_authToken=YOUR_GITHUB_TOKEN
   ```
   
   Then install:
   ```bash
   npm install @airoh-io/pulumi-yandex
   ```

## Option 3: Verdaccio (Self-Hosted)

### Setup Verdaccio Server

1. **Install Verdaccio**
   ```bash
   npm install -g verdaccio
   
   # Run server
   verdaccio
   # Runs on http://localhost:4873
   ```

2. **Configure Client**
   
   In your project or globally:
   ```bash
   npm set registry http://localhost:4873
   
   # Or add to package.json:
   # "publishConfig": {
   #   "registry": "http://your-verdaccio-server:4873"
   # }
   ```

3. **Create User**
   ```bash
   npm adduser --registry http://localhost:4873
   ```

4. **Publish**
   ```bash
   cd sdk/nodejs/bin
   npm publish --registry http://localhost:4873
   ```

5. **Install**
   ```bash
   npm install @airoh-io/pulumi-yandex --registry http://localhost:4873
   ```

## Option 4: Private Git Repository

### Simplest Option (No npm registry needed)

1. **Commit the Built Package**
   ```bash
   # Add build artifacts to repo (normally ignored)
   git add -f sdk/nodejs/bin
   git commit -m "Add built npm package"
   git push
   ```

2. **Install from Git**
   ```bash
   # In your Pulumi projects:
   npm install git+https://github.com/airoh-io/pulumi-yandex.git#master:sdk/nodejs/bin
   
   # Or with SSH:
   npm install git+ssh://git@github.com/airoh-io/pulumi-yandex.git#master:sdk/nodejs/bin
   
   # Or in package.json:
   {
     "dependencies": {
       "@airoh-io/pulumi-yandex": "git+https://github.com/airoh-io/pulumi-yandex.git#master:sdk/nodejs/bin"
     }
   }
   ```

## Recommended Approach

**For Production Use**: GitHub Packages (Option 2)
- ✅ Free
- ✅ Integrated with your GitHub repo
- ✅ Access control via GitHub permissions
- ✅ Easy CI/CD integration

**For Quick Testing**: Git Repository (Option 4)
- ✅ Simplest setup
- ✅ No additional services
- ✅ Works immediately

## Publishing Checklist

Before publishing:

- [ ] Update version in `provider/pkg/version/version.go`
- [ ] Rebuild SDK: `make build_nodejs`
- [ ] Test the package locally: `npm link` in bin directory
- [ ] Create a test Pulumi project and verify it works
- [ ] Update README.md with installation instructions
- [ ] Add CHANGELOG.md documenting changes
- [ ] Tag the release in git: `git tag v1.0.0`
- [ ] Publish the package
- [ ] Test installation from the registry

## Usage After Publishing

### In Pulumi TypeScript Projects

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as yandex from "@airoh-io/pulumi-yandex";

// Create resources
const vpc = new yandex.VpcNetwork("my-vpc", {
    name: "my-network",
    description: "My VPC network",
});

// Use new resources (Lockbox)
const secret = new yandex.LockboxSecret("my-secret", {
    name: "my-app-secret",
    folderId: "b1g8dn6s9q5e8j7k9m0n",
});

// Monitoring Dashboard (NEW!)
const dashboard = new yandex.MonitoringDashboard("my-dashboard", {
    name: "Application Metrics",
    folderId: "b1g8dn6s9q5e8j7k9m0n",
});

export const vpcId = vpc.id;
export const secretId = secret.id;
```

### package.json Example

```json
{
  "name": "my-pulumi-project",
  "dependencies": {
    "@pulumi/pulumi": "^3.0.0",
    "@airoh-io/pulumi-yandex": "^1.0.0"
  }
}
```

## Troubleshooting

### Issue: "Package not found"
- Verify you're logged in: `npm whoami`
- Check registry configuration: `npm config get registry`
- For GitHub Packages, verify your token has correct scopes

### Issue: "Forbidden"
- For private packages, ensure you have access to @airoh organization
- For GitHub Packages, verify your token is valid

### Issue: "Version already exists"
- Bump the version in `provider/pkg/version/version.go`
- Rebuild: `make build_nodejs`
- Publish again

## Automation with CI/CD

### GitHub Actions Example

```yaml
name: Publish npm Package

on:
  release:
    types: [created]

jobs:
  publish:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      - name: Setup Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.22'
      
      - name: Setup Node.js
        uses: actions/setup-node@v3
        with:
          node-version: '18'
          registry-url: 'https://npm.pkg.github.com'
      
      - name: Build SDK
        run: make build_nodejs
      
      - name: Publish to GitHub Packages
        run: |
          cd sdk/nodejs/bin
          npm publish
        env:
          NODE_AUTH_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

## Support

For issues or questions about the package:
- GitHub Issues: https://github.com/airoh-io/pulumi-yandex/issues
- Repository: https://github.com/airoh-io/pulumi-yandex

---

**Note**: The package includes 107 resources and 96 data sources with full TypeScript support!
