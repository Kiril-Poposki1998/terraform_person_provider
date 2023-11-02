terraform {
  required_providers {
    person = {
      version = "~> 1.0.0"
      source  = "kiril.com/personprovider/person"
    }
  }
}

provider "person" {
    session_id = ""
}

resource "person" "person1" {
  person_id = "1"
  name = "kiril"
  surname = "poposki"
}