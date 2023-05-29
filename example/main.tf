terraform {
  required_providers {
    gtm = {
      source  = "mirefly/google-tag-manager"
      version = "0.0.1"
    }
  }
}

provider "gtm" {
  credential_file = "./credentials-77b14e38b4dd.json"
  account_id      = "6105084028"
  container_id    = "119458552"
}

resource "gtm_workspace" "test" {
  name        = "test workspace"
  description = "This is a test workspace. Generated by terraform. Do not edit it."
}

resource "gtm_variable" "test_variable" {
  workspace_id = gtm_workspace.test.id
  name         = "test variable 1"
  type         = "v"
  notes        = "Generated by terraform. Do not edit it."
  parameter = [
    {
      key   = "name"
      type  = "template"
      value = "{{ parameters.alpha}}"
    }
  ]
}

resource "gtm_variable" "test_variable_2" {
  workspace_id = gtm_workspace.test.id
  name         = "test variable 2"
  type         = "v"
  notes        = "Generated by terraform. Do not edit it."
  parameter = [
    {
      key   = "name"
      type  = "template"
      value = "{{ parameters.beta}}"
    }
  ]
}

resource "gtm_tag" "test_tag_1" {
  workspace_id = gtm_workspace.test.id
  name         = "test tag 1"
  type         = "gaawe"
  notes        = "Generated by terraform. Do not edit it."
  parameter = [
    {
      key   = "eventName"
      type  = "template"
      value = "event1"
    },
    {
      key  = "eventParameters"
      type = "list"
      list = [{
        type = "map"
        map = [
          {
            type  = "template"
            key   = "name"
            value = "eventName"
          },
          {
            type  = "template"
            key   = "value"
            value = "eventValue"
        }]
      }]
    },
    {
      key   = "measurementId",
      type  = "template"
      value = "G-A2ABC2ABCD"
    }
  ]
  firing_trigger_id = [gtm_trigger.test_trigger_1.id]
}

resource "gtm_trigger" "test_trigger_1" {
  workspace_id = gtm_workspace.test.id
  name         = "test trigger 1"
  type         = "customEvent"
  notes        = "Generated by terraform. Do not edit it."
  custom_event_filter = [
    {
      type = "equals",
      parameter = [
        {
          type  = "template",
          key   = "arg0",
          value = "{{_event}}"
        },
        {
          "type"  = "template",
          "key"   = "arg1",
          "value" = "event-name"
        }
      ]
    }
  ]
}
