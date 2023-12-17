# RxGo
![CI](https://github.com/ReactiveX/RxGo/actions/workflows/ci.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/reactivex/rxgo)](https://goreportcard.com/report/github.com/reactivex/rxgo)
[![Join the chat at https://gitter.im/ReactiveX/RxGo](https://badges.gitter.im/ReactiveX/RxGo.svg)](https://gitter.im/ReactiveX/RxGo?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

Reactive Extensions for the Go Language

## ReactiveX

[ReactiveX](http://reactivex.io/), or Rx for short, is an API for programming with Observable streams. This is the official ReactiveX API for the Go language.

ReactiveX is a new, alternative way of asynchronous programming to callbacks, promises, and deferred. It is about processing streams of events or items, with events being any occurrences or changes within the system. A stream of events is called an [Observable](http://reactivex.io/documentation/contract.html).

An operator is a function that defines an Observable, how and when it should emit data. The list of operators covered is available [here](README.md#supported-operators-in-rxgo).

## RxGo

The RxGo implementation is based on the concept of [pipelines](https://blog.golang.org/pipelines). A pipeline is a series of stages connected by channels, where each stage is a group of goroutines running the same function.

![](doc/rx.png)