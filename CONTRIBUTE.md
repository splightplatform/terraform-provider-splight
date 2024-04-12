## Splight Terraform Provider


### Requirements

Install golang, terraform and delve (MacOS)

```bash
brew install go terraform delve
```

### Installation

Run

```bash
make install
```

and set your ~/.terraformrc as follows:

```hcl
provider_installation {
  dev_overrides {
      "splightplatform/splight" = "/Users/<you>/go/bin/"
  }
  direct {}
}
```

or any path configured for your go modules.

Remove the ```.terraform``` directory and the ```.terraform.lock.hcl``` file from your workspace folder.

The try out the provider your a configuration:

```bash
terraform init
terraform apply
```

Explore the examples folder for a complete file with all available resources.

### Debugging

Build the provider with debugging support:

```bash
make debug-start
```

You must do this each time you want to test new changes.

This will run the provider with debugging support for delve.

When the debugger starts you will see the following output:

```bash
❯ make debug-start
Type 'help' for list of commands.
(dlv)
```

Input ```continue``` or ```c``` to start the server:

```
(dlv) c
Provider started. To attach Terraform CLI, set the TF_REATTACH_PROVIDERS environment variable with the following:

	TF_REATTACH_PROVIDERS=<output>
```

Copy the enviroment variable and try applying changes:

```bash
TF_REATTACH_PROVIDERS=<output> terraform apply
```

Don't forget to set breakpoints in your code with:

```go
runtime.Breakpoint()
```

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
