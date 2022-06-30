# Prometheus
```
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
```
```
helm install prometheus prometheus-community/prometheus
```

```
kubectl expose service prometheus-server --type=NodePort --target-port=9090 --name=prometheus-server-np
```

# Grafana

```
helm repo add grafana https://grafana.github.io/helm-charts
```

```
helm install grafana stable/grafana
```
```
kubectl expose service grafana --type=NodePort --target-port=3000 --name=grafana-np
```

```
kubectl get secret --namespace default grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo
```
