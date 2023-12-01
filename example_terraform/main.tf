terraform {
  required_providers {
    person = {
      version = "~> 1.0.0"
      source  = "kiril.com/personprovider/person"
    }
  }
}

provider "person" {
    session_id = "371705ef-e836-4bfd-bfed-d67a5a17188d"
}

resource "person" "person1" {
  name = "kiril"
  surname = "dzajkovski"
}
