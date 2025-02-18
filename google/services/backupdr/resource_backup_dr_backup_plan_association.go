// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/backupdr/BackupPlanAssociation.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package backupdr

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceBackupDRBackupPlanAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceBackupDRBackupPlanAssociationCreate,
		Read:   resourceBackupDRBackupPlanAssociationRead,
		Delete: resourceBackupDRBackupPlanAssociationDelete,

		Importer: &schema.ResourceImporter{
			State: resourceBackupDRBackupPlanAssociationImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"backup_plan": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.ProjectNumberDiffSuppress,
				Description:      `The BP with which resource needs to be created`,
			},
			"backup_plan_association_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The id of backupplan association`,
			},
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The location for the backupplan association`,
			},
			"resource": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The resource for which BPA needs to be created`,
			},
			"resource_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The resource type of workload on which backupplan is applied`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time when the instance was created`,
			},
			"data_source": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Resource name of data source which will be used as storage location for backups taken`,
			},
			"last_successful_backup_consistency_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The point in time when the last successful backup was captured from the source`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The name of backup plan association resource created`,
			},
			"rules_config_info": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `Message for rules config info`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"last_backup_error": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: `google.rpc.Status object to store the last backup error`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"code": {
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: `The status code, which should be an enum value of [google.rpc.Code]`,
									},
									"message": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: `A developer-facing error message, which should be in English.`,
									},
								},
							},
						},
						"last_backup_state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `State of last backup taken.`,
						},
						"rule_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Backup Rule id fetched from backup plan.`,
						},
					},
				},
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time when the instance was updated.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceBackupDRBackupPlanAssociationCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	resourceProp, err := expandBackupDRBackupPlanAssociationResource(d.Get("resource"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("resource"); !tpgresource.IsEmptyValue(reflect.ValueOf(resourceProp)) && (ok || !reflect.DeepEqual(v, resourceProp)) {
		obj["resource"] = resourceProp
	}
	backupPlanProp, err := expandBackupDRBackupPlanAssociationBackupPlan(d.Get("backup_plan"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("backup_plan"); !tpgresource.IsEmptyValue(reflect.ValueOf(backupPlanProp)) && (ok || !reflect.DeepEqual(v, backupPlanProp)) {
		obj["backupPlan"] = backupPlanProp
	}
	resourceTypeProp, err := expandBackupDRBackupPlanAssociationResourceType(d.Get("resource_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("resource_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(resourceTypeProp)) && (ok || !reflect.DeepEqual(v, resourceTypeProp)) {
		obj["resourceType"] = resourceTypeProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{BackupDRBasePath}}projects/{{project}}/locations/{{location}}/backupPlanAssociations/?backup_plan_association_id={{backup_plan_association_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new BackupPlanAssociation: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for BackupPlanAssociation: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating BackupPlanAssociation: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/backupPlanAssociations/{{backup_plan_association_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = BackupDROperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating BackupPlanAssociation", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create BackupPlanAssociation: %s", err)
	}

	if err := d.Set("name", flattenBackupDRBackupPlanAssociationName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/backupPlanAssociations/{{backup_plan_association_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating BackupPlanAssociation %q: %#v", d.Id(), res)

	return resourceBackupDRBackupPlanAssociationRead(d, meta)
}

func resourceBackupDRBackupPlanAssociationRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{BackupDRBasePath}}projects/{{project}}/locations/{{location}}/backupPlanAssociations/{{backup_plan_association_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for BackupPlanAssociation: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("BackupDRBackupPlanAssociation %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading BackupPlanAssociation: %s", err)
	}

	if err := d.Set("name", flattenBackupDRBackupPlanAssociationName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading BackupPlanAssociation: %s", err)
	}
	if err := d.Set("backup_plan", flattenBackupDRBackupPlanAssociationBackupPlan(res["backupPlan"], d, config)); err != nil {
		return fmt.Errorf("Error reading BackupPlanAssociation: %s", err)
	}
	if err := d.Set("resource_type", flattenBackupDRBackupPlanAssociationResourceType(res["resourceType"], d, config)); err != nil {
		return fmt.Errorf("Error reading BackupPlanAssociation: %s", err)
	}
	if err := d.Set("create_time", flattenBackupDRBackupPlanAssociationCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading BackupPlanAssociation: %s", err)
	}
	if err := d.Set("update_time", flattenBackupDRBackupPlanAssociationUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading BackupPlanAssociation: %s", err)
	}
	if err := d.Set("data_source", flattenBackupDRBackupPlanAssociationDataSource(res["dataSource"], d, config)); err != nil {
		return fmt.Errorf("Error reading BackupPlanAssociation: %s", err)
	}
	if err := d.Set("rules_config_info", flattenBackupDRBackupPlanAssociationRulesConfigInfo(res["rulesConfigInfo"], d, config)); err != nil {
		return fmt.Errorf("Error reading BackupPlanAssociation: %s", err)
	}
	if err := d.Set("last_successful_backup_consistency_time", flattenBackupDRBackupPlanAssociationLastSuccessfulBackupConsistencyTime(res["lastSuccessfulBackupConsistencyTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading BackupPlanAssociation: %s", err)
	}

	return nil
}

func resourceBackupDRBackupPlanAssociationDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for BackupPlanAssociation: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{BackupDRBasePath}}projects/{{project}}/locations/{{location}}/backupPlanAssociations/{{backup_plan_association_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting BackupPlanAssociation %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "BackupPlanAssociation")
	}

	err = BackupDROperationWaitTime(
		config, res, project, "Deleting BackupPlanAssociation", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting BackupPlanAssociation %q: %#v", d.Id(), res)
	return nil
}

func resourceBackupDRBackupPlanAssociationImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/backupPlanAssociations/(?P<backup_plan_association_id>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<backup_plan_association_id>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<backup_plan_association_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/backupPlanAssociations/{{backup_plan_association_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenBackupDRBackupPlanAssociationName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRBackupPlanAssociationBackupPlan(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRBackupPlanAssociationResourceType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRBackupPlanAssociationCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRBackupPlanAssociationUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRBackupPlanAssociationDataSource(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRBackupPlanAssociationRulesConfigInfo(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"rule_id":           flattenBackupDRBackupPlanAssociationRulesConfigInfoRuleId(original["ruleId"], d, config),
			"last_backup_state": flattenBackupDRBackupPlanAssociationRulesConfigInfoLastBackupState(original["lastBackupState"], d, config),
			"last_backup_error": flattenBackupDRBackupPlanAssociationRulesConfigInfoLastBackupError(original["lastBackupError"], d, config),
		})
	}
	return transformed
}
func flattenBackupDRBackupPlanAssociationRulesConfigInfoRuleId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRBackupPlanAssociationRulesConfigInfoLastBackupState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRBackupPlanAssociationRulesConfigInfoLastBackupError(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["code"] =
		flattenBackupDRBackupPlanAssociationRulesConfigInfoLastBackupErrorCode(original["code"], d, config)
	transformed["message"] =
		flattenBackupDRBackupPlanAssociationRulesConfigInfoLastBackupErrorMessage(original["message"], d, config)
	return []interface{}{transformed}
}
func flattenBackupDRBackupPlanAssociationRulesConfigInfoLastBackupErrorCode(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRBackupPlanAssociationRulesConfigInfoLastBackupErrorMessage(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRBackupPlanAssociationLastSuccessfulBackupConsistencyTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandBackupDRBackupPlanAssociationResource(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBackupDRBackupPlanAssociationBackupPlan(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBackupDRBackupPlanAssociationResourceType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
