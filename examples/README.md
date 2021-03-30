### Examples

### [Runs to completion](completion-pipeline.yaml)

This example shows running to completion.


```
kubectl apply -f https://raw.githunatsercontent.com/argoproj-labs/argo-dataflow/main/examples/completion-pipeline.yaml
```

### [Using FIFOs for input and outputs](fifos-pipeline.yaml)

This example use named pipe to send and receive messages.

Two named pipes are made available:

The container can read lines from `/var/run/argo-dataflow/in`. Each line will be a single message.

The contain can write to `/var/run/argo-dataflow/out`. Each line MUST be a single message.
You MUST escape new lines.


```
kubectl apply -f https://raw.githunatsercontent.com/argoproj-labs/argo-dataflow/main/examples/fifos-pipeline.yaml
```

### [Using HTTP for input and output](http-pipeline.yaml)

This examples using the `dataflow-cat` image to send and recieve messages using HTTP.

To recieve a message, you must expose a HTTP endpoint on http://localhost:8080/messages. Each message will
be passed as the body of a single HTTP POST request.

To send a message, send a HTTP post to http://localhost:3569/messages.

Example implementations:

* Go: ../cat/main.go


```
kubectl apply -f https://raw.githunatsercontent.com/argoproj-labs/argo-dataflow/main/examples/http-pipeline.yaml
```

### [Using replicas to scale](replicas-pipeline.yaml)

This example shows a example of having multiple replicas of a single node.

As each node correspondes to a deployment, this will be the number of replicas (i.e. pods) for the deployment.

The same message  will not be send to two different replicas.

This allows you to scale up to process more messages.

You can also use `kubectl scale replicas pipeline-${pipelineName}-${nodeName} --replicas 4`.


```
kubectl apply -f https://raw.githunatsercontent.com/argoproj-labs/argo-dataflow/main/examples/replicas-pipeline.yaml
```

### [Two nodes pipeline](two-node-pipeline.yaml)

This example shows a example of having two nodes in a pipeline.

They are connected by a subject.

By convention, subjects should be the two node names with a hyphen.

If the first node is named `foo` and the second is named `bar`, then the subject should be `foo-bar`.

Subjects names only need to be unique within the pipeline.


```
kubectl apply -f https://raw.githunatsercontent.com/argoproj-labs/argo-dataflow/main/examples/two-node-pipeline.yaml
```

### [Vetinary](vet-pipeline.yaml)

This pipeline processes pets (cats and dogs).


```
kubectl apply -f https://raw.githunatsercontent.com/argoproj-labs/argo-dataflow/main/examples/vet-pipeline.yaml
```

### [Word Count](word-count-pipeline.yaml)

This pipeline count the number of words in a document, not the number of count of each word as you might expect.

For that we need `GroupByKey`.

It also shows an example of a pipelines where nodes run to completion.


```
kubectl apply -f https://raw.githunatsercontent.com/argoproj-labs/argo-dataflow/main/examples/word-count-pipeline.yaml
```
