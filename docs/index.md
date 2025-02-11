---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "google-tag-manager Provider"
subcategory: ""
description: |-
  
---

# google-tag-manager Provider



## Example Usage

```terraform
provider "gtm" {
  credential_file            = "../credentials-77b14e38b4dd.json"
  account_id                 = "6105084028"
  container_id               = "119458552"
  workspace_name             = "my-workspace"
  retry_limit                = 10
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String) GTM Account ID.
- `container_id` (String) GTM Container ID.
- `credential_file` (String) Path to the credential file.
- `workspace_name` (String) Workspace name.

### Optional

- `retry_limit` (Number) Number of times to retry requests when rate-limited before giving up.
