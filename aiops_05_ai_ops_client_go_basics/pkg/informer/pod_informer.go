package informer

import (
    "fmt"
    "time"

    corev1 "k8s.io/api/core/v1"
    "k8s.io/client-go/informers"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/cache"
    "k8s.io/client-go/util/workqueue"
    "k8s.io/klog/v2"
)

type PodInformer struct {
    clientset  kubernetes.Interface
    informer   cache.SharedIndexInformer
    workqueue  workqueue.RateLimitingInterface
}

func NewPodInformer(client kubernetes.Interface) *PodInformer {
    // 创建 informer factory
    factory := informers.NewSharedInformerFactory(client, time.Second*30)
    
    // 创建 pod informer
    podInformer := factory.Core().V1().Pods().Informer()
    
    // 创建工作队列
    queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())
    
    // 添加事件处理函数
    podInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
        AddFunc: func(obj interface{}) {
            key, err := cache.MetaNamespaceKeyFunc(obj)
            if err == nil {
                queue.Add(key)
                klog.Infof("Pod added: %s", key)
            }
        },
        UpdateFunc: func(old, new interface{}) {
            key, err := cache.MetaNamespaceKeyFunc(new)
            if err == nil {
                queue.Add(key)
                klog.Infof("Pod updated: %s", key)
            }
        },
        DeleteFunc: func(obj interface{}) {
            key, err := cache.MetaNamespaceKeyFunc(obj)
            if err == nil {
                queue.Add(key)
                klog.Infof("Pod deleted: %s", key)
            }
        },
    })
    
    return &PodInformer{
        clientset:  client,
        informer:   podInformer,
        workqueue:  queue,
    }
}

func (p *PodInformer) Run(stopCh <-chan struct{}) error {
    defer p.workqueue.ShutDown()

    // 启动 informer
    go p.informer.Run(stopCh)

    // 等待缓存同步
    if !cache.WaitForCacheSync(stopCh, p.informer.HasSynced) {
        return fmt.Errorf("failed to wait for caches to sync")
    }

    // 启动工作队列处理
    go p.runWorker()

    <-stopCh
    return nil
}

func (p *PodInformer) runWorker() {
    for p.processNextItem() {
    }
}

func (p *PodInformer) processNextItem() bool {
    obj, shutdown := p.workqueue.Get()
    if shutdown {
        return false
    }

    defer p.workqueue.Done(obj)

    key, ok := obj.(string)
    if !ok {
        p.workqueue.Forget(obj)
        return true
    }

    if err := p.syncHandler(key); err != nil {
        p.workqueue.AddRateLimited(key)
        return true
    }

    p.workqueue.Forget(obj)
    return true
}

func (p *PodInformer) syncHandler(key string) error {
    namespace, name, err := cache.SplitMetaNamespaceKey(key)
    if err != nil {
        return err
    }

    obj, exists, err := p.informer.GetIndexer().GetByKey(key)
    if err != nil {
        return err
    }

    if !exists {
        klog.Infof("Pod %s/%s has been deleted", namespace, name)
        return nil
    }

    pod := obj.(*corev1.Pod)
    klog.Infof("Processing pod %s/%s (phase: %s)", namespace, name, pod.Status.Phase)
    return nil
} 