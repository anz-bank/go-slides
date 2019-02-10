# Go training slides

This projects contains presentations rendered by [`present`](https://godoc.org/golang.org/x/tools/cmd/present).

[`content/`](content) contains the presentations following the [present syntax](https://godoc.org/golang.org/x/tools/present)

This projects is continuously deployed to Appengine at [`https://gotraining.appspot.com/`](https://gotraining.appspot.com) when merged into `master`.

## Run locally with

Run the present tool locally with

    go build .
    ./go-slides -base . -content content

## Manual Appengine deployment

Request access to GCP `gotraining` project from a [contributor](https://github.com/anz-bank/go-slides/graphs/contributors) or update [`app.yaml`](app.yaml) on your own fork with _your_ GCP project.

Execute

    gcloud app deploy

or test the coudbuild with

    gcloud builds submit --config cloudbuild.yaml

## Attribution

The training materials in this project as of February 2019 were collated by [@juliaogris](https://github.com/juliaogris) with many cues taken from [A Tour of Go](https://tour.golang.org/) and [Go by Example](https://gobyexample.com/).

Extra special thanks to [@camh-anz](https://github.com/camh-anz) for countless reviews, improvement suggestions and code contributions.

Many thanks also to [@anzdaddy](https://github.com/anzdaddy) for suggesting the pingserver comparison, to [@pentaphobe](https://github.com/pentaphobe) for the syntax colouring contribution and to [@anzboi](https://github.com/anzboi) for creating [go-samplerest](https://github.com/anz-bank/go-samplerest) - they and [@anz-rfc](https://github.com/anz-rfc) also contributed in many reviews.

## License

This work is copyright _Australia and New Zealand Banking Group Limited_ and licensed under a [Creative Commons Attribution 3.0 Unported License](https://creativecommons.org/licenses/by/3.0/).

