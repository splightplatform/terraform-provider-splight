## Splight Terraform Provider

:warning: Please remember to update the examples and documentation for each change in the provider resources and data sources.
Keeping these up to date is crucial, as the Pulumi provider uses this documentation as input.
For more details, see [Generate docs](#generate-docs).

### Requirements

Install golang, terraform and delve (MacOS)

```bash
brew install go terraform delve
```

### Installation

Run

```bash
make
```

and set your ~/.terraformrc as follows:

```hcl
provider_installation {
  dev_overrides {
      "splightplatform/splight" = "/Users/<you>/path/to/your/binary/terraform-provider-splight"
  }
  direct {}
}
```

The try out the provider your configuration:

```bash
terraform init
terraform apply
```

Explore the examples folder for a complete file with all available resources.

### Debugging

Build the provider with debugging support:

```bash
make debug
```

You must do this each time you want to test new changes.

This will run the provider with debugging support for delve.

When the debugger starts you will see the following output:

```bash
❯ make debug
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

You can set breakpoints in your code with:

```go
runtime.Breakpoint()
```

### Generate docs

To update the documentation, first manually update the examples. Then, run the following command to generate the updated docs:

```bash
make docs
```
