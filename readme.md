Go Configuration
================

```
import (

  "github.com/jonmorehouse/go-config/config"

)

func (Config config) Bootstrap() {

  config.env("", DEFAULT)

}

func main() {

  Config.Bootstrap()


}

```

* creates a config singleton
* implement your own bootstrap method on the interface
  



