# Delegatio

Delegatio is a framework that can be used to manage homework of, i.e., system security classes. The aim is to provide an infrastructure to let students work on problems independent of their hardware. The current architecture consists of three parts. 
1. Infrastructure is used to initialize the infrastructure and spawn a Kubernetes cluster. Currently, only local VMs with libvirt are supported. However, depending on the use case, we might add support to initialize a cluster in a public cloud.
2. Kubernetes is used to set up Kubernetes and deploy the necessary extensions. The extensions include storage (currently under development) and the CNI plugin. 
3. ssh (plan is to merge it into Kubernetes) let users connect to cluster pods using their ssh keys. Each key is assigned a unique identity to be able to grade the solutions in the future. 


## TODO
* Unittests
* Abstract storage 
* Refactor build system
* Merge ssh daemon into Kubernetes
* Webserver to deploy a website to generate ssh keys and sync them with the ssh daemon
* HA KV storage for the ssh daemon
* Support for multiple control planes
* ssh daemon vscode remote ssh support
* Harden Kubernetes Pods
