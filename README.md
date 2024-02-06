## Splight Terraform Provider

### How to install this Golang?

```
brew install go
```

### How to install this locally?

Check your architecture details with

```bash
$ terraform version
Terraform v1.4.6
on <YOUR_ARCHITECTURE>
+ provider local/splight/spl v<VERSION>

Your version of Terraform is out of date! The latest version
is 1.7.1. You can update by downloading from https://www.terraform.io/downloads.html
```

Then copy the provider to sources folder indicating `local/splight/spl` same as `source = "local/splight/spl"` in TF file.

```sh
make build
cp terraform-provider-spl_v<VERSION>  ~/.terraform.d/plugins/local/splight/spl/<VERSION>/<YOUR_ARCHITECTURE>
```

and start using it with any tf file.

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
