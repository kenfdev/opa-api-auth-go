# OPA API Authorization

A sample to show the difference between using OPA as the Policy Decision Point.

## How to

The master branch does not include OPA as the Policy Decision Point. OPA is added in the [add-opa branch](https://github.com/kenfdev/opa-api-auth-go/tree/add-opa).

You can compare the difference in [this PR](https://github.com/kenfdev/opa-api-auth-go/pull/1)

If you want to spin up the OPA version, do the following steps:

```bash
git checkout add-opa

docker-compose up
```

#### Requesting the API

The API is available in the `GET /finance/salary/:username` endpoint. You can `GET` to this endpoint but a JWT is needed to properly make a request.

You can easily create an JWT at [jwt.io](https://jwt.io/). This application assumes that the algorithm is HS256 and the secret is simply `secret`. Also, the payload should look something like this:

```json
{
  "sub": "1234567890",
  "iat": 1516239022,
  "user": "bob",
  "subordinates": ["alice"],
  "hr": false
}
```

The following is an example using jwt.io.

![](./docs/assets/jwt.png)

After you have created a JWT, put that JWT in your `Authorization` header with the prefix `Bearer `. The following is an example using `curl`.

```bash
curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwiaWF0IjoxNTE2MjM5MDIyLCJ1c2VyIjoiYm9iIiwic3Vib3JkaW5hdGVzIjpbImFsaWNlIl0sImhyIjpmYWxzZX0.KIGr6epjELT308JxjVbzJkZ__-7q7L5VxHUwHzMUAsw" http://localhost:1323/finance/salary/bob
```