# Developer Notes

A script suite for keeping developer notes.

Read more about my note keeping workflow in [this blog post](https://cochrane343.wordpress.com/2017/09/10/my-developer-notes/).

If you choose to fork this repository instead of simply cloning it, please keep in mind that the archive folder, which contains the notes by default, is part of the git repository. Therefore you are at risk of pushing possibly sensitive notes to a public Github repo.

## Scripts

The [new day script](https://github.com/cochrane343/notes/blob/master/scripts/new_day.go) creates a new note file in the archive folder, pre populates it with some template data and opens it in a text editor. Don't forget to cut the template section before committing the note after finishing writing it.

Example use and command alias:
```
go run new_day.go

alias newday='(cd ~/some/path/notes/scripts/ && go run new_day.go)'
```

The [last week script](https://github.com/cochrane343/notes/blob/master/scripts/last_week.go) prints a report generated from the most recent notes while skipping some entry types. This script can be used before starting development work to get oneself back into context and to refresh one's memory about loose ends from the last days.

Example use and command alias:
```
go run last_week.go

alias lastweek='(cd ~/some/path/notes/scripts/ && go run last_week.go)'
```
