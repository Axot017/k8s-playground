resource "keycloak_openid_client" "openid_client" {
  realm_id  = keycloak_realm.realm.id
  client_id = "app-client"

  name    = "App Client"
  enabled = true

  access_type = "PUBLIC"
  valid_redirect_uris = [
    "http://localhost:7777/*",
  ]

  standard_flow_enabled = true
}
