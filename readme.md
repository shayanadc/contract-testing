### Intro
Back in the day when I worked for an internet service provider, I picked up a valuable lesson about teaming up with partners to develop a service.

Basically, as a developer, you often have to exchange data with other parties to make your service or application work. Sometimes these parties are external to your company, and sometimes they're internal.

For instance, if you're sending order information to a delivery service, you've got to collaborate with an external party. And sometimes, you need product lists from another department or service within your own company.

In both cases, you've got to build your project based on their setup, but you don't have control over how they've implemented things. In other words, if they change their implementation, you've got to adapt, or your code will break.

### Main Idea
Sure, it's a reality that we don't have control over what happens outside our system, but could that be risky for our business?
Absolutely, it's inherent in working with external parties to be dependent. But how can we organize it in a more effective way.
I am going to address the problem of collaborating with third apis to develop a stable product in a practical way.

### My Experience
Let me start with the business idea from our company, In fact, we were working on building an integrated system to connect the end users to the internet. Actually, on that time based on your region, phone numbers and many other factors people needed to access to different system to get subscribtion and access to the internet which was really painful.

As a backend developer, our mission was building an interface to connect to those different third parties APIs with HTTP or message

We started to gather the text-based documentation of all of our third parties and implement a wrapper to send the proper message to them and get the relevant message.

During the development process we realized that some endpoints does not align with their documentation and we tried to cover all those with unit and functional tests and eventually we  released after 3 months of hard working.

It did not take more 3 days to face the first issue and then the next time before having enough time to fix the previous one.

After working like a fireman to turn off the out of control fire which you have to turn of something and before complete the process, do the same somewhere else we categorize the cause of the issues to find a proper generic solution for that.

### Observe The Issues

There were two reasons:

1. The behaviour of Third API had changed and we were expecting the changed response.

2. Our expectation to send/receive the request/response was not aligned with the current state of the third APIs. 
First, we came up with the idea of keeping the partners code more align with their documentation and so we asked them to use some tools like openAPI which helps you to reflect the code implementation in documentation

### Issue 1 Example
```
   { Id: "1" },
   { Id: 1 }
```
### Issue 2 Example
``` 
{ Id: "1", Amount: 2000 },
{ Id: 1, Amount: 2000, TransactionId: 56901 }
```

### First Attempt 

It always not easy to affect other team whether they are in your organization or not. But we successed to do that for some of the partners at least.

We started to write some behavioural test to cover the third api’s behaviour in our development process with Gherkin.

![image](https://github.com/shayanadc/contract-testing/assets/6294224/89319968-6e26-4696-a155-ae599ba9acbe)

```
  Scenario: Get The List Of Users
    When send the GET request to the user endpoint
    Then the response should be:
	```[
		{ Id: "1", Amount: 2000 },
		{ Id: "1", Amount: 2000 },
	]```
```

Added more test while we had unit test was challenging for our team to handle, organize but we did the hard job and implemented them.

While we thought everything should be good now, we released the new version but again we got the same error with the same reason which was really disappointing.

Again we categorized the cause to improve the solution…

### Second Attempt 

![image](https://github.com/shayanadc/contract-testing/assets/6294224/cc746446-e6ae-4899-8453-50e3191d7293)

Yes, they had testable documentation which really reflect their code implementation but they were free to change their implementation without informing us and we only aware of them when we failed.

Yes, we had another test layer (BDD) but it was based on our mock and our mock was based on the version of their APIs which was easily changable.

In fact, we realized that what we need is more than code level and we need to have a more efficient collaboration with our third APIs

In fact we needed to be informed by them before applying any changes to apply it in our internal code if it needs and keep our code more sync.

![image](https://github.com/shayanadc/contract-testing/assets/6294224/359075aa-04a0-409a-a9cf-99b15156c921)

We introduced a specialized position within our company tasked with the responsibility of reaching out to our partners. Their primary objective is to gather updates regarding any modifications made or planned for their APIs

### Why We Need
But we always heard some like excuses like : 

- Oh I did not know that you were using this…
- Oh I only correct the typo in json and it should not affect anyone… 

Distributed system does mean decoupled system.

In fact, we lessont that the it’s too late to detect the changes on your partner side, the part of their service which is related to you since your consuming it, because your customer might be get that error sooner.

![image](https://github.com/shayanadc/contract-testing/assets/6294224/43b50a60-ac56-4d98-905c-0af5fb504479)

## Unit Test vs Contract Test

Consider the scenario of testing a fire alarm in your house. There are different approaches you could take to ensure its functionality.

Firstly, you could wait until there's an actual fire in your home to see if the alarm reacts appropriately. However, by this point, it's too late to prevent potential damage or harm.

Alternatively, you could simply press the alarm's test button to check if it produces sound. While this method confirms that the alarm is capable of making noise, it doesn't necessarily replicate real-life conditions accurately.

The most effective approach, akin to a smoke test, involves simulating a realistic scenario by introducing smoke to the alarm's sensors. This method provides a thorough evaluation of the alarm's performance in a situation closely resembling an actual fire, ensuring its reliability when it truly matters.


### Coding

Imagine we have a microservice architecture with two services which communicate with each other with HTTP. Clinet needs to get the list of users from server and want to be confident to realease.

### Checkout Step-One

First, we need to have unit test in client service with mocking the provider response. Client service send the http request to server :

```
	uri := fmt.Sprintf("http://%s/users", host)

	resp, err := http.Get(uri)
```
and for the response we mock:
```
    []User{
		{
			ID:   "1",
			Name: "John",
		},
		{
			ID:   "2",
			Name: "Alice",
		},
	}
```

Now, we need to add interaction with pactflow API to create and publish the pact between the client and server.

```
	pact.
		AddInteraction().
		Given("User Lists").
		UponReceiving("User Collection is requested").
		WithRequest(dsl.Request{
			Method: "GET",
			Path:   dsl.Term("/users", "/users*"),
		}).
		WillRespondWith(dsl.Response{
			Status: 200,
			Body: dsl.Like([]User{
				{
					ID:   "1",
					Name: "John",
				},
				{
					ID:   "2",
					Name: "Alice",
				},
			}),
		})
```

for specifying, the path in the interaction, we can use regex to cover all possible path variations of user endpoint like users/active or users?name=xyz to simplify the mocking and avoiding the repetitive job.
```
	Path:   dsl.Term("/users", "/users*"),
```

From the server side, you only need to run your service against of the contract file which it creates after running the mock test of the client.

```
	pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL: "http://127.0.0.1:8080",
		PactURLs:        []string{"../client/pacts/example-client-example-server.json"},
	})
```

but as you see, the server test is running against of a local file but in microservice architecutre two different system is usually deployed in different environemnts. So we need to publish the pact file to pactflow server and run the verification against of the published pact url.

### Checkout Step-Two


we only need to call the pact broker to upload the contract after writing by the client side

```
	publisher.Publish(types.PublishRequest{})
```
and from the server side, instead of checking the interaction from the local file, check it with the broker:

```
	pact.VerifyProvider(t, types.VerifyRequest{})
```
