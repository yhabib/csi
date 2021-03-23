PROGRAM=ls.out

clean: 
	rm -rf $(PROGRAM)*

build: clean
	cc -Werror -Wextra -g ls.c -o $(PROGRAM) 

debug: build
	lldb ./$(PROGRAM)