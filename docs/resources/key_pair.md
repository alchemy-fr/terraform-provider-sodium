---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sodium_key_pair Resource - sodium"
subcategory: ""
description: |-
  Creates a Sodium formatted key pair.
  Generates a secure key pair and encodes it in base64
---

# sodium_key_pair (Resource)

Creates a Sodium formatted key pair.

Generates a secure key pair and encodes it in base64

## Example Usage

```terraform
resource "sodium_key_pair" "foo" {}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `id` (String) ID of the resource (public key).
- `public_key` (String, Sensitive) Public key data in base64 format.
- `secret_key` (String, Sensitive) Secret key data in base64 format.