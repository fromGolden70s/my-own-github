# My own github

This is a Version Control System that can track changes implemented in a program.
Pass arguments through a command line, e.g.

$go build main.go

$go run main.go config John123

•	--help -- prints the help page;

•	config [name] -- sets or outputs the name of a commit author;

•	add -- adds a file to the list of tracked files or outputs this list;

•	log -- shows all commits;

•	commit [comment] -- saves file changes and the author name;

•	checkout [commit ID] -- allows you to switch between commits and restore a previous file state.
