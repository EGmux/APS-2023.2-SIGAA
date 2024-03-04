_his post is part of_ _[The Software Architecture Chronicles](https://herbertograca.com/2017/07/03/the-software-architecture-chronicles/)__, a_ _[series of posts about Software Architecture](https://herbertograca.com/category/development/series/software-architecture/)__. In them, I write about what I’ve learned on Software Architecture, how I think of it, and how I use that knowledge. The contents of this post might make more sense if you read the previous posts in this series._

We learn how to code and we build some cool applications, and then we learn about architecture and how to make the application maintainable for several years…

However when we need to explain to someone else (new developer, product owner, investor, …) how the application works, we need something more… we need documentation.

But what documentation options do we have that can express the whole application building blocks and how it works?!

In this post I’m going to write about:

-   UML
-   4+1 Architectural view model
-   Architecture Decision Records
-   The C4 Model
-   Dependency diagrams
-   Application Map

## UML

There are several diagrams we can create using UML, and we can segregate them into two categories:

-   **Behavioural UML Diagram**
    -   [Activity Diagram](https://www.visual-paradigm.com/guide/uml-unified-modeling-language/what-is-activity-diagram/)
    -   [Use Case Diagram](https://www.visual-paradigm.com/guide/uml-unified-modeling-language/what-is-use-case-diagram/)
    -   [Interaction Overview Diagram](https://www.visual-paradigm.com/guide/uml-unified-modeling-language/what-is-interaction-overview-diagram/)
    -   [Timing Diagram](https://www.visual-paradigm.com/guide/uml-unified-modeling-language/what-is-timing-diagram/)
    -   [State Machine Diagram](https://www.visual-paradigm.com/guide/uml-unified-modeling-language/what-is-state-machine-diagram/)
    -   [Communication Diagram](https://www.visual-paradigm.com/guide/uml-unified-modeling-language/what-is-communication-diagram/)
    -   [Sequence Diagram](https://www.visual-paradigm.com/guide/uml-unified-modeling-language/what-is-sequence-diagram/)
-   **Structural UML Diagram**
    -   [Class Diagram](https://www.visual-paradigm.com/guide/uml-unified-modeling-language/uml-class-diagram-tutorial/)
    -   [Object Diagram](https://www.visual-paradigm.com/guide/uml-unified-modeling-language/what-is-object-diagram/)
    -   [Component Diagram](https://www.visual-paradigm.com/guide/uml-unified-modeling-language/what-is-component-diagram/)
    -   [Composite Structure Diagram](https://www.visual-paradigm.com/guide/uml-unified-modeling-language/what-is-composite-structure-diagram/)
    -   [Deployment Diagram](https://www.visual-paradigm.com/guide/uml-unified-modeling-language/what-is-deployment-diagram/)
    -   [Package Diagram](https://www.visual-paradigm.com/guide/uml-unified-modeling-language/what-is-package-diagram/)
    -   [Profile Diagram](https://www.visual-paradigm.com/guide/uml-unified-modeling-language/what-is-profile-diagram/)

I will not go into the details of each type of diagram because it would be too much to cover in this post, and there are plenty of resources out there documenting these diagram types. To know more about each of these types, you can check each of the links above that lead to some [Visual Paradigm guides](https://www.visual-paradigm.com/guide/), or check out this [blog post](https://tallyfy.com/uml-diagram/).

All in all, UML is cool, it’s very interesting, we can be very expressive with it, we can easily sketch some ideas with it and discuss it with colleagues.

However, to document a whole application architecture using UML we need to use several types of diagrams. Furthermore, if we try to use one single class diagram to express the whole application we are asking for trouble.

An example of good usage of an UML class diagram is to document design patterns:

![](https://herbertograca.files.wordpress.com/2019/07/strategy_1.png)

Source: [https://java-design-patterns.com/patterns/strategy](https://java-design-patterns.com/patterns/strategy/)

This is fine, this is actually great! It can express classes, interfaces, usability and inheritance relations, data and behaviours. It is also concise and readable, and because it’s small, it’s also fast to create.

However, the example below is not so useful… It’s very big, so it gets confusing and difficult to follow. Furthermore, it will take so much time to create it, that when we are finished, it will probably be outdated already because someone will have made changes to the code in the mean time.

![](https://herbertograca.files.wordpress.com/2019/07/class-diagram2.png)

Source: [https://knowhow.visual-paradigm.com](https://knowhow.visual-paradigm.com/)

So, we can and should use UML, but for the situations it should be used: to describe patterns, small portions of an application in detail, or high granularity views of the application with low detail (not using class diagrams).

But then the question remains, how do we document a full application?!

## 4+1 Architectural view model

The 4+1 Architectural view model was created by Philippe Kruchten and published, back in 1995, in his paper titled “[Architectural Blueprints—The “4+1” View Model of Software Architecture](https://www.win.tue.nl/~wstomv/edu/2ip30/references/Kruchten4+1.pdf)“.

This way of visualising a software application architecture is based on 5 views/perspectives of the application, telling us what diagrams can be used to document each of those views.

![](https://herbertograca.files.wordpress.com/2019/07/41_view_model-2.png)

1.  **_Logical/Structural view_**  
    Concerns itself with the functionality that is provided by the system and how the code is designed to provide such functionality;
2.  **_Implementation/Developer view_**  
    Portrays the static organisation of the code, the components, modules and packages;
3.  **_Process/Behaviour view_**  
    Focuses on the runtime behaviour of the system, how system processes communicate, concurrency, synchronisation, performance and so on;
4.  **_Deployment/Physical view_**  
    Illustrates the physical organisation of the application, its about “what code runs in what hardware”;
5.  **_Use Case/Scenario view_**  
    The architecture as a whole is explained with the help of a few use cases, which are simply sequences of interactions. Part of the architecture evolves from such use cases.

It’s important to note that the _4+1 architectural view-model_ does not mandate that we use _all_ mentioned diagrams, and not even all the views. We always need to understand the tools, and use no more and no less than what we need.

## Architecture Decision Records

The **Architecture Decision Records (ADR)** are actually not really about documenting the current, or future, state of an application architecture, but instead the reasons that led to it. They are specially important because they intend to tell others, and our future selves, **why the architecture is what it is**.

**An ADR is a log entry about the architecture decisions** that have been made and that lead to the state of the architecture as it is now or as it is intended to be in the future. They contain the **why** behind the the diagrams that describe the architecture.

To start, there are a few artefacts that we need to know:

-   **Architecturally-Significant Requirement (ASR)**: a requirement that has a measurable effect on a software system’s architecture;
-   **Architecture Decision (AD)**: a software design choice that addresses a significant requirement;
-   **Architecture Decision Record (ADR)**: a document that captures an important architectural decision made along with its context and consequences;
-   **Architecture Decision Log (ADL)**: the collection of all ADRs created and maintained for a particular project (or organisation);
-   **Architecture Knowledge Management (AKM)**: the higher sphere of all previous concepts.

I have seen a few templates for creating ADRs, and I saw nice things in several of them, so I created my own template. You can, and maybe should, create yours as well, one that makes sense to you and your team.

For me, the most important thing for a template is that it’s simple, and it has some documentation in it to help fill it in and even to help make pragmatical and unbiased decisions.

The best way to use an ADR is not simply as a document written after having a discussion and making a decision. The best is to use it as the starting point for the discussion, as an RFC (Request For Comments), which is an idea/proposal that we submit to the other members of the team/department requesting their input/opinion/approval. The intention is really to use it to start a discussion, brainstorm, make the best decision possible, and use the proposal document itself as the decision log entry (ADR). The fact that the ADR is written before hand, doesn’t mean that it is immutable, it must be updated/improved as the discussion unfolds. I find it specially important that all the options under consideration be written down with their pros and cons, as to spark discussion and a clear decision.

So, this is the template I came up with:

![](https://herbertograca.files.wordpress.com/2019/07/adr_template.png)

[https://docs.google.com/document/d/1Xe5erulKsdaha3uwU6tNWsAWxjQK-Rcrr\_iJeOCXuSQ/edit?usp=sharing](https://docs.google.com/document/d/1Xe5erulKsdaha3uwU6tNWsAWxjQK-Rcrr_iJeOCXuSQ/edit?usp=sharing)

Feel free to [copy it from google docs](https://docs.google.com/document/d/1Xe5erulKsdaha3uwU6tNWsAWxjQK-Rcrr_iJeOCXuSQ/edit?usp=sharing).

If you want to explore this subject more, I recommend heading to the [Joel Parker Henderson github repository about ADRs](https://github.com/joelparkerhenderson/architecture_decision_record).

## The C4 Model

The C4 model was introduced by Simon Brown, and it’s the best idea about software architecture documentation that I’ve come across so far. I’ll quickly explain the main idea in my own words, although using his own example diagrams.

The idea is to use 4 different granularity (or zoom) levels for documenting software architecture:

-   Level 1: System Context diagram
-   Level 2: Container diagram
-   Level 3: Component diagram
-   Level 4: Code diagram

### Level 1: System Context diagram

This is the highest granularity diagram. It has little detail but **its main goal is to describe the context in which the application is**. So, it will be composed by one single box for the whole application, and it will be surrounded by other boxes that refer to **the external systems and users the application interacts** with.

![](https://herbertograca.files.wordpress.com/2019/07/bigbankplc-systemcontext.png)

### Level 2: Container diagram

Now, we zoom into our application, the blue square in the diagram above which maps to the dashed square in the diagram below.

At this level of granularity, **we will see the containers of the application**, where a container is any independent technical piece of the application, for example a mobile app, an API or a database. It also documents the **major technologies** used and **how the containers communicate**.

![](https://herbertograca.files.wordpress.com/2019/07/bigbankplc-containers.png)

### Level 3: Component diagram

The component diagram shows us the components inside one container. In this context, each component is a module of the application, not restricted to domain wise modules (ie. billing, users, …) but also including purely functional modules (ie. email, sms, …). So this diagram **shows us the main cog wheels of a container and the relations between those cog wheels**.

![](https://herbertograca.files.wordpress.com/2019/07/bigbankplc-components.png)

### Level 4: Code

The most fine grained diagram, **aimed at describing the code structure inside a component**. For this level, we use an UML diagram with class level artefacts.

![](https://herbertograca.files.wordpress.com/2019/07/bigbankplc-classes.png)

To know more about it, you can read Simon Brown’s own explanations about it [here](https://c4model.com/) and [here](https://www.infoq.com/articles/C4-architecture-model/), or even watch him talk about it [here](https://www.youtube.com/watch?v=uvnZ46OvJLA).

## What is still missing?!

I think the C4 Model is a great way to document applications architecture, it is great to understand the architecture of the application to a certain level, but I still find it insufficient, although it took me some time to put my finger in what is missing.

There are three limitations I see in these diagrams:

1.  Save some exceptions, like Simon Brown’s [structurizr](https://structurizr.com/), they need to be manually made, not automated nor directly extracted from the code, which means they might not reflect the actual code but, instead, our current understanding of it;
2.  They don’t quite help us see what is wrong in our application codebase, in regards to promiscuous code relations and poor structure, which impacts modularity and encapsulation, essential to any engineering product;
3.  They don’t help us understand our codebase as a whole, what the application cog wheels can do and how they interact with each other.

I have found two categories of diagrams that can help us with that.

### Dependency diagrams

The dependency diagrams are useful to tell us about the dependencies that exist in the different types of code in our codebase.

Crucially important here is that these diagrams be automatically generated directly from the code, otherwise the diagram will reflect only what we think the code looks like, and if that was accurate we wouldn’t really have much need this type documentation.

Furthermore, maybe more important than the diagrams themselves is the ability to use these dependencies analysis to stop a build in the case of a break in our predefined dependency rules. So, the tool used to generate these diagrams should also be usable as a testing tool and included in our CI pipeline, just like unit tests are, preventing unwanted dependencies to reach production, which maintains and enforces modularity, which in turn helps reach high changeability rate and therefore high velocity of feature development.

Within this category of diagram, I find it useful to have 3 different types of diagram, to assert about different dependency types.

In the case of the examples I have below, they were all generated by [deptrac](https://github.com/sensiolabs-de/deptrac) for my [pet project (explicit-architecture-php)](https://github.com/hgraca/explicit-architecture-php), which I use for experimenting. You can find the configuration used to generate them in the repository root.

Do note, however, that I added the colours myself as to make it easier to read in this blog post. The colours represent different layers in the application, in accordance with the layers I wrote about in previous blog posts:

-   ![](https://herbertograca.files.wordpress.com/2019/07/infographic-16_9.png?w=960)
    
-   ![](https://herbertograca.files.wordpress.com/2019/06/more_than_concentric_layers.png?w=994)
    

#### Layer dependency diagram

The intention of this diagram is to visualize and make sure that the code in each layer can only depend on the layers inner or below it.

So, in the diagram below we can see, for example, that the Infrastructure layer, being one of the top outer layers, can depend on any other layer. On the other hand, the Domain layer, being the top center layer, can only depend on the layers below, namely the SharedKernel-Domain (which is part of the Domain as well) and the PhpExtension (whose code is used as if it was part of the language itself).

![](https://herbertograca.files.wordpress.com/2019/07/deptrac_layers-2.png)

_Layer dependencies diagram generated by_ [deptrac](https://github.com/sensiolabs-de/deptrac) _for_ [https://github.com/hgraca/explicit-architecture-php](https://github.com/hgraca/explicit-architecture-php)

#### Class dependency diagram

The Layer dependency diagram analyses the dependencies between layers, but within a layer there are still dependencies that must not occur.

The Class dependency diagram is useful to analyse the dependencies between the different types of class we have in our codebase, specially if they are in the same layer.

For example, if we want our events to be serializable, so that we can put them in a queue, we probably don’t want them to contain an entity because it would be problematic to unserialize it and persist it using an ORM. It would also not make sense for an event to depend on a service. With this type of diagram, or more accurately with the tool to test dependencies, we can easily detect such cases and prevent them from reaching production.

![](https://herbertograca.files.wordpress.com/2019/07/deptrac_class-3.png)

_Class dependencies diagram generated by_ [deptrac](https://github.com/sensiolabs-de/deptrac) _for_ [https://github.com/hgraca/explicit-architecture-php](https://github.com/hgraca/explicit-architecture-php)

#### Component dependency diagram

A component is a domain wise module, a module that contains both Application and Domain layers. A component can be, for example, “Billing” containing all its use cases and Domain logic.

Components can be mapped to DDD bounded contexts and/or Microservices, which means they must be completely decoupled, physically and temporally, from other components. If we have a monolithic application with fully decoupled components, it will be fairly easy (code wise) to transform it into a Microservice Architecture.

Furthermore, applying the same decoupling requirements to the other non domain wise modules, we can guarantee that we can easily replace any module.

The Component dependencies diagram is aimed at making sure that the application components and modules are decoupled.

Note, in the diagram below, how the modules in the same layer (nodes with the same colour) are all unaware of each other, at least directly.

Specially important is that the two components (User and Blog, in mid-blue colour) are decoupled. If this application had a Microservices Architecture, these two components would be the microservices.

![](https://herbertograca.files.wordpress.com/2019/07/deptrac_components.png)

Component dependencies diagram generated by [deptrac](https://github.com/sensiolabs-de/deptrac) for [https://github.com/hgraca/explicit-architecture-php](https://github.com/hgraca/explicit-architecture-php)

### Application Map

About a year ago, I realised something else I was also missing in these documentation options: All these diagrams, they tell us what are the building blocks of the application, which blocks interact with each other and how they are related, but **they don’t tell us** **what they do****, nor** **how and when they interact** **with each other.** For that we need to either know the application very well from the user perspective, or the codebase from the developer perspective. The previous diagrams don’t tell us what use cases we have in the application, nor what events are triggered by what use cases, nor what are the consequences of those events. If we show those diagrams to a Product Owner, he will find them mostly useless for his role.

So I came up with an idea for a new documentation diagram, which I call an _**Application Map**_, that can replace the C4 Model Component diagram.

The **_Application Map_** is aimed at being truly a map of the application, defining its “cities” (Components), its “local roads” (use cases), “highways” (events), etc.

The difference between modules and components is that a module is any modular piece of the application, while a component is a **domain wise** module of the application. So, while an ORM is a module of the application, it is not a component because it only deals with technical concerns. On the other hand, a “Billing” module is a component because it deals with domain concerns.

An Application Map starts by defining the components of the application, the domain wise modules, like “Billing”, “User”, “Company”, “Orders”, “Products”, and so on. In the case of a simple blog application, we could have two components, the “User” and the “Blog” components:

![](https://herbertograca.files.wordpress.com/2019/06/app_map_01.png)

In each of those components, we define what are the commands that can be issued to them. The “User” component can create and delete users, while the “Blog” component can create and delete posts, and create comments to a post.

![](https://herbertograca.files.wordpress.com/2019/06/app_map_02-1.png)

Next, in each component, we list any relevant services. These services are relevant because, for example, they trigger an event or are used directly by another component. This is important because the application map should make visible the connections between components as well as what they mean and any followup side effects, and for this we need to expose the services that wiring to other components and their names (which should express what they do).

![](https://herbertograca.files.wordpress.com/2019/06/app_map_03-2.png)

Following the services, we list all the event listeners in each component, even if they are not actually used, which is handy because then we can detect it and either fix whatever needs to be fixed or remove the unused code.

By listener I mean a class whose public methods are all independently triggered by only one type of event, they focus on the event.

![](https://herbertograca.files.wordpress.com/2019/07/listeners.png)

We will also list the event subscribers in each component, for exactly the same reasons as we list the listeners.

An event subscriber is similar to an event listener, except that its public methods are triggered by different events, they focus on a composite task, an example of a subscriber can be a class listening to different framework events in order to control when to start, commit or rollback the Request transaction.

![](https://herbertograca.files.wordpress.com/2019/07/subscribers.png?w=1024)

At this point, we have all the components and their capabilities in the map. This is very valuable because it tells us, or any non technical person, what each component can do.

However, it still doesn’t tell us how all these capabilities relate to each other, for example “what happens as a consequence of a user creating a blog post?”.

In order to achieve that, the first step is to list what happens in a component when a specific capability is triggered.

In the image below, we can see that deleting a post (“DeletePost”) will trigger the deletePost() method in the PostService, which is also triggered by a listener listening to the event that notifies that a user has been deleted. This tells us that our application deletes posts as a result of either a direct command from a user or when a post author has been deleted.

In the User component, we can see that when a post is created, its author is automatically subscribed to that post subjects (tags).

![](https://herbertograca.files.wordpress.com/2019/07/direct_calls.png?w=1024)

Now we have the information about the flow within a component, but we are still lacking the information about cross component flow, so lets add the events being triggered and listened to:

![](https://herbertograca.files.wordpress.com/2019/07/events-1.png?w=1024)

We can see, for example, that:

-   Deleting a user will trigger an event that will delete the users’ posts;
-   Creating a post will trigger the event that will lead to both subscribing the author to the posts’ subjects and increasing the authors rating;
-   Deleting a post, from whatever use case, triggers an event that will result in decreasing that authors’ rating.

With all this information in our map, we can navigate it. Any technical or non-technical person can clearly visualise what happens when any of the use cases of the application is triggered. This can help us clarify our code, and our idea of the application behaviour.

But, when used in a big application, this diagram will still have problems common to the previously mentioned diagrams:

1.  It’s an artefact that will take a lot of effort and time to get done and also to simply keep it up to date;
2.  We will still end up having a big diagram with a lot of lines on it, which is not the most readable.

To solve the first problem, we need to be able to generate the diagram from the code, on-demand. This will make it effortless to create such a diagram, remove the need for maintaining it, and make it virtually immediate to create it.

To solve the second problem, we need to be able to selectively generate only part of the diagram. For example by providing the name of the use case that we want to analyse, which would result in only generating the sections of the diagram that somehow are related to the given use case.

So we need a tool… which does not exist… yet!

Or does it?! 😀

Some time ago I started creating it, and I got to the point where only the component internal flow is missing, but it lists all the commands, services, listeners, subscribers and events. It is still very alpha because of the missing information, but also because it is not flexible when it comes to the code base it needs to analyse but, from the codebase of the company where I currently work at, it can generate something like this:

![](https://herbertograca.files.wordpress.com/2019/07/appmap00.png)

Example of an (incomplete) application map, as generated by [https://gitlab.com/hgraca/app-mapper](https://gitlab.com/hgraca/app-mapper)

If you are curious about the project, you can check it out [here](https://gitlab.com/hgraca/app-mapper), however be advised that it is still very alpha, its just a proof of concept and I haven’t worked on it for a few months already. If you feel it’s a worthy project and you have free time to contribute, let me know and I will try get you up to speed and create tasks that you can pick up to bring it to the next level.