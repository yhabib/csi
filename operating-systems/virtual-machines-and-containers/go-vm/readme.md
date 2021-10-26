# Containers in Go

**Important!!** we need to compile the code for Linux otherwise `Cloneflags` will not work: 
`GOOS=linux go run main.go ....`

## Examples

To see the output of the command we need to redirect its outputs to the std family. Now this example should work:

* `go run main.go run echo hello world`
* `go run main.go run echo ls /`

We can run a new shell within the container but not isolated yet:

* `go run main.go run /bin/zsh`

I can see and edit the parent's hostname for example:

* `hostname`
* `hostname LA`

We need Namespace to isolate the access of the child

$$

$$## Resources

*[Containers in GO](https://learning.oreilly.com/library/view/how-to-containerize/9781491982310/ch01.html#idm140204518951312)