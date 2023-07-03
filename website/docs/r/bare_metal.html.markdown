---
subcategory: "Bare Metal"
layout: "ksyun"
page_title: "ksyun: ksyun_bare_metal"
sidebar_current: "docs-ksyun-resource-bare_metal"
description: |-
  Provides a Bare Metal resource.
---

# ksyun_bare_metal

Provides a Bare Metal resource.

#

## Example Usage

```hcl
resource "ksyun_bare_metal" "default" {
  host_name              = "test"
  host_type              = "MI-I2"
  image_id               = "eb8c0428-476e-49af-8ccb-9fad2455a54c"
  key_id                 = "9c45b560-e51d-4aee-9e99-0e292476692d"
  network_interface_mode = "single"
  raid                   = "Raid1"
  availability_zone      = "cn-beijing-6b"
  security_agent         = "classic"
  cloud_monitor_agent    = "classic"
  subnet_id              = "d2fdc1b5-0280-4ca7-920b-0bd0453c130c"
  security_group_ids     = ["7e2f45b5-e79d-4612-a7fc-fe74a50b639a"]
  system_file_type       = "EXT4"
  container_agent        = "supported"
  force_re_install       = false
}
```

## Argument Reference

The following arguments are supported:

* `availability_zone` - (Required, ForceNew) The Availability Zone.
* `host_type` - (Required, ForceNew) The Bare Metal Host Type (e.g. CAL-III).
* `image_id` - (Required) ID of the image.
* `key_id` - (Required) The certificate id of the Bare Metal.
* `security_group_ids` - (Required) The security_group_id set of the Bare Metal primary network interface.
* `subnet_id` - (Required) The subnet id of the Bare Metal primary network interface.
* `bond_attribute` - (Optional) The bond attribute of the Bare Metal. Valid Values:'bond0','bond1'.Default is 'bond1'. Only effective when network_interface_mode is bond4.
* `cloud_monitor_agent` - (Optional) The cloud monitor agent choice of the Bare Metal.Valid Values:'classic','no'.Default is 'no'.
* `computer_name` - (Optional) The computer name of the Bare Metal.
* `container_agent` - (Optional) Whether to support KCE cluster, valid values: 'supported', 'unsupported'.
* `data_disk_catalogue_suffix` - (Optional) The data disk catalogue suffix of the Bare Metal.Valid Values:'NoSuffix','NaturalNumber','NaturalNumberFromZero'.Default is 'NaturalNumber'.
* `data_disk_catalogue` - (Optional) The data disk catalogue of the Bare Metal.Valid Values:'/DATA/disk','/data'.Default is '/DATA/disk'.
* `data_file_type` - (Optional) The data disk file type of the Bare Metal.Valid Values:'EXT4','XFS'.Default is 'XFS'.
* `dns1` - (Optional) The dns1 of the Bare Metal primary network interface.
* `dns2` - (Optional) The dns2 of the Bare Metal primary network interface.
* `extension_dns1` - (Optional) The dns1 of the Bare Metal extension network interface.Only effective when network_interface_mode is dual.
* `extension_dns2` - (Optional) The dns2 of the Bare Metal extension network interface.Only effective when network_interface_mode is dual.
* `extension_private_ip_address` - (Optional) The private ip address of the Bare Metal extension network interface.Only effective when network_interface_mode is dual.
* `extension_security_group_ids` - (Optional) The security_group_id set of the Bare Metal extension network interface.Max is 3.Only effective when network_interface_mode is dual and Required.
* `extension_subnet_id` - (Optional) The subnet id of the Bare Metal primary extension interface.Only effective when network_interface_mode is dual and Required.
* `force_re_install` - (Optional) Indicate whether to reinstall system.
* `host_name` - (Optional) The name of the Bare Metal.Default is 'ksc_epc'.
* `hyper_threading` - (Optional) The HyperThread status of the Bare Metal. Valid Values:'Open','Close','NoChange'.Default is 'NoChange'.
* `network_interface_mode` - (Optional) The network interface mode of the Bare Metal. Valid Values:'bond4','single','dual'.Default is 'bond4'.When bond4->single,single->bond4,dual->single,dual->bond4 can modify,otherwise is ForceNew.
* `nvme_data_disk_catalogue_suffix` - (Optional) The nvme data disk catalogue suffix of the Bare Metal.Valid Values:'NoSuffix','NaturalNumber','NaturalNumberFromZero'.
* `nvme_data_disk_catalogue` - (Optional) The nvme data disk catalogue of the Bare Metal.Valid Values:'/DATA/disk','/data'.
* `nvme_data_file_type` - (Optional) The nvme data file type of the Bare Metal.Valid Values:'EXT4','XFS'.
* `password` - (Optional) The password of the Bare Metal.
* `path` - (Optional) The path of the Bare Metal.Only effective on modify and host type is COLO.
* `private_ip_address` - (Optional) The private ip address of the Bare Metal primary network interface.
* `project_id` - (Optional) The project id of the Bare Metal.Default is '0'.
* `raid_id` - (Optional) The Raid template id of Bare Metal.Conflict raid. If you don't set raid,raid_id is Required. If you want to use raid_id,you must in user white list.
* `raid` - (Optional) The Raid type of the Bare Metal. Valid Values:'Raid0','Raid1','Raid5','Raid10','Raid50','SRaid0'. Conflict raid_id. If you don't set raid_id,raid is Required.
* `security_agent` - (Optional) The security agent choice of the Bare Metal. Valid Values:'classic','no'. Default is 'no'.
* `server_ip` - (Optional) The pxe server ip of the Bare Metal.Only effective on modify and host type is COLO.
* `system_file_type` - (Optional) The system disk file type of the Bare Metal.Valid Values:'EXT4','XFS'.Default is 'EXT4'.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `extension_network_interface_id` - ID of the extension network interface.
* `network_interface_id` - ID of the primary network interface.


## Import

Bare Metal can be imported using the id, e.g.

```
$ terraform import ksyun_bera_metal.default 67b91d3c-c363-4f57-b0cd-xxxxxxxxxxxx
```

