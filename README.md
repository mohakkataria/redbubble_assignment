# Redbubble Coding Test

## Requirements
http://take-home-test.herokuapp.com/full-stack-engineer

## How to run it
* First thing first, clone or download the repo to your local machine.
* Install Go.

### CLI
Run the following command in your repo directory:
~~~
> go install
> ./main {endpoint} {output_dir}
~~~
`{endpoint}` is the API url, while `{output_dir}` is the directory containing the generated files.

### Strategy
In this project, images are supplied by XML, but in order to account for extensibility for the source to be a JSON file or a CSV file, `Provider Interface` is provided to implement different source providers to convert image data to a collection of domain objects `Image`.
Similarly, `Repository Interface` gives flexibility to implement a repository service to retrieve data. At the moment, the image objects are stored in memory, but the repository could be an implementation of a database like MySQL or something else altogether.

## Testing
The implementation is TDD-driven. You can verify the tests as follows:

~~~
$ go test ./... --cover
~~~

## GoReportCard
https://goreportcard.com/report/github.com/mohakkataria/redbubble_assignment