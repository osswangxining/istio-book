# one-big-payload-backend-sample

1. Follow the steps belwo to prepare the sample application deployment.

    -  try to clean the folder target by running `./mvnw clean`
    -  build the package by running `./mvnw package`
    -  make sure that `one-big-payload-backend-sample-0.1.0.jar` is created under the folder target and then build the docker image by running `docker build -t <image> .` , e.g. `docker build -t registry.cn-beijing.aliyuncs.com/aliyun-asm/bigpayloadtest:v0.1 .`
    -  push to Docker image repository by running `docker push registry.cn-beijing.aliyuncs.com/aliyun-asm/bigpayloadtest:v0.1`
    -  replace the image in deploy.yaml using the above.

    Now the YAML manifest is ready to push into the Kubernetes cluster.  

2. Make sure that the service and the pod are running like below.

    ```
    kubectl get pod |grep bigpayload
    bigpayloadtest-7bf88b76dc-jkj74   2/2     Running   0          2m
    ```

    ```
    kubectl get service |grep bigpayload
    bigpayloadtest   ClusterIP   172.21.11.62    <none>        8080/TCP   2m23s
    ```

3. create the Gateway as below.
    ```
    apiVersion: networking.istio.io/v1alpha3
    kind: Gateway
    metadata:
    name: mysamplegateway
    namespace: istio-system
    spec:
    selector:
        istio: ingressgateway
    servers:
        - hosts:
            - '*'
        port:
            name: http
            number: 80
            protocol: HTTP

    ```    

4. create VirtualService as below.
    ```
    apiVersion: networking.istio.io/v1alpha3
    kind: VirtualService
    metadata:
    name: bigpayload-vs
    namespace: istio-system
    spec:
    gateways:
        - mysamplegateway
    hosts:
        - '*'
    http:
        - route:
            - destination:
                host: bigpayloadtest.default.svc.cluster.local
                port:
                number: 8080
    ```

5. create 10MB test file
```
dd if=/dev/zero of=output.file bs=1024 count=10240
```

6. run the testing as below
   From [envoy official website](https://www.envoyproxy.io/docs/envoy/latest/api-v2/config/filter/network/http_connection_manager/v2/http_connection_manager.proto):
   > The maximum request headers size for incoming connections. If unconfigured, the default max request headers allowed is 60 KiB. Requests that exceed this limit will receive a 431 response. The max configurable limit is 96 KiB, based on current implementation constraints..

   make sure your curl version is above 7.55.0, then you can use -H @filename to add http headers as below.
```

curl -svo /dev/null -F 'file=@output.file' -H 'Cache-Control: no-cache' -H @bigheader.txt http://<ingressgateway-ipaddress>/scan 
```

1. check the log from sample application. You should see as below.

```
2020.05.25.11.52.15
File Name is: output.file
File Name is: 10240KB or about 10MB
```