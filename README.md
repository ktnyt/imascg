![CircleCI](https://circleci.com/gh/ktnyt/imascg.svg?style=shield&circle-token=2feafa7c5c3bf0ac40aa8172b36f460edd7a542b)
[![Go Report Card](https://goreportcard.com/badge/github.com/ktnyt/imascg)](https://goreportcard.com/report/github.com/ktnyt/imascg)

# imascg
Useful blazingly fast APIs for developing IM@S CG related tools.

## API Documentation
Beware: these APIs are subject to major changes in the future. The current API is served in a flat namespace but future versions will be placed under version paths. These flat paths will either be deprecated or aliased to the newest version in the future.

### Characters
#### Synopsis
The `/characters` resource endpoint provides access to registered character objects. Character objects have the following schema:

```yaml
title: Character
type: object
properties: 
  id:
    title: Character ID
    type: string
  name:
    title: Character Name
    type: string
  type:
    title: Character Type
    type: string
  readings:
    title: Character Readings
    type: array
    items:
      type: string
```

A character ID is a four digit decimal value string which MSD represents the character type. 346 Production idols are assigned one type out of 'cute', 'cool', or 'pasn' which has a corresponding ID with MSD of '0', '1', or '2'. Other characters are assigned the type 'rest' and has an ID with MSD of '3'.

#### GET `/characters`
- Returns a list of registered characters.
- Options:
  - ?search (optional): filters characters based on the values supplied.
- Response:
  - 200: a list of filtered character instances.

#### POST `/characters`
- Adds a character to the database.
- Options:
  - body (required): a character instance without an ID.
- Response:
  - 200: the added character instance.
  - 400

#### GET `/characters/:id`
- Returns a character for the given ID.
- Options:
  - :id (required): ID of the character to return.
- Response:
  - 200: a matching character instance.
  - 404

#### PATCH `/characters/:id`
- Partially updates information about a character for the given ID.
- Options:
  - :id (required): ID of the character to update.
  - body (required):
    - name (optional)
    - type (optional)
    - readings (optional)
- Response:
  - 200: the updated character instance.
  - 400
  - 404

#### PUT `/characters/:id`
- Replaces information about a character for the given ID.
- Options:- :id (required): ID of the character to replace.
  - body (required): a character instance without an ID.
- Response:
  - 200: the replaced character instance.
  - 400
  - 404

#### DELETE `/characters/:id`
- Deletes the character instance for the given ID.
- Options:
  - :id (required): ID of the character to delete.
- Response:
  - 200: empty
  - 404

### Units
#### Synopsis
The `/units` resource endpoint provides access to registered unit objects. Unit objects have the following schema:

```yaml
title: Unit
type: object
properties: 
  id:
    title: Unit ID
    type: string
  name:
    title: Unit Name
    type: string
  members:
    title: Unit Members
    type: array
    items:
      type: string
  readings:
    title: Unit Readings
    type: array
    items:
      type: string
```

#### GET `/units`
- Returns a list of registered units.
- Options:
  - ?search (optional): filters units based on the values supplied.
- Response:
  - 200: a list of filtered unit instances.

#### POST `/units`
- Adds a unit to the database.
- Options:
  - body (required): a unit instance without an ID.
- Response:
  - 200: the added unit instance.
  - 400

#### GET `/units/:id`
- Returns a unit for the given ID.
- Options:
  - :id (required): ID of the unit to return.
- Response:
  - 200: a matching unit instance.
  - 404

#### PATCH `/units/:id`
- Partially updates information about a unit for the given ID.
- Options:
  - :id (required): ID of the unit to update.
  - body (required):
    - name (optional)
    - members (optional)
    - readings (optional)
- Response:
  - 200: the updated unit instance.
  - 400
  - 404

#### PUT `/units/:id`
- Replaces information about a unit for the given ID.
- Options:- :id (required): ID of the unit to replace.
  - body (required): a unit instance without an ID.
- Response:
  - 200: the replaced unit instance.
  - 400
  - 404

#### DELETE `/units/:id`
- Deletes the unit instance for the given ID.
- Options:
  - :id (required): ID of the unit to delete.
- Response:
  - 200: empty
  - 404

### Calltable
The `/calltable` resource endpoint provides access to the calltable. Calltable entries have the following schema:

```yaml
title: Calltable Entry
type: object
properties:
  id:
    title: Calltable Entry ID
    type: string
  caller:
    title: Caller
    type: string
  callee:
    title: Callee
    type: string
  called:
    title: Called
    type: string
  remark:
    title: Remark
    type: string
```

The ID of a calltable entry is defined as the caller's id + the callee's id + one digit decimal number. This is based on the assumption that no one person calls another person with more than 10 different names. This assumption may be proven wrong and the id naming convention may change in the future.

#### GET `/calltable`
- Returns a list of registered calltable entries.
- Options:
  - caller (optional): filters the callers by given character ID.
  - callee (optional): filters the callees by given character ID.
  - called (optional): filters the called values with regular expression.
  - remark (optional): filters the remark values with regular expression.
- Response:
  - 200: a list of filtered calltable instances.

#### POST `/calltable`
- Adds a calltable entry to the database.
- Options:
  - body (required): a calltable entry instance without an ID.
- Response:
  - 200: the added calltable entry instance.
  - 400

#### GET `/calltable/:id`
- Returns a calltable entry for the given ID.
- Options:
  - :id (required): ID of the calltable entry to return.
- Response:
  - 200: a matching calltablee ntry instance.
  - 404

#### PATCH `/calltable/:id`
- Partially updates information about a calltable entry for the given ID.
- Options:
  - :id (required): ID of the calltable entry to update.
  - body (required):
    - caller (optional)
    - callee (optional)
    - called (optional)
    - remark (optional)
- Response:
  - 200: the updated calltable entry instance.
  - 400
  - 404

#### PUT `/calltable/:id`
- Replaces information about a calltable entry for the given ID.
- Options:- :id (required): ID of the calltable entry to replace.
  - body (required): a calltable entry instance without an ID.
- Response:
  - 200: the replaced calltable entry instance.
  - 400
  - 404

#### DELETE `/calltable/:id`
- Deletes the calltable entry instance for the given ID.
- Options:
  - :id (required): ID of the calltable entry to delete.
- Response:
  - 200: empty
  - 404
