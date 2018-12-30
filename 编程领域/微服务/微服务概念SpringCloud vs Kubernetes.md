# 微服务概念SpringCloud vs Kubernetes

原文链接 [微服务Spring Cloud与Kubernetes比较](https://www.jdon.com/48605)

Spring Cloud或Kubernetes都宣称它们是开发运行微服务的最好环境，哪个更好？答案是两个都是，但他们拥有各自不同的特征方式。

## 1. 背影故事

最近，Lukyanchikov发表了一篇使用Spring Cloud和Docker建立微服务架构的[文章](https://dzone.com/articles/microservice-architecture-with-spring-cloud-and-do)。 它给出了如何使用Spring Cloud创建一个简单的基于微服务的系统所需的全面概述。 为构建一个可扩展到数十或数百个服务的伸缩弹性的微服务系统，必须借助具有宽泛的构建时间和运行能力的工具集进行集中管理和管理。 Spring Cloud包括实现功能性服务（如统计服务，帐户服务和通知服务）和支持基础设施服务（如日志分析，配置服务器，服务发现，授权服务）。

这些服务涵盖了系统的运行时各个方面，但不涉及打包，持续集成，扩展，高可用性和自我修复，这在MSA(微服务架构)世界中也非常重要。 假设大多数Java开发人员熟悉Spring Cloud，在本文中，我们将绘制一个平行图，通过解决这些额外的问题，了解Kubernetes与Spring Cloud的关系。

有关MSA的好处是，它有一种架构风格很好理解的好处，微服务可实现强大的模块边界，具有独立的部署和技术多样性。 但代价是开发分布式系统成本和显著运营开销 。一个关键的成功因素是使用各种工具解决这些问题.

## 2. 微服务关注点

微服务关注方面有: 配置管理、服务发现与负载平衡、弹性和失败冗余、API管理、安全服务、中央集中日志、集中测量、分布式跟踪、调度部署和自动扩展self Healing等几个方面。

根据这些观点，得出Spring Cloud和Kubernetes两个平台映射: 

|         概念         | Spring Cloud                         | Kubernetes                             |
|:--------------------:|:-------------------------------------|:---------------------------------------|
|       配置管理       | 配置服务器、Consul和Netflix Archaius | ConfigMap&Secrets                      |
|       服务发现       | Netflix Eureka,Hashicorp Consul      | Service&Ingress Resource               |
|       负载平衡       | Netflix Ribbon                       | Service                                |
|       API网关        | Netflix Zuul                         | Service&Ingress Resource               |
|       安全服务       | SpringCloud Security                 | (无)                                   |
|     中央集中日志     | ELK Stack（LogStash）                | EFKstack（Fluentd）                    |
|       集中测量       | Netflix Spectator& Atlas             | Heapster、Prometheus、Grafana。        |
|      分布式跟踪      | SpringCloud Sleuth，Zipkin           | OpenTracing、Zipkin                    |
|    弹性和失败冗余    | Netflix Hystrix、Turbine&Ribbon      | Health Check&resource isolation        |
| 自动扩展self Healing | (无)                                 | Health Check、SelfHealing、Autoscaling |
| 打包, 部署和调度部署 | Spring Boot；Docker                  | Rkt、Kubernetes Scheduler&Deployment   |
|     任务工作管理     | Spring Batch                         | Jobs&Scheduled Jobs                    |
|       单个应用       | Spring Cloud Cluster                 | Kubernetes Pods                        |

Spring Cloud有一套丰富的集成良好的Java库，作为应用程序栈一部分解决所有运行时问题。 因此，微服务本身通过库和运行时作为代理来执行客户端服务发现，负载平衡，配置更新，度量跟踪等。诸如单例集群服务和批处理作业的模式也在JVM中进行管理。

Kubernetes是多语言的，不仅针对Java平台，并以通用的方式为所有语言解决分布式计算的挑战。 它提供应用程序栈外部的配置管理，服务发现，负载平衡，跟踪，度量，单例，平台调度作业等平台级别功能。 该应用系统不需要任何库或代理程序用于客户端逻辑，它可以用任何语言编写。

在某些方面，这两个平台都依赖类似的第三方工具。例如，ELK和EFK堆栈，跟踪库等一些库，如Hystrix和Spring Boot，在这两种环境中同样有用。

有些情况下这两个平台是互补的，并且可以结合在一起，创造一个更加强大的解决方案,例如，Spring Boot提供了用于构建单个JAR应用程序包的Maven插件。结合Docker和Kubernetes的声明性部署和调度功能，使微服务运行变得轻而易举。 类似地，Spring Cloud具有应用程序库，用于使用Hystrix（断路器）和Ribbon（用于负载平衡）创建弹性的，容错的微服务。 但是单单这是不够的，当它与Kubernetes的健康检查，进程重新启动和自动扩展功能相结合时，微服务成为一个真正的抗脆弱的系统。

## 3. 长处和短处

由于两个平台不具有直接的可比性特征，下面是逐项总结其优点和缺点。

### 3.1 Spring Cloud

Spring Cloud为开发人员提供了快速构建分布式系统中的一些常见模式的工具，例如配置管理，服务发现，断路器，路由等。它是为Java开发人员使用，构建在Netflix OSS库之上的。

**优势**

1. Spring Platform提供的统一编程模型和Spring Boot的快速应用程序创建能力为开发人员提供了巨大的微服务开发体验。 例如，使用很少的注释，您可以创建一个配置服务器，并且几乎没有更多的注释，您可以获得客户端库来配置您的服务。

2. 有丰富的库选择，覆盖大多数运行时关注。由于所有库都是用Java编写的，它提供了多种功能，更好的控制和精细调整选项。

3. 不同的Spring Cloud库彼此完全集成。例如，Feign客户端还将使用Hystrix用于断路器，并且Ribbon用于负载平衡请求。 一切都是注释驱动的，使其易于为Java开发人员开发。

**弱点**

1. Spring Cloud的一个主要优点是它的缺点 - 它仅限于Java。MSA的强大动力是在需要时交换各种技术栈，库，甚至语言的能力。 只是使用Spring Cloud是不可能的。 如果您想要使用Spring Cloud / Netflix OSS基础架构服务（如配置管理，服务发现或负载平衡），那么解决方案就不那么优雅。 在Netflix的 Prana项目通过基于HTTP暴露Java客户端实现了sidecar模式，使其可能让非JVM语言运行在NetflixOSS生态系统中，但它不是很优雅。

2. Java开发人员关心Java应用程序并需要处理太多与开发无关的事情。每个微服务需要运行各种客户端以进行配置检索，服务发现和负载平衡。虽然很容易设置，但这并不会降低对环境的构建时间和运行时依赖性。例如，开发人员可以使用@EnableConfigServer创建一个配置服务器，但这只是开心的假象。 每当开发人员想要运行单个微服务时，他们需要启动并运行Config Server。对于受控环境，开发人员必须考虑使Config Server高度可用，并且由于它可以由Git或Svn支持，因此它们需要一个共享文件系统。 类似地，对于服务发现，开发人员需要首先启动Eureka服务器。 为了创建一个受控的环境，他们需要在每个AZ上使用多个实例实现集群。像开发人员一样，除了实现所有功能服务之外，Java开发人员还必须构建和管理一个非平凡的微服务平台。

3. Spring Cloud在微服务发展过程只有很短历程，开发人员还需要考虑自动化部署，调度，资源管理，过程隔离，自我修复，构建管道等，以获得完整的微服务体验。 对于这点，我认为这是不公平的比较，应该比较 Spring Cloud + Cloud Foundry (or Docker Swarm) 和Kubernetes。但这也意味着对于一个完整的端到端微服务体验，Spring Cloud必须补充一个像Kubernetes本身这样的应用程序平台。

### 3.2 Kubernetes

Kubernetes是一个用于自动化部署，扩展和管理容器化应用程序的开源系统。 它是多种语言并且提供用于供应，运行，扩展和管理分布式系统的操作系统。

**优势**

1. Kubernetes是一个多语言和语言不可知的容器管理平台，能够运行云本地和传统的容器化应用程序。其提供的服务（如配置管理，服务发现，负载平衡，测量指标收集和日志聚合）可供各种语言使用。 这允许在一个组织中有一个平台，可以被多个团队（包括使用Spring的Java开发人员）使用，并提供多种用途：应​​用程序开发，测试环境，构建环境（运行源代码控制系统，构建服务器，工件存储库）等。

2. 与Spring Cloud相比，Kubernetes解决了更广泛的MSA问题。 除了提供运行时服务，Kubernetes也可以让你规定的环境中，设置资源限制，RBAC，管理应用程序生命周期，启用自动缩放和自我修复（几乎表现得像一个抗脆弱平台）。

3. Kubernetes技术基于Google 15年的研发经验和管理容器的经验。此外，有近1000个提交者，它是Github上最活跃的开源社区之一。

**弱点**

1. Kubernetes是多语言的，因此它的服务是通用的，并不针对不同的平台（如Spring Cloud for JVM）进行优化。 例如，配置作为环境变量或安装的文件系统传递到应用程序。 它没有Spring Cloud Config提供的奇特的配置更新功能。

2. Kubernetes不是一个以开发人员为中心的平台。 它旨在由DevOps的IT人员使用。因此，Java开发人员需要学习一些新的概念，并开放学习解决问题的新方法。手动安装高度可用的Kubernetes集群有一个显著操作的开销。

3. Kubernetes仍然是一个相对较新的平台（2岁），它仍然积极发展和成长。因此，每个版本都添加了很多新功能，可能很难跟上。 好消息是，这已经被考虑到，API将是可扩展和向后兼容的。

------

最好的两个世界

正如你所看到的，这两个平台在某些领域有优势，在其他领域有待改进。 Spring Cloud是一个快速开始的开发者友好平台，而Kubernetes是DevOps友好的，具有更陡峭的学习曲线，但涵盖了更广泛的微服务关注点。

这两个框架涉及不同范围的MSA关注，他们以一种根本不同的方式去实现。 Spring Cloud方法试图解决JVM中的每个MSA挑战，而Kubernetes方法试图通过在平台层面解决为开发人员解决问题。 Spring Cloud在JVM内部非常强大，Kubernetes在管理这些JVM方面功能强大。结合他们，并从两个项目的最好的部分受益。

有了这样的组合，Spring提供了应用程序打包，而Docker和Kubernetes提供了部署和调度。 Spring通过Hystrix线程池提供应用程序防火墙，Kubernetes通过资源，进程和命名空间隔离提供防火墙。Spring为每个微服务提供健康端点，Kubernetes执行健康检查和流量路由到健康的服务。 Spring负责外部化和更新配置，Kubernetes将配置分发到每个微服务。