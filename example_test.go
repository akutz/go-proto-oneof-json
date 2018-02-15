package example_test

import (
	"bytes"
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"

	"github.com/akutz/go-proto-oneof-json"
)

func TestParentMarshalJSON(t *testing.T) {
	testMarshalToJSON(t, &example.Parent{
		Name: "Andrew",
		Dependents: []*example.Dependent{
			&example.Dependent{
				Name: "E",
				Child: &example.Dependent_Daughter{
					Daughter: &example.Dependent_Female{
						Attributes: []string{"my one true love"},
					},
				},
			},
		},
	})
}

func TestCreateDependencyRequestMarshalJSON(t *testing.T) {
	testMarshalToJSON(t, &example.CreateDependencyRequest{
		ParentName: "Andrew",
		Dependents: []*example.Dependent{
			&example.Dependent{
				Name: "A",
				Child: &example.Dependent_Son{
					Son: &example.Dependent_Male{},
				},
			},
		},
	})
}

func TestCreateVolumeRequestMarshalJSON(t *testing.T) {
	testMarshalToJSON(t, &csi.CreateVolumeRequest{
		VolumeCapabilities: []*csi.VolumeCapability{
			&csi.VolumeCapability{
				AccessMode: &csi.VolumeCapability_AccessMode{
					Mode: csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER,
				},
				AccessType: &csi.VolumeCapability_Block{
					Block: &csi.VolumeCapability_BlockVolume{},
				},
			},
		},
	})
}

func TestValidateVolumeCapabilitiesRequestMarshalToJSON(t *testing.T) {
	testMarshalToJSON(t, &csi.ValidateVolumeCapabilitiesRequest{
		VolumeCapabilities: []*csi.VolumeCapability{
			&csi.VolumeCapability{
				AccessMode: &csi.VolumeCapability_AccessMode{
					Mode: csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER,
				},
				AccessType: &csi.VolumeCapability_Mount{
					Mount: &csi.VolumeCapability_MountVolume{
						FsType:     "ext4",
						MountFlags: []string{"nosuid", "uid=500"},
					},
				},
			},
		},
	})
}

func testMarshalToJSON(t *testing.T, obj proto.Message) {
	// Marshal obj to JSON using the protobuf "jsonpb" package.
	enc := &jsonpb.Marshaler{
		Indent:      "  ",
		EnumsAsInts: true,
	}
	buf := &bytes.Buffer{}
	if err := enc.Marshal(buf, obj); err != nil {
		t.Fatalf("failed to marshal obj to json: %v", err)
	}

	// Print the marshalled JSON.
	t.Logf("obj=%s", buf)

	// Unnmarshal the JSON into obj using the protobuf "jsonpb" package.
	obj.Reset()
	if err := jsonpb.Unmarshal(buf, obj); err != nil {
		t.Fatalf("failed to unmarshal obj from json: %v", err)
	}

	// Print the object after the JSON has been unmarshalled into it.
	t.Logf("obj=%s", obj)
}
