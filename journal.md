**2014.3.3**

Trying travis.yml

**2014.2.26**

Have a simple working assembler. TODO: add some constants, and string decl. like:

    ; data segments
    .string msg "Hello, world.\n\000'
    .uint8 endl '\n'
    .int16 start 0x2000
    .int32 magic 0x32323322

    .func main ; just a label that declares a label namespace, nothing special
    ; then our code here

