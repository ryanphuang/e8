**Road Plan**

- Build a MIPS-like very simple virtual machine
- Build an assembler
- Build a compiler for a Go-like language
- Reimplement all the above in Javascript
- Build an online platform where people can submit interfaces, test cases, and
  implementations
- Test cases works like executables (`package main`, but maybe `package test`
  is better). A test case link with other interface and implementation modules,
  and generate test scores as output. It could be testing the actual implementation, or 
  just defining the public interfaces of a type of module.
- Interfaces and implementations are general library modules.
- A module can only access other modules via publicly defined interfaces,
  methods, functions, types, consts.
- Users write modules under version control (perferable the website works like
  something similar to github, but more language aware, where tests are automatically
  performed.)
- This is just for single vm open source testing
- For more complex stuff, we can have multiple vm testing, where all calls are
  performed over network RPC calls

**Why**

- This could be an online platform for ACM training
- This could be an online arena/platform for game AI fighting
- This could be used for collaborative programming
- This encourages users to collaborate by assembling modules via clearly
  defined interfaces and test them with testcases.
- Ideally, all the modules and interfaces could be easily reused, since they
  would likely be small, and hence easy for human to read, understand, modify
  and maintain.

**TODO**

- Refactor the instruction assembly parsing and formatting into a separate package.
- Multi-section support in assembly.
- Const immediates support in assembly.
- Data section support in assembly.
