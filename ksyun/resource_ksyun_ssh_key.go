/*
Provides an SSH key resource.

# Example Usage

```hcl

	resource "ksyun_ssh_key" "default" {
	  key_name="ssh_key_tf"
	  public_key="ssh-rsa xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	}

```

# Import

SSH key can be imported using the id, e.g.

```
$ terraform import ksyun_ssh_key.default xxxxxxxxxxxx
```
*/
package ksyun

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceKsyunSSHKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceKsyunSSHKeyCreate,
		Read:   resourceKsyunSSHKeyRead,
		Update: resourceKsyunSSHKeyUpdate,
		Delete: resourceKsyunSSHKeyDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"key_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "name of the key.",
			},
			"key_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID of the key.",
			},
			"public_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "public key.",
			},
			"private_key": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "private key.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "creation time of the key.",
			},
		},
	}
}
func resourceKsyunSSHKeyCreate(d *schema.ResourceData, meta interface{}) (err error) {
	sksService := SksService{meta.(*KsyunClient)}
	err = sksService.CreateKey(d, resourceKsyunSSHKey())
	if err != nil {
		return fmt.Errorf("error on creating ssh key %q, %s", d.Id(), err)
	}
	return resourceKsyunSSHKeyRead(d, meta)
}

func resourceKsyunSSHKeyRead(d *schema.ResourceData, meta interface{}) (err error) {
	sksService := SksService{meta.(*KsyunClient)}
	err = sksService.ReadAndSetKey(d, resourceKsyunSSHKey())
	if err != nil {
		return fmt.Errorf("error on reading ssh key %q, %s", d.Id(), err)
	}
	return err
}

func resourceKsyunSSHKeyUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	sksService := SksService{meta.(*KsyunClient)}
	err = sksService.ModifyKey(d, resourceKsyunSSHKey())
	if err != nil {
		return fmt.Errorf("error on updating ssh key %q, %s", d.Id(), err)
	}
	return resourceKsyunSSHKeyRead(d, meta)
}

func resourceKsyunSSHKeyDelete(d *schema.ResourceData, meta interface{}) (err error) {
	sksService := SksService{meta.(*KsyunClient)}
	err = sksService.RemoveKey(d)
	if err != nil {
		return fmt.Errorf("error on deleting ssh key %q, %s", d.Id(), err)
	}
	return err
}
