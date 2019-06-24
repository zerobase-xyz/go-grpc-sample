#!/bin/bash

kubectl apply -f test.yaml 2>> err_text.txt
kubectl apply -f namespace.yaml 2>> err_namespace.txt
kubectl apply -f alb_ingress.yaml 2>> err_alb_ingress.txt
kubectl apply -f app.yaml -n not-inject-envoy 2>> err_not_inject_envoy.txt
kubectl apply -f app.yaml -n inject-envoy 2>> err_inject_envoy.txt
kubectl apply -f app.yaml -n use_istio_ingressgw 2>> err_use_istio_ingressgw.txt
kubectl apply -f istio_gw.yaml 2>> err_istio_gw.txt
