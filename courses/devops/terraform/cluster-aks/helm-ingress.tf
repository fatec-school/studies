resource "kubernetes_namespace_v1" "ingress" {
  metadata {
    name = "ingress-nginx"
  }
}

resource "helm_release" "nginx_release" {
  name       = "ingress-nginx"
  namespace  = kubernetes_namespace_v1.ingress.metadata[0].name
  repository = "https://kubernetes.github.io/ingress-nginx"
  chart      = "ingress-nginx"

  values = [
    <<EOF
    controller:
        replicaCount: 2
        nodeSelector:
          "kubernetes.io/os": linux
        service:
          type: LoadBalancer
          externalTrafficPolicy: Local
          loadBalancerIP: ${azurerm_public_ip.ingress_ip.ip_address}

        defaultBackend:
            nodeSelector:
              "kubernetes.io/os": linux
    EOF
  ]

  depends_on = [azurerm_public_ip.ingress_ip]
}
