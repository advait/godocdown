/*
Command godocdown extracts and generates Go documentation in a GitHub-friendly Markdown format.

This program is targeted at providing nice-looking documentation for GitHub. With this in
mind, it generates GitHub Flavored Markdown (http://github.github.com/github-flavored-markdown/) by
default. This can be changed with the use of the "plain" flag to generate standard Markdown.

	$ go get github.com/robertkrimen/godocdown/godocdown

	$ godocdown /path/to/package > README.markdown

	# Generate documentation for the package/command in the current directory
	$ godocdown > README.markdown

	# Generate standard Markdown
	$ godocdown -plain . 

Installation

	go get github.com/robertkrimen/godocdown/godocdown

Example

http://github.com/robertkrimen/godocdown/blob/master/example.markdown

Usage

The following options are accepted:

	-heading="TitleCase1Word"
	Heading detection method: 1Word, TitleCase, Title, TitleCase1Word, ""

	-no-template=false
	Disable template processing

	-plain=false
	Emit standard Markdown, rather than Github Flavored Markdown (the default)

*/
package documentation
