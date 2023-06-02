// Generated by *go generate* - DO NOT EDIT
/*
Copyright 2022-23. projectsveltos.io. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package crd

var EventSourceFile = "../../config/crd/bases/lib.projectsveltos.io_eventsources.yaml"
var EventSourceCRD = []byte(`---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.11.3
  creationTimestamp: null
  name: eventsources.lib.projectsveltos.io
spec:
  group: lib.projectsveltos.io
  names:
    kind: EventSource
    listKind: EventSourceList
    plural: eventsources
    singular: eventsource
  scope: Cluster
  versions:
  - name: v1alpha1
    schema:
      openAPIV3Schema:
        description: EventSource is the Schema for the EventSource API
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: EventSourceSpec defines the desired state of EventSource
            properties:
              collectResources:
                default: false
                description: CollectResources indicates whether matching resources
                  need to be collected and added to EventReport.
                type: boolean
              group:
                description: Group of the resource deployed in the Cluster.
                type: string
              kind:
                description: Kind of the resource deployed in the Cluster.
                minLength: 1
                type: string
              labelFilters:
                description: LabelFilters allows to filter resources based on current
                  labels.
                items:
                  properties:
                    key:
                      description: Key is the label key
                      type: string
                    operation:
                      description: Operation is the comparison operation
                      enum:
                      - Equal
                      - Different
                      type: string
                    value:
                      description: Value is the label value
                      type: string
                  required:
                  - key
                  - operation
                  - value
                  type: object
                type: array
              namespace:
                description: Namespace of the resource deployed in the  Cluster. Empty
                  for resources scoped at cluster level.
                type: string
              script:
                description: Script is a text containing a lua script. Must return
                  struct with field "matching" representing whether object is a match.
                type: string
              version:
                description: Version of the resource deployed in the Cluster.
                type: string
            required:
            - group
            - kind
            - version
            type: object
        type: object
    served: true
    storage: true
`)
