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

var AccessRequestFile = "../../manifests/apiextensions.k8s.io_v1_customresourcedefinition_accessrequests.lib.projectsveltos.io.yaml"
var AccessRequestCRD = []byte(`apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.16.3
  name: accessrequests.lib.projectsveltos.io
spec:
  group: lib.projectsveltos.io
  names:
    kind: AccessRequest
    listKind: AccessRequestList
    plural: accessrequests
    singular: accessrequest
  scope: Namespaced
  versions:
  - name: v1alpha1
    schema:
      openAPIV3Schema:
        description: AccessRequest is the Schema for the accessrequest API
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
            description: AccessRequestSpec defines the desired state of AccessRequest
            properties:
              controlPlaneEndpoint:
                description: |-
                  ControlPlaneEndpoint represents the endpoint used to communicate with the
                  management cluster controlplane endpoint. It will be used when generating the
                  kubeconfig.
                properties:
                  host:
                    description: The hostname on which the API server is serving.
                    type: string
                  port:
                    description: The port on which the API server is serving.
                    format: int32
                    type: integer
                required:
                - host
                - port
                type: object
              name:
                description: |-
                  Name is the name of the service account created
                  for this AccessRequest
                type: string
              namespace:
                description: |-
                  Namespace is the namespace of the service account created
                  for this AccessRequest
                type: string
              type:
                description: Type represent the type of the request
                enum:
                - SveltosAgent
                - Different
                type: string
            required:
            - controlPlaneEndpoint
            - name
            - namespace
            - type
            type: object
          status:
            description: AccessRequestStatus defines the status of AccessRequest
            properties:
              failureMessage:
                description: FailureMessage provides more information if an error
                  occurs.
                type: string
              secretRef:
                description: SecretRef points to the Secret containing Kubeconfig
                properties:
                  apiVersion:
                    description: API version of the referent.
                    type: string
                  fieldPath:
                    description: |-
                      If referring to a piece of an object instead of an entire object, this string
                      should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].
                      For example, if the object reference is to a container within a pod, this would take on a value like:
                      "spec.containers{name}" (where "name" refers to the name of the container that triggered
                      the event) or if no container name is specified "spec.containers[2]" (container with
                      index 2 in this pod). This syntax is chosen only to have some well-defined way of
                      referencing a part of an object.
                    type: string
                  kind:
                    description: |-
                      Kind of the referent.
                      More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
                    type: string
                  name:
                    description: |-
                      Name of the referent.
                      More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
                    type: string
                  namespace:
                    description: |-
                      Namespace of the referent.
                      More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
                    type: string
                  resourceVersion:
                    description: |-
                      Specific resourceVersion to which this reference is made, if any.
                      More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency
                    type: string
                  uid:
                    description: |-
                      UID of the referent.
                      More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids
                    type: string
                type: object
                x-kubernetes-map-type: atomic
            type: object
        type: object
    served: true
    storage: false
    subresources:
      status: {}
  - name: v1beta1
    schema:
      openAPIV3Schema:
        description: AccessRequest is the Schema for the accessrequest API
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
            description: AccessRequestSpec defines the desired state of AccessRequest
            properties:
              controlPlaneEndpoint:
                description: |-
                  ControlPlaneEndpoint represents the endpoint used to communicate with the
                  management cluster controlplane endpoint. It will be used when generating the
                  kubeconfig.
                properties:
                  host:
                    description: The hostname on which the API server is serving.
                    type: string
                  port:
                    description: The port on which the API server is serving.
                    format: int32
                    type: integer
                required:
                - host
                - port
                type: object
              name:
                description: |-
                  Name is the name of the service account created
                  for this AccessRequest
                type: string
              namespace:
                description: |-
                  Namespace is the namespace of the service account created
                  for this AccessRequest
                type: string
              type:
                description: Type represent the type of the request
                enum:
                - SveltosAgent
                - Different
                type: string
            required:
            - controlPlaneEndpoint
            - name
            - namespace
            - type
            type: object
          status:
            description: AccessRequestStatus defines the status of AccessRequest
            properties:
              failureMessage:
                description: FailureMessage provides more information if an error
                  occurs.
                type: string
              secretRef:
                description: SecretRef points to the Secret containing Kubeconfig
                properties:
                  apiVersion:
                    description: API version of the referent.
                    type: string
                  fieldPath:
                    description: |-
                      If referring to a piece of an object instead of an entire object, this string
                      should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].
                      For example, if the object reference is to a container within a pod, this would take on a value like:
                      "spec.containers{name}" (where "name" refers to the name of the container that triggered
                      the event) or if no container name is specified "spec.containers[2]" (container with
                      index 2 in this pod). This syntax is chosen only to have some well-defined way of
                      referencing a part of an object.
                    type: string
                  kind:
                    description: |-
                      Kind of the referent.
                      More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
                    type: string
                  name:
                    description: |-
                      Name of the referent.
                      More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
                    type: string
                  namespace:
                    description: |-
                      Namespace of the referent.
                      More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
                    type: string
                  resourceVersion:
                    description: |-
                      Specific resourceVersion to which this reference is made, if any.
                      More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency
                    type: string
                  uid:
                    description: |-
                      UID of the referent.
                      More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids
                    type: string
                type: object
                x-kubernetes-map-type: atomic
            type: object
        type: object
    served: true
    storage: true
    subresources:
      status: {}
`)
