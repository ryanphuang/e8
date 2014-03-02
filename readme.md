**Road Plan**

- Build a MIPS-like very simple virtual machine
- Build an assembler
- Build a compiler for a Go-like language `e8go`
- Port the assembler and compiler to `e8go`.
- Reimplement the VM in Javascript. (Maybe NaCl for golang will come out at that time.)
- Write a small OS in `e8go`, so that is runs in the browser.
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

See more on my [Motivation](https://github.com/h8liu/e8/wiki/Motivation) page.

**Design**

- [RISC](https://github.com/h8liu/e8/wiki/RISC-Specification): The MIPS-like
  simple instructions set that `e8` uses only has less than 40 instructions,
  which means `e8` CPU will be very easy to port. In fact, I have already
  ported the core to Javascript: [`e8js`](https://github.com/h8liu/e8js).
- System Page (wiki page coming soon): all IOs in `e8` will be memory mapped,
  so there is no need for special instructions like `in` and `out`. Basic
  system functionality will be mapped to Page 0.  Future fancy hardware will be
  mapped to the following small-id pages in the address space.
- Multi-core and ring protection (future plan): `e8` will not have protection
  rings (e.g. kernel mode and user mode). Instead, it will use an approach
  similar to ARM's TrustZone, where there will be a previledged VM that can
  manipulate other child VMs' execution state and page tables.
- Interrupts (future plan): Inside a VM, there will be events, but there will
  be no interrupt handlers. A VM can suspend itself and wake up on an interrupt
  event, so a VM can be event driven, but code execution will not be forced to
  suspend and continue at another program counter. A previledged VM, however,
  can simulate interrupt handling on its child VMs, if it is desired to.
- Language support: The project (simulator, assembler, compiler) is written in
  golang, and I plan to implement a subset of golang that compiles to `e8`, the
  assembler and the compiler will be ported to that subset language later in
  the future, and there will be an advanced language compiler that runs in `e8`
  where it can compile itself.

**TODO**

- Const immediates support in assembly.
- Data section support in assembly.
- Assemble a project of multiple files.
