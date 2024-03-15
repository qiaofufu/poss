all: one two

one two:
	echo $@


test: $(wildcard *.go)
	printf $?