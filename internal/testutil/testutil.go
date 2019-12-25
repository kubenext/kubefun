package testutil

import (
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/utils/pointer"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// LoadObjectFromFile loads a file from the `testdata` directory. It will
// assign it a `default` namespace if one is not set.
func LoadObjectFromFile(t *testing.T, objectFile string) runtime.Object {
	data, err := ioutil.ReadFile(filepath.Join("testdata", objectFile))
	require.NoError(t, err)

	object, _, err := scheme.Codecs.UniversalDeserializer().Decode(data, nil, nil)
	require.NoError(t, err, "unable to decode serialized data")

	accessor := meta.NewAccessor()
	namespace, err := accessor.Namespace(object)
	require.NoError(t, err)
	if namespace == "" {
		require.NoError(t, accessor.SetNamespace(object, "default"))
	}

	return object
}

// LoadUnstructuredFromFile loads an object from a file in the in `testdata` directory.
// It will assign a `default` namespace if one is not set. This helper does not support
// multiple objects in a YAML file.
func LoadUnstructuredFromFile(t *testing.T, objectFile string) *unstructured.Unstructured {
	f, err := os.Open(filepath.Join("testdata", objectFile))
	require.NoError(t, err)

	defer func() {
		require.NoError(t, f.Close())
	}()

	d := yaml.NewYAMLOrJSONDecoder(f, 4096)

	ext := runtime.RawExtension{}
	require.NoError(t, d.Decode(&ext))

	obj, _, err := unstructured.UnstructuredJSONScheme.Decode(ext.Raw, nil, nil)
	require.NoError(t, err)

	u, ok := obj.(*unstructured.Unstructured)
	require.True(t, ok, "object is not an unstructured object")

	return u
}

// LoadTypedObjectFromFile loads a file from the `testdata` directory. It will
// assign it a `default` namespace if one is not set.
func LoadTypedObjectFromFile(t *testing.T, objectFile string, into runtime.Object) {
	data, err := ioutil.ReadFile(filepath.Join("testdata", objectFile))
	require.NoError(t, err)

	gvk := into.GetObjectKind().GroupVersionKind()
	object, _, err := scheme.Codecs.UniversalDeserializer().Decode(data, &gvk, into)
	require.NoError(t, err)

	accessor := meta.NewAccessor()
	namespace, err := accessor.Namespace(object)
	require.NoError(t, err)
	if namespace == "" {
		require.NoError(t, accessor.SetNamespace(object, "default"))
	}
}

// ToOwnerReferences converts an object to owner references.
func ToOwnerReferences(t *testing.T, object runtime.Object) []metav1.OwnerReference {
	objectKind := object.GetObjectKind()
	apiVersion, kind := objectKind.GroupVersionKind().ToAPIVersionAndKind()

	accessor := meta.NewAccessor()
	name, err := accessor.Name(object)
	require.NoError(t, err)

	uid, err := accessor.UID(object)
	require.NoError(t, err)

	return []metav1.OwnerReference{
		{
			APIVersion: apiVersion,
			Kind:       kind,
			Name:       name,
			UID:        uid,
			Controller: pointer.BoolPtr(true),
		},
	}
}

// Time generates a test time
func Time() time.Time {
	return time.Unix(1547211430, 0)
}

// RequireErrorOrNot or not is a helper that requires an error or not.
func RequireErrorOrNot(t *testing.T, wantErr bool, err error) {
	if wantErr {
		require.Error(t, err)
		return
	}
	require.NoError(t, err)
}
