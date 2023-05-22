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

resource "scalebase_customer_staff" "example" {
  customer_id = data.scalebase_customer.example.id
  optional_id = "example"
  name = {
    first_name = "<俺、𩸽"
    last_name  = "🐟"
  }
  email_address = "example+hoge.fuga@example.com"
  phone_number  = "03-1111-2222"
  department    = "日本三名園"
  title         = "兼六園"
  address = {
    zip_code      = "920-0936"
    country       = "COUNTRY_JP"
    prefecture    = "PREFECTURE_ISHIKAWA"
    city          = "金沢市"
    address_lines = ["兼六町１", "徽軫灯籠の霞ヶ池"]
  }
}
