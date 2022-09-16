# outbox

# Usage

Create a collection on mongodb with this fields

```json
{
  "id": "fc848030-d72e-4a9c-bf61-55cafdb76454",
  "payload": {
      "id": "58452f68-705b-4b2e-8685-fc929e750588",
      "name": "Guilherme",
      "age": 27
  },
  "topic": "users",
  "event": "user_saved",
  "checked": false
}
```