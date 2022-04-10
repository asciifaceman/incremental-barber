# Design Doc Sortof

A short incremental

A new pandemic is afoot! But this one is a bit weird...

The only side-effect of the H virus is hair that grows phenomenally fast. Due to this, most people require haircuts at least weekly, and the barber industry takes off like a rocket.

# General gist
Cut peoples hair

Quality of cut depends on tools

earn a base rate + tip

Customer satisfaction determines tip, defined by tools and environment

improve environment by getting better shops 

add more barbers to handle more cuts at once

# Customer

Each customer will have a few values that determine the outcome

```
type Customer struct {
    Name string
    Satisfaction float
    Comfort float
}
```

# Tools unlockable

In no particular order. Each tool would have a modifier assigned

* Tools
    * Scissors (starting)
    * Cape (+comfort)
    * Powder (+comfort)
    * Stereo (+comfort)
    * Padded Chair (+comfort)
    * Brush (+comfort)
    * Comb (+quality)
    * Sink (+quality)
    * Electric Trimmers (+quality)
    * Razor (+quality)
    * 
* Shop
    * Booking App (+autowork)
    * Fishtank (+global comfort)