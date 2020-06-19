# dec2hex

This is a simple utility for rapidly converting decimal numbers (base10) to
their hexadecimal (base16) representation. By default the command adds a newline
to the output, and this can be omitted by providing the `-n` flag.

If you'd like to convert from hexadecimal to decimal, either create a copy of
the binary named `h2d` or symlink it:

```
ln -sf ./dec2hex ./h2d
```

I personally have two symlinks: `d2h` for a short-handed dec2hex, and `h2d` for
a short-handed hex2dec.
