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
you will mostly use these commands the `save` and `access` commands
the `save` command will save the entery at the top of your clipboard history
you will have to provide it with a label as an argument


you can also set an optional priority for the entery using the `--priority` option
and giving it a number the number entered will be in this range 1 min - 3 max
1 means its not very important 3 means its very important


by default the entery will be saved with a priority of 1 unless you set it manually
to something else using the option, Example


`saveClip save <label> --priority 3`


this will save the entery at the top of your clipboard with the label given and a priority
level of 3, *note that all the operations like removing or viewing which will be discussed later
use the label as eitehr arguemnts or options, so its recommended to save with a short one word descriptive 
label*


*note that we will call the entery which means that the saved entery with the label and priority will 
be called a **Clip** so the saved entery at the top example is called a Clip*


as for the `access` command this will put a clip to your clipboard history, by default it will put
the last entered clip to your clipboard history, you could override this behaviour by providing it with
a label to a clip you want to put to your clipboard history using the `--label` option, Example


`saveClip access`


this will put the last entered clip to your clipboard history


`saveClip access -l <label>`

this will put the clip that has the entered label to your clipboard history

## save command
view the [overview](#overview) section at the top

## access command
view the [overview](#overview) section at the top