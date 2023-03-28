terraform {
  required_providers {
    scalebase = {
      source = "kzmake/scalebase"
    }
  }
}

provider "scalebase" {}

resource "scalebase_customer" "example" {
  optional_id = "example"
  name        = "example"
}
