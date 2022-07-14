package image

import (
	"testing"

	igntypes "github.com/coreos/ignition/v2/config/v3_2/types"

	hiveext "github.com/openshift/assisted-service/api/hiveextension/v1beta1"
	"github.com/openshift/assisted-service/api/v1beta1"
	"github.com/openshift/assisted-service/models"
	hivev1 "github.com/openshift/hive/apis/hive/v1"
	"github.com/openshift/installer/pkg/types/agent"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Unable to test Generate because bootstrap.AddStorageFiles
// returns error in unit test:
//   open data/agent/files: no such file or directory
// Unit test working directory is ./pkg/asset/agent/image
// While normal execution working directory is ./data
// func TestIgnition_Generate(t *testing.T) {}

func TestIgnition_getTemplateData(t *testing.T) {
	clusterImageSet := &hivev1.ClusterImageSet{
		ObjectMeta: metav1.ObjectMeta{
			Name: "openshift-v4.10.0",
		},
		Spec: hivev1.ClusterImageSetSpec{
			ReleaseImage: "quay.io:443/openshift-release-dev/ocp-release:4.10.0-rc.1-x86_64",
		},
	}
	pullSecret := "pull-secret"
	nodeZeroIP := "192.168.111.80"
	agentClusterInstall := &hiveext.AgentClusterInstall{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-agent-cluster-install",
			Namespace: "cluster0",
		},
		Spec: hiveext.AgentClusterInstallSpec{
			APIVIP:       "192.168.111.2",
			SSHPublicKey: "ssh-rsa AAAAmyKey",
			ProvisionRequirements: hiveext.ProvisionRequirements{
				ControlPlaneAgents: 3,
				WorkerAgents:       5,
			},
		},
	}
	releaseImage := "quay.io:443/openshift-release-dev/ocp-release:4.10.0-rc.1-x86_64"
	releaseImageMirror := "virthost.ostest.test.metalkube.org:5000/localimages/local-release-image"
	mirrorRegistriesMount := "-v /etc/assisted/mirror/registries.conf:/etc/containers/registries.conf"
	caBundleMount := "-v /etc/assisted/mirror/ca-bundle.crt:/etc/pki/ca-trust/extracted/pem/tls-ca-bundle.pem"
	infraEnvID := "random-infra-env-id"

	releaseImageList, err := releaseImageList(clusterImageSet.Spec.ReleaseImage, "x86_64")
	assert.NoError(t, err)
	templateData := getTemplateData(pullSecret, nodeZeroIP, releaseImageList, releaseImage, releaseImageMirror, mirrorRegistriesMount, caBundleMount, agentClusterInstall, infraEnvID)
	assert.Equal(t, "http", templateData.ServiceProtocol)
	assert.Equal(t, "http://"+nodeZeroIP+":8090/", templateData.ServiceBaseURL)
	assert.Equal(t, pullSecret, templateData.PullSecret)
	assert.Equal(t, "", templateData.PullSecretToken)
	assert.Equal(t, nodeZeroIP, templateData.NodeZeroIP)
	assert.Equal(t, nodeZeroIP+":8090", templateData.AssistedServiceHost)
	assert.Equal(t, agentClusterInstall.Spec.APIVIP, templateData.APIVIP)
	assert.Equal(t, agentClusterInstall.Spec.ProvisionRequirements.ControlPlaneAgents, templateData.ControlPlaneAgents)
	assert.Equal(t, agentClusterInstall.Spec.ProvisionRequirements.WorkerAgents, templateData.WorkerAgents)
	assert.Equal(t, releaseImageList, templateData.ReleaseImages)
	assert.Equal(t, releaseImage, templateData.ReleaseImage)
	assert.Equal(t, releaseImageMirror, templateData.ReleaseImageMirror)
	assert.Equal(t, mirrorRegistriesMount, templateData.MirrorRegistriesMount)
	assert.Equal(t, caBundleMount, templateData.CaBundleMount)
	assert.Equal(t, infraEnvID, templateData.InfraEnvID)
}

func TestIgnition_addStaticNetworkConfig(t *testing.T) {
	cases := []struct {
		Name                string
		staticNetworkConfig []*models.HostStaticNetworkConfig
		expectedError       string
		expectedFileList    []string
	}{
		{
			Name: "default",
			staticNetworkConfig: []*models.HostStaticNetworkConfig{
				{
					MacInterfaceMap: models.MacInterfaceMap{
						{LogicalNicName: "eth0", MacAddress: "52:54:01:aa:aa:a1"},
					},
					NetworkYaml: "interfaces:\n- ipv4:\n    address:\n    - ip: 192.168.122.21\n      prefix-length: 24\n    enabled: true\n  mac-address: 52:54:01:aa:aa:a1\n  name: eth0\n  state: up\n  type: ethernet\n",
				},
			},
			expectedError: "",
			expectedFileList: []string{
				"/etc/assisted/network/host0/eth0.nmconnection",
				"/etc/assisted/network/host0/mac_interface.ini",
				"/usr/local/bin/pre-network-manager-config.sh",
			},
		},
		{
			Name:                "no-static-network-configs",
			staticNetworkConfig: []*models.HostStaticNetworkConfig{},
			expectedError:       "",
			expectedFileList:    nil,
		},
		{
			Name: "error-processing-config",
			staticNetworkConfig: []*models.HostStaticNetworkConfig{
				{
					MacInterfaceMap: models.MacInterfaceMap{
						{LogicalNicName: "eth0", MacAddress: "52:54:01:aa:aa:a1"},
					},
					NetworkYaml: "interfaces:\n- ipv4:\n    address:\n    - ip: bad-ip\n      prefix-length: 24\n    enabled: true\n  mac-address: 52:54:01:aa:aa:a1\n  name: eth0\n  state: up\n  type: ethernet\n",
				},
			},
			expectedError:    "'bad-ip' does not appear to be an IPv4 or IPv6 address",
			expectedFileList: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			config := igntypes.Config{}
			err := addStaticNetworkConfig(&config, tc.staticNetworkConfig)

			if tc.expectedError != "" {
				assert.Regexp(t, tc.expectedError, err.Error())
			} else {
				assert.NoError(t, err)
			}

			var fileList []string
			for _, file := range config.Storage.Files {
				fileList = append(fileList, file.Node.Path)
			}
			assert.Equal(t, tc.expectedFileList, fileList)
		})
	}
}

func TestRetrieveRendezvousIP(t *testing.T) {
	rawConfig := `interfaces: 
  - ipv4: 
      address: 
        - ip: "192.168.122.21"`
	cases := []struct {
		Name                 string
		agentConfig          *agent.Config
		nmStateConfigs       []*v1beta1.NMStateConfig
		expectedRendezvousIP string
	}{
		{
			Name: "valid-agent-config-provided-with-RendezvousIP",
			agentConfig: &agent.Config{
				Spec: agent.Spec{
					RendezvousIP: "192.168.122.21",
				},
			},
			expectedRendezvousIP: "192.168.122.21",
		},
		{
			Name: "no-agent-config-provided-so-read-from-nmstateconfig",
			nmStateConfigs: []*v1beta1.NMStateConfig{
				{
					Spec: v1beta1.NMStateConfigSpec{
						NetConfig: v1beta1.NetConfig{
							Raw: []byte(rawConfig),
						},
					},
				},
			},
			expectedRendezvousIP: "192.168.122.21",
		},
	}
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			rendezvousIP, err := retrieveRendezvousIP(tc.agentConfig, tc.nmStateConfigs)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedRendezvousIP, rendezvousIP)
		})
	}

}