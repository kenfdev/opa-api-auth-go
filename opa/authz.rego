package httpapi.authz

# io.jwt.decode_verify
# https://www.openpolicyagent.org/docs/latest/language-reference/#tokens
token = t {
  [valid, _, payload] = io.jwt.decode_verify(input.token, { "secret": "secret" })
  t := {
    "valid": valid,
    "payload": payload
  }
}

default allow = false

# Allow users to get their own salaries.
allow {
  token.valid

  some username
  input.method == "GET"
  input.path = ["finance", "salary", username]
  token.payload.user == username
}

# Allow managers to get their subordinate's salaries.
allow {
  token.valid

  some username
  input.method == "GET"
  input.path = ["finance", "salary", username]
  token.payload.subordinates[_] == username
}

# Allow HR members to get anyone's salary.
allow {
  token.valid

  input.method == "GET"
  input.path = ["finance", "salary", _]
  token.payload.hr == true
}