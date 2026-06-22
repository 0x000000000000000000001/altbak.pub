---
description: Clean les imports
---

Cleanes les imports inutiles. E.g. 

src/Inter/Ui/Router/Menu/Core/Search/Search.purs:14:1

  14  import Inter.Ui.Router.Menu.Type (Action(..), Slots)
      ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

  The import of type Action from module Inter.Ui.Router.Menu.Type includes data constructors but only the type is used
  It could be replaced with:
    import Inter.Ui.Router.Menu.Type (Action, Slots)

Pour cela, tu pourrais avoir besoin de compiler le code (voir .agents/workflows/compile.md)