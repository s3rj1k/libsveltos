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

var ReloaderFile = "../../manifests/apiextensions.k8s.io_v1_customresourcedefinition_reloaders.lib.projectsveltos.io.yaml"
var ReloaderCRD = []byte(`apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.15.0
  name: reloaders.lib.projectsveltos.io
spec:
  group: lib.projectsveltos.io
  names:
    kind: Reloader
    listKind: ReloaderList
    plural: reloaders
    singular: reloader
  scope: Cluster
  versions:
  - name: v1alpha1
    schema:
      openAPIV3Schema:
        description: Reloader is the Schema for the Reloader API
        properties:
          apiVersion:
            description: |-
              APIVersion defines the versioned schema of this representation of an object.
              Servers should convert recognized schemas to the latest internal value, and
              may reject unrecognized values.
              More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
            type: string
          kind:
            description: |-
              Kind is a string value representing the REST resource this object represents.
              Servers may infer this from the endpoint the client submits requests to.
              Cannot be updated.
              In CamelCase.
              More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
            type: string
          metadata:
            type: object
          spec:
            description: ReloaderSpec defines the desired state of Reloader
            properties:
              reloaderInfo:
                items:
                  description: |-
                    ReloaderInfo represents a resource that need to be reloaded
                    if any mounted ConfigMap/Secret changes.
                  properties:
                    kind:
                      description: 'Kind of the resource. Supported kinds are: Deployment
                        StatefulSet DaemonSet.'
                      enum:
                      - Deployment
                      - StatefulSet
                      - DaemonSet
                      type: string
                    name:
                      description: Name of the referenced resource.
                      minLength: 1
                      type: string
                    namespace:
                      description: Namespace of the referenced resource.
                      minLength: 1
                      type: string
                    value:
                      type: string
                  required:
                  - kind
                  - name
                  - namespace
                  type: object
                type: array
            type: object
        type: object
    served: true
    storage: false
  - name: v1beta1
    schema:
      openAPIV3Schema:
        description: Reloader is the Schema for the Reloader API
        properties:
          apiVersion:
            description: |-
              APIVersion defines the versioned schema of this representation of an object.
              Servers should convert recognized schemas to the latest internal value, and
              may reject unrecognized values.
              More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
            type: string
          kind:
            description: |-
              Kind is a string value representing the REST resource this object represents.
              Servers may infer this from the endpoint the client submits requests to.
              Cannot be updated.
              In CamelCase.
              More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
            type: string
          metadata:
            type: object
          spec:
            description: ReloaderSpec defines the desired state of Reloader
            properties:
              reloaderInfo:
                items:
                  description: |-
                    ReloaderInfo represents a resource that need to be reloaded
                    if any mounted ConfigMap/Secret changes.
                  properties:
                    kind:
                      description: 'Kind of the resource. Supported kinds are: Deployment
                        StatefulSet DaemonSet.'
                      enum:
                      - Deployment
                      - StatefulSet
                      - DaemonSet
                      type: string
                    name:
                      description: Name of the referenced resource.
                      minLength: 1
                      type: string
                    namespace:
                      description: Namespace of the referenced resource.
                      minLength: 1
                      type: string
                    value:
                      type: string
                  required:
                  - kind
                  - name
                  - namespace
                  type: object
                type: array
            type: object
        type: object
    served: true
    storage: true
`)
