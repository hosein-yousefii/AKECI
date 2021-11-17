# AKECI - Ansible, Kubernetes cluster, Envoy, Contour, Ingress
[![GitHub license](https://img.shields.io/github/license/hosein-yousefii/docker-ansible)](https://github.com/hosein-yousefii/docker-ansible/blob/master/LICENSE)
![LinkedIn](https://shields.io/badge/style-hoseinyousefii-black?logo=linkedin&label=LinkedIn&link=https://www.linkedin.com/in/hoseinyousefi)

Implementing **Kubernetes-cluster (Kind), Service proxy (Envoy), Ingress (Contour), Cert-Manager, Let's Encrypt and a simple Hello world's service** with an ansible playbook on a single-node in order to show how these services work with your application.

## What is kind?

kind is a tool for running local Kubernetes clusters using Docker container “nodes”.
kind was primarily designed for testing Kubernetes itself, but may be used for local development or CI.

## What is Contour?

Contour is an open source Kubernetes ingress controller providing the control plane for the Envoy edge and service proxy.

## What is Envoy?

Envoy is an open source edge and service proxy, designed for cloud-native applications.

## What is Cert-Manager?

cert-manager adds certificates and certificate issuers as resource types in Kubernetes clusters, and simplifies the process of obtaining, renewing and using those certificates.

It can issue certificates from a variety of supported sources, including Let’s Encrypt, HashiCorp Vault, and Venafi as well as private PKI.

## What is Let's Encrypt??

To enable HTTPS on your website, you need to get a certificate (a type of file) from a Certificate Authority (CA). Let’s Encrypt is a CA. In order to get a certificate for your website’s domain from Let’s Encrypt, you have to demonstrate control over the domain.


# Get started:

To clarify the repository, There is a **my-go-app** directory which contains our simple hello world's app, written in go, and a Dockerfile to create an Image for the app. Also you can put this directory to a separated repository to enable your **CI/CD** which I've putted the **GitlabCi** configuration on this directory too.

Other yaml files in the main repository are different related services which should be implemented. You should have a **Public domain and a Sub-domain** for testing this senario.
So, You just need to change the **email-address** in the  **letsencrypt-prod.yaml** to your related address then, chnage the **hosts** in the **ingress.yaml** to your sub-domain.

Moreover, there is a script **"check-pods-certificate.sh"** which will append to the crontab in order to renew the certificate of our program in a timely manner.

Anyway, ingress controller consists of a Pod that watches the Kubernetes Control Plane for new and updated Ingress Resource objects and An Ingress Resource is essentially a list of traffic routing rules for backend Services. for example, routing requests that hit **test.sample.com** to the backend Kubernetes Service **go-app**.

As You can see the entire implementation here.

![](1.png)


# HOW TO


After editing yaml files You are able to use **ansible-deploy.yml** to do the implementation automatically.

In order to deploy the cluster and related pods, you should insure that **docker and kubectl** are installed, then just configure your ansible hosts and run:

```bash
ansible-playbook -i hosts.ini ansible-deploy.yml
```

# How to contribute:

You are able to automate docker and kubectl installation. However, I have a repository to install and upgrade docker.[docker-ansible](https://github.com/hosein-yousefii/docker-ansible.git)

Another thing is that implementing multi nodes cluster would be a good idea.



Copyright 2021 Hosein Yousefi <yousefi.hosein.o@gmail.com>





</details>
