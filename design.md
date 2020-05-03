api/
    - define Task/Project types that only contain info needed for API calls
    - create an api.Client type? - DONE
        - initialization would check for TODOIST_TOKEN

command/
    - define Command interface? - DONE
        - Name
        - Validate
        - Execute(*Context, *api.Client, args ...string)
    - define Parser type - DONE
        - keeps map[cmdName]Command
        - Parse(name) Command

todolist/
    - define Task/Project types - skeletons done

app/
    - define App type to wrap grumble.App

TODO:
    - how to hook up commands into main application?