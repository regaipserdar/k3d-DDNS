# Kubernetes Cluster with Dynamic DNS using K3d and Cloudflare

This project aims to create a Kubernetes cluster with dynamic DNS (DDNS) using K3d and Cloudflare, allowing external access without a static IP address. 

## Overview

Kubernetes is a powerful container orchestration platform, but it can be challenging to provide external access to services running within a cluster when you don't have a static IP address. This project demonstrates how to set up a Kubernetes cluster using K3d and configure dynamic DNS using Cloudflare to overcome this limitation.

## Prerequisites

Before getting started, make sure you have the following prerequisites in place:

- [K3d](https://k3d.io/) installed on your local machine
- A [Cloudflare](https://www.cloudflare.com/) account
- Docker or Rancher 
- Helm Chard must be installed on your local machine
- A registered domain or subdomain managed by Cloudflare
- API token with the necessary permissions for Cloudflare DNS management

## Usage

1. **Create a Kubernetes Cluster**

   Use K3d to create a Kubernetes cluster. For example:

    ```bash
   k3d cluster create recon
    ```


1.1  **Create a Kubernetes Cluster**
![Cluster-Create](/static/images/cluster-create.png)

## Set up ExternalDNS

1. **Install ExternalDNS with Helm:**

   Use Helm to install ExternalDNS into your Kubernetes cluster. Make sure Helm is installed and configured before proceeding.

   ```bash
   helm repo add bitnami https://charts.bitnami.com/bitnami
   helm install external-dns bitnami/external-dns --namespace external-dns --set provider=cloudflare --set cloudflare.apiToken=YOUR_CLOUDFLARE_API_TOKEN --set txtOwnerId=my-identifier

    ```

    Replace *YOUR_CLOUDFLARE_API_TOKEN* with your Cloudflare API token. You can customize the txtOwnerId to a unique identifier.


