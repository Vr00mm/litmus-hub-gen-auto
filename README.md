# Litmus helm

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

Auto Hub Generator for Litmus
## Usage

### PreRequisites

- A Connection to github
- Kubernetes >= 1.17

### Installation Steps

The following steps will help you to generate privates hub.

#### Step-1: Download the hub-gen-auto latest release

```bash
curl https://github.com/vr00mm/litmus-
```

#### Step-2: Execute the hub-gen-auto

- The hub will be generate into new "build" directory.

**Note**: An active connection to https://github.com/litmuschaos/chaos-chart/* is required.

```bash
hub-gen-auto
```

#### Step-3: Push the hub to your chaos-chart private repository

Thats all, the hub is generated and you just have to push the subdirectory named with cluster as hub to your repo.
No configuration is available yet ( it take all default value for each experiments )

### Incoming

- Customize experiments parameters globaly (from cli) (like container-runtime, experiment library etc)
- Customize experiments parameters (from an input manifest) (like icons, default duration chaos etc etc)
- Customize component experiments (chaos duration, percent affect, etc)

- Improve documentation generation and links (can be cool to get the litmusportal chaoscenter direct link)
- Auto push hubs to git
- Auto add hubs to litmus project

## License

[Apache 2.0 License](./LICENSE).
