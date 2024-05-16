As a developer, when developing a feature, you often need to collaborate and communicate with other systems/services to acquire the necessary information that isn't available on your end.

This means your application needs to communicate with another application or service to request data in the manner dictated by the server for all its consumers, including yourself.

This communication occurs via the HTTP protocol, and the request structure and type determine the response you'll receive.

However, on the flip side, it means you must develop and design parts of your application based on other applications over which you have no control. Consequently, if they change their implementation, you'll need to adapt or perish...

What we typically do is attempt to adapt to changes because we've learned that they're inevitable. In my opinion, all software development methodologies aim to support the idea of reducing the cost of changes.

Microservice architecture is more common than ever today. The primary goal of microservice architecture is to segregate concerns, meaning each service maintains its API. Changes to these services may not be under your control, yet you will be affected by them, even if the service providers are unaware of who is consuming their data.

Let's consider some possible ways to handle this situation, step by step, and weigh their pros and cons.

Imagine you need to work with another service as a client, such as a mobile application or a backend server, which requires connecting to another server to fetch data.

You typically begin by requesting API documentation from the server and then develop your application based on that documentation.

If you employ a Test-Driven Development (TDD) approach, you'll start by mocking the third-party API's behavior in your tests. Once you're confident in handling the communication between your application and the third-party API, you'll implement your logic and deploy the application to production.

Now, let's briefly consider the server-side as well. They have their own unit tests that run against their API to ensure everything is functioning correctly.

However, after development, when everything appears to be working properly, you may suddenly find that some parts of the application don't function.

There are usually two main reasons for failures in communicating with other services:

- You expect something that no longer aligns with the current state.
- Your partner expects to receive something, but you haven't sent the proper message.

We typically handle this situation with the mindset that change is inevitable, and we aim to minimize its impact.

Firstly, we may request that the data provider use auto-generated tools for their API specification to ensure compatibility between the API docs and the real implementation. However, this doesn't entirely solve the problem since the provider can still break communication by changing the specification unexpectedly.

Alternatively, we can employ Behavior-Driven Development (BDD) to be more flexible against changes. With Gherkin, we can create scenario-based test cases, enabling quicker debugging when there are issues with the third-party API. However, this approach relies on provider mocks, and if they change, you'll need to update your mocks. Additionally, if they don't notify you of changes, it's your customers who will encounter issues.

The eventual decision we often make is End-to-End (E2E) testing, where you can test server behavior in the real world, like with Postman collections or Cypress. However, E2E tests can be expensive and brittle since you need to run tests in the real or simulated server environment, making test implementation difficult to establish and fragile. E2E tests are brittle because they're not conducted in isolated environments.

So, what's the better solution? Let's recap what we want:

- We want to be informed of server changes that could disrupt our communication.
- We want to test communication in isolation.

Contract testing is a technique that ensures confidence in the established communication between two parties by testing the agreement or API contract between them before each deployment.

Both the consumer and provider need to generate the contract, which should be tested by the other party before deploying any changes to production.

There are various implementations of contract testing, but here I'll explain Pactflow, an open-source tool supporting most popular programming languages, which is consumer-based.

The consumer uses the Pactflow SDK to generate the contract, a JSON file, by running mock tests. The server-side then tests this contract file against its API.

Pushing changes is only possible when both sides have passed their contract tests.

In microservice architecture, since the client and server are typically not in the same repository or machine, we need to keep the contract in the cloud accessible to both parties, a feature provided by Pact. Pact broker also handles finding the corresponding contract of each provider that needs to be committed to run and also the Work In Progress (WIP) branch contract that should be run against the server API but should stop the pipeline if necessary.
