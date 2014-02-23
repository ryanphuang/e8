** Road Plan **

- Build a MISP-like very simple virtual machine
- Build an assembler
- Build a compiler for a Go-like language
- Reimplement all the above in Javascript
- Build an online platform where people can submit interfaces, test cases, and implementations
- Each module import will have its own data, but no global shared variable
- Test cases are like executable test modules
- Interfaces and implementations are general lib modules
- Can only access other modules via publicly defined interfaces, methods, functions, types, consts
- This is just for single vm open source testing
- For more complex stuff, we can have multiple vm testing, where all calls are performed over RPC calls