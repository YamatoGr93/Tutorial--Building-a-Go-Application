### **README.md**

```markdown
# Go Task Manager

This is a simple command-line task manager built using Go. It allows users to add, list, and delete tasks. The application is a basic demonstration of Go's functionality, focusing on slices, functions, and basic input/output.

## Features

- Add tasks
- List tasks
- Delete tasks
- Command-line interface

## Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16+)

## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/your-username/go-task-manager.git
cd go-task-manager
```

### 2. Initialize the Go module

If you haven't done this already, initialize the Go module in the project folder:

```bash
go mod init go-task-manager
```

### 3. Run the application

Use the following command to run the application:

```bash
go run main.go
```

### 4. Use the Task Manager

Once the program is running, you will see a menu like this:

```
Task Manager
1. Add a task
2. List tasks
3. Delete a task
4. Exit
```

- To **add a task**, type the task description after selecting option 1.
- To **list tasks**, choose option 2, and all tasks will be displayed.
- To **delete a task**, choose option 3 and provide the task number.

### Example:

```
Task Manager
1. Add a task
2. List tasks
3. Delete a task
4. Exit
Enter choice: 1
Enter task: Complete Go tutorial

Task Manager
1. Add a task
2. List tasks
3. Delete a task
4. Exit
Enter choice: 2
1. Complete Go tutorial
```

## Future Improvements

- Save tasks to a file for persistence across sessions.
- Add error handling for invalid input.
- Support for editing tasks.
