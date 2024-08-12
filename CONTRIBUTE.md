## Splight Terraform Provider

:warning: Please remember to update the examples and documentation for each change in the provider resources and data sources.
Keeping these up to date is crucial, as the Pulumi provider uses this documentation as input.
For more details, see [Generate docs](#generate-docs).

### Requirements

Install golang, goreleaser, terraform and delve (MacOS)

```bash
brew install go terraform delve goreleaser
```

### Installation

Run

```bash
make provider
```

and set your ~/.terraformrc as follows:

```hcl
provider_installation {
  dev_overrides {
      "splightplatform/splight" = "/Users/<you>/path/to/your/binary"
  }
  direct {}
}
```

or any path configured for your go modules.

Remove the ```.terraform``` directory and the ```.terraform.lock.hcl``` file from your workspace folder.

The try out the provider your a configuration:

```bash
terraform apply
```

Explore the examples folder for a complete file with all available resources.

### Debugging

Build the provider with debugging support:

```bash
make dlv
```

You must do this each time you want to test new changes.

This will run the provider with debugging support for delve.

When the debugger starts you will see the following output:

```bash
❯ make dlv
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

### Generate docs

To update the documentation, first manually update the examples. Then, run the following command to generate the updated docs:

```bash
make generate-docs:
```
