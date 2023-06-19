terraform {
  required_providers {
    scalebase = {
      source = "kzmake/scalebase"
    }
  }
}

provider "scalebase" {}

data "scalebase_contract" "example" {
  optional_id = "example"
}

data "scalebase_customer" "example" {
  optional_id = "example"
}

data "scalebase_custom_field_master" "contract_string" {
  name       = "文字列"
  field_type = "contract"
}

data "scalebase_custom_field_master" "contract_select" {
  name       = "選択リスト"
  field_type = "contract"
}

data "scalebase_custom_field_master" "contract_date" {
  name       = "日付"
  field_type = "contract"
}

data "scalebase_custom_field_master" "customer_string" {
  name       = "文字列"
  field_type = "customer"
}

data "scalebase_custom_field_master" "customer_select" {
  name       = "選択リスト"
  field_type = "customer"
}

data "scalebase_custom_field_master" "customer_date" {
  name       = "日付"
  field_type = "customer"
}

resource "scalebase_resource" "contract_example" {
  srn = {
    type = "contract"
    id   = data.scalebase_contract.example.id
  }
  custom_fields = [
    {
      id     = data.scalebase_custom_field_master.contract_string.id
      string = "もじれつ"
    },
    {
      id     = data.scalebase_custom_field_master.contract_select.id
      select = "せんたく"
    },
    {
      id   = data.scalebase_custom_field_master.contract_date.id
      date = "2023-01-03"
    }
  ]
}

resource "scalebase_resource" "customer_example" {
  srn = {
    type = "customer"
    id   = data.scalebase_customer.example.id
  }
  custom_fields = [
    {
      id     = data.scalebase_custom_field_master.customer_string.id
      string = "もじれつ"
    },
    {
      id     = data.scalebase_custom_field_master.customer_select.id
      select = "せんたく"
    },
    {
      id   = data.scalebase_custom_field_master.customer_date.id
      date = "2023-01-03"
    }
  ]
}
