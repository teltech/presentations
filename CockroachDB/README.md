# Start Kuberenetes cluster
`gcloud container clusters create cockroachdb`

# Create resources
`kubectl create -f https://raw.githubusercontent.com/mlesar/CockroachDB_on_Kubernetes/master/cockroachdb-statefulset.yaml`

# Initialize cluster
`kubectl create -f https://raw.githubusercontent.com/mlesar/CockroachDB_on_Kubernetes/master/cluster-init.yaml`

# Connect to the cluster
`kubectl run cockroachdb -it --image=cockroachdb/cockroach --rm --restart=Never -- sql --insecure --host=cockroachdb-public`

# Forward port for the Admin
`kubectl port-forward cockroachdb-0 8080`

# Scale cluster
`kubectl scale statefulset cockroachdb --replicas=4`

# Node status
```
kubectl run cockroachdb -it --image=cockroachdb/cockroach --rm --restart=Never \
-- node status --insecure --host=cockroachdb-public
```

# Decomission node
```
kubectl run cockroachdb -it --image=cockroachdb/cockroach --rm --restart=Never \
-- node decommission 1 --insecure --host=cockroachdb-public
```

# Load test
* https://www.cockroachlabs.com/docs/stable/cockroach-workload.html
```
kubectl run cockroachdb -it --image=cockroachdb/cockroach --rm --restart=Never \
-- workload run kv --init --drop --tolerate-errors --max-rate=100 --duration=2m 'postgresql://root@cockroachdb-public:26257?sslmode=disable'

kubectl run cockroachdb -it --image=cockroachdb/cockroach --rm --restart=Never \
-- workload run bank --init --drop --tolerate-errors --max-rate=1000 --duration=1m 'postgresql://root@cockroachdb-public:26257?sslmode=disable'

kubectl run cockroachdb -it --image=cockroachdb/cockroach --rm --restart=Never \
-- workload run tpcc --init --drop --ramp=30s --warehouses=1000 --duration=2m --split --scatter --tolerate-errors 'postgresql://root@cockroachdb-public:26257?sslmode=disable'
```


# Clean
`kubectl delete pods,statefulsets,services,persistentvolumeclaims,persistentvolumes,poddisruptionbudget,jobs,rolebinding,clusterrolebinding,role,clusterrole,serviceaccount -l app=cockroachdb`

# Delete cluster
`gcloud container clusters delete cockroachdb`