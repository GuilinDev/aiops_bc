package main

import (
    "context"
    "flag"
    "path/filepath"
    "time"

    "github.com/yourusername/aiops_05_ai_ops_client_go_basics/pkg/informer"
    "k8s.io/client-go/dynamic"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/clientcmd"
    "k8s.io/client-go/util/homedir"
    "k8s.io/klog/v2"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/apimachinery/pkg/runtime/schema"
)

func main() {
    var kubeconfig *string
    if home := homedir.HomeDir(); home != "" {
        kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
    } else {
        kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
    }
    flag.Parse()

    // 创建 config
    config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
    if err != nil {
        klog.Fatalf("Error building kubeconfig: %s", err.Error())
    }

    // 创建 clientset
    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        klog.Fatalf("Error building kubernetes clientset: %s", err.Error())
    }

    // 创建 dynamic client
    dynamicClient, err := dynamic.NewForConfig(config)
    if err != nil {
        klog.Fatalf("Error building dynamic client: %s", err.Error())
    }

    // 创建并运行 Pod Informer
    podInformer := informer.NewPodInformer(clientset)
    stopCh := make(chan struct{})
    defer close(stopCh)

    go func() {
        if err := podInformer.Run(stopCh); err != nil {
            klog.Fatalf("Error running pod informer: %s", err.Error())
        }
    }()

    // 使用 dynamic client 获取 AIOps 资源
    gvr := schema.GroupVersionResource{
        Group:    "aiops.geektime.com",
        Version:  "v1alpha1",
        Resource: "aiops",
    }

    // 监听 AIOps 资源
    for {
        list, err := dynamicClient.Resource(gvr).Namespace("default").List(context.TODO(), metav1.ListOptions{})
        if err != nil {
            klog.Errorf("Error listing AIOps resources: %s", err.Error())
        } else {
            klog.Infof("Found %d AIOps resources", len(list.Items))
            for _, item := range list.Items {
                klog.Infof("AIOps: %s", item.GetName())
            }
        }
        time.Sleep(time.Second * 30)
    }
} 