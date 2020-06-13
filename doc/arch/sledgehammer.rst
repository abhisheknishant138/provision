.. Copyright (c) 2017 RackN Inc.
.. Licensed under the Apache License, Version 2.0 (the "License");
.. Digital Rebar Provision documentation under Digital Rebar master license
.. index::
  pair: Digital Rebar Provision; Sledgehammer

Sledgehammer Overview
=====================

Sledgehammer is our system discovery/inventory/configuration in-memory
OS image.  It is based on the latest CentOS 7.x live CD, with any
unneeded content stripped out, tools we usually use preinstalled, and
repacked to me more space and network bandwidth efficient.
https://github.com/digitalrebar/provision-content/tree/v4/sledgehammer-builder is
the part of the Git repo that contains the code needed to build Sledgehammer.

Sledgehammer Build
------------------

RackN provides pre-built Sledgehammer images that should function
properly on the vast majority of x86_64 based servers. However, if you
need to include support for custom network hardware, you can rebuild
Sledgehammer to have it include the extra drivers that are needed.

This process should mostly consist of forking the sledgehammer-builder
content bundle and modifying its install steps to add whatever extra
drivers and packages you need, then clean up any extra install dependencies.

Care should be taken to ensure that any extra drivers you install are compatible
with UEFI Secure Boot if you want to be able to operate on systems with
UEFI Secure Boot enabled.

Sledgehammer Design
-------------------

Sledgehammer is delivered in four parts:

- The CentOS signed UEFI bootloader components.  These are required to
  allow sledgehammer to boot on UEFI systems with Secure Boot enabled.

- vmlinuz0: This is the initial kernel that the system PXE boots to
  when booting to Sledgehammer. This file is usually around 5 - 6 megs
  in size.

- stage1.img: This is the first-stage initramfs image.  It is
  responsible for loading any required network drivers and then
  loading the second-stage initramfs over HTTP.  This is a standard
  Linux initramfs, and is usually around 15 megs in size, mainly due
  to including drivers and firmware for large amounts of network
  hardware.

- stage2.img: This is the second-stage initramfs.  It contains the
  final running Sledgehammer image as a highly compressed squashfs
  image.  This image is usually between 400 and 500 megs in size.

The two separate stages allows us to minimize the amount of data that
gets loaded over TFTP and over the relatively inefficient network
stack that the PXE layer in the BIOS and NIC firmware uses.

How Sledgehammer Boots
----------------------

1. The system PXE boots from a dr-provision server, either into the
   discovery bootenv, the sledgehammer bootenv, or some other boot
   environment that boots into Slegehammer.  This causes the system to
   load vmlinuz0 and stage1.img and boot to vmlinuz0.

2. stage1.img loads network drivers for all available network adaptors
   present in the system, finds the one that the system PXE booted
   from, uses DHCP to get the IP address to use, then fetches the
   second-stage image from the location specified by the
   provisioner.web kernel parameter.

3. stage1.img then mounts the squashfs that stage2.img contains,
   mounts an in-memory overlayfs on top of that image to allow normal
   read-write operations, mounts a few other required filesystems, and
   then pivots over to systemd on the stage2 image to allow for a
   normal Centos 7 bootup.

4. Once the network has been configured, The Sledgehammer service
   starts up.  It verifies that the system has an IP address on the
   expected network interface, then fetches the machine-agnostic startup
   script from the dr-provision static file server, verifies that it
   is sane, and executes that script.

5. The machine-agnostic startup script (which must be provided by the
   discovery bootenv and any other bootenv that boots unknown machines
   into Sledgehammer) examines the options that the kernel was booted
   with, sets the system hostname, fetches a copy of drpcli from the
   provisioner, and determines whether it needs to create a machine in
   dr-provision for this system.  If it was not, a new machine is
   created, and has its initial Stage and BootEnv set to the default
   Stage and the default BootEnv.  The machine-specific startup script
   is then downloaded, validated, and executed.

6. The machine-specific startup script (which must be provided the
   sledgehammer bootenv and any bootenv that boots known machines to
   Sledgehammer) then starts the machine agent, which starts executing
   tasks on the machine.

How Sledgehammer Identifies a Machine
-------------------------------------

There are several methods dr-provision can use to identify what machine
it is running on, and determine whether it needs to creater a new Machine.

Using drpcli machines whoami
~~~~~~~~~~~~~~~~~~~~~~~~~~~~

This method, which was introduced in v4.3, uses a combination of hopefully
unique attributes pulled from the system to identify the system.  It gathers
the following information:

1. The system UUID, as provided by the DMI information.
   If present and used, it will have a weight of 50 points.

2. The system serial number, as provided by the DMI information.
   If present and used, it will have a weight of 25 points.

3. The chassis serial number, as provided by the DMI information.
   If present and used, it will have a weight of 25 points

4. The serial number of all the DIMMs installed in the system.
   If present and used, they will have a collective weight of 100 points.

5. The MAC addresses of all the physical nics in the system.
   If present and used, they will have a collective weight of 100 points.

On the client side, this information is gathered, hashes generated for everything
except the MAC addresses, and posted to dr-provision.  On the server side, the fingerprint
is compared to the saved fingerprints for all systems.  The request is considered to match
the system with the highest score that is greater than or equal to 100 points.

If drpocli detects it is running under a hypervisor, it will leave everything in the fingerprint
blank except for the MAC addresses, as vms have even less of a guarantee of having unique
DMI information than physical systems do.

If whoami does not identify the machine, or if it not available on the server side,
sledgehammer will fall back to determining what the machine is based on the machine
UUID sent on the kernel commandline.

Using the kernel command line
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

This is the default fallback method for when whoami is not available or (for now)
when whoami fails to identify a machine.  This method relies on all the
boot environments rendering appropriate MAC address and/or IP address based paths
for each machine.
