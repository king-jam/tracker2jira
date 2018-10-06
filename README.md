<!--
This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this file,
You can obtain one at http://mozilla.org/MPL/2.0/.

Copyright (c) 2017, MPL James R. King james.r.king4[at]gmail.com
-->
## TODO: Work in progress. I do not recommend using this in production at all.

# **Tracker 2 Jira** [![Build Status](https://travis-ci.org/king-jam/tracker2jira.svg?branch=master)](https://travis-ci.org/king-jam/tracker2jira)  [![Go Report Card](https://goreportcard.com/badge/github.com/king-jam/tracker2jira)](https://goreportcard.com/report/github.com/king-jam/tracker2jira)

**Under Active Development, not all features are currently supported**

**Tracker 2 Jira** is a unidirectional mirror service for those who want to use Pivotal Tracker (public or private) but are required to use Atlassian Jira internal to an organization where traditional hooks don't work or custom workflows exist.

----

1. [Overview](#overview)
2. [Getting Started](#getting-started)
   * [System Requirements](#system-requirements)
   * [Installation](#installation)
   * [Running Tracker2JIRA](#running-t2j)
   * [Tracker2JIRA Configuration](#t2j-config)
3. [Documentation](#documentation)
4. [Contributing](#contributing)
5. [License](#license)

## Overview

* Mirror epic attributes (State/Stories/Progress) - Under Development
* Mirror story attributes (State/Comments/Attachments) - Under Development
* Mirror iterations (Pointing/Timelines) - Under Development

## Getting Started

### System Requirements

None, currently Tracker 2 Jira is self-contained. It uses BoltDB as the backing store with future goals to support configuration of the storage directory. We are leveraging Docker's libkv library for now so Bolt, ETCD, Consul, or Zookeeper are supported.

### Installation

### Running Tracker2JIRA

### Tracker2JIRA Configuration

## Documentation

## Contributing

## License
Tracker2JIRA is Open Source software released under the [Mozzila Public License 2.0](LICENSE).
