/*
This data source provides a list of VPN tunnels.

# Example Usage

```hcl

	data "ksyun_vpn_tunnels" "default" {
	  output_file="output_result"
	  ids=[]
	}

```
*/
package ksyun

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceKsyunVpnTunnels() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKsyunVpnTunnelsRead,

		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of VPN tunnel IDs, all the resources belong to this region will be retrieved if the ID is `\"\"`.",
			},

			"vpn_gateway_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of vpn gateway ids.",
			},

			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A regex string to filter results by name.",
			},

			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results (after running `terraform plan`).",
			},

			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Total number of resources that satisfy the condition.",
			},
			"vpn_tunnels": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "It is a nested type which documented below.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "VPN tunnel ID.",
						},

						"vpn_tunnel_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "VPN tunnel ID.",
						},

						"state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "VPN tunnel state.",
						},

						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "VPN tunnel type.",
						},

						"vpn_gre_ip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "VPN gre IP.",
						},

						"customer_gre_ip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Customer gre IP.",
						},

						"ha_vpn_gre_ip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "HA VPN gre IP.",
						},

						"ha_customer_gre_ip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "HA Customer gre IP.",
						},

						"vpn_gateway_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "VPN gateway ID.",
						},

						"customer_gateway_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Customer gateway ID.",
						},

						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "VPN tunnel name.",
						},

						"vpn_tunnel_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "VPN tunnel name.",
						},

						"pre_shared_key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "pre shared key.",
						},

						"ike_authen_algorithm": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IKE authen algorithm.",
						},

						"ike_dh_group": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "IKE dh group.",
						},

						"ike_encry_algorithm": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IKE encry algorithm.",
						},

						"ipsec_encry_algorithm": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPsec encry algorithm.",
						},

						"ipsec_authen_algorithm": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPsec authen algorithm.",
						},

						"ipsec_life_time_traffic": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "IPsec lifetime traffic.",
						},

						"ipsec_life_time_second": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "IPsec lifetime second.",
						},

						"extra_cidr_set": {
							Type:        schema.TypeList,
							Computed:    true,
							Optional:    true,
							Description: "A list of extra cidr.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cidr_block": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "cidr block.",
									},
								},
							},
						},

						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "creation time.",
						},
					},
				},
			},
		},
	}
}
func dataSourceKsyunVpnTunnelsRead(d *schema.ResourceData, meta interface{}) error {
	vpcService := VpcService{meta.(*KsyunClient)}
	return vpcService.ReadAndSetVpnTunnels(d, dataSourceKsyunVpnTunnels())
}
