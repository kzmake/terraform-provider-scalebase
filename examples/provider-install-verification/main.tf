terraform {
  required_providers {
    scalebase = {
      source = "kzmake/scalebase"
    }
  }
}

provider "scalebase" {}

data "scalebase_customer" "example" {
  optional_id = "test"
}

output "example_customer" {
  value = data.scalebase_customer.example
}
