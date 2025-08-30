# Code Style Guidelines

## prettier

## Variables

Variables name must be in camel_case.

```js
let phone_number = "";
```

Variable name must have proper naming and is grammatically correct. Exception from this is the short form or acronym of the word.

```js
let username = ""; // Correct
let uname = ""; // Incorrect

let user_repo; // user_repository is also acceptable
```

Constants variables must be written in CAPITAL LETTERS and \_ is used for spaces.

```js
const TOKEN = "abc";
const ROOT_DIR = "/";
```

## Classes and Structs

Classes and structs attributes must be written in pascalCase. The annotation for the struct can be in

```go
type User struct {
  username      string  `json:"username"`
  email         string  `json:"email"`
  emailVerified bool    `json:"email_verified"`
}
```

```ts
type User = {
  username: string;
  email: string;
  emailVerified: boolean;
};
```
