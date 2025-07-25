# saveClip
a clipboard manager tool to save only the things you care about —*this is a project I already made in python and I am remaking in golang for better performance and personal practice*


install with `go install github.com/Omar-Arabi1/saveClip@latest`


this app works on both windows, macOS and Linux

## table of contents:
- [overview](#overview)
- [save command](#save-command)
- [access command](#access-command)
- [remove command](#remove-command)
- [view command](#view-command)
- [search command](#search-command)
- [notes](#notes)
- [dependencies](#dependencies)

## overview:
you will mostly use these commands the `save` and `access` commands
the `save` command will save the entery at the top of your clipboard history
you will have to provide it with a label as an argument


you can also set an optional priority for the entery using the `--priority` option
and giving it a number the number entered will be in this range 1 min - 3 max
1 means its not very important 3 means its very important


by default the entery will be saved with a priority of 1 unless you set it manually
to something else using the option, the tool will provide the clip with a creation date
YYYY-MM-DD this will have some use cases in other commands, Example


`saveClip save <label> --priority 3`


this will save the entery at the top of your clipboard with the label given and a priority
level of 3, *note that all the operations like [removing](#remove-command) or [viewing](#view-command) which will be discussed later
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

## save command:
view the [overview](#overview) section at the top

## access command:
view the [overview](#overview) section at the top

## remove command:
this command will allow you to remove a clip from your list by providing it with a label as an argument
you could also use one of two options `--all` this will remove the entire list at once while `--remove-priority`
will remove all the clips with the given priority, Example


`saveClip remove <label>`


this will remove the clip at the given label


`saveClip remove --all`


this will remove all the clips at once


`saveClip remove --remove-priority 3`


this will remove all clips with the priority of 3

## view command:
this will show to you all clips, enteries, creation dates, labels and priorities
it takes in nothing, but you could provide it with one of two options `--priority`
this takes in `highest` or `lowest` and it will show to you the list sorted by priority


it also takes in a `--date` option which takes in `newest` or `oldest` this will show to you
the list sorted by date, Example


`saveClip view`


this will show to you all the clips with the order they are saved in


`saveClip view --priority highest`


this will show to you all the clips sorted from highest to lowest by priority


`saveClip view --date newest`


this will show to you all the clips sorted from newest to oldest by date

## search command:
this will allow you to search through all your labels with a given query as an argument
the search will be a fuzzy search taking only part of the label and showing to you the label and 
what is saved inside of it,


the command could also take a `--filter` option which takes in a date YYYY-MM-DD which will show to
you only results that match the query and has the creation date you entered


`saveClip search <query>`


this will show to you all the labels that match the query and what is saved inside of them


`saveClip search <query> --filter 2025-07-25`


this will show to you all the labels that match the query and have the creation date 2025-07-25

## notes:
the app saves your clips inside of a hidden file in the home directory named clips.json

## dependencies:
this app doesn't have dependencies on windows or macOS, but for linux there are dependencies
if you are on X11 you must install one of the following `libx11-dev` or `xorg-dev` or `libX11-devel`, but
if you are on wayland the app isn't supported so you must use XWayland with one of the dependencies on top installed
and make sure the DISPLAY variable is set, to make sure it is already set run `echo $DISPLAY` if it shows :0 or :1
it is set if not it will not show anything

---

## Thanks for installing the tool