terraform {
  required_providers {
    scalebase = {
      source = "kzmake/scalebase"
    }
  }
}

provider "scalebase" {}

locals {
  csv_data  = file("${path.module}/customers.csv")
  customers = csvdecode(local.csv_data)
}

resource "scalebase_customer" "example" {
  for_each    = { for customer in local.customers : customer.optional_id => customer }
  optional_id = each.value.optional_id
  name        = each.value.name
}
