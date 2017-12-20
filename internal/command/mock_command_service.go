package command

import (
	"context"
	"github.com/davejohnston/disco/internal/api"
	"github.com/davejohnston/disco/test"
	google_protobuf1 "github.com/golang/protobuf/ptypes/empty"
	"github.com/satori/go.uuid"
	"log"
)

// MockCommandSvc is a mock implementation of the GRPC service
// defined in api/controller.proto
// It uses fixtures in the test/fixtures directories to simulate
// real data.
type MockCommandSvc struct {
	ProviderList []*api.Provider
	ResourceList map[string]*api.Resource
}

// NewMockControllerSvc creates an instance of the
// mock controller.  It initializes the fixtures.
func NewMockCommandSvc() *MockCommandSvc {
	s := new(MockCommandSvc)
	s.ProviderList = []*api.Provider{}
	s.ResourceList = make(map[string]*api.Resource)
	s.loadFixtures()
	return s
}

// loadFixtures reads data from external json files and
// sets up the internal data structures to support the mock
// service.  It is called during a NewMockControllSvc() invocation
func (s *MockCommandSvc) loadFixtures() {
	test.LoadFixtureFromFile("test/fixtures/providers.json", &s.ProviderList)
	test.LoadFixtureFromFile("test/fixtures/resources.json", &s.ResourceList)
}

// GetProviders returns a list of available providers
// A Provider is something that takes input and returns resources
// e.g. a AWS Provider could take input in the form of account details
// and return all the EC2 instances available under that account
func (s MockCommandSvc) GetProviders(ctx context.Context, empty *google_protobuf1.Empty) (*api.ProviderList, error) {
	return &api.ProviderList{Providers: s.ProviderList}, nil
}

// GetProvider given the provider ID returns the provider
// The Provider detail includes a description of the provider, icons and other useful information.
func (s MockCommandSvc) GetProvider(context context.Context, id *api.ID) (*api.Provider, error) {

	for _, provider := range s.ProviderList {
		if provider.Uuid == id.Uuid {
			return provider, nil
		}
	}

	return &api.Provider{}, nil
}

// GetResources
func (s MockCommandSvc) GetResources(context context.Context, empty *google_protobuf1.Empty) (*api.ResourceList, error) {
	resources := make([]*api.Resource, len(s.ProviderList))
	for _, value := range s.ResourceList {
		resources = append(resources, value)
	}
	return &api.ResourceList{Resources: resources}, nil
}

// GetResource
func (s MockCommandSvc) GetResource(context context.Context, id *api.ID) (*api.Resource, error) {
	if resource, ok := s.ResourceList[id.Uuid]; ok {
		return resource, nil
	}
	return &api.Resource{}, nil
}

// CreateResource
func (s MockCommandSvc) CreateResource(context context.Context, resource *api.Resource) (*api.Resource, error) {
	resource.Uuid = uuid.NewV4().String()
	s.ResourceList[resource.Uuid] = resource
	return resource, nil
}

// UpdateResource
func (s MockCommandSvc) UpdateResource(context context.Context, resource *api.Resource) (*api.Resource, error) {
	return resource, nil
}

// DeleteResource
func (s MockCommandSvc) DeleteResource(context context.Context, id *api.ID) (*api.Resource, error) {
	return nil, nil
}

// GetSecrets
func (s MockCommandSvc) GetSecrets(empty *google_protobuf1.Empty, stream api.Command_GetSecretsServer) error {
	return nil
}

// GetSecret
func (s MockCommandSvc) GetSecret(context context.Context, id *api.ID) (*api.Secret, error) {
	return nil, nil
}

// CreateSecret
func (s MockCommandSvc) CreateSecret(context context.Context, secret *api.Secret) (*api.Secret, error) {
	log.Printf("Secrets: %+v", secret)
	return secret, nil
}

// UpdateSecret
func (s MockCommandSvc) UpdateSecret(context context.Context, secret *api.Secret) (*api.Secret, error) {
	return secret, nil
}

// DeleteSecret
func (s MockCommandSvc) DeleteSecret(context context.Context, secret *api.ID) (*api.Secret, error) {
	return nil, nil
}

// GetJobs
func (s MockCommandSvc) GetJobs(empty *google_protobuf1.Empty, stream api.Command_GetJobsServer) error {
	return nil
}

// GetJob
func (s MockCommandSvc) GetJob(context context.Context, id *api.ID) (*api.Job, error) {
	return nil, nil
}
