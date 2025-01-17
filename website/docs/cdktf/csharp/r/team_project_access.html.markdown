---
layout: "tfe"
page_title: "Terraform Enterprise: tfe_team_project_access"
description: |-
  Associate a team to permissions on a project.
---

# tfe_team_project_access

Associate a team to permissions on a project.

## Example Usage

Basic usage:

```csharp
using Constructs;
using HashiCorp.Cdktf;
/*Provider bindings are generated by running cdktf get.
See https://cdk.tf/provider-generation for more details.*/
using Gen.Providers.Tfe;
class MyConvertedCode : TerraformStack
{
    public MyConvertedCode(Construct scope, string name) : base(scope, name)
    {
        var tfeProjectTest = new Project.Project(this, "test", new ProjectConfig {
            Name = "myproject",
            Organization = "my-org-name"
        });
        var tfeTeamAdmin = new Team.Team(this, "admin", new TeamConfig {
            Name = "my-admin-team",
            Organization = "my-org-name"
        });
        var tfeTeamProjectAccessAdmin =
        new TeamProjectAccess.TeamProjectAccess(this, "admin_2", new TeamProjectAccessConfig {
            Access = "admin",
            ProjectId = Token.AsString(tfeProjectTest.Id),
            TeamId = Token.AsString(tfeTeamAdmin.Id)
        });
        /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
        tfeTeamProjectAccessAdmin.OverrideLogicalId("admin");
    }
}
```

## Argument Reference

The following arguments are supported:

* `TeamId` - (Required) ID of the team to add to the project.
* `ProjectId` - (Required) ID of the project to which the team will be added.
* `Access` - (Required) Type of fixed access to grant. Valid values are `Admin`, `Maintain`, `Write`, or `Read`.

## Attributes Reference

* `Id` The team project access ID.

## Import

Team project accesses can be imported; use the project team access ID as the import ID. For
example:

```shell
terraform import tfe_team_project_access.admin tprj-2pmtXpZa4YzVMTPi
```

<!-- cache-key: cdktf-0.17.0-pre.15 input-6077af17aff449bdcfb81f649aba3c085dccfd367de7cedd899696f602bd0c83 -->