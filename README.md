

# postgress-gorm-operator
Created to play around Kubebuilder and Kubernetes operators .
Kubernetes Operator with Postgres DB and using the GORM framework . As of today it reads the PostgresStore CR and then creates entry in the database "quests" in table quests if the entry is not already present .

References : 
https://book.kubebuilder.io/quick-start

https://tamerlan.dev/creating-go-rest-api/ 


Setup the Postgress DB :
docker run --name postgres_db_operator  -p 5433:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=quests -d postgres


1. kubebuilder init --domain venkateshredhat.com --repo github.com/venkateshsredhat/postgress-gorm-operator
2. kubebuilder create api \                                                                                
--group postgressgroup \
--version v1 \
--kind PostgresStore \
--resource true \
--controller true \
--namespaced true
3. make manifests

make install
/home/ves/Desktop/demo/postgress-gorm-operator/bin/controller-gen rbac:roleName=manager-role crd webhook paths="./..." output:crd:artifacts:config=config/crd/bases
Downloading sigs.k8s.io/kustomize/kustomize/v5@v5.4.2
go: downloading sigs.k8s.io/kustomize/kustomize/v5 v5.4.2
go: downloading sigs.k8s.io/kustomize/kyaml v0.17.1
go: downloading sigs.k8s.io/kustomize/api v0.17.2
go: downloading sigs.k8s.io/kustomize/cmd/config v0.14.1
go: downloading gopkg.in/evanphx/json-patch.v4 v4.12.0
/home/ves/Desktop/demo/postgress-gorm-operator/bin/kustomize build config/crd | kubectl apply -f -
customresourcedefinition.apiextensions.k8s.io/postgresstores.postgressgroup.venkateshredhat.com created





Log Output :

make run
/home/ves/Desktop/demo/postgress-gorm-operator/bin/controller-gen rbac:roleName=manager-role crd webhook paths="./..." output:crd:artifacts:config=config/crd/bases
/home/ves/Desktop/demo/postgress-gorm-operator/bin/controller-gen object:headerFile="hack/boilerplate.go.txt" paths="./..."
go fmt ./...
go vet ./...
go run ./cmd/main.go
2024-08-16T17:53:37+05:30	INFO	setup	starting manager
2024-08-16T17:53:37+05:30	INFO	starting server	{"name": "health probe", "addr": "[::]:8081"}
2024-08-16T17:53:37+05:30	INFO	Starting EventSource	{"controller": "postgresstore", "controllerGroup": "postgressgroup.venkateshredhat.com", "controllerKind": "PostgresStore", "source": "kind source: *v1.PostgresStore"}
2024-08-16T17:53:37+05:30	INFO	Starting Controller	{"controller": "postgresstore", "controllerGroup": "postgressgroup.venkateshredhat.com", "controllerKind": "PostgresStore"}
2024-08-16T17:53:37+05:30	INFO	Starting workers	{"controller": "postgresstore", "controllerGroup": "postgressgroup.venkateshredhat.com", "controllerKind": "PostgresStore", "worker count": 1}
Already present in database {26 hello}

2024/08/16 17:54:21 /home/ves/Desktop/demo/postgress-gorm-operator/internal/controller/postgresstore_controller.go:70 record not found
[0.135ms] [rows:0] SELECT * FROM "quests" WHERE id = 36 ORDER BY "quests"."id" LIMIT 1
quest not found :  record not found
Quest created successfully
2024-08-16T17:54:21+05:30	ERROR	Reconciler error	{"controller": "postgresstore", "controllerGroup": "postgressgroup.venkateshredhat.com", "controllerKind": "PostgresStore", "PostgresStore": {"name":"postgresstore-sample","namespace":"default"}, "namespace": "default", "name": "postgresstore-sample", "reconcileID": "78439260-2245-4d3f-adfd-fdf774ad03e1", "error": "record not found"}
sigs.k8s.io/controller-runtime/pkg/internal/controller.(*Controller).reconcileHandler
	/home/ves/go/pkg/mod/sigs.k8s.io/controller-runtime@v0.18.4/pkg/internal/controller/controller.go:324
sigs.k8s.io/controller-runtime/pkg/internal/controller.(*Controller).processNextWorkItem
	/home/ves/go/pkg/mod/sigs.k8s.io/controller-runtime@v0.18.4/pkg/internal/controller/controller.go:261
sigs.k8s.io/controller-runtime/pkg/internal/controller.(*Controller).Start.func2.2
	/home/ves/go/pkg/mod/sigs.k8s.io/controller-runtime@v0.18.4/pkg/internal/controller/controller.go:222
Already present in database {36 hello}



kubectl apply -k config/samples/
postgresstore.postgressgroup.venkateshredhat.com/postgresstore-sample configured


Step by step DB output :

docker exec -it postgres_db_operator /bin/bash
root@9f85d1c039d3:/# psql
psql: error: connection to server on socket "/var/run/postgresql/.s.PGSQL.5432" failed: FATAL:  role "root" does not exist
root@9f85d1c039d3:/# psql -h localhost -U postgres quests
psql (16.0 (Debian 16.0-1.pgdg120+1))
Type "help" for help.

1. quests=# \dt
Did not find any relations.

2. quests=# \dt
         List of relations
 Schema |  Name  | Type  |  Owner   
--------+--------+-------+----------
 public | quests | table | postgres
(1 row)


3. 
quests=# select * from quests;
 id | title 
----+-------
(0 rows)

quests=# select * from quests;
 id | title 
----+-------
 26 | hello
(1 row)

4.
quests=# select * from quests;
 id | title 
----+-------
 26 | hello
 36 | hello
(2 rows)



## Getting Started

### Prerequisites
- go version v1.22.0+
- docker version 17.03+.
- kubectl version v1.11.3+.
- Access to a Kubernetes v1.11.3+ cluster.

### To Deploy on the cluster
**Build and push your image to the location specified by `IMG`:**

```sh
make docker-build docker-push IMG=<some-registry>/postgress-gorm-operator:tag
```

**NOTE:** This image ought to be published in the personal registry you specified.
And it is required to have access to pull the image from the working environment.
Make sure you have the proper permission to the registry if the above commands don’t work.

**Install the CRDs into the cluster:**

```sh
make install
```

**Deploy the Manager to the cluster with the image specified by `IMG`:**

```sh
make deploy IMG=<some-registry>/postgress-gorm-operator:tag
```

> **NOTE**: If you encounter RBAC errors, you may need to grant yourself cluster-admin
privileges or be logged in as admin.

**Create instances of your solution**
You can apply the samples (examples) from the config/sample:

```sh
kubectl apply -k config/samples/
```

>**NOTE**: Ensure that the samples has default values to test it out.

### To Uninstall
**Delete the instances (CRs) from the cluster:**

```sh
kubectl delete -k config/samples/
```

**Delete the APIs(CRDs) from the cluster:**

```sh
make uninstall
```

**UnDeploy the controller from the cluster:**

```sh
make undeploy
```

## Project Distribution

Following are the steps to build the installer and distribute this project to users.

1. Build the installer for the image built and published in the registry:

```sh
make build-installer IMG=<some-registry>/postgress-gorm-operator:tag
```

NOTE: The makefile target mentioned above generates an 'install.yaml'
file in the dist directory. This file contains all the resources built
with Kustomize, which are necessary to install this project without
its dependencies.

2. Using the installer

Users can just run kubectl apply -f <URL for YAML BUNDLE> to install the project, i.e.:

```sh
kubectl apply -f https://raw.githubusercontent.com/<org>/postgress-gorm-operator/<tag or branch>/dist/install.yaml
```

## Contributing
// TODO(user): Add detailed information on how you would like others to contribute to this project

**NOTE:** Run `make help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

