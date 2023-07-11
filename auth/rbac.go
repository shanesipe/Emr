var userRoles = map[string]string{
  "user1": "doctor",
  "user2": "nurse",
  "user3": "admin",
}

func HasPermission(username string, permission string) bool {

  role := userRoles[username]

  if role == "admin" {
    return true
  }

  if role == "doctor" && permission == "read:patients" {
    return true
  }

  if role == "nurse" && permission == "read:appointments" {
    return true 
  }

  return false
}
