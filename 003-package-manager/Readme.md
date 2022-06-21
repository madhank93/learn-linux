### What is a Package ?

- An archive file containing the binary executable, configuration file and information about the dependencies. (ex: Installing chrome requires few other dependencies to be installed first)

### What is a Package manager ?

- Is a tool that allows users to install, remove, upgrade, configure and manage software packages on an operating system
- Manges and resolves all required dependencies
- upgrading the software can be easily done

Note:

- Package manager comes builtin in every Linux distribution. In Ubuntu, **APT (Advanced Package Tool)** package manager available along with **APT-GET** package manager. APT is the preferred and user friendly to use
- Package manager always gets the applications list from the repository
- APT configurations are available at `/etc/apt`
- Some times applications (browsers, IDEs) might not be available in APT package manager,
  - as an alternative **Snap package manager** or Ubuntu Software Center
  - another approach is to Add repository to APT using `add-apt-repository` command which will add it to `/etc/apt/source.list`
- Debian based linux distribution (Ubuntu, Debian, Mint) uses **APT** package manager
- Red Hat based linux distribution (RHEL, CentOS, Fedora) uses **YUM** package manager
