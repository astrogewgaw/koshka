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
contacting me directly (details on my home page).

### Usage

Right now, the only way to run **Кошка** is:

```bash
git clone https://github.com/astrogewgaw/koshka
cd koshka
go run .
```

where I have assumed that you have the latest version of Go installed. If not, you can
check out how to install Go [**here**][go_install].

### Preview

<img
    align="center"
    src="koshka.gif"
    alt="Кошка Preview GIF"
/>

### Acknowledgements

This project could not have been built without the [**psrqpy**][psrqpy] package, written in
Python and developed by [**Matt Pitkin**][mattpitkin], or the entire ecosystem of libraries
created by the [**Charm Project**][charm], such as the [**Bubbletea**][bubbletea] library,
[**Bubbles**][bubbles] and [**Lipgloss**][lipgloss]. I am also thankful for the amazing
[**bubble-table**][bubble-table] library, developed by [**Brandon Fulljames**][Evertras],
that gave me a super easy way to display the data in beautifully formatted, paginated
tables, right there in the terminal.

[charm]: https://charm.sh
[gitmoji]: https://gitmoji.dev
[Evertras]: https://github.com/Evertras
[go_install]: https://go.dev/doc/install
[mattpitkin]: https://github.com/mattpitkin
[psrqpy]: https://github.com/mattpitkin/psrqpy
[bubbles]: https://github.com/charmbracelet/bubbles
[lipgloss]: https://github.com/charmbracelet/lipgloss
[atnf]: https://www.atnf.csiro.au/people/pulsar/psrcat
[bubbletea]: https://github.com/charmbracelet/bubbletea
[bubble-table]: https://github.com/Evertras/bubble-table
[stars]: https://img.shields.io/github/stars/astrogewgaw/koshka?style=for-the-badge
[license]: https://img.shields.io/github/license/astrogewgaw/koshka?style=for-the-badge
[gitmoji-badge]: https://img.shields.io/badge/gitmoji-%20😜%20😍-FFDD67.svg?style=for-the-badge
