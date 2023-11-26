resource "keycloak_realm" "realm" {
  realm   = var.realm
  enabled = true

  // Login
  registration_allowed           = false
  registration_email_as_username = true
  edit_username_allowed          = true
  reset_password_allowed         = true
  remember_me                    = true
  login_with_email_allowed       = true
  duplicate_emails_allowed       = false
}
