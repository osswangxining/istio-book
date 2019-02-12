Introduction
In a document published previously this year (cited in the "Related Research" section of this
document), IDC examined the service mesh, explaining its nature, purpose, use cases, and business
value. The document was intended to help enterprises determine whether a service mesh, such as the
open source project Istio, was something they ought to consider or adopt as they embraced containers
and microservices.
In this document, IDC examines the Istio ecosystem, focusing on how major vendors are positioning
themselves and their portfolios in relation to Istio, an open source service mesh. The document is not
intended to be a comprehensive listing of all vendors within the Istio ecosystem, which is evolving
quickly. Other vendors not profiled in this document will emerge as contributors to Istio and as vendors
of Istio-related offerings.
At this point, let's briefly restate the context in which Istio has emerged. It all started with containers
and microservices, the latter an emerging architecture that structures applications as loosely coupled
services that simplify and speed development, deployment, and updates. Microservices are predicated
on Linux-based containers, which are orchestrated with tools such as a Kubernetes. Each function of
the application is handled as an independent service (a microservice) that can be modified, updated,
or taken down without affecting the rest of the application.
By orchestrating containers, Kubernetes makes it possible for microservices deployments to achieve
resilience and scale, but runtime in production environments introduces connectivity challenges,
particularly as microservices-based applications scale and proliferate and service-to-service
interactions become more complex. Indeed, such applications can encompass hundreds of services,
with each service potentially including thousands of instances that are subject to changes occasioned
by orchestration systems such as Kubernetes.
This is where the need for a service mesh, such as Istio, comes into play, addressing the need for
secure service-to-service communication, authentication, traffic management, observability, and other
facets of cloud-native networking for microservices. Indeed, in containerized microservice runtime
environments, networking can be a daunting challenge, shifting a critical burden from the lower layers
of the network (Layer 2 through Layer4) up to Layer 7, the application layer. Higher-layer network and
security services (often called application delivery services) are essential because IP networking, the
foundation of client/server networking, does not provide adequately for microservice availability,
elasticity, reliability, responsiveness, and security.
The term service mesh is used to describe the network of microservices, facilitating interactions
between the services that constitute applications. As a service mesh, Istio defines itself as an open
means of connecting, securing, controlling, and observing microservices. It controls the flow of traffic
and API calls between services; automatically secures services through managed authentication,
authorization, and encryption; delivers control through application and enforcement of policies, while
ensuring that resources are allocated properly; and provides observability through automated tracing,
monitoring, and logging of services.
Before Istio, in early cloud-native microservices deployments (such as those at Netflix), these
challenges were met through the development of purpose-built, language-specific application libraries.
However, the operational complexity and cost of developing bespoke libraries across languages,
frameworks, and runtimes is prohibitive for most organizations, especially those with heterogenous
applications and polyglot programming languages.
running on virtual machines. It can be deployed on-premises and in public
clouds.
Architecturally, Istio is logically divided into a decoupled data plane and control plane. Its data plane is
composed of intelligent proxies, represented by the open source Envoy project. Each proxy is
deployed as a sidecar, which sits alongside the application service it supports in the same Kubernetes
pod. Envoy sidecar proxies mediate and control all interservice network communication and work in
conjunction with Mixer, a general-purpose policy and telemetry hub. Capabilities inherent to Envoy
proxies include service discovery, load balancing, TLS termination, HTTP/2 and gRPC proxying, circuit
breakers, health checks, staged rollouts, fault injection, and sophisticated metrics. The flexibility of
sidecar proxies allows operators to incrementally add Istio capabilities to a deployment without having
to re-architect or rewrite codes.
The control plane manages and configures the sidecar proxies to route microservices traffic. The
control plane also configures Mixer components to enforce policies and gather telemetry for insights
into the overall behavior of the service mesh. A platform-independent component, Mixer, enforces
access control and usage policies across the service mesh in addition to collecting telemetry from the
Envoy sidecar proxies.
Also part of the control plane, Pilot provides service discovery for the Envoy sidecars and trafficmanagement
capabilities for intelligent routing, such as A/B tests and canary deployments. It also
provides resiliency through time-outs, retries, and circuit breakers. Pilot controls traffic behavior by
converting high-level routing rules (or policy) into Envoy-specific configurations that are propagated to
the Envoy sidecars at runtime.
Meanwhile, Citadel provides service-to-service and end-user authentication with built-in identify and
credential management. It allows for policies to be based on service identity rather than on network
controls.
Although not depicted in the Istio architectural diagram (see Figure 2), Galley has been developed to
validate user-authored API configuration on behalf of other control-plane components. The intent is for
Galley to become Istio's top-level configuration ingestion, processing, and distribution component,
insulating other Istio components from the details of obtaining user configurations from platforms such
as Kubernetes.
Strictly speaking, Istio and Envoy together constitute the service mesh, with Istio providing the control
plane and Envoy providing the data plane. By decoupling the control plane from the data plane —
similar to the architectural principles of software-defined networking (SDN) — Istio allows for
substitution of the Envoy proxy-based data plane with alternatives, such as NGINX or Linkerd or
vendor-supplied data planes.
Istio is intended not only to reduce the complexity involved in running microservices at scale but also to
abstract developers from the underlying infrastructure. Platform operators will manage the Istio control
plane, but developers can be abstracted completely from Istio, although some capabilities and features
could be made available to them as desired.
Industry Dynamics
Google Cloud, IBM, and Lyft (source of the Envoy sidecar proxy) originally developed and launched
Istio in 2017, but others have joined the effort since then and contributed to it, including Pivotal, Cisco,
Red Hat, and VMware.
Istio only recently had its 1.0 release, but with development continuing at a brisk pace, Istio is
expected to gain greater prominence and increased adoption. Early users of Istio include webtechs
and cloud services, and enterprise involved in financial services and healthcare. Many telcos are also
keenly interested in Istio, seeing it as a potential linchpin of their telco-cloud and network-as-a-service
initiatives both in their datacenters and at the intelligent edge (uCPE, 5G, etc.).
Among other open source projects that complement Istio is Knative, which was started by engineers
from Google, Pivotal, and other industry players. Knative is a collection of components that extend
Kubernetes and Istio, designed to simplify how functions-as-a-service (FaaS or "serverless") run in any
cloud. Another open source project that complements Istio is Prometheus, which provides monitoring
and alerting; Prometheus can collect and leverage Istio metrics.
Vendor Examples
Vendors active in the Istio ecosystem include Aspen Mesh (F5 Networks), Avi Networks, Cisco,
Google Cloud, IBM, NGINX, Pivotal, and Red Hat. Those vendors' Istio strategies and positioning are
briefly outlined in this document, although others are part of the ecosystem, and additional vendors are
likely to enter the space as the technology matures and gains commercial traction. These are early
days for Istio, and the situation is likely to remain dynamic and fluid. IDC will closely monitor the
developments.
Aspen Mesh (F5 Networks)
Part of F5 Networks, Aspen Mesh believes cloud-native environments with more than 20 services
reach a point of complexity at which services meshes, such as Istio, become increasingly necessary. It
is seeing many customers advancing through what Aspen Mesh calls a "container journey," first
showing an interest in containers and microservices, and then deploying them tentatively before
adopting them more extensively.
Aspen Mesh believes that customers will be looking for advanced Istio capabilities in the control plane
as well as in rich observability. It is in these areas, as well as data-plane capabilities, where vendors
will be able to add value and differentiate. Aspen Mesh also believes that there may be need for
hardware-accelerated components to provide offload for certain processing-intensive security and
network functions.
As for observability, Aspen Mesh believes this will be critical challenge for customers seeking to
quickly troubleshoot and remediate application layer performance insights through actionable, realtime
insights. Observability should also provide insights that result in prompt and even proactive
security postures.
Aspen Mesh has already had Istio in field trials, gaining valuable feedback from customers — including
those in financial services, healthcare, telco, and ecommerce/retail — and it has now moved to beta.
Early indications from customers suggest that, aside from the features and functionality discussed
previously, many customers are looking for a simple, intuitive dashboard that can help them not only
deploy but manage and monitor real-time events for actionable insights.
Avi Networks
Avi Networks is a vendor of software-defined application delivery controllers (ADCs) and application
delivery services. The vendor's flagship Avi Vantage software-defined platform features a decoupled
control plane and proxy-based data plane that preceded and anticipated the formal emergence of open
source service meshes such as Istio.
To be sure, in its "Elastic Application Services Fabric," Avi had articulated a concept that bears striking
resemblance to an enterprise-class service mesh. As such, Avi now is taking a keen interest in Istio,
which it sees as providing benefits for both developers and operators. Avi sees developers benefiting
from Istio's support for granular CI/CD, circuit breakers, error injection, and rate limiters. For operators,
Avi sees Istio value accruing from its ability to deliver metrics, application mapping, API tracing, and
request logging. This is all in addition to Istio's ability to elastically and securely scale to support cloudnative
environments.
Avi, like Cisco (which is a notable Avi technology partner and strategic investor), seeks to extend Istio
beyond Kubernetes environments, supporting not only containers but virtual machines, bare metal,
and multiple public clouds. To make that happen, Avi will take its current Avi Engine proxy-based data
plane and deploy it as sidecars that can support Istio in Kubernetes and OpenShift deployments. The
company will also extend the Istio service mesh so that it can run across multiple clusters or clouds,
with the capacity to support any infrastructure or cloud.
Avi will also focus heavily on the security and observability capabilities of Istio, as well as traffic-routing
features and application security, all through support for native Istio APIs.
Cisco
Cisco sees Istio playing an integral in hybrid computing across multiple public and private clouds.
Toward that end, Cisco has worked with the broader Istio community to define and develop a model for
extending a single Istio control plane across multiple Kubernetes clusters. That capability is an alpha
feature in Istio 1.0, and Cisco expects further developments on that front, as the Istio community and
the vendors active within it explore ways for Istio to support multicloud and multiple application
environments, including bare metal and virtual machines (VMs) in addition to containers and public
cloud.
Cisco believes Istio provides valuable answers to questions that have been asked about how
networking will evolve to accommodate cloud-native environments comprising containers and
microservices. Some of these questions were fundamental and related to whether the industry even
had the right model for microservices networking. At first, Cisco came at the problem from the
perspective of the network stack, asking whether a new layer might be required to facilitate easier
networking for microservices. Then, working from the primacy of developers and their applications,
Cisco realized that the network stack wasn't necessarily that right way to frame the challenge. Instead,
it was about how the network can serve applications. Service meshes such as Istio are well aligned
with that outlook.
Cisco has developed an extensive hybrid cloud partnership with Google Cloud. The objective of the
partnership is to provide enterprises with the ability to deploy, manage, and secure applications and
services across on-premises environments and the Google Cloud Platform. A core piece of the offering
is the Cisco Hybrid Cloud Platform for Google Cloud, which represents the on-premises application
environment that complements the off-premises Google Cloud. The Cisco on-premises platform includes the Cisco Container Platform, Cisco HyperFlex (HCI), Cisco Nexus 9K switches with Cisco
ACI. Running between and across the hybrid environments are Cisco Stealthwatch Cloud, Cisco
CloudCenter, the Cisco CSR1000v virtual router, and Istio.
Cisco wants to give customers the choice of deploying the Istio control plane on-premises or in the
public cloud. What's more, although Cisco is partnering with Google Cloud, it has a broader multicloud
view of how Istio can be deployed. It also sees use cases for Istio at the edge. In addition, Cisco will
look at how some of its other products, including AppDynamics, can work in conjunction with Istio to
take actions based on observability, and it believes that policy is an area where it will be able to
differentiate. It also believes that customers across various use cases and verticals will insist on choice
and diversity in the data plane.
Google Cloud
As the progenitor of Kubernetes and an original contributor to Istio, Google Cloud is a prominent player
in the Istio community. Google sees the Istio service mesh as integral to the resiliency, scalability, and
security of Kubernetes application environments. Google notes that Istio's decoupled control plane and
sidecar-proxy data plane make it possible to put networking and security functions in front of the
applications, enabling organizations to think about those functions as application services, thus
bringing cloud-native networking solidly into the as-a-service sphere.
In addition to working close with Cisco in the hybrid-cloud partnership, Google offers Istio on its Google
Cloud Platform (GCP). Google emphasizes that at least a dozen companies are running Istio in
production, including several on GCP. Google is focusing on Istio's ease of use on GCP, and it
recently introduced an alpha release of Managed Istio, which is automatically installed and upgraded
on customers' Kubernetes Engine clusters as part of the Cloud Services Platform. Google's Managed
Istio integrates with Stackdriver and Apigee, and it is positioned as providing the visibility, security, and
control that enterprise customers require to run microservices in hybrid environments.
Working within the Istio community, Google Cloud intends to add features and continually improve
Istio's usability.
IBM
Early on, IBM perceived Istio as a means of scaling networks of microservices. IBM had seen how
Netflix and others had addressed the need for microservices networking by taking a language-specific
approach, but IBM sought to address the challenge in a language-agnostic way that didn't force
changes to application libraries.
IBM says many customers believe that Istio isn't required until they reach a certain scale of
microservices, even though Istio can provide value in contexts in which two or three services are
running. Regardless, IBM believes it becomes difficult to manage a microservices network when
customers reach a threshold of 25 microservices.
IBM supports Istio running on top of IBM Cloud Kubernetes Services (IKS) and the IBM Cloud Private
platform. IBM is allowing customers to incrementally adopt Istio at their pace, selecting modular
features and functionalities that suit their current environments and workloads and extending the
breadth and depth of their Istio deployments as warranted.
IBM envisions Istio as a key component of its hybrid cloud strategy. As such, it is focusing on providing
federated Istio service mesh capabilities across clusters that span the public IBM Cloud Kubernetes Service and on-premises IBM Cloud Private environments. IBM is also working on integrating its IBM
API Connect with Istio.
In addition, IBM continues to contribute to Istio in areas that involve networking improvements,
incremental adoption, multiple environments, performance, and scalability, and it is involved with the
Knative project, which has synergies with Istio in the context of serverless capabilities running on
Kubernetes.
NGINX
In NGINX's view, the inherent complexity of microservices applications has driven discussion of and
interest in service meshes such as Istio. That's a valid driver, of course, but NGINX (which delivers
NGINX Plus as a commercial offering) believes that many organizations would have to incur significant
costs to successfully adopt Istio, which remains a nascent open source project. These costs of
adopting Istio include the acquisition of new skills, the evolving nature of the project, the need for
additional instrumentation, the need to support more containers (in the form of sidecars), additional
points of potential failure, and the need for more extensive troubleshooting.
As such, NGINX contends that the cost of managing Istio might actually outweigh its benefits in
complexity mitigation. Indeed, NGINX suggests that most container environments today are not unduly
complex. NGINX believes the complexity threshold that necessitates a service mesh such as Istio is in
the range of hundreds of services, involving multiple teams working collaboratively to deliver services.
Before organizations reach that threshold, NGINX believes they should give thought to how they can
limit "container sprawl," which is analogues to the "VM sprawl" that emerged during the halcyon days
of virtualization. NGINX is working to help its customers address that problem, effectively deferring the
point at which they deploy a service mesh.
Nonetheless, NGINX plans to support Istio 1.0, and it points out that its Engine Mesh is already
compatible with Istio. NGINX believes the release of Istio 1.0 could represent a pivot point for
increased interest, and it already has noted more Istio queries from its ecommerce customers.
Pivotal
Pivotal has made contributions to Istio, and it is making Istio a core component in other open source
projects and in its own commercial offerings. Pivotal will be targeting the combination of Kubernetes
and Istio at a range of customers in the Forbes Global 2000.
Among Pivotal's contributions to Istio is the Gateway Resource, which enables operators to configure
an edge-gateway router when requirements demand functionality other than that provided by the
sidecar proxy at the data plane. Pivotal also has developed the Mesh Control Protocol (MCP) API for
Pilot. It is designed to make it easier to roll out configuration changes across Istio components while
making it easier to configure Pilot. One of Pivotal's objectives was to make integration between Cloud
Foundry and Istio simpler. Pivotal believes that MCP will become a standard API that Cloud Foundry,
Kubernetes, and Istio teams can all use, with MCP managing multiple configuration profiles for
components across platforms.
Pivotal is comprehensively integrating Istio with Cloud Foundry, continuing to make contributions to
Istio, and should be expected to play a prominent role in the Istio community. Red Hat
Red Hat has contributed and continues to contribute to Istio, and Red Hat is developing a product road
map for the technology.
One example is Red Hat's integration of Istio with OpenShift. In late September, Red Hat released its
first technology preview of the Istio-based Red Hat OpenShift Service Mesh. The preview provides
OpenShift customers with the ability to deploy and consume the Istio platform on OpenShift clusters.
Full support and general availability of Istio on OpenShift is scheduled for next year.
Red Hat sees broad applicability for Istio across a range of use cases and market segments, including
numerous enterprises verticals and telco clouds. Red Hat is not focusing exclusively on the telco cloud
opportunity as it relates to Istio, but it believes telecom operators represent a potentially enthusiastic
source of Istio adoption. The majority of telco services today are predicated on monolithic applications,
many of which run on virtual machines, but Red Hat foresees telcos aggressively refactoring their
applications for containers and microservices. As they do so, telcos will see a need for Istio, not just in
datacenters but increasingly at the 5G edge.
The telco shift to microservices will not be without challenges for the telcos themselves and for the
vendors that have supplied them with network appliances and virtual network functions (VNFs) that are
encompassed within their NFV strategies. The shift will be toward Container Network Function
Infrastructure (CNFI), but today, there is no blueprint governing how, for example, a wireless operator
would assemble and service chain containerized network functions (CNFs). That said, Istio is seen as
part of the solution, allowing management of services as integrated bundles that can be automatically
and elastically modified and scaled as needed.
In enterprise verticals, Red Hat is seeing considerable early interest in Istio from large financial
services customers, but the vendor sees interest building across its enterprise installed base of Istio
users.
Other Vendors
Other vendors are active in the Istio community, and IDC expects to follow them closely. In fact, IDC
anticipates that all the major ADC vendors — F5 Networks, Citrix NetScaler, and A10 Networks, among
others — will offer significant support for service meshes in general and Istio specifically. VMware,
which has made contributions to Istio, is also likely to become a more prominent player. IDC will
monitor and document notable market and technology developments relating to Istio.

ADVICE FOR THE TECHNOLOGY SUPPLIERS AND SERVICES PROVIDERS
Istio has emerged to provide the sort of higher-layer connectivity and networking that will be
indispensable to containers and microservices, with particular emphasis on Kubernetes clusters. IDC
believes that multicluster support will make Istio an inherently multicloud offering, and that vendors — if
not the community — will extend it to non-Kubernetes and non-containerized environments.
Istio will reside at a fluid intersection of technologies, open source projects, and clouds. As such, it will
draw the interest and adoption of a wide constellation of vendors, including major IT infrastructure
stalwarts, established ADC vendors, cloud-native startups, and open source mainstays. Alliances and
partnerships will form, and IDC also believes that Istio will prove as an impetus for M&A activity. As an
open source project with an active and engaged community, Istio will require that vendors carefully consider how and where they intend to achieve value-based differentiation without alienating the Istio
open source community.
IDC recommends that ADC vendors have a strategy for service meshes, and particularly for Istio. The
service mesh represents the evolution of application delivery infrastructure into software-defined
application delivery services that are modular and composable. In some respects, one might even view
the service mesh as the disaggregation of the appliance-centric ADC into discrete network and
security services that better aligned to support the agility, elasticity, and portability of cloud-native
microservices. ADC vendors must adapt and devise strategies that allow them to provide differentiated
value within the context of Istio.
IDC also believes that networking vendors that aren't currently part of the ADC market, such as Cisco
and VMware, will move aggressively into service meshes such as Istio. This could result in M&A
activity involving vendors of lower-layer network infrastructure acquiring Layer 4–7 vendors. IDC
advises networking vendors to think about whether their embrace of software-based value and
multicloud necessitates a move up the stack to Layer 7 networking for microservices.
IDC also advises IaaS public clouds that haven't articulated a position on Istio and service meshes to
do so. Multicluster Istio is likely to become increasingly important in a hybrid IT and multicloud, and
public cloud IaaS providers will need to develop services and partnerships that address enterprise
demand.