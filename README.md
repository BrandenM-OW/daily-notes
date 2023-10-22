<a name="readme-top"></a>

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]



<br />
  <p align="center">
    Daily notes is a CLI tool for generating note files based on templates and can preserve items day over day.
    <br />
    <a href="https://github.com/BrandenM-OW/sage/issues">Report Bug</a>
    ·
    <a href="https://github.com/BrandenM-OW/sage/issues">Request Feature</a>
  </p>
</div>



<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#license">License</a></li>
  </ol>
</details>



## About The Project

Daily notes allows you to maintain a customized template in Markdown to generate your notes from and allows you to preseve items from the previous day.

For example:

The base template will include a Tasks section:



## Tasks
- [x] Task 1
- [ ] Task 2

When you preseve with the -p flag, any unchecked tasks will be brought to the next day.

Any line prefixed with a -p in the note will also be preserved. 

For example:
```sh
- item 1
-p preserved item 3
- item 3
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>


### Built With

* <a href="https://github.com/spf13/cobra">Cobra CLI</a>



## Getting Started

To get started you will need the following installed on your system:


### Prerequisites

* Go


### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/BrandenM-OW/daily-notes.git
   ```
2. Build the project 
   ```sh
   cd path/to/repo/
   go build main.go
   ```
3. Add the binary to your system path if you so choose.

4. Binary downloads to come so no build will be required.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



## Usage

### init
```sh
daily-notes init
```

Creates the following directory structure:
```sh
root/
├── config.yaml
├── notes
│  └── month
│    └── week
│      └── day.md
└── templates
	└── daily.md
```


### New

```sh
daily-notes new
daily-notes new -p // Preserves notes from last day
daily-notes new -d 1 // Creates a note for tomorrow
```

Adds the following file:
```sh
root/
└── notes
    └── month
        └── week
            └── day.md
```

The day.md file will be created based on the template specified in the config.yaml file.

If the preserve flag is set, the last day.md file will be used as a template for the new day.md file.
	- All unchecked tasks will be copied over
	- All notes will be copied over that are not marked with the -p flag


<p align="right">(<a href="#readme-top">back to top</a>)</p>



## Roadmap

- [ ] Adding binary downloads to repo.
- [ ] Template flag to use other templates.
- [ ] Better preseve flag options to allow you to pick a day to preseve from.

See the [open issues](hhttps://github.com/BrandenM-OW/daily-notes/issues) for a full list of proposed features (and known issues).



## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>


[contributors-shield]: https://img.shields.io/github/contributors/BrandenM-OW/daily-notes.svg?style=for-the-badge
[contributors-url]: hhttps://github.com/BrandenM-OW/daily-notes/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/BrandenM-OW/daily-notes.svg?style=for-the-badge
[forks-url]: hhttps://github.com/BrandenM-OW/daily-notes/network/members
[stars-shield]: https://img.shields.io/github/stars/BrandenM-OW/daily-notes.svg?style=for-the-badge
[stars-url]: hhttps://github.com/BrandenM-OW/daily-notes/stargazers
[issues-shield]: https://img.shields.io/github/issues/BrandenM-OW/daily-notes.svg?style=for-the-badge
[issues-url]: hhttps://github.com/BrandenM-OW/daily-notes/issues
[license-shield]: https://img.shields.io/github/license/BrandenM-OW/daily-notes.svg?style=for-the-badge
[license-url]: hhttps://github.com/BrandenM-OW/daily-notes/blob/master/LICENSE.txt
