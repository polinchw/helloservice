# Helloservice

A Helloservice microservices.

The following repository will create you a Kubernetes cluster using:

1. [GKE](https://www.digitalocean.com/) for the underlying infrastructure.

After standing up the Kubernetes cluster we will deploy the following applications:

1. [Helloservice (Golang microservices application)](docs/1.md)

We will then proceed to cause Prometheus to fire alerts to Slack based upon interaction with Fotia.

## Prerequisites

A list of prerequisities for Mac can be found [here](docs/1.mac-prerequisities.md)

A list of prerequisities for Windows can be found [here](docs/2.windows-prerequisities.md)

## Cluster creation

A list of steps to build and provision the Kubernetes cluster can be found [here](docs/3.create-cluster.md)

## Helm chart deployment

A list of steps to deploy the necessary Helm charts can be found [here](docs/4.deploy-helloservice.md)

## Cleanup (from your local machine)

To remove all the droplets from digital ocean execute the following:

```
$ make destroy-cluster
```