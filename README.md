<div align="center">
<img alt="Кошка" align="center" src="header.png"/>
<br/>
<h2>Meow-ster of cat-alogs!</h2>
</div>
<br/>
<div align="justify">

*Кошка* (pronounced *kóška*; means *cat* in Russian)  is a work-in-progress application
that I am making for fun. It aims to be an app to interact with the catalogs for pulsars,
FRBs (**F**ast **R**adio **B**ursts) and RRATs (**R**otational **RA**dio **T**ransients).
It will packaged as a binary that can be run from anywhere, and will have a terminal user
interface (that is, a TUI) which will allow the user to search the above databases right
from the command line.

At the moment, there is not much here. There is a simple dummy TUI that does not do much,
and some functions that allows one to search the pulsar database via the pulsar's `JNAME`
property. Right now only the pulsar database has been setup (and can be found in the `data`
directory.

To check it out, just clone the repository:

```bash
git clone https://github.com/astrogewgaw/koshka
```

and then run:

```bash
go run .
```

Of course, you need to have Go installed on your system.

</div>
