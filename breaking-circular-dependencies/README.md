# go-breaking-circular-dependencies

This is just a PoC about my learning journey in Software Engineering with Golang.

It is related to a common problem when implementing packages to be provided to end-users.

Actually, Golang deny the compilation when encounter an `import cycle not allowed` compilation error.

Here's a possible implementation of packages dependent on each other.