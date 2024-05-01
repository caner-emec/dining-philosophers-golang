# dining-philosophers-golang

Dining Philosophers Problem With Golang.

![](image.png)

## **Problem:**

Implement the dining philosopher’s problem with the following constraints/modifications.

- There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.

- Each philosopher should eat only 3 times.

- The philosophers pick up the chopsticks in any order, not lowest-numbered first.

- In order to eat, a philosopher must get permission from a host which executes in its own goroutine.

- The host allows no more than 2 philosophers to eat concurrently.

- Each philosopher is numbered, 1 through 5.

- When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.

- When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.

</br>

## Example Output:

**command:**

```sh
go run philosophers.go
```

**output:**

```sh
Starting to eat 4
Starting to eat 1
Finishing to eat 1
Starting to eat 2
Finishing to eat 4
Starting to eat 0
Finishing to eat 0
Starting to eat 4
Finishing to eat 2
Starting to eat 1
Finishing to eat 1
Finishing to eat 4
Starting to eat 2
Finishing to eat 2
Starting to eat 0
Starting to eat 3
Finishing to eat 3
Finishing to eat 0
Starting to eat 1
Starting to eat 4
Finishing to eat 4
Starting to eat 3
Finishing to eat 1
Starting to eat 0
Finishing to eat 0
Finishing to eat 3
Starting to eat 3
Finishing to eat 3
Starting to eat 2
Finishing to eat 2
```
