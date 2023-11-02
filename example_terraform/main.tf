terraform {
  required_providers {
    person = {
      version = "~> 1.0.0"
      source  = "kiril.com/personprovider/person"
    }
  }
}

provider "person" {
    session_id = "9fc264e5-b4b8-43fb-a006-65ae9dea24b0"
}

resource "person" "person1" {
  id = 1
  name = "kiril"
  surname = "poposki"
}