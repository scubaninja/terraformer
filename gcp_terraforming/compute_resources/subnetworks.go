// AUTO-GENERATED CODE. DO NOT EDIT.
package computeTerrforming

import (
	"context"
	"log"
	"os"
	"strings"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
	"waze/terraformer/gcp_terraforming/gcp_generator"
	"waze/terraformer/terraform_utils"
)

var subnetworksIgnoreKey = map[string]bool{
	"^id$":                 true,
	"^self_link$":          true,
	"^fingerprint$":        true,
	"^label_fingerprint$":  true,
	"^creation_timestamp$": true,

	"gateway_address": true,
}

var subnetworksAllowEmptyValues = map[string]bool{}

var subnetworksAdditionalFields = map[string]string{
	"project": os.Getenv("GOOGLE_CLOUD_PROJECT"),
}

type SubnetworksGenerator struct {
	gcp_generator.BasicGenerator
}

// Run on subnetworksList and create for each TerraformResource
func (SubnetworksGenerator) createResources(subnetworksList *compute.SubnetworksListCall, ctx context.Context, region, zone string) []terraform_utils.TerraformResource {
	resources := []terraform_utils.TerraformResource{}
	if err := subnetworksList.Pages(ctx, func(page *compute.SubnetworkList) error {
		for _, obj := range page.Items {
			resources = append(resources, terraform_utils.NewTerraformResource(
				obj.Name,
				obj.Name,
				"google_compute_subnetwork",
				"google",
				nil,
				map[string]string{
					"name":    obj.Name,
					"project": os.Getenv("GOOGLE_CLOUD_PROJECT"),
					"region":  region,
				},
			))
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return resources
}

// Generate TerraformResources from GCP API,
// from each subnetworks create 1 TerraformResource
// Need subnetworks name as ID for terraform resource
func (g SubnetworksGenerator) Generate(zone string) ([]terraform_utils.TerraformResource, map[string]terraform_utils.ResourceMetaData, error) {
	region := strings.Join(strings.Split(zone, "-")[:len(strings.Split(zone, "-"))-1], "-")
	project := os.Getenv("GOOGLE_CLOUD_PROJECT")
	ctx := context.Background()

	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}

	subnetworksList := computeService.Subnetworks.List(project, region)

	resources := g.createResources(subnetworksList, ctx, region, zone)
	metadata := terraform_utils.NewResourcesMetaData(resources, subnetworksIgnoreKey, subnetworksAllowEmptyValues, subnetworksAdditionalFields)
	return resources, metadata, nil

}
