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

var backendServicesIgnoreKey = map[string]bool{
	"^id$":                 true,
	"^self_link$":          true,
	"^fingerprint$":        true,
	"^label_fingerprint$":  true,
	"^creation_timestamp$": true,

	"region": true,
}

var backendServicesAllowEmptyValues = map[string]bool{}

var backendServicesAdditionalFields = map[string]string{
	"project": os.Getenv("GOOGLE_CLOUD_PROJECT"),
}

type BackendServicesGenerator struct {
	gcp_generator.BasicGenerator
}

// Run on backendServicesList and create for each TerraformResource
func (BackendServicesGenerator) createResources(backendServicesList *compute.BackendServicesListCall, ctx context.Context, region, zone string) []terraform_utils.TerraformResource {
	resources := []terraform_utils.TerraformResource{}
	if err := backendServicesList.Pages(ctx, func(page *compute.BackendServiceList) error {
		for _, obj := range page.Items {
			resources = append(resources, terraform_utils.NewTerraformResource(
				obj.Name,
				obj.Name,
				"google_compute_backend_service",
				"google",
				nil,
				map[string]string{
					"name":    obj.Name,
					"project": os.Getenv("GOOGLE_CLOUD_PROJECT"),
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
// from each backendServices create 1 TerraformResource
// Need backendServices name as ID for terraform resource
func (g BackendServicesGenerator) Generate(zone string) ([]terraform_utils.TerraformResource, map[string]terraform_utils.ResourceMetaData, error) {
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

	backendServicesList := computeService.BackendServices.List(project)

	resources := g.createResources(backendServicesList, ctx, region, zone)
	metadata := terraform_utils.NewResourcesMetaData(resources, backendServicesIgnoreKey, backendServicesAllowEmptyValues, backendServicesAdditionalFields)
	return resources, metadata, nil

}
