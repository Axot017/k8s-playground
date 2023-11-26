terraform {
  required_providers {
    keycloak = {
      source  = "mrparkers/keycloak"
      version = ">= 4.0.0"
    }
  }
}

provider "keycloak" {
  client_id = "admin-cli"
  username  = var.keycloak_user
  password  = var.keycloak_password
  url       = var.keycloak_url
}

module "app" {
  source = "../module"

  realm = "local"
}
