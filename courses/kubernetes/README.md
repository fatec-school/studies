# Kubernetes: From Basics to Advanced (Udemy)

This directory contains notes, code snippets, and experiments from the Udemy course **Kubernetes: Do Básico ao Avançado**, which focuses on understanding and mastering Kubernetes from fundamental concepts to advanced cluster operations. The course explores Kubernetes internals, emphasizing how the platform works under the hood and how to operate production-ready clusters.

## Course Summary

This course covers the Kubernetes ecosystem from foundational concepts to advanced operational practices, including:

* **Kubernetes Architecture Fundamentals:** Control Plane and Data Plane.
* **Cluster Creation:** Using tools such as `kind`, `kubeadm`, and `kubespray`.
* **State Management:** Managing cluster state and backups using `etcd` and `etcdctl`.
* **Manifests:** Writing and organizing Kubernetes manifests using **YAML**.
* **Organization:** Resource isolation and organization using **Namespaces**.
* **Workloads & Pods:**
  * Working with **Pods** (init containers, multi-container pods, and lifecycle hooks).
  * Managing workloads with **Deployments**, **StatefulSets**, **DaemonSets**, **Jobs**, and **CronJobs**.
* **Networking & Services:**
  * Configuring **Services** (ClusterIP, NodePort, LoadBalancer, ExternalName).
  * Exposing applications using **Ingress Controllers**.
  * Understanding CNI plugins, DNS resolution with **CoreDNS**, and traffic flow.
* **Configuration & Storage:**
  * Managing configuration and secrets with **ConfigMaps** and **Secrets**.
  * Persistent storage concepts: **Persistent Volumes (PV)**, **Persistent Volume Claims (PVC)**, and **StorageClasses**.
* **Scaling:** Application scaling using **Horizontal Pod Autoscaler (HPA)** and **Vertical Pod Autoscaler (VPA)** concepts.
* **Security:** Access control using **RBAC** and **TLS certificates**.
* **Advanced Scheduling:** Strategies such as **Node Affinity**, **Pod Affinity**, **Taints**, and **Tolerations**.
* **Maintenance & Operations:**
  * Cluster lifecycle management, including **Kubernetes upgrades**.
  * Real-world **troubleshooting** techniques for common cluster and workload issues.

The course focuses on a deep technical understanding of Kubernetes internals and operations, preparing students to manage container orchestration platforms in real-world **DevOps** and **SRE** environments.
