# Project Documentation

- [Project Documentation](#project-documentation)
- [Home](#home)
    - [Kubernetes Cluster with Dynamic DNS using K3d and Cloudflare](#kubernetes-cluster-with-dynamic-dns-using-k3d-and-cloudflare)
  - [Prerequisites](#prerequisites)
  - [Basic Usage](#basic-usage)
  - [Set up ExternalDNS](#set-up-externaldns)
- [NodePort Usage](/nodeport/README.md)
- [LoadBalancer Usage](/loadbalancer/README.md)
- [Ingress Usage](/Ingress/README.md)

# Home

Welcome to the documentation for the project. Here, you'll find detailed information on setting up a Kubernetes cluster with dynamic DNS using K3d and Cloudflare.
### Kubernetes Cluster with Dynamic DNS using K3d and Cloudflare

This project aims to create a Kubernetes cluster with dynamic DNS (DDNS) using K3d and Cloudflare, allowing external access without a static IP address.

## Prerequisites

Before getting started, make sure you have the following prerequisites in place:

- [K3d](https://k3d.io/) installed on your local machine
- A [Cloudflare](https://www.cloudflare.com/) account
- [Docker](https://www.docker.com/) or [Rancher](https://rancher.com/)
- [Helm Chart](https://helm.sh/) must be installed on your local machine
- A registered domain or subdomain managed by Cloudflare
- API token with the necessary permissions for Cloudflare DNS management


## Basic Usage
NodePort can be used to quickly and easily create externally accessible services within Kubernetes. In this section, we explain its basic usage, but for more detailed and complex scenarios, you can refer to other documents.

1. **Create a Kubernetes Cluster**

   Use K3d to create a Kubernetes cluster. For example:

    ```bash
   k3d cluster create recon
    ```


1.1  **Create a Kubernetes Cluster**
![Cluster-Create](/static/images/cluster-create.png)

```
$ kubectl cluster-info
Kubernetes control plane is running at https://0.0.0.0:53401
CoreDNS is running at https://0.0.0.0:53401/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy
Metrics-server is running at https://0.0.0.0:53401/api/v1/namespaces/kube-system/services/https:metrics-server:https/proxy
```

## Set up ExternalDNS

1. **Install ExternalDNS with Helm:**

   Use Helm to install ExternalDNS into your Kubernetes cluster. Make sure Helm is installed and configured before proceeding.

   ```bash
   helm repo add bitnami https://charts.bitnami.com/bitnami
   helm install external-dns bitnami/external-dns --namespace external-dns --set provider=cloudflare --set cloudflare.apiToken=YOUR_CLOUDFLARE_API_TOKEN --set txtOwnerId=my-identifier
    ```
     Replace *YOUR_CLOUDFLARE_API_TOKEN* with your Cloudflare API token. You can customize the txtOwnerId to a unique identifier.

> **_Note:_** Namespace is important because it determines in which namespace ExternalDNS will run within your Kubernetes cluster. Namespace provides isolation and organization in Kubernetes, allowing you to separate resources for different projects or applications.
>
> The choice of which namespace to run ExternalDNS in depends on the requirements of your project and the preferences of your organization. You can select any namespace, but in most cases, placing ExternalDNS in the namespace of a specific application or project can enhance isolation and organization.
>
> Furthermore, when determining the namespace in which ExternalDNS operates, ensure it aligns with the zone in your Cloudflare account for which DNS records need to be updated. This alignment is essential to ensure the accurate synchronization of DNS records.


