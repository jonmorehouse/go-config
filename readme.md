Go Configuration
================

Bootstrap / Basic Usage
-----------------------

* automatically set multiple config vars from environment variables
* convert strings to integers were safe and applicable
* ex:

```
import "github.com/jonmorehouse/go-config/config"

func main() {

  // bootstrap several config elements from environment
  // convert to string/int as needed
  config.Bootstrap(["HOST", "PORT"])

  value, err := config.Get("HOST").(string)
  
  // return just the integer value 
  intValue := config.Value("PORT").(int)

}
```

Fine Tuning Config with Set
---------------------------

* declare default value
* converts to type as determined from default value
* ex:

```
import "github.com/jonmorehouse/go-config/config"

func main() {

  // default value is an int64
  value := int64(5555)

  config.Set("HOST", value)

  config.Get("HOST").(int64)

}
```

