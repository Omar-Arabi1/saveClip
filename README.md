# saveClip
a clipboard manager tool to save only the things you care about —*this is a project I already made in python and I am remaking in golang for better performance and personal practice*

## table of contents:
- [overview](#overview)
- [save command](#save-command)
- [access command](#access-command)
- [remove command](#remove-command)
- [view command](#view-command)
- [search command](#search-command)
- [notes](#notes)

## overview:
you will mostly do these commands the `save` and `access` commands
the `save` command will save the entery at the top of your clipboard history
you will have to provide it with a label as an argument


you can also set an optional priority for the entery using the `--priority` option
and giving it a number the number entered will be in this range 1 min - 3 max
1 means its not very important 3 means its very important


by default the entery will be saved with a priority of 1 unless you set it manually
to something else using the option, Example


`saveClip save <label> --priority 3`


this will save the entery at the top of your clipboard with the label given and a priority
level of 3, *note that all the operations like searching or viewing which will be discussed later
so its recommended to save with a short one word descriptive label*


*note that we will call the entery which means that the saved entery with the label and priority will 
be called a **Clip** so the saved entery at the top example is called a Clip*
