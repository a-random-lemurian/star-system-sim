# star-system-sim

Also known as `spaceport`, though for some weird reason there is a name
discrepancy.

It is a library that simulates the flow of ships through a series of star
systems.

## Architecture

|---|---|
| Module | Purpose
|---|---|
| enumerate | Simply allows you to view all possible default name options.
| metrics | Reports on all that is occurring in the simulation.
| shipinfogen | Services for creating ships.
| shipnamegen | Services for generating ship names. Comes with a default name set.
| starsystem | The "meat" of the simulation, where ships move between systems through a bunch of channels.
