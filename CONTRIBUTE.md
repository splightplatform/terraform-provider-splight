## Splight Terraform Provider

### Installation

Install golang and terraform (MacOS)

```bash
brew install go terraform
```

Run

```bash
make install
```

You must do this each time you want to test new changes.

This will install the plugin inside the terraform folder of your home directory.

```bash
❯ tree ~/.terraform.d
.terraform.d
├── checkpoint_cache
├── checkpoint_signature
└── plugins
    └── local
        └── splight
            └── spl
                └── 0.1.5 <-- From the 'version' file
                    └── darwin_arm64 <-- Depending on your platform
                        └── terraform-provider-spl_v0.1.5 <-- Compiled binary

7 directories, 3 files
```

It will also delete your .terraform.lock.hcl and the symlink inside the .terraform folder in your workspace.

### Usage

Create a main.tf file as follows:

```hcl
terraform {
  required_providers {
    spl = {
      source  = "local/splight/spl"
      version = "<VERSION>"
    }
  }
}

provider "spl" {
  hostname = var.address
  token   = var.token
}

resource "spl_asset" "AssetTest" {
  name = "Asset1"
  description = "Created with Terraform"
}
```

then

```bash
terraform init
```

You can see the cached provider in your workspace pointing to the plugin we built previously.

```bash
❯ tree .terraform
.terraform
└── providers
    └── local
        └── splight
            └── spl
                └── 0.1.5
                    └── darwin_arm64 -> /Users/user/.terraform.d/plugins/local/splight/spl/0.1.5/darwin_arm64

7 directories, 0 files
```

Finally run:

```bash
terraform plan
```

Check the examples folder for a complete file with all the available resources.

### Import resources

```sh
terraform import <STATE_REFERENCE> <RESOURCE_ID>
```

Examples

```sh
terraform import -var-file variables.tfvars spl_secret.SecretImportTest 3e408b18-79df-465b-850d-6629088224de
terraform import -var-file variables.tfvars spl_asset.AssetImportTest 4e408b18-79df-465b-850d-6629088224de
```

### Generate docs

```
go generate
```
