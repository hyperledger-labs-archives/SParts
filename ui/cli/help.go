package main

/* This file provides the help text for the main help as well as
 * each command  (e.g., config, supplier, )
 */

/*
 * NOTICE:
 * =======
 *  Copyright (c) 2018 Wind River Systems, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at:
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software  distributed
 * under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES
 * OR CONDITIONS OF ANY KIND, either express or implied.
 */

// --help
const _HELP_CONTENT = `
usage:  sparts <command> [<args>]

e.g., 
  sparts --help
  sparts add --help
  sparts add notices.pdf
  sparts config --list 

These are the <commands>:

General
-------
  about   - Provide an overview about this tool
  help    - Display this message

Manage a sparts work space
--------------------------
  dir	  - Display full paths for the .sparts directory and global config file
  init	  - Create an empty working directory or reinitialize an existing one
 
  delete  - Delete the current sparts working space data.
  version - Version of this utility

Info on the current ledger (see also: sparts help everyday)
  config  - Obtain local configuration info (e.g., ledger name and address, ...	)
  ping    - Ping current ledger node

Creating Compliance Artifacts
-----------------------------
  add     - add an artifact to the staging area prior to commit
  remove  - Remove an artifact to the staging area prior to commit

Ledger Actions
--------------
  supplier - Perform supplier operations
  part     - Perform supplier operations

For more details about each command use the --help option.
For example: 
  sparts add --help
  sparts part --help
  sparts supplier --help
  sparts about
`

const _ABOUT_SPARTS_HELP = `
                      About SParts CLI

The SParts Coomand Line Interface (CLI) allows one to:
  i)   create and register parts and compliance aritfacts for 
       parts on a sparts specified supply chain network;
  ii)  check information that has been registered on the
       ledger;
  iii) create a compliance envelope from a collection of 
       compliance artifacts;

Use the following option to get more details about the different
sparts commands:
   --help

For Example:
  sparts --help  
  sparts status --help (to obtain help details about the 'status' command)
`

// add --help
const _ADD_HELP_CONTENT = `usage: sparts add [--openchain] [<filepath>…​]

Directories will be recursively traverse and all the indentified files will
be added. 
  --openchain, -oc  flag all the artifacts as as being prepared under an OpenChain 
                    conforming program

Examples:
   sparts add notices.pdf  licenses.spdx
   sparts add link=https://github.com/zephyrproject/zephyr/tree/v1.11-branch
   sparts add --openchain notices.pdf   // flag artifacts as being prepared 
                                        // under an OpenChain conforming program
   sparts add --openchain               // flag all artifacts in staging table as being  
                                        // prepared under an OpenChain conforming program
`

const _ALIAS_HELP_CONTENT = `usage: sparts alias <name> <value>
Examples:
   sparts alias --set ibm 8879f843-bd68-4ebd-6cee-3e7fe91e3bcd /* company uuid */

   sparts alias --set debian-14.1.1 0f3d2681-1272-4aaa-7dea-658942dcecfe /* part uuid */

   sparts alias --list  /* List all the alais definitions */
`

const _ARTIFACT_HELP_CONTENT = `usage: sparts artifact [--part <uuid>] <art1> <art2> ... <artN>
Examples:
   sparts artifact --add busybox1.24.2.spdx  // does not assign to part
   sparts 
   sparts artifact --add --part id=zephyr1.12 zephyr1.12.spdx 
   sparts artifact --add --part id=zephyr1.12 notices.txt zephyr1.12.tar.gz
`

// compare --help
const _COMPARE_HELP_CONTENT = `usage: sparts compare  [options]
Examples: 
    sparts compare --help
    sparts compare --dir artifacts1/ --dir artifacts2/ 
    sparts compare --dir artifacts1/ --part <uuid>
    sparts compare --part <uuid1> --part <uuid2>`

// config --help
const _CONFIG_HELP_CONTENT = `usage: sparts config [<options>]

Config file location
  --global		use global config file
  --local		use working directory config file

Action
  --get			get value: name [value-regex]
  -l, --list		list all

Set values:
  <name> value	e.g., config --global user.email bob237@gmail.com
`

// delete --help
const _DELETE_HELP_CONTENT = `  usage: sparts delete
  It will result in the deletion of all the sparts work space data.
  No arguments are expected.`

// dir --help
const _DIRECTORY_HELP_CONTENT = `usage: sparts dir
  Displays the full paths for 
     (1) the .sparts directory and
     (2) the global configuration file
  The SParts directory (1) contains the configuration and repository 
  information for a specific part/repo. The global configuration 
  file (2) holds global data for all all part repos. 
`

// part --help
const _ENVELOPE_HELP_CONTENT = `usage: sparts envelope [<options>]
  -h, --help : display Envelope help
  -l, --list : List envelopes for current working directory.
    e.g., 
      sparts envelope --list
      
  -c, --create:  Create envelopes for the listed directories. 
    e.g.,
        sparts envelope --create usb-driver/
  For Example:
    sparts envelope --create
    sparts envelope --list 
    sparts focus --focus id=zephyr-compliance
`

// focus [none|part|envelope|both]
const _FOCUS_HELP_CONTENT = `usage: sparts focus [none|part|envelope|both]
Focus command allow one to direct sparts to collect artifacts for a envelope 
and part. One can specific one, the other or both. 
  Examples:
    sparts focus --envelope
    sparts focus --part
    sparts focus --both
    sparts focus --none
`

// init --help
const _INIT_HELP_CONTENT = `usage: sparts init
  Create an empty SParts repository or reinitialize an existing one
`

// network help
const _NETWORK_HELP_CONTENT = `usage: sparts network [<options>]
-h, --help : display help
  -l, --list :list all parts in network - e.g., 
      sparts part --list

Examples:
  sparts network --help
  sparts network --list
`

// part --help
const _PART_HELP_CONTENT = `usage: sparts part [<options>]
  -h, --help : display help
  -l, --list :list all parts in network - e.g., 
      sparts part --list
  --all : list all parts in network - e.g., 
      sparts part --list --all
      
  -c, --create  create new part, returns UUID ('short_id' and 'url' are optional)
    e.g.,
        sparts part --create name="ABC" short_id="short_id" url="www.abc.com"
        sparts supplier --create name="ABC"

  --get Get part info. Used local config supplier uuid if not specified.
    e.g.,
        sparts part --get
        sparts part --get <uuid>

  --set uuid=<uuid> - Set the part to work on. This will determine which part
    to assign the artifacts to with commands 'add' and 'commit'. Use 'uuid=none' to 
    clear the part.
    e.g., 
      sparts part --set uuid=5f2a332b-2283-4af4-79ca-6cf50af88273
      sparts part --set uuid=none
      sparts part --list        // list my supplier parts
      sparts part --list --all  // list all the parts on ledger
      sparts part --get id=zephyr1.12   // list details about part one.
`
const _PUSH_HELP_CONTENT = `  usage: sparts push envelope ledger
  Examples:
    sparts push envelope ledger
    sparts push artifacts ledger
  `

const _REMOVE_HELP_CONTENT = `usage: sparts remove  -all|<id>+
  Exampes:
    sparts remove 2
    sparts --all
`

const _STATUS_HELP_CONTENT = `usage: sparts status [--view]
Provides a list of artifacts in the staging area prior to posting to 
the ledger.
  Exampes:
    sparts status
    sparts status --view 4     // view artifact details with table id 4
`

// supplier --help
const _SUPPLIER_HELP_CONTENT = `usage: sparts supplier [<options>]
  -h, --help    display help
  -l, --list	list all suppliers in network
    e.g., 
      sparts supplier --list
  -c, --create  create new supplier, returns UUID ('short_id' and 'url' are optional)
    e.g.,
        sparts supplier --create name="ABC" short_id="short_id" url="www.abc.com"
        sparts supplier --create name="ABC"
  --get Get supplier info. Used local config supplier uuid if not specified.
    e.g.,
        sparts --get
        sparts --get <uuid>
`

// sparts version
const _VERSION_HELP_CONTENT = `usage: sparts version [--all]
  Display version for this instance of sparts.

  Option(s):
    -h, --help : display this help message.
    --all, -a : list version for utility and utility sub components.

  Examples: 
    sparts version
    sparts version --all
    sparts version --help
 `
