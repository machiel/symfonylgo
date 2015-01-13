symfonylgo
============================

A small logfile splitter on date written in go. Usage:

`symfonylgo logfilename`

For example, running `symfonylgo prod.log` will produce:

	prod.2015-01-10.log
	prod.2015-01-11.log
	prod.2015-01-12.log
	prod.2015-01-13.log

-------------------

Even though you can use `logrotate` to achieve this, or simply use default linux tools like `awk` and `sed`, I wanted a quick tool that would easily split my logfiles into files based on their dates. Besides that it gave me a chance to play around with Go for a bit :).

This is a very raw first version, without any tests. Please create a PR if you'd like to add anything.

You can download the executable here, for Linux i386. https://github.com/Machiel/symfonylgo/releases/tag/v0.1
