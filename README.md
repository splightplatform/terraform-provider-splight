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
+ provider local/splight/spl v0.1.0

Your version of Terraform is out of date! The latest version
is 1.7.1. You can update by downloading from https://www.terraform.io/downloads.html
```

Then copy the provider to sources folder indicating `local/splight/spl` same as `source = "local/splight/spl"` in TF file.

```bash
make build
cp terraform-provider-spl_v0.1.0  ~/.terraform.d/plugins/local/splight/spl/0.1.0/<YOUR_ARCHITECTURE>
```

and start using it with any tf file.

In case you have the provider already installed run

```
rm -rf .terraform .terraform.lock.hcl
```

### How to create a sample main.tf file.

Here is a sample TF file to test the provider

```
terraform {
  required_providers {
    spl = {
      source  = "local/splight/spl"
      version = "0.1.0"
    }
  }
}

provider "spl" {
  address = var.address
  port    = var.port
  token   = var.token
}

resource "spl_asset" "AssetTest" {
  name = "Asset1"
  description = "Created with Terraform"
}
```

Just run

```
terraform init
terraform apply
```
