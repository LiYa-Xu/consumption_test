package pkg

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"

	//"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
	"net/http"
	//"os"
)

var (
	oofsGVR = schema.GroupVersionResource{
		Group:    "build.dev",
		Version:  "v1alpha1",
		Resource: "buildruns",
	}
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	//kubeConfig := os.Getenv("KUBECONFIG")
	//if kubeConfig == "" {
	//	fmt.Println("cannot get cluster kube config, please export environment variable KUBECONFIG")
	//	os.Exit(1)
    //}
	//config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	//if err != nil {
	//	fmt.Println("failed to build config:", err)
	//	os.Exit(1)
	//}
	config, _ := rest.InClusterConfig()
	klog.InitFlags(nil)

	klog.Infof("Starting test")
	//config, _ := rest.InClusterConfig()

	dynClient, errClient := dynamic.NewForConfig(config)
	if errClient != nil {
		klog.Fatalf("Error received creating client %v", errClient)
	}

	crdClient := dynClient.Resource(oofsGVR)

	crd, errCrd := crdClient.Get("kaniko-golang-buildrun-liya-02", metav1.GetOptions{})
	if errCrd != nil {
		klog.Fatalf("Error getting CRD %v", errCrd)
	}
	klog.Infof("Got CRD: %v", crd)
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
