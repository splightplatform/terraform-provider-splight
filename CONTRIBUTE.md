## Splight Terraform Provider

### Installation

Install golang, terraform and delve (MacOS)

```bash
brew install go terraform delve
```

Run

```bash
make debug-start
```

You must do this each time you want to test new changes.

This will run the provider with debugging support for delve

When the debugger starts you will see the following output:

```bash
❯ make debug-start
go mod tidy
gofmt -w .
go build -gcflags="all=-N -l" -o terraform-provider-spl_$(terraform version | grep -o '^on [^\s]\+' | cut -d ' ' -f2)_$(cat version)_debug
dlv exec terraform-provider-spl_$(terraform version | grep -o '^on [^\s]\+' | cut -d ' ' -f2)_$(cat version)_debug -- -debug
Type 'help' for list of commands.
(dlv)
```

Input 'continue' or 'c' to start the server:

```
(dlv) c
{"@level":"debug","@message":"plugin address","@timestamp":"2024-04-08T14:30:09.697517-03:00","address":"/var/folders/p0/8t8w97s96xb7gy2rtmkc_w4m0000gn/T/plugin3729862057","network":"unix"}
Provider started. To attach Terraform CLI, set the TF_REATTACH_PROVIDERS environment variable with the following:

	TF_REATTACH_PROVIDERS=<output>
```

Copy the env var and try applying changes:

```
	TF_REATTACH_PROVIDERS=<output> terraform apply
```

### Usage

To utilize the Splight Terraform Provider, create a ```main.tf``` file with the following content:

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

and a ```variables.tf``` file

```hcl
variable "spl_secret" {
  type      = string
  sensitive = true
}

variable "spl_api_token" {
  type      = string
  sensitive = true
}

variable "spl_hostname" {
  type = string

}
```

then

```bash
terraform init
```

You can see the cached provider in your workspace pointing to the plugin we built previously.

```
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

Finally run

```bash
terraform plan
```

Explore the examples folder for a complete file with all available resources.

### Import resources

```bash
terraform import <STATE_REFERENCE> <RESOURCE_ID>
```

Examples

```bash
terraform import -var-file variables.tfvars spl_secret.SecretImportTest 3e408b18-79df-465b-850d-6629088224de
terraform import -var-file variables.tfvars spl_asset.AssetImportTest 4e408b18-79df-465b-850d-6629088224de
```

### Generate docs

```bash
go generate
```