<div align="center">
<img
    alt="Кошка"
    align="center"
    src="https://raw.githubusercontent.com/astrogewgaw/logos/main/rasters/koshka.png"
/>
<br/>
<h3>Meow-ster of cat-alogs 🐱 !</h3>

![License][license]
![GitHub Stars][stars]
[![Gitmoji][gitmoji-badge]][gitmoji]

</div>
<br/>
<div align="justify">

**Кошка** (pronounced *koshka*; means **cat** in Russian) is a terminal user interface
(TUI) for searching through catalogs of radio transients: pulsars, FRBs, & RRATs. I
am making this for fun, so expect to find cat puns everywhere. Currently, I have only
written code around the [**ATNF Pulsar Database**][atnf], but I expect to add support
for other objects soon. Since this is a Go binary, it should run almost anywhere, but
the binary itself has to be compiled with the latest release of Go. If (and only if)
there is sufficient interest, I will release this project properly, so that people can
install it via their respective package managers. So, if you use this and find it fun,
let me know by staring the repository :star:, spreading the word :speech_balloon:, and/or
contacting me directly.

Right now, the only way to run **Кошка** is:

```bash
go run .
```

where I have assumed that you have the latest version of Go installed. If not, you can check
out how to install Go [**here**][go_install].

[gitmoji]: https://gitmoji.dev
[go_install]: https://go.dev/doc/install
[atnf]: https://www.atnf.csiro.au/people/pulsar/psrcat
[stars]: https://img.shields.io/github/stars/astrogewgaw/koshka?style=for-the-badge
[license]: https://img.shields.io/github/license/astrogewgaw/koshka?style=for-the-badge
[gitmoji-badge]: https://img.shields.io/badge/gitmoji-%20😜%20😍-FFDD67.svg?style=for-the-badge
