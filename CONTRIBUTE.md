## Splight Terraform Provider

### Installation

Install golang and terraform (MacOS)

```
brew install go terraform
```

Run

```bash
make install
```

You must do this each time you want to test new changes.

If you already have the provider installed from the registry, delete it from the lockfile with:

```sh
rm -rf .terraform .terraform.lock.hcl
```

### Usage

Create a main.tf file as follows:

```
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

```sh
terraform init
terraform apply
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
