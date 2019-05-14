# Go training slides

This projects contains presentations rendered by [`present`](https://godoc.org/golang.org/x/tools/cmd/present).

[`content/`](content) contains the presentations following the [present syntax](https://godoc.org/golang.org/x/tools/present)

This projects is continuously deployed to Appengine at [`https://gotraining.appspot.com/`](https://gotraining.appspot.com) when merged into `master`.

## Run locally with

Run the present tool locally with

    go build .
    ./go-slides -base . -content content


## Appengine deployment

The "production" slides are deployed to the `gotraining` project. There is
limited access to this project - slides are automatically deployed when
the master branch of this repository is updated (usually merging a PR).

The `gotraining-testing` project is available with laxer permissions to
allow you to deploy pre-release versions of the slides for testing, and
should be used to deploy your PR version for review.

You should set up one or both of these projects with the `gcloud`
command so that you can deploy PR versions and/or manually deploy the
production version as necessary:

    gcloud config configurations create gotraining-testing
    gcloud config set account your@email.address
    gcloud config set project gotraining-testing

You with need to authenticate the new configuration:

    gcloud auth login
    
If you are having SSL validation issues (due to a corporate network or
otherwise), you can disable SSL validation as a last resort:

    gcloud config set auth/disable_ssl_validation

### Deployment for PR review

When you have a PR open for review, you can deploy the changes under
review as an appengine "service" so it can be deployed alongside the
master/default version of the slides as well as other open PRs. It can
easily be cleaned up when your PR is closed.

1. Edit `app.yaml` and add `service: prNN` to the end of the file
   (top-level yaml key) where `NN` is your pull request number. You
   could put it at the top of the file if that is easier.

1. Deploy to app engine: `gcloud --configuration=gotraining-testing app deploy`

1. In your browser, go to https://prNN-dot-gotraining-testing.appspot.com/

1. Review the changes as your reviewer will see them

1. When done, delete the app engine service:
     `gcloud --configuration=gotraining-testing app services delete prNN`

1. Undo changes to `app.yaml`: `git checkout -- app.yaml`

Please make sure you do not commit a version of `app.yaml` with your
`service:` line in it. This should only ever be local to your workspace.

You should also make sure your workspace is clean of changes so what you
deploy is what is on the PR branch in github. run `git describe --all
--dirty` and if that outputs a string with `-dirty` on the end, you have
uncommitted changes. Stash them before you deploy (`git stash -u`) and
unstash after (`git stash pop`). Note that you should check this before
modifying the `app.yaml` file as that will make your workspace dirty.
You can run `git status` and `git diff` if you need to and ensure the
only dirtyness is the `service: prNN` line in `app.yaml`.

### Manual production Appengine deployment

Request access to GCP `gotraining` project from a [contributor](https://github.com/anz-bank/go-slides/graphs/contributors). 

Add a new gcloud configuration as described in the [Appengine deployment](#appengine-deployment) section, using `gotraining` instead of `gotraining-testing`.

With the `gotraining` configuration active, Execute

    gcloud app deploy

or test the coudbuild with

    gcloud builds submit --config cloudbuild.yaml

## Attribution

The Go course materials in this project were collated by [@juliaogris](https://github.com/juliaogris) with many cues taken from [A Tour of Go](https://tour.golang.org/) and [Go by Example](https://gobyexample.com/).

Extra special thanks to [@camh-anz](https://github.com/camh-anz) for countless reviews, improvement suggestions and code contributions.

Many thanks also to [@anzdaddy](https://github.com/anzdaddy) for suggesting the pingserver comparison, to [@pentaphobe](https://github.com/pentaphobe) for the syntax colouring contribution and to [@anzboi](https://github.com/anzboi) for creating [go-samplerest](https://github.com/anz-bank/go-samplerest) - they and [@anz-rfc](https://github.com/anz-rfc) also contributed in many reviews.

[@anzboi](https://github.com/anzboi), [@camh-anz](https://github.com/camh-anz), [@linuxmonk](https://github.com/linuxmonk), [@Joshcarp](https://github.com/Joshcarp), [@juliaogris](https://github.com/juliaogris) and [@rokane](https://github.com/rokane) collaborated on an updated version of the materials for a follow up course.

## License

This work is copyright _Australia and New Zealand Banking Group Limited_ and licensed under a [Creative Commons Attribution 3.0 Unported License](https://creativecommons.org/licenses/by/3.0/).
