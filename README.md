
# Auth tutorial

This is the example to the buffalo auth tutorial.

## Installation

After cloning this repository you need to run:

```
npm install
```

Go to GitHub settings and [Register a new OAuth application](https://github.com/settings/applications/new).

Create a file named `.env` in the root directory of this repository. Add following lines to the file:

```
GITHUB_KEY=YOUR_KEY
GITHUB_SECRET=YOUR_SECRET
```


## Starting the Application

Buffalo ships with a command that will watch your application and automatically rebuild the Go binary and any assets for you. To do that run the "buffalo dev" command:

	$ buffalo dev

If you point your browser to [http://localhost:3000](http://localhost:3000) you should see a "Welcome to Buffalo!" page.

**Congratulations!** You now have your Buffalo application up and running.

## What Next?

We recommend you heading over to [http://gobuffalo.io](http://gobuffalo.io) and reviewing all of the great documentation there.

Good luck!

[Powered by Buffalo](http://gobuffalo.io)