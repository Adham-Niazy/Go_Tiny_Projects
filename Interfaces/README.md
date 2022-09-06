# Interfaces 
_look at the two files before reading this_

While using this line `type bot interface` our program has a new type called `bot`, And within this interface we will define this => `getGreeting() string`, So If you are any type in this program with a function called `getGreeting` and you return a string then you are now an honorary member of type `bot`.

Now you can call this function => `printGreeting` that will receive a `bot` without being interested what kind of bot this is (Spanish, English, Russian,... etc)


# Important Notes about Interfaces

1. **Interfaces are not generic types**: Other languages have generic' 'types-go
(famously) does not.

2. **Interfaces are 'implicit'**: We don't manually have to say that our custom type satisfies some interface.

3. **Interfaces are a contract to help us manage types**: GARBAGE IN->GARBAGE OUT. If our custom type's implementation of a function is broken then interfaces wont help us!

4. **Interfaces are tough. Step#1 is understanding how to read them**: Understand how to read interfaces in the standard lib. Writing your own interfaces is tough and requires experience.