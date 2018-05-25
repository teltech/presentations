# JWT

## Presentation

+ [PDF](IntroToTddInIonic.pdf)

## Tutorial

Below is the full tutorial copy pasted for archive reasons, but you can also read it [here](https://dev.to/hitman666/introduction-to-tdd-in-js-and-ionic-framework-1-4080).

![](https://thepracticaldev.s3.amazonaws.com/i/fve6sjwma42pyhgs3vzv.png)

## TL;DR

In this rather long post, I'm going to give you an introduction to Test Driven Development in Ionic. First, I'm going to cover some basic theoretical concepts and then we'll see how to apply this to few examples. First in plain JavaScript and then finally in Ionic.

At the end of this tutorial, you'll have a clear path on how to start practicing TDD in your JavaScript and Ionic applications. Also, at the bottom, you'll see a full 'resources dump' of all the resources that I've gone through in trying to learn about TDD myself.

_The presentation slides, in case someone is interested, can be viewed [here](http://www.nikola-breznjak.com/portfolio/IntroToTddInIonic.pdf)._

## Let's answer some tough questions
How many of you actually test your code? Don't worry; this is a rhetorical question, you don't need to raise your hands.

Well, if we're honest here, in my case (since I'm writing mostly JavaScript lately) up until recently I was practicing a so-called **CLTDD**. Which, of course, stands for **console.log TDD**.

We all know we should do *something* to make this better, but far too often we do it like this gentleman here:

![](https://i.imgur.com/zu1Wx05m.png)

Ok, jokes aside, let me try to emphasize on why testing may actually be useful to you. Just think about the following questions:

+ Have you ever fixed a bug, only to find that it broke something in another part of the system?
+ Have you ever been afraid to touch a complicated piece of code for fear that you might break?
+ Have you ever found a piece of code that you’re pretty sure wasn’t being used anymore and should be deleted, but you left it there *just in case*?

![](https://i.imgur.com/dDh41WOm.jpg)

Well, if the answer to any of these questions is yes, then you'll see value in what TDD can bring to the table if practiced correctly.

## What is TDD?
Since most of us here are developers I bet you've heard about **unit testing**. However, unit testing is not the same thing as TDD. Unit tests are a **type of test**. TDD is a **coding technique**. Meaning that if you write unit tests, you don't actually consequently do TDD.

> TDD is an approach to writing software where you **write tests before you write application code**. The basic steps are:

+ **Red**  - write a test and make sure it fails
+ **Green** - write the easiest possible code to make the test pass
+ **Refactor** - simplify/refactor the application code, making sure that all the tests still pass

![](https://i.imgur.com/dJKgmChm.jpg)

At this point you may be like:

> Wait, now I have to write code to test the code that I still haven't written?!"

Yes, you write more code, but [studies have shown](http://evidencebasedse.com/?q=node/78) objectively that **good test coverage with TDD can reduce bug density by 40% - 80%**.

## Why bother with tests?

So, why would you want to test your code in the first place? Isn't it enough that you have a deadline approaching, and now you should spend your precious time writing a test, instead of the actual application code?

Well, as features and codebases grow, manual QA becomes more expensive, time-consuming, and error-prone. 

Say for example if you remove some function from the code, do you remember all of its potential side-effects? Probably not. But with unit tests, you don't even have to. **If you removed something that is a requirement somewhere else, that unit test will fail, and you'll know you did something wrong**.

So basically, we test our code to verify that it behaves as we expect it to. As a result of this process, you'll find you have better feature documentation for yourself and other developers.

Also, as [James Sinclair](http://jrsinclair.com/articles/2016/gentle-introduction-to-javascript-tdd-intro/) argues, practicing TDD forces one to think, as you have to think first and then write a test. Also, it makes debugging easier and programming more fun.

## 5 Common Misconceptions About TDD & Unit Tests

There are 5 Common Misconceptions About TDD & Unit Tests based on [Eric Elliot](https://medium.com/javascript-scene/5-common-misconceptions-about-tdd-unit-tests-863d5beb3ce9#.f1xankzgr).

+ TDD is too Time Consuming
+ You Can’t Write Tests Until You Know the Design, & You Can’t Know the Design Until You Implement the Code
+ You Have to Write All Tests Before You Start the Code
+ Red, Green, and ALWAYS Refactor?
+ Everything Needs Unit Tests

Also, he holds a rather strong point about mocking in TDD:
> Here’s a tip that will change your life: `Mocking is a code smell.`
	
## Demo time
OK, enough with the theory, now let's see some code!

![](https://i.imgur.com/PxpH624m.jpg)

### Prerequisites
To be able to follow this tutorial you need to have [Node.js](https://nodejs.org/en/) installed. Also, via **npm** you'll need to install globally the following packages:

+ [Karma](https://karma-runner.github.io/1.0/index.html)
+ [Jasmine](http://jasmine.github.io/)
+ [PhantomJS](http://phantomjs.org/)

I picked Karma as an environment for running the tests and Jasmine for the actual test cases because these frameworks seem to me as the most reliable for this task and seem to be in widespread use. However, keep in mind that there are many other options. Few worth mentioning are [Mocha](https://mochajs.org/), [Chai](http://chaijs.com/), [Sinon](http://sinonjs.org/), [Tape](https://github.com/substack/tape), etc.

![](https://i.imgur.com/CseGOCbm.png)

What I would like to add here is that these days (especially in the JavaScript world) you have a vast number of options. **Choosing one option and actually starting is way better than endlessly weighing the options.**

With Jasmine, we'll be using a so-called `Behaviour Driven Development (BDD)` style to write the tests. This is a variation on TDD where tests are written in the form:

+ describe `[thing]`
+ it should `[do something]`

The `[thing]` can be a module, class, or a function. Jasmine includes built-in functions like `describe()` and `it()` to make writing in this style possible. Also, Jasmine offers some other cool stuff like spies, which we won't cover here, but you can learn more about it from the [official documentation](http://jasmine.github.io/2.4/introduction.html).

## The JavaScript demo
In this demo, I'll show you a simple step by step TDD approach to building a simple calculator library. This will be a simple file with just two functions (`add` and `sub`). This will be nothing fancy; it's just to illustrate how this process would go.

### Folder structure and dependencies
Let's start by creating a new folder called `jstdd` and inside it a folder `app`:

`mkdir jstdd && cd jstdd && mkdir app && cd app`

Also, create an `index.js` file inside the `app` folder:

`touch index.js`

Next, execute `npm init` in the `jstdd` directory. This will create a `package.json` file for us, where all the other dependencies (which we'll install shortly) will be saved to. On every question in the `npm init` command you can safely press `ENTER` by leaving the default values.

Next, install all the needed dependencies:

`npm install karma karma-jasmine jasmine-core karma-phantomjs-launcher --save-dev`

For those who aren't too familiar with Node and npm, with the `--save-dev` switch we save these dependencies to our `package.json` file that was created with the aforementioned `npm init` command.

Next, create a new folder called `tests` and a file `index.spec.js` inside it:

`mkdir tests && cd tests && touch index.spec.js`

### Setting up Karma
Basically, we have everything set up now. But, before we actually start writing our tests, we have to configure Karma. So, in the root of our application (folder `jstdd`) we have to execute

`karma init`

The answers to the questions should be:

+  use Jasmine as a testing framework
+  don't use Require.js
+  use PhantomJS instead of Chrome (use `TAB` key on your keyboard to switch between options). This is because we want to run our tests in the console
+  use `app/*.js` and `tests/*.spec.js` when asked for source files and test files. We can use glob patterns, meaning that star (*) matches anything
+  when asked for which files to exclude, just skip by pressing `ENTER`
+  finally, choose `yes` to have Karma watch all the files and run the tests on change

With this process being done, Karma generated the `karma.conf.js` file, which (without the comments) should look like this:

```
module.exports = function(config) {
    config.set({
        basePath: '',
        frameworks: ['jasmine'],

        files: [
            'app/*.js',
            'tests/*.spec.js'
        ],

        exclude: [],
        preprocessors: {},
        reporters: ['spec'],

        port: 9876,
        colors: true,
        logLevel: config.LOG_INFO,

        autoWatch: true,
        browsers: ['PhantomJS'],
        singleRun: false,

        concurrency: Infinity
    });
};
```

### Finally let's write some tests
At this point we have everything set up and we can start writing our tests. We will write our tests in `index.spec.js` file.

To remind you, our goal here is to create a simple calculator library. So, we start by writing a test.

When we're using Jasmine to test our code we group our tests together with what Jasmine calls a `test suite`. We begin our test suite by calling Jasmine's global `describe` function.

So we're going to write (in `index.spec.js` file):

```
describe ("Calculator", function (){

});
```

This function takes two parameters: a string and a function. The string serves as a title and the function is the code that implements our test.


Within this describe block we'll add so-called **specs**. Within our `it` block is where we put our expectations that test our code.

So, for example, the first thing that we're going to test is that we indeed have an `add` function:

```
it('should have an add function', function() {
    expect(add).toBeDefined();
});
```

Don't worry about the syntax; that can be easily learned by going through [Jasmine's documentation](http://jasmine.github.io/2.4/introduction.html). And, besides, the good news is that all of the test tools have more or less similar syntax.

Ok, so we wrote our test, but now what? Well, we run the test in the terminal by running `karma start`.

You should see something like:

![](https://i.imgur.com/1J2vYXP.png)

And, what do we see here? We see that we have a failing test. So, what do we do now? We move to the next step, and we make the test pass in the simplest possible way. So, how are we going to do that? We write an `add` function in the `index.js` file:

`function add() {}`

And now we have a passing test. Great. Can we refactor (3rd step) something? Most probably not at this stage, therefore we move onward.

So what's the next thing we expect from our `add` function? Well, we expect that, for example, if we pass numbers 1 and 2 to it, that it will return number 3. So how do we write a test for this? Well, exactly as we said. So:

```
it ("should return 3 when passed 1, 2", function (){
	expect(3).toEqual(add(1,2));
});
```

Now we have a failing test and we go and fix it. At this point we ask ourselves:

> What's the fastest way to pass this test?

Well, the answer to this questions is to return 3 from our function:

```
function add(){
	return 3;
}
```

And, yet again we have a passing test. 

However, say we want to make another test where we say that we expect 5 when passed in 3 and 2:

```
it ("should return 5 when passed 3, 2", function (){
    expect(5).toEqual(add(3,2));
});
```

Well, one way we could make this pass is to check for the parameters and create some switch cases... But, as you can see this is growing and, to be honest, it's not the way should do things, so we refactor.

So, the rule of thumb, the third step is REFACTOR and make sure the tests are still passing.

In the moment of inspiration we write (in `index.js` file):

```
function add (a, b){
	return a + b;
}
```

and with that, we now have a passing test and refactored code.

### Making the output prettier

At this point, it may not be so nicely presented what all specs we have as passing. And, if you want to see that, you can install:

```
npm install karma-spec-reporter --save-dev
npm install jasmine-spec-reporter --save-dev
```

And then, in the `karma.conf.js` file just change the reporter to `spec`, like this:

```
reporters: ['spec']
```

Now when we run `karma start` we will have a nice output like:

```
Calculator
    ✓ should have an add function
    ✓ should return 3 when passed 1, 2
    ✓ should return 5 when passed 3, 2

PhantomJS 2.1.1 (Mac OS X 0.0.0): Executed 3 of 3 SUCCESS (0.002 secs / 0.002 secs)
TOTAL: 3 SUCCESS
```

Just a quick note on how to skip a certain test, by adding x before it:

```
xit ("should return 5 when passed 3, 2", function (){
    expect(5).toEqual(add(3,2));
});
```

Karma then reports this in the console log:

```
Calculator
    ✓ should have an add function
    ✓ should return 3 when passed 1, 2
    - should return 5 when passed 3, 2
```

indicating that the last test was skipped.

### Full source and test code listing

Just for reference, this is how the `index.spec.js` file would look like when we add the tests for the `sub` function:

```
describe ("Calculator", function (){
	
	describe ("add function", function (){
		it('should have an add function', function() {
		    expect(add).toBeDefined();
		});

		it ("should return 3 when passed 1, 2", function (){
		    expect(3).toEqual(add(1,2));
		});

		it ("should return 5 when passed 3, 2", function (){
		    expect(5).toEqual(add(3,2));
		});
	});

	describe ("sub function", function (){
		it('should have an sub function', function() {
		    expect(sub).toBeDefined();
		});

		it ("should return -1 when passed 1, 2", function (){
		    expect(-1).toEqual(sub(1,2));
		});

		it ("should return 1 when passed 3, 2", function (){
		    expect(1).toEqual(sub(3,2));
		});
	});

});
```

This is the contents of the `index.js` file:

```
function add(a, b) {
	return a + b;
}

function sub(a, b) {
	return a - b;
}
```

And, this is what Karma would output to the console once run at this point:

```
Calculator
    add function
      ✓ should have an add function
      ✓ should return 3 when passed 1, 2
      ✓ should return 5 when passed 3, 2
    sub function
      ✓ should have an sub function
      ✓ should return -1 when passed 1, 2
      ✓ should return 1 when passed 3, 2
```

If you want to take a look at the whole code, you can [fork it on Github](https://github.com/Hitman666/jstdd).

### Wallaby
This all is pretty cool and you can have your terminal opened up and see how your test turns green. However, as with everything these days, there are better tools out there. One such tool is [Wallabyjs](https://wallabyjs.com/docs/intro/how-it-works.html). And, let me just show you what it can do.

First, you have to install Wallaby for your editor. They support Visual Studio Code, Atom, Submlime, Webstorm, etc.

After you've installed it, you have to set its config file. Let's create a new file and name it `wallaby.js` and place it in the root of our app. Copy/Paste the following code into it:

```
module.exports = function (wallaby) {
  return {
    files: [
      'app/*.js'
    ],

    tests: [
      'tests/*.spec.js'
    ],
    debug: true
  };
};
```

_You may have to restart your editor at this point_. At this point, you just run Wallaby from within your editor. In Sublime it's done by pressing `CMD + SHIFT + P` and selecting `Wallaby.js: Start`. There is also a handy shortcut in sublime: `CMD + .` followed by `CMD + R`.

As you will see, you now have information about your tests passing (green rectangles on the left-hand side) or failing inside the actual editor:

![](https://i.imgur.com/QHs7BuB.png)

There are actually a lot more features to Wallaby, which I'll leave to you to explore. I'm not affiliated with them in any way; I just happen to like it. But, just so you don't say I didn't mention it; as every great tool, it has its price. And, if you're contemplating (or even complaining) about whether or not you should pay for certain software, please read this awesome post by Ambrose Little on [How Much Is Your Productivity Worth?](https://ambroselittle.svbtle.com/how-much-is-your-productivity-worth).

Ok, so this was the JavaScript tutorial. Let's now take a look how would we setup up Jasmine and Karma in the Ionic framework application.

## The Ionic framework demo
_You need to have Ionic and Cordova packages installed globally with npm in order to follow this part of the tutorial. You can learn more about how to do that in [Ionic Framework: A definitive 10,000 word guide](http://www.nikola-breznjak.com/blog/javascript/ionic/ionic-framework-a-definitive-10000-word-guide/)._

### Starting a new project and installing prerequisites
First, we start a new Ionic project:

`ionic start ionic-tdd tabs`

Next, we go inside this folder and install the necessary prerequisites.

```
cd ionic-tdd
npm install karma karma-jasmine karma-phantomjs-launcher jasmine-core --save-dev
```

### Setting up Karma
Please make sure you have Karma installed globally from the previous JavaScript section. If you don't you can do this simply with:

`npm install -g karma-cli`

Also, at this point, we have to run `npm install` to install all the prerequisites from the Ionic `package.json` file. 

Finally, we need to install `angular-mocks` with bower:

`bower install angular-mocks --save-dev`

since we'll use that to mock certain Angular controllers.

Once this is done we create a new folder in our project's root directory. Let's call it `tests`:

`mkdir tests`

Also, let's run `karma init` command (run this command in your terminal, once in the root directory of your project).

You can follow the same instructions for Karma as in the JavaScript section, just don't enter the location of the source and test files, we'll add them separatelly.

Now we have to open up the `karma.conf.js` file and add our source and test files:
	
```
files: [
        'www/lib/angular/angular.js',
        'www/js/*.js',
        'www/lib/angular-mocks/angular-mocks.js',
        'tests/*.spec.js'
],
browsers: ['PhantomJS']
```

In the next step, we'll configure our `gulpfile.js` file, so that we'll be able to run our test via [Gulp](http://gulpjs.com/), since Ionic uses it as it's task runner. We import Karma at the top of the file:

`var karmaServer = require('karma').Server;`

And we write a new task called `test`:
	
```
gulp.task('test', function(done) {
    new karmaServer({
        configFile: __dirname + '/karma.conf.js',
        singleRun: false
    }).start();
});
```

Now, we can run `gulp` with the `test` parameter like this: `gulp test`.

### Testing the controller
First, let's create a new `tests/controllers.spec.js` file in the `tests` folder.

Please note that this now isn't a TDD approach, since we already have the code in our controller written. But, if you ever come to a project that hasn't got unit tests this is what you'll be doing. Plus, all the refactoring to make the code testable, but that's a different story for some other time...

We start by writing our describe function:

```
describe('Controllers', function(){

});
```

Next, since this is Angular, we'll have a local scope variable (`var scope`). And before each test, we have to load the `starter.controller` module:

`beforeEach(module('starter.controllers'));`

How do we know we have to set this module? Well, if you take a look at the `controllers.js` file, you'll see the name of the module there on the top as `starter.controllers`.

Also, we need to inject Angular's scope variable and set the controller.

```
beforeEach(inject(function($rootScope, $controller) {
	scope = $rootScope.$new();
   	$controller('AccountCtrl', {$scope: scope});
}));
```

To put this all in one place, you should have a `controllers.spec.js` file that looks like this:

```
describe('Controllers', function(){
	var scope;

	beforeEach(module('starter.controllers'));

	beforeEach(inject(function($rootScope, $controller) {
	    scope = $rootScope.$new();
	    $controller('AccountCtrl', {$scope: scope});
	}));
});
```

This is a boilerplate code that you'll have to write in every test, so though it may seem strange at first, it becomes something you don't think about after you've worked with it for some time.

Again, if you wonder how we came to the `AccountCtrl`, just take a look at the `controllers.js` file and the name of the controller we're trying to test.

Finally, we come to our test. And, say we want to test if the `enableFriends` property on the `settings` object is set to `true`, we would write a test like this:
	
```
it('should have enableFriends property set to true', function(){
    expect(scope.settings.enableFriends).toEqual(true);
});
```

Now we run our tests with `gulp test` and we can see our test is passing.

### Testing the service/factory

Now we're going to write a test for our factory `Chats`. As you can see, the factory has three functions for getting all chats (that are currently hard-coded), removing a chat and getting a specific chat.

First, we'll create a new file in the `tests` folder called `services.spec.js` and add our `describe` function:
	
```
describe('Chats Unit Tests', function(){

});
```

Next, we're going to set the module and inject the Chats factory:

```
var Chats;
beforeEach(module('starter.services'));

beforeEach(inject(function (_Chats_) {
    Chats = _Chats_;
}));
```

Now, we can write our first test, and well, let's first test if our Chats factory is defined:

```
it('can get an instance of my factory', inject(function(Chats) {
    expect(Chats).toBeDefined();
}));
```

Then, we can check if it returns five chats

```
it('has 5 chats', inject(function(Chats) {
    expect(Chats.all().length).toEqual(5);
}));
```

If at this point, we also want to see a nicer spec reports, we should kill the currently running gulp process. Install the required packages:

```
npm install karma-spec-reporter --save-dev
npm install jasmine-spec-reporter --save-dev
```

adjust the `karma.conf.js` file:

`reporters: ['spec'],`

and rerun gulp with `gulp test`.

To put this all in one place, you should have `services.spec.js` file that looks like this:

```
describe('Chats Unit Tests', function(){
	var Chats;
	beforeEach(module('starter.services'));

	beforeEach(inject(function (_Chats_) {
	    Chats = _Chats_;
	}));

	it('can get an instance of my factory', inject(function(Chats) {
	    expect(Chats).toBeDefined();
	}));

	it('has 5 chats', inject(function(Chats) {
	    expect(Chats.all().length).toEqual(5);
	}));
});
```

If you want to take a look at the whole code, you can [fork it on Github](https://github.com/Hitman666/ionic-tdd).

### Wallaby
If you want to try Wallaby in Ionic you just need to create the `wallaby.js` file and set the configuration:

```
module.exports = function (wallaby) {
  return {
    files: [
      	'www/lib/angular/angular.js',
        'www/js/*.js',
        'www/lib/angular-mocks/angular-mocks.js',
    ],

    tests: [
      	'tests/*.spec.js'
    ],
    debug: true
  };
};
```

## Conclusion
My personal takeaway from this so far is that even if you don't adopt this whole TDD mantra, I'm urging you to start using Unit tests at least, as you've seen how valuable they can be. As for the whole TDD mantra, I'm yet to see how all this pans out, as I feel that adopting this properly requires a certain discipline until implemented properly.

Of course, all this is just a tip of the iceberg. I just touched the Unit tests and what Jasmine can do as your test environment. I hope that some time from now I'll be able to share with you some best practices and some advanced techniques. Until then, I hope this was useful to some of you to at least get you going.

Demo projects are on Github:

+ [JavaScript demo]()
+ [Ionic framework demo](https://github.com/Hitman666/ionic-tdd)

And yes, take the red pill ;)

![](https://i.imgur.com/SzkbXjDm.png)

---

In case someone is interested, below is my path to the ever so slightly awesome TDD regarding the read materials and the notes I collected along the way.

## [Treehouse course](https://teamtreehouse.com/library/javascript-unit-testing)

+ Use E2E test sparringly (this is in line with the [Google post](http://googletesting.blogspot.hr/2015/04/just-say-no-to-more-end-to-end-tests.html))
+ suits and specs
+ `mocha --reporter nyan`
+ `"scripts": {"test":mocha, "test:watch":"mocha --watch ./test ./"}`
+ `npm run test:watch`

## Books on the topic

+ [Test Driven Development, Kent Beck](http://amzn.to/1SGPsks)
+ [Refactoring: Improving the Design of Existing Code](http://amzn.to/1VUpnBV)
+ [Ionic in action - chapter about TDD in Ionic](http://amzn.to/1NqbdGh)

## Blog posts

### [Introduction to JS TDD](http://jrsinclair.com/articles/2016/gentle-introduction-to-javascript-tdd-intro/)

Advantages of TDD:

+ It forces one to think
+ It makes debugging easier
+ It makes coding more fun

TDD is an approach to writing software where you **write tests before you write application code**. The basic steps are:

+ **Red**  - write a test and make sure it fails
+ **Green** - write the simplest, easiest possible code to make the test pass
+ **Refactor** - optimize and/or simplify the application code, making sure that all the tests still pass

You have to think first, then write a test.

```
// flickr-fetcher-spec.js
'use strict';
var expect = require('chai').expect;

describe('FlickrFetcher', function() {
    it('should exist', function() {
        var FlickrFetcher = require('./flickr-fetcher.js');
        expect(FlickrFetcher).to.not.be.undefined;
    });
});
```

We are using a `Behaviour Driven Development (BDD)` style to write the tests. This is a variation on TDD where tests are written in the form:

+ Describe `[thing]`
+ It should `[do something]`

The `[thing]` can be a module, or a class, or a method, or a function. Mocha includes built-in functions like `describe()` and `it()` to make writing in this style possible.

No module code until there’s a failing test. So what do I do? I write another test.

The rule of thumb is, use **equal** when comparing numbers, strings, or booleans, and use **eql** when comparing arrays or objects. Note: `eql` is named `deepEqual` in some other testing frameworks. However, note that Jasmine has only `toEqual`.

### [Introduction to JS TDD Part 2](http://jrsinclair.com/articles/2016/gentle-introduction-to-javascript-tdd-ajax/)

The `fakeFetcher()` function I’ve used to replace `$.getJSON()` is known as a **stub**. A stub is a piece of code that has the same API and behavior as the ‘real’ code, but with much-reduced functionality. Usually, this means **returning static data** instead of interacting with some external resource.

Typical stubs might replace things like:

+ Queries to a relational database
+ Interaction with the file system
+ Accepting user input
+ Complex computations that take a long time to calculate

### [TDD should be fun](http://jrsinclair.com/articles/2016/tdd-should-be-fun/)

+ functional tests (E2E)
+ integration tests, more often than E2E

### [The ever so slightly famous Eric Elliot on the subject of JS testing](http://www.sitepoint.com/javascript-testing-unit-functional-integration/)

+ Unit tests, integration tests, and functional tests are all types of automated tests which form essential cornerstones of continuous delivery, a development methodology that allows you to safely ship changes to production in days or hours rather than months or years.
+ The cost of a bug that makes it into production is many times larger than the cost of a bug caught by an automated test suite. In other words, TDD has an overwhelmingly positive ROI.
+ You don’t choose between unit tests, functional tests, and integration tests. Use all of them, and make sure you can run each type of test suite in isolation from the others.

---

+ **Unit tests**
	+ ensure that individual components of the app work as expected. Assertions test the **component API**
+ **Integration tests**
	+ ensure that component collaborations work as expected. Assertions may test component API, UI, or side-effects (such as database I/O, logging, etc…)
+ **Functional tests**
	+ ensure that the app works as expected from the user’s perspective. Assertions primarily test the **user interface**

> For example, your app may need to route URLs to route handlers. A unit test may be written against the URL parser to ensure that the relevant components of the URL are parsed correctly. Another unit test might ensure that the router calls the correct handler for a given URL. 
However, if you want to test that when a specific URL is posted to, a corresponding record gets added to the database, that would be an integration test, not a unit test.

Yes, you write more code, but [studies have shown](http://evidencebasedse.com/?q=node/78) objectively that **good test coverage with TDD can reduce bug density by 40% - 80%**.

Another two posts from him:

[5 Common Misconceptions About TDD & Unit Tests](https://medium.com/javascript-scene/5-common-misconceptions-about-tdd-unit-tests-863d5beb3ce9#.f1xankzgr)

+ TDD is too Time Consuming. The Business Team Would Never Approve
+ You Can’t Write Tests Until You Know the Design, & You Can’t Know the Design Until You Implement the Code
+ You Have to Write All Tests Before You Start the Code
+ Red, Green, and ALWAYS Refactor?
+ Everything Needs Unit Tests

> Here’s a tip that will change your life: `Mocking is a code smell.`

[5 Questions Every Unit Test Must Answer](https://medium.com/javascript-scene/what-every-unit-test-needs-f6cd34d9836d#.9uudcx44o)

+ What’s in a good test failure bug report?
+ What were you testing?
+ What should it do?
+ What was the output (actual behavior)?
+ What was the expected output (expected behavior)?


### Few general good blog posts

+ [Google's take on E2E, Integration and Unit tests](http://googletesting.blogspot.hr/2015/04/just-say-no-to-more-end-to-end-tests.html)
+ [TDD is dead, long live testing](http://david.heinemeierhansson.com/2014/tdd-is-dead-long-live-testing.html)
+ [Test-Driven Development Isn't Testing](https://www.stickyminds.com/article/test-driven-development-isnt-testing)
+ [Triangulation in TDD](http://elnur.pro/triangulation-in-testing/)
+ [Introduction to Test Driven Development in JavaScript](http://tutorials.pluralsight.com/front-end-javascript/introduction-to-test-driven-development-in-javascript)
+ [Making your functions pure](http://alistapart.com/article/making-your-javascript-pure)
+ [Writing great unit tests](http://blog.stevensanderson.com/2009/08/24/writing-great-unit-tests-best-and-worst-practises/)
	+ Unit testing is not about finding bugs but it is excellent when refactoring
+ [Testing services in Angular for fun and Profit](http://nathanleclaire.com/blog/2014/04/12/unit-testing-services-in-angularjs-for-fun-and-for-profit/)
	+ If there was a way to reduce the number of defects in the code you write (or manage), improve the quality and time to market of deliverables, and make things easier to maintain for those who come after you- would you do it?
	+ How many times have you heard some variant on, “Writing tests isn’t as important as delivering finished code?” If you’re like me, it’s way too many, and god help you if you’re working with no tests at all. Programmers are human and we all make mistakes. So test your code. The number of times testing my code has helped me catch unforeseen issues before they became flat-out bugs, prevent future regressions, or simply architect better is pretty amazing. And this is coming from a guy who used to hate writing tests for the code. Hated it.
	+ Jasmine is a Behavior-Driven-Development framework, which is sort of a roundabout way of saying that our tests include descriptions of the sections that they are testing and what they are supposed to do.
	+ You can create stubbed objects quite easily in JavaScript, so if there’s no need to introduce the extra complexity of a spy, then do so.
	+ Always code as if the person who ends up maintaining your code is a violent psychopath who knows where you live.
+ [One Weird Trick That Will Change The Way You Code Forever: Javascript TDD](http://jrsinclair.com/articles/2016/one-weird-trick-that-will-change-the-way-you-code-forever-javascript-tdd/)
	+ Have you ever fixed a bug, only to find that it broke something horribly in another part of the system? And you had no idea until the client called support in a panic?
	+ Have you ever been afraid to touch a complicated piece of code for fear that you might break it and never be able to fix it again? … Even though you wrote it?
	+ Have you ever found a piece of code that you’re pretty sure wasn’t being used anymore and should be deleted? But you left it there just in case?
	+ **TDD is not about testing**. It is a way of thinking and coding that just-so-happens to involve tests.
	+ TDD is not the same thing as unit tests. Unit tests are a type of test. TDD is a **coding technique**.
		+ Red—write a little test that doesn’t work, perhaps doesn’t even compile at first
		+ Green—make the test work quickly, committing whatever sins necessary in the process
		+ Refactor—eliminate all the duplication created in just getting the test to work



## Finally, Ionic (Angular) related TDD posts

### [How To Write Automated Tests For Your Ionic App](http://gonehybrid.com/how-to-write-automated-tests-for-your-ionic-app-part-1/)
+ In the example for Unit Tests we saw that we need to mock the dependencies. For Integration Tests, depending on which units you want to test together, you could still mock certain dependencies or none at all.

### [TDD with ionic](https://kapware.com/tdd-with-ionic/)
+ Short tutorial showcasing how to run Karma with Jasmine

### [Unit Testing Your Ionic Framework App](http://mcgivery.com/unit-testing-ionic-app/)

	This tutorial was actually great (which I can't say for the previous two) and I've learned the most out of it and finally set up a test environment.

	Fun fact: I added `npm install --save-dev karma-nyan-reporter` and now am running my tests like this: `karma start tests/my.conf.js --reporters nyan
	
## Some other AngularJS TDD blog posts

+ [Unit Testing an AngularJS Ionic App with Codeship Continuous Integration, Jasmine, and Karma](http://blog.pdsullivan.com/posts/2014/12/05/ionic-codeship-unit-testing-ci.html)
+ [Unit Testing Best Practices in AngularJS](http://andyshora.com/unit-testing-best-practices-angularjs.html)
+ [Official AngularJS Unit Testing Guide](https://docs.angularjs.org/guide/unit-testing)
	+ Underscore notation: The use of the underscore notation (e.g.: `_$rootScope_`) is a convention widespread in AngularJS community to keep the variable names clean in your tests. That's why the $injector strips out the leading and the trailing underscores when matching the parameters. The underscore rule applies only if the name starts and ends with exactly one underscore, otherwise no replacing happens. 
+ [Add Karma And Jasmine To An Existing Ionic Project](https://novas1r1.wordpress.com/2014/10/27/add-karma-and-jasmine-to-an-existing-ionic-project/)
+ [Unit testing AngularJS applications](https://www.airpair.com/angularjs/posts/unit-testing-angularjs-applications)
+ [Testing AngularJS with Jasmine and Karma](https://scotch.io/tutorials/testing-angularjs-with-jasmine-and-karma-part-1)


## My notes

+ `npm install phantomjs-prebuilt` was needed in order to get Karma running with PhantomJS.

+ Had to change the actual Angular mocks 1.5.1 error in the code ([https://github.com/angular/angular.js/issues/14251](https://github.com/angular/angular.js/issues/14251)).

At this point, the tests finally passed!

## Tools
[Wallabyjs](https://wallabyjs.com/docs/intro/how-it-works.html) - An awesome tool