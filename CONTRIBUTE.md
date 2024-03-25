## Splight Terraform Provider

### Installation

Install golang (MacOS)

```
brew install go
```

Run

```bash
make install
```

You must run this each time you want to test new changes.

In case you have the provider already installed run

```sh
rm -rf .terraform .terraform.lock.hcl
```

### How to create a sample main.tf file.

Here is a sample TF file to test the provider

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

Just run

```sh
terraform init
terraform apply
```

If you want a complete example just go inside example folder and see a complete file with all resources

### How to import resources

```sh
terraform import <STATE_REFERENCE> <RESOURCE_ID>
```

Examples

```sh
terraform import -var-file variables.tfvars spl_secret.SecretImportTest 3e408b18-79df-465b-850d-6629088224de
terraform import -var-file variables.tfvars spl_asset.AssetImportTest 4e408b18-79df-465b-850d-6629088224de
```

### How to regenerate documantation

```
go generate
```
