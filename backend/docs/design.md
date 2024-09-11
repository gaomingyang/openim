# Design of Open IM

## system architect
This system should be distributed and scalable

The user interface can be varied. We should provide sdk for user customization their interface.

## Main Modules

### Group chat
The first version, all user can join group and send message to it.
In the latter version, user should be approved by group organizer,then join the group successful.

some apis
* create group
* join group
* init group chat( create ws connection)
* send group message
* mute group message
* quite group 


### One-to-one Chat
Only after add each other as friends, can user send message to other one-to-one



db move to server