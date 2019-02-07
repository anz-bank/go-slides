# Go training slides

This projects contains presentations rendered go [`present`](https://godoc.org/golang.org/x/tools/cmd/present).

[`content/*.slides`](content) contains the presentations following the [present syntax](https://godoc.org/golang.org/x/tools/present)

It is continuously deployed to Appengine at [`https://gotraining.appspot.com/`](https://gotraining.appspot.com) when merged into `master`.

## Run locally with

Run the present tool locally with

    go build .
    ./go-slides -base . -content content

## Manual Appengine deployment

Request access rights to the GCP `gotraining` project from a [contributor](https://github.com/anz-bank/go-slides/graphs/contributors) or update [`app.yaml`](app.yaml) on your own fork with your own GCP project.

Execute

    gcloud app deploy

Test the coudbuild with

    gcloud builds submit --config cloudbuild.yaml

## Attribution

The training materials in this project as of February 2019 were collated by @juliaogris with many cues taken from [A Tour of Go](https://tour.golang.org/) and [Go by Example](https://gobyexample.com/).

Extra special thanks to @camh-anz for countless reviews, improvement suggestions and code contributions.

Many thanks also to @anzdaddy for suggesting the pingserver comparison, to @pentaphobe for the syntax colouring contribution and to @anzboi for creating [go-samplerest](https://github.com/anz-bank/go-samplerest) - they and @anz-rfc also contributed in many reviews.
