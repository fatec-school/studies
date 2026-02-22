resource "kubernetes_namespace_v1" "keda" {
  metadata {
    name = "keda"
  }
}

resource "helm_release" "keda" {
  name       = "keda"
  namespace  = kubernetes_namespace_v1.keda.metadata[0].name
  repository = "https://kedacore.github.io/charts"
  chart      = "keda"

  values = [
    <<EOF
        replicaCount: 1
    EOF
  ]

  dependency_update = kubernetes_namespace_v1.keda
}
