terraform {
  required_providers {
    scalebase = {
      source = "kzmake/scalebase"
    }
  }
}

provider "scalebase" {}

locals {
  csv_data  = file("${path.module}/resources.csv")
  resources = csvdecode(local.csv_data)
  contracts = toset([for each in local.resources : each if each["リソース種別"] == "contract"])
  customers = toset([for each in local.resources : each if each["リソース種別"] == "customer"])
}

data "scalebase_contract" "example" {
  for_each    = { for v in local.contracts : v["リソース管理ID"] => v }
  optional_id = each.value["リソース管理ID"]
}

data "scalebase_customer" "example" {
  for_each    = { for v in local.customers : v["リソース管理ID"] => v }
  optional_id = each.value["リソース管理ID"]
}


data "scalebase_custom_field_master" "contract_cfm1" {
  name       = "文字列"
  field_type = "contract"
}

data "scalebase_custom_field_master" "contract_cfm2" {
  name       = "選択リスト"
  field_type = "contract"
}

data "scalebase_custom_field_master" "contract_cfm3" {
  name       = "日付"
  field_type = "contract"
}

data "scalebase_custom_field_master" "customer_cfm1" {
  name       = "文字列"
  field_type = "customer"
}

data "scalebase_custom_field_master" "customer_cfm2" {
  name       = "選択リスト"
  field_type = "customer"
}

data "scalebase_custom_field_master" "customer_cfm3" {
  name       = "日付"
  field_type = "customer"
}

resource "scalebase_resource" "contract_example" {
  for_each = { for v in local.contracts : v["リソース管理ID"] => v }

  srn = {
    type = "contract"
    id   = data.scalebase_contract.example[each.value["リソース管理ID"]].id
  }
  custom_fields = [
    {
      id     = data.scalebase_custom_field_master.contract_cfm1.id
      string = each.value["文字列"] != "" ? each.value["文字列"] : null
    },
    {
      id     = data.scalebase_custom_field_master.contract_cfm2.id
      select = each.value["選択リスト"] != "" ? each.value["選択リスト"] : null
    },
    {
      id   = data.scalebase_custom_field_master.contract_cfm3.id
      date = each.value["日付"] != "" ? each.value["日付"] : null
    }
  ]
}

resource "scalebase_resource" "customer_example" {
  for_each = { for v in local.customers : v["リソース管理ID"] => v }

  srn = {
    type = "customer"
    id   = data.scalebase_customer.example[each.value["リソース管理ID"]].id
  }
  custom_fields = [
    {
      id         = data.scalebase_custom_field_master.customer_cfm1.id
      string = each.value["文字列"] != "" ? each.value["文字列"] : null
    },
    {
      id     = data.scalebase_custom_field_master.customer_cfm2.id
      select = each.value["選択リスト"] != "" ? each.value["選択リスト"] : null
    },
    {
      id   = data.scalebase_custom_field_master.customer_cfm3.id
      date = each.value["日付"] != "" ? each.value["日付"] : null
    }
  ]
}
