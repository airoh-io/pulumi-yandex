# Quick Start - @airoh/pulumi-yandex

## Package Information
- **Name**: `@airoh/pulumi-yandex`
- **Repository**: https://github.com/airoh-io/pulumi-yandex
- **Resources**: 107 (46 new)
- **Data Sources**: 96 (39 new)
- **Terraform Provider**: v0.160.0
- **Pulumi Bridge**: v3.114.0

## Installation

### Option 1: GitHub Packages (Free & Recommended)

**Setup (one time):**
```bash
# Create .npmrc in your home directory
echo "@airoh:registry=https://npm.pkg.github.com" >> ~/.npmrc
echo "//npm.pkg.github.com/:_authToken=YOUR_GITHUB_TOKEN" >> ~/.npmrc
```

**Install:**
```bash
npm install @airoh/pulumi-yandex
```

### Option 2: From Git Repository (Simplest)

**In your package.json:**
```json
{
  "dependencies": {
    "@airoh/pulumi-yandex": "git+https://github.com/airoh-io/pulumi-yandex.git#master:sdk/nodejs/bin"
  }
}
```

### Option 3: Local Development

```bash
cd /path/to/pulumi-yandex/sdk/nodejs/bin
npm link

# In your project
npm link @airoh/pulumi-yandex
```

## Usage Example

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as yandex from "@airoh/pulumi-yandex";

// Create a VPC Network
const network = new yandex.VpcNetwork("my-network", {
    name: "production-network",
    description: "Production VPC",
});

// Create a Subnet
const subnet = new yandex.VpcSubnet("my-subnet", {
    networkId: network.id,
    v4CidrBlocks: ["10.0.1.0/24"],
    zone: "ru-central1-a",
});

// Use NEW resources: Lockbox Secret
const secret = new yandex.LockboxSecret("app-secret", {
    name: "my-app-secret",
    folderId: "b1g8dn6s9q5e8j7k9m0n",
});

// NEW: Monitoring Dashboard
const dashboard = new yandex.MonitoringDashboard("metrics", {
    name: "Application Metrics",
    folderId: "b1g8dn6s9q5e8j7k9m0n",
});

// NEW: Backup Policy
const backupPolicy = new yandex.BackupPolicy("vm-backup", {
    name: "daily-backup",
    compression: "NORMAL",
});

// Export outputs
export const networkId = network.id;
export const subnetId = subnet.id;
export const secretId = secret.id;
```

## New Resources Available (46 total)

### Security & Compliance
- `AuditTrailsTrail` - Audit logging
- `LockboxSecret`, `LockboxSecretVersion` - Secret management
- `SmartcaptchaCaptcha` - CAPTCHA protection
- `SwsSecurityProfile`, `SwsWafProfile`, `SwsAdvancedRateLimiterProfile` - Web security

### Backup & Recovery
- `BackupPolicy`, `BackupPolicyBindings` - VM backup management

### Certificates
- `CmCertificate` - Certificate Manager

### Compute
- `ComputeFilesystem` - Network filesystems
- `ComputeGpuCluster` - GPU clusters
- `ComputeSnapshotSchedule` - Automated snapshots

### Networking
- `VpcGateway` - NAT and egress gateways
- `VpcPrivateEndpoint` - Private service connections

### Databases
- `MdbKafkaConnector`, `MdbKafkaUser` - Kafka management
- `MdbMysqlDatabase`, `MdbMysqlUser` - MySQL management
- `MdbPostgresqlCluster`, `MdbPostgresqlDatabase`, `MdbPostgresqlUser` - PostgreSQL

### Monitoring
- `MonitoringDashboard` - Custom dashboards

### Organization
- `OrganizationmanagerGroup*` - Group management
- `OrganizationmanagerOsLoginSettings` - OS login configuration
- `OrganizationmanagerUserSshKey` - SSH key management

### Serverless
- `ServerlessEventrouterBus`, `ServerlessEventrouterConnector`, `ServerlessEventrouterRule`

### IAM
- `IamWorkloadIdentityFederatedCredential`, `IamWorkloadIdentityOidcFederation`

### YDB (Yandex Database)
- `YdbTable`, `YdbTableChangefeed`, `YdbTableIndex`, `YdbTopic`

### Testing
- `LoadtestingAgent` - Load testing

## Publishing the Package

### To GitHub Packages:

```bash
# 1. Add to sdk/nodejs/package.json:
{
  "publishConfig": {
    "registry": "https://npm.pkg.github.com"
  }
}

# 2. Rebuild
make build_nodejs

# 3. Publish
cd sdk/nodejs/bin
npm publish
```

### To Private npm Registry:

```bash
cd sdk/nodejs/bin
npm login
npm publish --access restricted  # For private packages
# or
npm publish --access public      # For public packages
```

## Building the Provider

Before using the SDK, build the provider plugin:

```bash
# Build provider binary
make provider

# The binary will be at: bin/pulumi-resource-yandex
```

## Minimal Project Setup

```bash
# Create new Pulumi project
mkdir my-yandex-project && cd my-yandex-project
pulumi new typescript -y

# Install the package
npm install @airoh/pulumi-yandex

# Configure Yandex Cloud
pulumi config set yandex:token YOUR_TOKEN
pulumi config set yandex:cloudId YOUR_CLOUD_ID
pulumi config set yandex:folderId YOUR_FOLDER_ID

# Deploy
pulumi up
```

## package.json Example

```json
{
  "name": "my-yandex-infrastructure",
  "version": "1.0.0",
  "dependencies": {
    "@pulumi/pulumi": "^3.0.0",
    "@airoh/pulumi-yandex": "^1.0.0"
  },
  "devDependencies": {
    "@types/node": "^18.0.0",
    "typescript": "^5.0.0"
  }
}
```

## Configuration

Set these environment variables or Pulumi config:

```bash
# Required
export YC_TOKEN="your-token"
export YC_CLOUD_ID="your-cloud-id"
export YC_FOLDER_ID="your-folder-id"

# Optional
export YC_ZONE="ru-central1-a"
export YC_REGION="ru-central1"
```

Or via Pulumi config:
```bash
pulumi config set yandex:token YOUR_TOKEN --secret
pulumi config set yandex:cloudId YOUR_CLOUD_ID
pulumi config set yandex:folderId YOUR_FOLDER_ID
```

## TypeScript Support

Full TypeScript definitions are included:
- ✅ IntelliSense/Autocomplete
- ✅ Type checking
- ✅ Documentation hints
- ✅ 205 type definition files

## Troubleshooting

### "Cannot find module '@airoh/pulumi-yandex'"

**For GitHub Packages:**
```bash
# Verify .npmrc is configured
cat ~/.npmrc | grep airoh

# Should show:
# @airoh:registry=https://npm.pkg.github.com
# //npm.pkg.github.com/:_authToken=...
```

**For local development:**
```bash
cd /path/to/pulumi-yandex/sdk/nodejs/bin
npm link

cd /path/to/your-project
npm link @airoh/pulumi-yandex
```

### "pulumi-resource-yandex: not found"

Build the provider:
```bash
cd /path/to/pulumi-yandex
make provider
```

The binary should be in `bin/pulumi-resource-yandex`. Pulumi will find it automatically if you run `pulumi` commands from the same directory.

## Documentation

- **Publishing Guide**: See `PUBLISHING_NPM.md`
- **Update Summary**: See `UPDATE_SUMMARY.md`
- **SDK Details**: See `SDK_GENERATION_COMPLETE.md`

## Support

- **Issues**: https://github.com/airoh-io/pulumi-yandex/issues
- **Repository**: https://github.com/airoh-io/pulumi-yandex

---

**Updated**: September 30, 2025  
**Version**: 0.1.0 (based on Terraform Provider v0.160.0)
