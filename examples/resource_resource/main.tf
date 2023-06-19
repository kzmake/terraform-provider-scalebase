terraform {
  required_providers {
    scalebase = {
      source = "kzmake/scalebase"
    }
  }
}

provider "scalebase" {}

data "scalebase_customer" "example" {
  optional_id = "example"
}

data "scalebase_custom_field_master" "str" {
  name       = "文字列"
  field_type = "customer"
}

data "scalebase_custom_field_master" "select" {
  name       = "選択リスト"
  field_type = "customer"
}

data "scalebase_custom_field_master" "date" {
  name       = "日付"
  field_type = "customer"
}

resource "scalebase_resource" "example" {
  srn = {
    type = "customer"
    id   = data.scalebase_customer.example.id
  }
  custom_fields = [
    {
      id     = data.scalebase_custom_field_master.str.id
      string = "aaaa"
    },
    {
      id     = data.scalebase_custom_field_master.select.id
      select = "cccc"
    },
    {
      id   = data.scalebase_custom_field_master.date.id
      date = "2023-01-03"
    }
  ]
}
